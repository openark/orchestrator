package memberlist

import (
	"log"
	"net"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTransport_Join(t *testing.T) {
	net := &MockNetwork{}

	t1 := net.NewTransport()

	c1 := DefaultLANConfig()
	c1.Name = "node1"
	c1.Transport = t1
	c1.Logger = testLogger(t)
	m1, err := Create(c1)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	m1.setAlive()
	m1.schedule()
	defer m1.Shutdown()

	c2 := DefaultLANConfig()
	c2.Name = "node2"
	c2.Transport = net.NewTransport()
	c2.Logger = testLogger(t)
	m2, err := Create(c2)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	m2.setAlive()
	m2.schedule()
	defer m2.Shutdown()

	num, err := m2.Join([]string{t1.addr.String()})
	if num != 1 {
		t.Fatalf("bad: %d", num)
	}
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if len(m2.Members()) != 2 {
		t.Fatalf("bad: %v", m2.Members())
	}
	if m2.estNumNodes() != 2 {
		t.Fatalf("bad: %v", m2.Members())
	}

}

func TestTransport_Send(t *testing.T) {
	net := &MockNetwork{}

	t1 := net.NewTransport()
	d1 := &MockDelegate{}

	c1 := DefaultLANConfig()
	c1.Name = "node1"
	c1.Transport = t1
	c1.Delegate = d1
	c1.Logger = testLogger(t)
	m1, err := Create(c1)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	m1.setAlive()
	m1.schedule()
	defer m1.Shutdown()

	c2 := DefaultLANConfig()
	c2.Name = "node2"
	c2.Transport = net.NewTransport()
	c2.Logger = testLogger(t)
	m2, err := Create(c2)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	m2.setAlive()
	m2.schedule()
	defer m2.Shutdown()

	num, err := m2.Join([]string{t1.addr.String()})
	if num != 1 {
		t.Fatalf("bad: %d", num)
	}
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if err := m2.SendTo(t1.addr, []byte("SendTo")); err != nil {
		t.Fatalf("err: %v", err)
	}

	var n1 *Node
	for _, n := range m2.Members() {
		if n.Name == c1.Name {
			n1 = n
			break
		}
	}
	if n1 == nil {
		t.Fatalf("bad")
	}

	if err := m2.SendToUDP(n1, []byte("SendToUDP")); err != nil {
		t.Fatalf("err: %v", err)
	}
	if err := m2.SendToTCP(n1, []byte("SendToTCP")); err != nil {
		t.Fatalf("err: %v", err)
	}
	if err := m2.SendBestEffort(n1, []byte("SendBestEffort")); err != nil {
		t.Fatalf("err: %v", err)
	}
	if err := m2.SendReliable(n1, []byte("SendReliable")); err != nil {
		t.Fatalf("err: %v", err)
	}
	time.Sleep(100 * time.Millisecond)

	expected := []string{"SendTo", "SendToUDP", "SendToTCP", "SendBestEffort", "SendReliable"}

	msgs1 := d1.getMessages()

	received := make([]string, len(msgs1))
	for i, bs := range msgs1 {
		received[i] = string(bs)
	}
	// Some of these are UDP so often get re-ordered making the test flaky if we
	// assert send ordering. Sort both slices to be tolerant of re-ordering.
	require.ElementsMatch(t, expected, received)
}

type testCountingWriter struct {
	t *testing.T
	numCalls *int32
}

func (tw testCountingWriter) Write(p []byte) (n int, err error) {
	atomic.AddInt32(tw.numCalls, 1)
	if !strings.Contains(string(p), "memberlist: Error accepting TCP connection") {
		tw.t.Error("did not receive expected log message")
	}
	tw.t.Log("countingWriter:", string(p))
	return len(p), nil
}

// TestTransport_TcpListenBackoff tests that AcceptTCP() errors in NetTransport#tcpListen()
// do not result in a tight loop and spam the log. We verify this here by counting the number
// of entries logged in a given time period.
func TestTransport_TcpListenBackoff(t *testing.T) {

	// testTime is the amount of time we will allow NetTransport#tcpListen() to run
	// This needs to be long enough that to verify that maxDelay is in force,
	// but not so long as to be obnoxious when running the test suite.
	const testTime = 4 * time.Second

	var numCalls int32
	countingWriter := testCountingWriter{t, &numCalls}
	countingLogger := log.New(countingWriter, "test", log.LstdFlags)
	transport := NetTransport{
		streamCh: make(chan net.Conn),
		logger: countingLogger,
	}
	transport.wg.Add(1)

	// create a listener that will cause AcceptTCP calls to fail
	listener, _ := net.ListenTCP("tcp", nil)
	listener.Close()
	go transport.tcpListen(listener)

	// sleep (+yield) for testTime seconds before asking the accept loop to shut down
	time.Sleep(testTime)
	atomic.StoreInt32(&transport.shutdown, 1)

	// Verify that the wg was completed on exit (but without blocking this test)
	// maxDelay == 1s, so we will give the routine 1.25s to loop around and shut down.
	c := make(chan struct{})
	go func() {
		defer close(c)
		transport.wg.Wait()
	}()
	select {
	case <- c:
	case <-time.After(1250 * time.Millisecond):
		t.Error("timed out waiting for transport waitgroup to be done after flagging shutdown")
	}

	// In testTime==4s, we expect to loop approximately 12 times (and log approximately 11 errors),
	// with the following delays (in ms):
	//   0+5+10+20+40+80+160+320+640+1000+1000+1000 == 4275 ms
	// Too few calls suggests that the minDelay is not in force; too many calls suggests that the
	// maxDelay is not in force or that the back-off isn't working at all.
	// We'll leave a little flex; the important thing here is the asymptotic behavior.
	// If the minDelay or maxDelay in NetTransport#tcpListen() are modified, this test may fail
	// and need to be adjusted.
	require.True(t, numCalls > 8)
	require.True(t, numCalls < 14)

	// no connections should have been accepted and sent to the channel
	require.Equal(t, len(transport.streamCh), 0)
}

