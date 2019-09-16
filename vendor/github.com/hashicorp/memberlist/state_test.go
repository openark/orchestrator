package memberlist

import (
	"bytes"
	"fmt"
	"net"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func HostMemberlist(host string, t *testing.T, f func(*Config)) *Memberlist {
	t.Helper()

	c := DefaultLANConfig()
	c.Name = host
	c.BindAddr = host
	c.BindPort = 0 // choose a free port
	c.Logger = testLoggerWithName(t, host)
	if f != nil {
		f(c)
	}

	m, err := newMemberlist(c)
	if err != nil {
		t.Fatalf("failed to get memberlist: %s", err)
	}
	return m
}

func TestMemberList_Probe(t *testing.T) {
	addr1 := getBindAddr()
	addr2 := getBindAddr()

	m1 := HostMemberlist(addr1.String(), t, func(c *Config) {
		c.ProbeTimeout = time.Millisecond
		c.ProbeInterval = 10 * time.Millisecond
	})
	defer m1.Shutdown()

	bindPort := m1.config.BindPort

	m2 := HostMemberlist(addr2.String(), t, func(c *Config) {
		c.BindPort = bindPort
	})
	defer m2.Shutdown()

	a1 := alive{
		Node:        addr1.String(),
		Addr:        []byte(addr1),
		Port:        uint16(m1.config.BindPort),
		Incarnation: 1,
		Vsn:         m1.config.BuildVsnArray(),
	}
	m1.aliveNode(&a1, nil, true)
	a2 := alive{
		Node:        addr2.String(),
		Addr:        []byte(addr2),
		Port:        uint16(m2.config.BindPort),
		Incarnation: 1,
		Vsn:         m2.config.BuildVsnArray(),
	}
	m1.aliveNode(&a2, nil, false)

	// should ping addr2
	m1.probe()

	// Should not be marked suspect
	n := m1.nodeMap[addr2.String()]
	if n.State != stateAlive {
		t.Fatalf("Expect node to be alive")
	}

	// Should increment seqno
	if m1.sequenceNum != 1 {
		t.Fatalf("bad seqno %v", m2.sequenceNum)
	}
}

func TestMemberList_ProbeNode_Suspect(t *testing.T) {
	addr1 := getBindAddr()
	addr2 := getBindAddr()
	addr3 := getBindAddr()
	addr4 := getBindAddr()
	ip1 := []byte(addr1)
	ip2 := []byte(addr2)
	ip3 := []byte(addr3)
	ip4 := []byte(addr4)

	m1 := HostMemberlist(addr1.String(), t, func(c *Config) {
		c.ProbeTimeout = time.Millisecond
		c.ProbeInterval = 10 * time.Millisecond
	})
	defer m1.Shutdown()

	bindPort := m1.config.BindPort

	m2 := HostMemberlist(addr2.String(), t, func(c *Config) {
		c.BindPort = bindPort
	})
	defer m2.Shutdown()
	m3 := HostMemberlist(addr3.String(), t, func(c *Config) {
		c.BindPort = bindPort
	})
	defer m3.Shutdown()
	a1 := alive{Node: addr1.String(), Addr: ip1, Port: uint16(bindPort), Incarnation: 1, Vsn: m1.config.BuildVsnArray()}
	m1.aliveNode(&a1, nil, true)
	a2 := alive{Node: addr2.String(), Addr: ip2, Port: uint16(bindPort), Incarnation: 1, Vsn: m2.config.BuildVsnArray()}
	m1.aliveNode(&a2, nil, false)
	a3 := alive{Node: addr3.String(), Addr: ip3, Port: uint16(bindPort), Incarnation: 1, Vsn: m3.config.BuildVsnArray()}
	m1.aliveNode(&a3, nil, false)
	a4 := alive{Node: addr4.String(), Addr: ip4, Port: uint16(bindPort), Incarnation: 1, Vsn: m1.config.BuildVsnArray()}
	m1.aliveNode(&a4, nil, false)

	n := m1.nodeMap[addr4.String()]
	m1.probeNode(n)

	// Should be marked suspect.
	if n.State != stateSuspect {
		t.Fatalf("Expect node to be suspect")
	}
	time.Sleep(10 * time.Millisecond)

	// One of the peers should have attempted an indirect probe.
	if m2.sequenceNum != 1 && m3.sequenceNum != 1 {
		t.Fatalf("bad seqnos %v, %v", m2.sequenceNum, m3.sequenceNum)
	}
}

func TestMemberList_ProbeNode_Suspect_Dogpile(t *testing.T) {
	cases := []struct {
		name          string
		numPeers      int
		confirmations int
		expected      time.Duration
	}{
		{"n=2, k=3 (max timeout disabled)", 1, 0, 500 * time.Millisecond},
		{"n=3, k=3", 2, 0, 500 * time.Millisecond},
		{"n=4, k=3", 3, 0, 500 * time.Millisecond},
		{"n=5, k=3 (max timeout starts to take effect)", 4, 0, 1000 * time.Millisecond},
		{"n=6, k=3", 5, 0, 1000 * time.Millisecond},
		{"n=6, k=3 (confirmations start to lower timeout)", 5, 1, 750 * time.Millisecond},
		{"n=6, k=3", 5, 2, 604 * time.Millisecond},
		{"n=6, k=3 (timeout driven to nominal value)", 5, 3, 500 * time.Millisecond},
		{"n=6, k=3", 5, 4, 500 * time.Millisecond},
	}

	for i, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// Create the main memberlist under test.
			addr := getBindAddr()

			m := HostMemberlist(addr.String(), t, func(c *Config) {
				c.ProbeTimeout = time.Millisecond
				c.ProbeInterval = 100 * time.Millisecond
				c.SuspicionMult = 5
				c.SuspicionMaxTimeoutMult = 2
			})
			defer m.Shutdown()

			bindPort := m.config.BindPort

			a := alive{Node: addr.String(), Addr: []byte(addr), Port: uint16(bindPort), Incarnation: 1, Vsn: m.config.BuildVsnArray()}
			m.aliveNode(&a, nil, true)

			// Make all but one peer be an real, alive instance.
			var peers []*Memberlist
			for j := 0; j < c.numPeers-1; j++ {
				peerAddr := getBindAddr()

				peer := HostMemberlist(peerAddr.String(), t, func(c *Config) {
					c.BindPort = bindPort
				})
				defer peer.Shutdown()

				peers = append(peers, peer)

				a = alive{Node: peerAddr.String(), Addr: []byte(peerAddr), Port: uint16(bindPort), Incarnation: 1, Vsn: m.config.BuildVsnArray()}
				m.aliveNode(&a, nil, false)
			}

			// Just use a bogus address for the last peer so it doesn't respond
			// to pings, but tell the memberlist it's alive.
			badPeerAddr := getBindAddr()
			a = alive{Node: badPeerAddr.String(), Addr: []byte(badPeerAddr), Port: uint16(bindPort), Incarnation: 1, Vsn: m.config.BuildVsnArray()}
			m.aliveNode(&a, nil, false)

			// Force a probe, which should start us into the suspect state.
			m.probeNodeByAddr(badPeerAddr.String())

			if m.getNodeState(badPeerAddr.String()) != stateSuspect {
				t.Fatalf("case %d: expected node to be suspect", i)
			}

			// Add the requested number of confirmations.
			for j := 0; j < c.confirmations; j++ {
				from := fmt.Sprintf("peer%d", j)
				s := suspect{Node: badPeerAddr.String(), Incarnation: 1, From: from}
				m.suspectNode(&s)
			}

			// Wait until right before the timeout and make sure the timer
			// hasn't fired.
			fudge := 25 * time.Millisecond
			time.Sleep(c.expected - fudge)

			if m.getNodeState(badPeerAddr.String()) != stateSuspect {
				t.Fatalf("case %d: expected node to still be suspect", i)
			}

			// Wait through the timeout and a little after to make sure the
			// timer fires.
			time.Sleep(2 * fudge)

			if m.getNodeState(badPeerAddr.String()) != stateDead {
				t.Fatalf("case %d: expected node to be dead", i)
			}
		})
	}
}

/*
func TestMemberList_ProbeNode_FallbackTCP(t *testing.T) {
	addr1 := getBindAddr()
	addr2 := getBindAddr()
	addr3 := getBindAddr()
	addr4 := getBindAddr()
	ip1 := []byte(addr1)
	ip2 := []byte(addr2)
	ip3 := []byte(addr3)
	ip4 := []byte(addr4)

	var probeTimeMax time.Duration
	m1 := HostMemberlist(addr1.String(), t, func(c *Config) {
		c.ProbeTimeout = 10 * time.Millisecond
		c.ProbeInterval = 200 * time.Millisecond
		probeTimeMax = c.ProbeInterval + 20*time.Millisecond
	})
	defer m1.Shutdown()

	bindPort := m1.config.BindPort

	m2 := HostMemberlist(addr2.String(), t, func(c *Config) {
		c.BindPort = bindPort
	})
	defer m2.Shutdown()

	m3 := HostMemberlist(addr3.String(), t, func(c *Config) {
		c.BindPort = bindPort
	})
	defer m3.Shutdown()

	m4 := HostMemberlist(addr4.String(), t, func(c *Config) {
		c.BindPort = bindPort
	})
	defer m4.Shutdown()

	a1 := alive{Node: addr1.String(), Addr: ip1, Port: uint16(bindPort), Incarnation: 1}
	m1.aliveNode(&a1, nil, true)
	a2 := alive{Node: addr2.String(), Addr: ip2, Port: uint16(bindPort), Incarnation: 1}
	m1.aliveNode(&a2, nil, false)
	a3 := alive{Node: addr3.String(), Addr: ip3, Port: uint16(bindPort), Incarnation: 1}
	m1.aliveNode(&a3, nil, false)

	// Make sure m4 is configured with the same protocol version as m1 so
	// the TCP fallback behavior is enabled.
	a4 := alive{
		Node:        addr4.String(),
		Addr:        ip4,
		Port:        uint16(bindPort),
		Incarnation: 1,
		Vsn: []uint8{
			ProtocolVersionMin,
			ProtocolVersionMax,
			m1.config.ProtocolVersion,
			m1.config.DelegateProtocolMin,
			m1.config.DelegateProtocolMax,
			m1.config.DelegateProtocolVersion,
		},
	}
	m1.aliveNode(&a4, nil, false)

	// Isolate m4 from UDP traffic by re-opening its listener on the wrong
	// port. This should force the TCP fallback path to be used.
	var err error
	if err = m4.udpListener.Close(); err != nil {
		t.Fatalf("err: %v", err)
	}
	udpAddr := &net.UDPAddr{IP: ip4, Port: 9999}
	if m4.udpListener, err = net.ListenUDP("udp", udpAddr); err != nil {
		t.Fatalf("err: %v", err)
	}

	// Have node m1 probe m4.
	n := m1.nodeMap[addr4.String()]
	startProbe := time.Now()
	m1.probeNode(n)
	probeTime := time.Now().Sub(startProbe)

	// Should be marked alive because of the TCP fallback ping.
	if n.State != stateAlive {
		t.Fatalf("expect node to be alive")
	}

	// Make sure TCP activity completed in a timely manner.
	if probeTime > probeTimeMax {
		t.Fatalf("took to long to probe, %9.6f", probeTime.Seconds())
	}

	// Confirm at least one of the peers attempted an indirect probe.
	time.Sleep(probeTimeMax)
	if m2.sequenceNum != 1 && m3.sequenceNum != 1 {
		t.Fatalf("bad seqnos %v, %v", m2.sequenceNum, m3.sequenceNum)
	}

	// Now shutdown all inbound TCP traffic to make sure the TCP fallback
	// path properly fails when the node is really unreachable.
	if err = m4.tcpListener.Close(); err != nil {
		t.Fatalf("err: %v", err)
	}
	tcpAddr := &net.TCPAddr{IP: ip4, Port: 9999}
	if m4.tcpListener, err = net.ListenTCP("tcp", tcpAddr); err != nil {
		t.Fatalf("err: %v", err)
	}

	// Probe again, this time there should be no contact.
	startProbe = time.Now()
	m1.probeNode(n)
	probeTime = time.Now().Sub(startProbe)

	// Node should be reported suspect.
	if n.State != stateSuspect {
		t.Fatalf("expect node to be suspect")
	}

	// Make sure TCP activity didn't cause us to wait too long before
	// timing out.
	if probeTime > probeTimeMax {
		t.Fatalf("took to long to probe, %9.6f", probeTime.Seconds())
	}

	// Confirm at least one of the peers attempted an indirect probe.
	time.Sleep(probeTimeMax)
	if m2.sequenceNum != 2 && m3.sequenceNum != 2 {
		t.Fatalf("bad seqnos %v, %v", m2.sequenceNum, m3.sequenceNum)
	}
}

func TestMemberList_ProbeNode_FallbackTCP_Disabled(t *testing.T) {
	addr1 := getBindAddr()
	addr2 := getBindAddr()
	addr3 := getBindAddr()
	addr4 := getBindAddr()
	ip1 := []byte(addr1)
	ip2 := []byte(addr2)
	ip3 := []byte(addr3)
	ip4 := []byte(addr4)

	var probeTimeMax time.Duration
	m1 := HostMemberlist(addr1.String(), t, func(c *Config) {
		c.ProbeTimeout = 10 * time.Millisecond
		c.ProbeInterval = 200 * time.Millisecond
		probeTimeMax = c.ProbeInterval + 20*time.Millisecond
	})
	defer m1.Shutdown()

	bindPort := m1.config.BindPort

	m2 := HostMemberlist(addr2.String(), t, func(c *Config) {
		c.BindPort = bindPort
	})
	defer m2.Shutdown()

	m3 := HostMemberlist(addr3.String(), t, func(c *Config) {
		c.BindPort = bindPort
	})
	defer m3.Shutdown()

	m4 := HostMemberlist(addr4.String(), t, func(c *Config) {
		c.BindPort = bindPort
	})
	defer m4.Shutdown()

	a1 := alive{Node: addr1.String(), Addr: ip1, Port: uint16(bindPort), Incarnation: 1}
	m1.aliveNode(&a1, nil, true)
	a2 := alive{Node: addr2.String(), Addr: ip2, Port: uint16(bindPort), Incarnation: 1}
	m1.aliveNode(&a2, nil, false)
	a3 := alive{Node: addr3.String(), Addr: ip3, Port: uint16(bindPort), Incarnation: 1}
	m1.aliveNode(&a3, nil, false)

	// Make sure m4 is configured with the same protocol version as m1 so
	// the TCP fallback behavior is enabled.
	a4 := alive{
		Node:        addr4.String(),
		Addr:        ip4,
		Port:        uint16(bindPort),
		Incarnation: 1,
		Vsn: []uint8{
			ProtocolVersionMin,
			ProtocolVersionMax,
			m1.config.ProtocolVersion,
			m1.config.DelegateProtocolMin,
			m1.config.DelegateProtocolMax,
			m1.config.DelegateProtocolVersion,
		},
	}
	m1.aliveNode(&a4, nil, false)

	// Isolate m4 from UDP traffic by re-opening its listener on the wrong
	// port. This should force the TCP fallback path to be used.
	var err error
	if err = m4.udpListener.Close(); err != nil {
		t.Fatalf("err: %v", err)
	}
	udpAddr := &net.UDPAddr{IP: ip4, Port: 9999}
	if m4.udpListener, err = net.ListenUDP("udp", udpAddr); err != nil {
		t.Fatalf("err: %v", err)
	}

	// Disable the TCP pings using the config mechanism.
	m1.config.DisableTcpPings = true

	// Have node m1 probe m4.
	n := m1.nodeMap[addr4.String()]
	startProbe := time.Now()
	m1.probeNode(n)
	probeTime := time.Now().Sub(startProbe)

	// Node should be reported suspect.
	if n.State != stateSuspect {
		t.Fatalf("expect node to be suspect")
	}

	// Make sure TCP activity didn't cause us to wait too long before
	// timing out.
	if probeTime > probeTimeMax {
		t.Fatalf("took to long to probe, %9.6f", probeTime.Seconds())
	}

	// Confirm at least one of the peers attempted an indirect probe.
	time.Sleep(probeTimeMax)
	if m2.sequenceNum != 1 && m3.sequenceNum != 1 {
		t.Fatalf("bad seqnos %v, %v", m2.sequenceNum, m3.sequenceNum)
	}
}

func TestMemberList_ProbeNode_FallbackTCP_OldProtocol(t *testing.T) {
	addr1 := getBindAddr()
	addr2 := getBindAddr()
	addr3 := getBindAddr()
	addr4 := getBindAddr()
	ip1 := []byte(addr1)
	ip2 := []byte(addr2)
	ip3 := []byte(addr3)
	ip4 := []byte(addr4)

	var probeTimeMax time.Duration
	m1 := HostMemberlist(addr1.String(), t, func(c *Config) {
		c.ProbeTimeout = 10 * time.Millisecond
		c.ProbeInterval = 200 * time.Millisecond
		probeTimeMax = c.ProbeInterval + 20*time.Millisecond
	})
	defer m1.Shutdown()

	bindPort := m1.config.BindPort

	m2 := HostMemberlist(addr2.String(), t, func(c *Config) {
		c.BindPort = bindPort
	})
	defer m2.Shutdown()

	m3 := HostMemberlist(addr3.String(), t, func(c *Config) {
		c.BindPort = bindPort
	})
	defer m3.Shutdown()

	m4 := HostMemberlist(addr4.String(), t, func(c *Config) {
		c.BindPort = bindPort
	})
	defer m4.Shutdown()

	a1 := alive{Node: addr1.String(), Addr: ip1, Port: uint16(bindPort), Incarnation: 1}
	m1.aliveNode(&a1, nil, true)
	a2 := alive{Node: addr2.String(), Addr: ip2, Port: uint16(bindPort), Incarnation: 1}
	m1.aliveNode(&a2, nil, false)
	a3 := alive{Node: addr3.String(), Addr: ip3, Port: uint16(bindPort), Incarnation: 1}
	m1.aliveNode(&a3, nil, false)

	// Set up m4 so that it doesn't understand a version of the protocol
	// that supports TCP pings.
	a4 := alive{
		Node:        addr4.String(),
		Addr:        ip4,
		Port:        uint16(bindPort),
		Incarnation: 1,
		Vsn: []uint8{
			ProtocolVersionMin,
			ProtocolVersion2Compatible,
			ProtocolVersion2Compatible,
			m1.config.DelegateProtocolMin,
			m1.config.DelegateProtocolMax,
			m1.config.DelegateProtocolVersion,
		},
	}
	m1.aliveNode(&a4, nil, false)

	// Isolate m4 from UDP traffic by re-opening its listener on the wrong
	// port. This should force the TCP fallback path to be used.
	var err error
	if err = m4.udpListener.Close(); err != nil {
		t.Fatalf("err: %v", err)
	}
	udpAddr := &net.UDPAddr{IP: ip4, Port: 9999}
	if m4.udpListener, err = net.ListenUDP("udp", udpAddr); err != nil {
		t.Fatalf("err: %v", err)
	}

	// Have node m1 probe m4.
	n := m1.nodeMap[addr4.String()]
	startProbe := time.Now()
	m1.probeNode(n)
	probeTime := time.Now().Sub(startProbe)

	// Node should be reported suspect.
	if n.State != stateSuspect {
		t.Fatalf("expect node to be suspect")
	}

	// Make sure TCP activity didn't cause us to wait too long before
	// timing out.
	if probeTime > probeTimeMax {
		t.Fatalf("took to long to probe, %9.6f", probeTime.Seconds())
	}

	// Confirm at least one of the peers attempted an indirect probe.
	time.Sleep(probeTimeMax)
	if m2.sequenceNum != 1 && m3.sequenceNum != 1 {
		t.Fatalf("bad seqnos %v, %v", m2.sequenceNum, m3.sequenceNum)
	}
}
*/

func TestMemberList_ProbeNode_Awareness_Degraded(t *testing.T) {
	addr1 := getBindAddr()
	addr2 := getBindAddr()
	addr3 := getBindAddr()
	addr4 := getBindAddr()
	ip1 := []byte(addr1)
	ip2 := []byte(addr2)
	ip3 := []byte(addr3)
	ip4 := []byte(addr4)

	var probeTimeMin time.Duration
	m1 := HostMemberlist(addr1.String(), t, func(c *Config) {
		c.ProbeTimeout = 10 * time.Millisecond
		c.ProbeInterval = 200 * time.Millisecond
		probeTimeMin = 2*c.ProbeInterval - 50*time.Millisecond
	})
	defer m1.Shutdown()

	bindPort := m1.config.BindPort

	m2 := HostMemberlist(addr2.String(), t, func(c *Config) {
		c.BindPort = bindPort
		c.ProbeTimeout = 10 * time.Millisecond
		c.ProbeInterval = 200 * time.Millisecond
	})
	defer m2.Shutdown()

	m3 := HostMemberlist(addr3.String(), t, func(c *Config) {
		c.BindPort = bindPort
		c.ProbeTimeout = 10 * time.Millisecond
		c.ProbeInterval = 200 * time.Millisecond
	})
	defer m3.Shutdown()

	a1 := alive{Node: addr1.String(), Addr: ip1, Port: uint16(bindPort), Incarnation: 1, Vsn: m1.config.BuildVsnArray()}
	m1.aliveNode(&a1, nil, true)
	a2 := alive{Node: addr2.String(), Addr: ip2, Port: uint16(bindPort), Incarnation: 1, Vsn: m2.config.BuildVsnArray()}
	m1.aliveNode(&a2, nil, false)
	a3 := alive{Node: addr3.String(), Addr: ip3, Port: uint16(bindPort), Incarnation: 1, Vsn: m3.config.BuildVsnArray()}
	m1.aliveNode(&a3, nil, false)

	vsn4 := []uint8{
		ProtocolVersionMin, ProtocolVersionMax, ProtocolVersionMin,
		1, 1, 1,
	}
	// Node 4 never gets started.
	a4 := alive{Node: addr4.String(), Addr: ip4, Port: uint16(bindPort), Incarnation: 1, Vsn: vsn4}
	m1.aliveNode(&a4, nil, false)

	// Start the health in a degraded state.
	m1.awareness.ApplyDelta(1)
	if score := m1.GetHealthScore(); score != 1 {
		t.Fatalf("bad: %d", score)
	}

	// Have node m1 probe m4.
	n := m1.nodeMap[addr4.String()]
	startProbe := time.Now()
	m1.probeNode(n)
	probeTime := time.Now().Sub(startProbe)

	// Node should be reported suspect.
	if n.State != stateSuspect {
		t.Fatalf("expect node to be suspect")
	}

	// Make sure we timed out approximately on time (note that we accounted
	// for the slowed-down failure detector in the probeTimeMin calculation.
	if probeTime < probeTimeMin {
		t.Fatalf("probed too quickly, %9.6f", probeTime.Seconds())
	}

	// Confirm at least one of the peers attempted an indirect probe.
	if m2.sequenceNum != 1 && m3.sequenceNum != 1 {
		t.Fatalf("bad seqnos %v, %v", m2.sequenceNum, m3.sequenceNum)
	}

	// We should have gotten all the nacks, so our score should remain the
	// same, since we didn't get a successful probe.
	if score := m1.GetHealthScore(); score != 1 {
		t.Fatalf("bad: %d", score)
	}
}

func TestMemberList_ProbeNode_Wrong_VSN(t *testing.T) {
	addr1 := getBindAddr()
	addr2 := getBindAddr()
	addr3 := getBindAddr()
	addr4 := getBindAddr()
	ip1 := []byte(addr1)
	ip2 := []byte(addr2)
	ip3 := []byte(addr3)
	ip4 := []byte(addr4)

	m1 := HostMemberlist(addr1.String(), t, func(c *Config) {
		c.ProbeTimeout = 10 * time.Millisecond
		c.ProbeInterval = 200 * time.Millisecond
	})
	defer m1.Shutdown()

	bindPort := m1.config.BindPort

	m2 := HostMemberlist(addr2.String(), t, func(c *Config) {
		c.BindPort = bindPort
		c.ProbeTimeout = 10 * time.Millisecond
		c.ProbeInterval = 200 * time.Millisecond
	})
	defer m2.Shutdown()

	m3 := HostMemberlist(addr3.String(), t, func(c *Config) {
		c.BindPort = bindPort
		c.ProbeTimeout = 10 * time.Millisecond
		c.ProbeInterval = 200 * time.Millisecond
	})
	defer m3.Shutdown()

	a1 := alive{Node: addr1.String(), Addr: ip1, Port: uint16(bindPort), Incarnation: 1, Vsn: m1.config.BuildVsnArray()}
	m1.aliveNode(&a1, nil, true)
	a2 := alive{Node: addr2.String(), Addr: ip2, Port: uint16(bindPort), Incarnation: 1, Vsn: m2.config.BuildVsnArray()}
	m1.aliveNode(&a2, nil, false)
	a3 := alive{Node: addr3.String(), Addr: ip3, Port: uint16(bindPort), Incarnation: 1, Vsn: m3.config.BuildVsnArray()}
	m1.aliveNode(&a3, nil, false)

	vsn4 := []uint8{
		0, 0, 0,
		0, 0, 0,
	}
	// Node 4 never gets started.
	a4 := alive{Node: addr4.String(), Addr: ip4, Port: uint16(bindPort), Incarnation: 1, Vsn: vsn4}
	m1.aliveNode(&a4, nil, false)

	// Start the health in a degraded state.
	m1.awareness.ApplyDelta(1)
	if score := m1.GetHealthScore(); score != 1 {
		t.Fatalf("bad: %d", score)
	}

	// Have node m1 probe m4.
	n, ok := m1.nodeMap[addr4.String()]
	if ok || n != nil {
		t.Fatalf("expect node a4 to be not taken into account, because of its wrong version")
	}
}

func TestMemberList_ProbeNode_Awareness_Improved(t *testing.T) {
	addr1 := getBindAddr()
	addr2 := getBindAddr()
	ip1 := []byte(addr1)
	ip2 := []byte(addr2)

	m1 := HostMemberlist(addr1.String(), t, func(c *Config) {
		c.ProbeTimeout = 10 * time.Millisecond
		c.ProbeInterval = 200 * time.Millisecond
	})
	defer m1.Shutdown()

	bindPort := m1.config.BindPort

	m2 := HostMemberlist(addr2.String(), t, func(c *Config) {
		c.BindPort = bindPort
	})
	defer m2.Shutdown()

	a1 := alive{Node: addr1.String(), Addr: ip1, Port: uint16(bindPort), Incarnation: 1, Vsn: m1.config.BuildVsnArray()}
	m1.aliveNode(&a1, nil, true)
	a2 := alive{Node: addr2.String(), Addr: ip2, Port: uint16(bindPort), Incarnation: 1, Vsn: m2.config.BuildVsnArray()}
	m1.aliveNode(&a2, nil, false)

	// Start the health in a degraded state.
	m1.awareness.ApplyDelta(1)
	if score := m1.GetHealthScore(); score != 1 {
		t.Fatalf("bad: %d", score)
	}

	// Have node m1 probe m2.
	n := m1.nodeMap[addr2.String()]
	m1.probeNode(n)

	// Node should be reported alive.
	if n.State != stateAlive {
		t.Fatalf("expect node to be suspect")
	}

	// Our score should have improved since we did a good probe.
	if score := m1.GetHealthScore(); score != 0 {
		t.Fatalf("bad: %d", score)
	}
}

func TestMemberList_ProbeNode_Awareness_MissedNack(t *testing.T) {
	addr1 := getBindAddr()
	addr2 := getBindAddr()
	addr3 := getBindAddr()
	addr4 := getBindAddr()
	ip1 := []byte(addr1)
	ip2 := []byte(addr2)
	ip3 := []byte(addr3)
	ip4 := []byte(addr4)

	var probeTimeMax time.Duration
	m1 := HostMemberlist(addr1.String(), t, func(c *Config) {
		c.ProbeTimeout = 10 * time.Millisecond
		c.ProbeInterval = 200 * time.Millisecond
		probeTimeMax = c.ProbeInterval + 50*time.Millisecond
	})
	defer m1.Shutdown()

	bindPort := m1.config.BindPort

	m2 := HostMemberlist(addr2.String(), t, func(c *Config) {
		c.BindPort = bindPort
		c.ProbeTimeout = 10 * time.Millisecond
		c.ProbeInterval = 200 * time.Millisecond
	})
	defer m2.Shutdown()

	a1 := alive{Node: addr1.String(), Addr: ip1, Port: uint16(bindPort), Incarnation: 1, Vsn: m1.config.BuildVsnArray()}
	m1.aliveNode(&a1, nil, true)
	a2 := alive{Node: addr2.String(), Addr: ip2, Port: uint16(bindPort), Incarnation: 1, Vsn: m1.config.BuildVsnArray()}
	m1.aliveNode(&a2, nil, false)

	vsn := m1.config.BuildVsnArray()
	// Node 3 and node 4 never get started.
	a3 := alive{Node: addr3.String(), Addr: ip3, Port: uint16(bindPort), Incarnation: 1, Vsn: vsn}
	m1.aliveNode(&a3, nil, false)
	a4 := alive{Node: addr4.String(), Addr: ip4, Port: uint16(bindPort), Incarnation: 1, Vsn: vsn}
	m1.aliveNode(&a4, nil, false)

	// Make sure health looks good.
	if score := m1.GetHealthScore(); score != 0 {
		t.Fatalf("bad: %d", score)
	}

	// Have node m1 probe m4.
	n := m1.nodeMap[addr4.String()]
	startProbe := time.Now()
	m1.probeNode(n)
	probeTime := time.Now().Sub(startProbe)

	// Node should be reported suspect.
	if n.State != stateSuspect {
		t.Fatalf("expect node to be suspect")
	}

	// Make sure we timed out approximately on time.
	if probeTime > probeTimeMax {
		t.Fatalf("took to long to probe, %9.6f", probeTime.Seconds())
	}

	// We should have gotten dinged for the missed nack. Note that the code under
	// test is waiting for probeTimeMax and then doing some other work before it
	// updates the awareness, so we need to wait some extra time. Rather than just
	// add longer and longer sleeps, we'll retry a few times.
	time.Sleep(probeTimeMax)
	retry(t, 5, 10*time.Millisecond, func(failf func(string, ...interface{})) {
		if score := m1.GetHealthScore(); score != 1 {
			failf("expected health score to decrement on missed nack. want %d, "+
				"got: %d", 1, score)
		}
	})
}

func TestMemberList_ProbeNode_Awareness_OldProtocol(t *testing.T) {
	addr1 := getBindAddr()
	addr2 := getBindAddr()
	addr3 := getBindAddr()
	addr4 := getBindAddr()
	ip1 := []byte(addr1)
	ip2 := []byte(addr2)
	ip3 := []byte(addr3)
	ip4 := []byte(addr4)

	var probeTimeMax time.Duration
	m1 := HostMemberlist(addr1.String(), t, func(c *Config) {
		c.ProbeTimeout = 10 * time.Millisecond
		c.ProbeInterval = 200 * time.Millisecond
		probeTimeMax = c.ProbeInterval + 20*time.Millisecond
	})
	defer m1.Shutdown()

	bindPort := m1.config.BindPort

	m2 := HostMemberlist(addr2.String(), t, func(c *Config) {
		c.BindPort = bindPort
	})
	defer m2.Shutdown()

	m3 := HostMemberlist(addr3.String(), t, func(c *Config) {
		c.BindPort = bindPort
	})
	defer m3.Shutdown()

	a1 := alive{Node: addr1.String(), Addr: ip1, Port: uint16(bindPort), Incarnation: 1}
	m1.aliveNode(&a1, nil, true)
	a2 := alive{Node: addr2.String(), Addr: ip2, Port: uint16(bindPort), Incarnation: 1}
	m1.aliveNode(&a2, nil, false)
	a3 := alive{Node: addr3.String(), Addr: ip3, Port: uint16(bindPort), Incarnation: 1}
	m1.aliveNode(&a3, nil, false)

	// Node 4 never gets started.
	a4 := alive{Node: addr4.String(), Addr: ip4, Port: uint16(bindPort), Incarnation: 1}
	m1.aliveNode(&a4, nil, false)

	// Make sure health looks good.
	if score := m1.GetHealthScore(); score != 0 {
		t.Fatalf("bad: %d", score)
	}

	// Have node m1 probe m4.
	n := m1.nodeMap[addr4.String()]
	startProbe := time.Now()
	m1.probeNode(n)
	probeTime := time.Now().Sub(startProbe)

	// Node should be reported suspect.
	if n.State != stateSuspect {
		t.Fatalf("expect node to be suspect")
	}

	// Make sure we timed out approximately on time.
	if probeTime > probeTimeMax {
		t.Fatalf("took to long to probe, %9.6f", probeTime.Seconds())
	}

	// Confirm at least one of the peers attempted an indirect probe.
	time.Sleep(probeTimeMax)
	if m2.sequenceNum != 1 && m3.sequenceNum != 1 {
		t.Fatalf("bad seqnos %v, %v", m2.sequenceNum, m3.sequenceNum)
	}

	// Since we are using the old protocol here, we should have gotten dinged
	// for a failed health check.
	if score := m1.GetHealthScore(); score != 1 {
		t.Fatalf("bad: %d", score)
	}
}

func TestMemberList_ProbeNode_Buddy(t *testing.T) {
	addr1 := getBindAddr()
	addr2 := getBindAddr()
	ip1 := []byte(addr1)
	ip2 := []byte(addr2)

	m1 := HostMemberlist(addr1.String(), t, func(c *Config) {
		c.ProbeTimeout = time.Millisecond
		c.ProbeInterval = 10 * time.Millisecond
	})
	defer m1.Shutdown()

	bindPort := m1.config.BindPort

	m2 := HostMemberlist(addr2.String(), t, func(c *Config) {
		c.BindPort = bindPort
	})
	defer m2.Shutdown()

	a1 := alive{Node: addr1.String(), Addr: ip1, Port: uint16(bindPort), Incarnation: 1, Vsn: m1.config.BuildVsnArray()}
	a2 := alive{Node: addr2.String(), Addr: ip2, Port: uint16(bindPort), Incarnation: 1, Vsn: m2.config.BuildVsnArray()}

	m1.aliveNode(&a1, nil, true)
	m1.aliveNode(&a2, nil, false)
	m2.aliveNode(&a2, nil, true)

	// Force the state to suspect so we piggyback a suspect message with the ping.
	// We should see this get refuted later, and the ping will succeed.
	n := m1.nodeMap[addr2.String()]
	n.State = stateSuspect
	m1.probeNode(n)

	// Make sure a ping was sent.
	if m1.sequenceNum != 1 {
		t.Fatalf("bad seqno %v", m1.sequenceNum)
	}

	// Check a broadcast is queued.
	if num := m2.broadcasts.NumQueued(); num != 1 {
		t.Fatalf("expected only one queued message: %d", num)
	}

	// Should be alive msg.
	if messageType(m2.broadcasts.orderedView(true)[0].b.Message()[0]) != aliveMsg {
		t.Fatalf("expected queued alive msg")
	}
}

func TestMemberList_ProbeNode(t *testing.T) {
	addr1 := getBindAddr()
	addr2 := getBindAddr()
	ip1 := []byte(addr1)
	ip2 := []byte(addr2)

	m1 := HostMemberlist(addr1.String(), t, func(c *Config) {
		c.ProbeTimeout = time.Millisecond
		c.ProbeInterval = 10 * time.Millisecond
	})
	defer m1.Shutdown()

	bindPort := m1.config.BindPort

	m2 := HostMemberlist(addr2.String(), t, func(c *Config) {
		c.BindPort = bindPort
	})
	defer m2.Shutdown()

	a1 := alive{Node: addr1.String(), Addr: ip1, Port: uint16(bindPort), Incarnation: 1}
	m1.aliveNode(&a1, nil, true)
	a2 := alive{Node: addr2.String(), Addr: ip2, Port: uint16(bindPort), Incarnation: 1}
	m1.aliveNode(&a2, nil, false)

	n := m1.nodeMap[addr2.String()]
	m1.probeNode(n)

	// Should be marked alive
	if n.State != stateAlive {
		t.Fatalf("Expect node to be alive")
	}

	// Should increment seqno
	if m1.sequenceNum != 1 {
		t.Fatalf("bad seqno %v", m1.sequenceNum)
	}
}

func TestMemberList_Ping(t *testing.T) {
	addr1 := getBindAddr()
	addr2 := getBindAddr()
	ip1 := []byte(addr1)
	ip2 := []byte(addr2)

	m1 := HostMemberlist(addr1.String(), t, func(c *Config) {
		c.ProbeTimeout = 1 * time.Second
		c.ProbeInterval = 10 * time.Second
	})
	defer m1.Shutdown()

	bindPort := m1.config.BindPort

	m2 := HostMemberlist(addr2.String(), t, func(c *Config) {
		c.BindPort = bindPort
	})
	defer m2.Shutdown()

	a1 := alive{Node: addr1.String(), Addr: ip1, Port: uint16(bindPort), Incarnation: 1}
	m1.aliveNode(&a1, nil, true)
	a2 := alive{Node: addr2.String(), Addr: ip2, Port: uint16(bindPort), Incarnation: 1}
	m1.aliveNode(&a2, nil, false)

	// Do a legit ping.
	n := m1.nodeMap[addr2.String()]
	addr, err := net.ResolveUDPAddr("udp", net.JoinHostPort(addr2.String(), strconv.Itoa(bindPort)))
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	rtt, err := m1.Ping(n.Name, addr)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if !(rtt > 0) {
		t.Fatalf("bad: %v", rtt)
	}

	// This ping has a bad node name so should timeout.
	_, err = m1.Ping("bad", addr)
	if _, ok := err.(NoPingResponseError); !ok || err == nil {
		t.Fatalf("bad: %v", err)
	}
}

func TestMemberList_ResetNodes(t *testing.T) {
	m := GetMemberlist(t, func(c *Config) {
		c.GossipToTheDeadTime = 100 * time.Millisecond
	})
	defer m.Shutdown()

	a1 := alive{Node: "test1", Addr: []byte{127, 0, 0, 1}, Incarnation: 1, Vsn: m.config.BuildVsnArray()}
	m.aliveNode(&a1, nil, false)
	a2 := alive{Node: "test2", Addr: []byte{127, 0, 0, 2}, Incarnation: 1, Vsn: m.config.BuildVsnArray()}
	m.aliveNode(&a2, nil, false)
	a3 := alive{Node: "test3", Addr: []byte{127, 0, 0, 3}, Incarnation: 1, Vsn: m.config.BuildVsnArray()}
	m.aliveNode(&a3, nil, false)
	d := dead{Node: "test2", Incarnation: 1}
	m.deadNode(&d)

	m.resetNodes()
	if len(m.nodes) != 3 {
		t.Fatalf("Bad length")
	}
	if _, ok := m.nodeMap["test2"]; !ok {
		t.Fatalf("test2 should not be unmapped")
	}

	time.Sleep(200 * time.Millisecond)
	m.resetNodes()
	if len(m.nodes) != 2 {
		t.Fatalf("Bad length")
	}
	if _, ok := m.nodeMap["test2"]; ok {
		t.Fatalf("test2 should be unmapped")
	}
}

func TestMemberList_NextSeq(t *testing.T) {
	m := &Memberlist{}
	if m.nextSeqNo() != 1 {
		t.Fatalf("bad sequence no")
	}
	if m.nextSeqNo() != 2 {
		t.Fatalf("bad sequence no")
	}
}

func ackHandlerExists(t *testing.T, m *Memberlist, idx uint32) bool {
	t.Helper()

	m.ackLock.Lock()
	_, ok := m.ackHandlers[idx]
	m.ackLock.Unlock()

	return ok
}

func TestMemberList_setProbeChannels(t *testing.T) {
	m := &Memberlist{ackHandlers: make(map[uint32]*ackHandler)}

	ch := make(chan ackMessage, 1)
	m.setProbeChannels(0, ch, nil, 10*time.Millisecond)

	require.True(t, ackHandlerExists(t, m, 0), "missing handler")

	time.Sleep(20 * time.Millisecond)

	require.False(t, ackHandlerExists(t, m, 0), "non-reaped handler")
}

func TestMemberList_setAckHandler(t *testing.T) {
	m := &Memberlist{ackHandlers: make(map[uint32]*ackHandler)}

	f := func([]byte, time.Time) {}
	m.setAckHandler(0, f, 10*time.Millisecond)

	require.True(t, ackHandlerExists(t, m, 0), "missing handler")

	time.Sleep(20 * time.Millisecond)

	require.False(t, ackHandlerExists(t, m, 0), "non-reaped handler")
}

func TestMemberList_invokeAckHandler(t *testing.T) {
	m := &Memberlist{ackHandlers: make(map[uint32]*ackHandler)}

	// Does nothing
	m.invokeAckHandler(ackResp{}, time.Now())

	var b bool
	f := func(payload []byte, timestamp time.Time) { b = true }
	m.setAckHandler(0, f, 10*time.Millisecond)

	// Should set b
	m.invokeAckHandler(ackResp{0, nil}, time.Now())
	if !b {
		t.Fatalf("b not set")
	}

	require.False(t, ackHandlerExists(t, m, 0), "non-reaped handler")
}

func TestMemberList_invokeAckHandler_Channel_Ack(t *testing.T) {
	m := &Memberlist{ackHandlers: make(map[uint32]*ackHandler)}

	ack := ackResp{0, []byte{0, 0, 0}}

	// Does nothing
	m.invokeAckHandler(ack, time.Now())

	ackCh := make(chan ackMessage, 1)
	nackCh := make(chan struct{}, 1)
	m.setProbeChannels(0, ackCh, nackCh, 10*time.Millisecond)

	// Should send message
	m.invokeAckHandler(ack, time.Now())

	select {
	case v := <-ackCh:
		if v.Complete != true {
			t.Fatalf("Bad value")
		}
		if bytes.Compare(v.Payload, ack.Payload) != 0 {
			t.Fatalf("wrong payload. expected: %v; actual: %v", ack.Payload, v.Payload)
		}

	case <-nackCh:
		t.Fatalf("should not get a nack")

	default:
		t.Fatalf("message not sent")
	}

	require.False(t, ackHandlerExists(t, m, 0), "non-reaped handler")
}

func TestMemberList_invokeAckHandler_Channel_Nack(t *testing.T) {
	m := &Memberlist{ackHandlers: make(map[uint32]*ackHandler)}

	nack := nackResp{0}

	// Does nothing.
	m.invokeNackHandler(nack)

	ackCh := make(chan ackMessage, 1)
	nackCh := make(chan struct{}, 1)
	m.setProbeChannels(0, ackCh, nackCh, 10*time.Millisecond)

	// Should send message.
	m.invokeNackHandler(nack)

	select {
	case <-ackCh:
		t.Fatalf("should not get an ack")

	case <-nackCh:
		// Good.

	default:
		t.Fatalf("message not sent")
	}

	// Getting a nack doesn't reap the handler so that we can still forward
	// an ack up to the reap time, if we get one.
	require.True(t, ackHandlerExists(t, m, 0), "handler should not be reaped")

	ack := ackResp{0, []byte{0, 0, 0}}
	m.invokeAckHandler(ack, time.Now())

	select {
	case v := <-ackCh:
		if v.Complete != true {
			t.Fatalf("Bad value")
		}
		if bytes.Compare(v.Payload, ack.Payload) != 0 {
			t.Fatalf("wrong payload. expected: %v; actual: %v", ack.Payload, v.Payload)
		}

	case <-nackCh:
		t.Fatalf("should not get a nack")

	default:
		t.Fatalf("message not sent")
	}

	require.False(t, ackHandlerExists(t, m, 0), "non-reaped handler")
}

func TestMemberList_AliveNode_NewNode(t *testing.T) {
	ch := make(chan NodeEvent, 1)
	m := GetMemberlist(t, func(c *Config) {
		c.Events = &ChannelEventDelegate{ch}
	})
	defer m.Shutdown()

	a := alive{Node: "test", Addr: []byte{127, 0, 0, 1}, Incarnation: 1, Vsn: m.config.BuildVsnArray()}
	m.aliveNode(&a, nil, false)

	if len(m.nodes) != 1 {
		t.Fatalf("should add node")
	}

	state, ok := m.nodeMap["test"]
	if !ok {
		t.Fatalf("should map node")
	}

	if state.Incarnation != 1 {
		t.Fatalf("bad incarnation")
	}
	if state.State != stateAlive {
		t.Fatalf("bad state")
	}
	if time.Now().Sub(state.StateChange) > time.Second {
		t.Fatalf("bad change delta")
	}

	// Check for a join message
	select {
	case e := <-ch:
		if e.Node.Name != "test" {
			t.Fatalf("bad node name")
		}
	default:
		t.Fatalf("no join message")
	}

	// Check a broad cast is queued
	if m.broadcasts.NumQueued() != 1 {
		t.Fatalf("expected queued message")
	}
}

func TestMemberList_AliveNode_SuspectNode(t *testing.T) {
	ch := make(chan NodeEvent, 1)
	ted := &toggledEventDelegate{
		real: &ChannelEventDelegate{ch},
	}
	m := GetMemberlist(t, func(c *Config) {
		c.Events = ted
	})
	defer m.Shutdown()

	a := alive{Node: "test", Addr: []byte{127, 0, 0, 1}, Incarnation: 1, Vsn: m.config.BuildVsnArray()}
	m.aliveNode(&a, nil, false)

	// Listen only after first join
	ted.Toggle(true)

	// Make suspect
	state := m.nodeMap["test"]
	state.State = stateSuspect
	state.StateChange = state.StateChange.Add(-time.Hour)

	// Old incarnation number, should not change
	m.aliveNode(&a, nil, false)
	if state.State != stateSuspect {
		t.Fatalf("update with old incarnation!")
	}

	// Should reset to alive now
	a.Incarnation = 2
	m.aliveNode(&a, nil, false)
	if state.State != stateAlive {
		t.Fatalf("no update with new incarnation!")
	}

	if time.Now().Sub(state.StateChange) > time.Second {
		t.Fatalf("bad change delta")
	}

	// Check for a no join message
	select {
	case <-ch:
		t.Fatalf("got bad join message")
	default:
	}

	// Check a broad cast is queued
	if m.broadcasts.NumQueued() != 1 {
		t.Fatalf("expected queued message")
	}
}

func TestMemberList_AliveNode_Idempotent(t *testing.T) {
	ch := make(chan NodeEvent, 1)
	ted := &toggledEventDelegate{
		real: &ChannelEventDelegate{ch},
	}
	m := GetMemberlist(t, func(c *Config) {
		c.Events = ted
	})
	defer m.Shutdown()

	a := alive{Node: "test", Addr: []byte{127, 0, 0, 1}, Incarnation: 1, Vsn: m.config.BuildVsnArray()}
	m.aliveNode(&a, nil, false)

	// Listen only after first join
	ted.Toggle(true)

	// Make suspect
	state := m.nodeMap["test"]
	stateTime := state.StateChange

	// Should reset to alive now
	a.Incarnation = 2
	m.aliveNode(&a, nil, false)
	if state.State != stateAlive {
		t.Fatalf("non idempotent")
	}

	if stateTime != state.StateChange {
		t.Fatalf("should not change state")
	}

	// Check for a no join message
	select {
	case <-ch:
		t.Fatalf("got bad join message")
	default:
	}

	// Check a broad cast is queued
	if m.broadcasts.NumQueued() != 1 {
		t.Fatalf("expected only one queued message")
	}
}

type toggledEventDelegate struct {
	mu      sync.Mutex
	real    EventDelegate
	enabled bool
}

func (d *toggledEventDelegate) Toggle(enabled bool) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.enabled = enabled
}

// NotifyJoin is invoked when a node is detected to have joined.
// The Node argument must not be modified.
func (d *toggledEventDelegate) NotifyJoin(n *Node) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.enabled {
		d.real.NotifyJoin(n)
	}
}

// NotifyLeave is invoked when a node is detected to have left.
// The Node argument must not be modified.
func (d *toggledEventDelegate) NotifyLeave(n *Node) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.enabled {
		d.real.NotifyLeave(n)
	}
}

// NotifyUpdate is invoked when a node is detected to have
// updated, usually involving the meta data. The Node argument
// must not be modified.
func (d *toggledEventDelegate) NotifyUpdate(n *Node) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.enabled {
		d.real.NotifyUpdate(n)
	}
}

// Serf Bug: GH-58, Meta data does not update
func TestMemberList_AliveNode_ChangeMeta(t *testing.T) {
	ch := make(chan NodeEvent, 1)
	ted := &toggledEventDelegate{
		real: &ChannelEventDelegate{ch},
	}

	m := GetMemberlist(t, func(c *Config) {
		c.Events = ted
	})
	defer m.Shutdown()

	a := alive{
		Node:        "test",
		Addr:        []byte{127, 0, 0, 1},
		Meta:        []byte("val1"),
		Incarnation: 1,
		Vsn:         m.config.BuildVsnArray()}
	m.aliveNode(&a, nil, false)

	// Listen only after first join
	ted.Toggle(true)

	// Make suspect
	state := m.nodeMap["test"]

	// Should reset to alive now
	a.Incarnation = 2
	a.Meta = []byte("val2")
	m.aliveNode(&a, nil, false)

	// Check updates
	if bytes.Compare(state.Meta, a.Meta) != 0 {
		t.Fatalf("meta did not update")
	}

	// Check for a NotifyUpdate
	select {
	case e := <-ch:
		if e.Event != NodeUpdate {
			t.Fatalf("bad event: %v", e)
		}
		if e.Node != &state.Node {
			t.Fatalf("bad event: %v", e)
		}
		if bytes.Compare(e.Node.Meta, a.Meta) != 0 {
			t.Fatalf("meta did not update")
		}
	default:
		t.Fatalf("missing event!")
	}

}

func TestMemberList_AliveNode_Refute(t *testing.T) {
	m := GetMemberlist(t, nil)
	defer m.Shutdown()

	a := alive{Node: m.config.Name, Addr: []byte{127, 0, 0, 1}, Incarnation: 1, Vsn: m.config.BuildVsnArray()}
	m.aliveNode(&a, nil, true)

	// Clear queue
	m.broadcasts.Reset()

	// Conflicting alive
	s := alive{
		Node:        m.config.Name,
		Addr:        []byte{127, 0, 0, 1},
		Incarnation: 2,
		Meta:        []byte("foo"),
		Vsn:         m.config.BuildVsnArray(),
	}
	m.aliveNode(&s, nil, false)

	state := m.nodeMap[m.config.Name]
	if state.State != stateAlive {
		t.Fatalf("should still be alive")
	}
	if state.Meta != nil {
		t.Fatalf("meta should still be nil")
	}

	// Check a broad cast is queued
	if num := m.broadcasts.NumQueued(); num != 1 {
		t.Fatalf("expected only one queued message: %d",
			num)
	}

	// Should be alive mesg
	if messageType(m.broadcasts.orderedView(true)[0].b.Message()[0]) != aliveMsg {
		t.Fatalf("expected queued alive msg")
	}
}

func TestMemberList_AliveNode_Conflict(t *testing.T) {
	m := GetMemberlist(t, func(c *Config) {
		c.DeadNodeReclaimTime = 10 * time.Millisecond
	})
	defer m.Shutdown()

	nodeName := "test"
	a := alive{Node: nodeName, Addr: []byte{127, 0, 0, 1}, Port: 8000, Incarnation: 1, Vsn: m.config.BuildVsnArray()}
	m.aliveNode(&a, nil, true)

	// Clear queue
	m.broadcasts.Reset()

	// Conflicting alive
	s := alive{
		Node:        nodeName,
		Addr:        []byte{127, 0, 0, 2},
		Port:        9000,
		Incarnation: 2,
		Meta:        []byte("foo"),
		Vsn:         m.config.BuildVsnArray(),
	}
	m.aliveNode(&s, nil, false)

	state := m.nodeMap[nodeName]
	if state.State != stateAlive {
		t.Fatalf("should still be alive")
	}
	if state.Meta != nil {
		t.Fatalf("meta should still be nil")
	}
	if bytes.Equal(state.Addr, []byte{127, 0, 0, 2}) {
		t.Fatalf("address should not be updated")
	}
	if state.Port == 9000 {
		t.Fatalf("port should not be updated")
	}

	// Check a broad cast is queued
	if num := m.broadcasts.NumQueued(); num != 0 {
		t.Fatalf("expected 0 queued messages: %d", num)
	}

	// Change the node to dead
	d := dead{Node: nodeName, Incarnation: 2}
	m.deadNode(&d)
	m.broadcasts.Reset()

	state = m.nodeMap[nodeName]
	if state.State != stateDead {
		t.Fatalf("should be dead")
	}

	time.Sleep(m.config.DeadNodeReclaimTime)

	// New alive node
	s2 := alive{
		Node:        nodeName,
		Addr:        []byte{127, 0, 0, 2},
		Port:        9000,
		Incarnation: 3,
		Meta:        []byte("foo"),
		Vsn:         m.config.BuildVsnArray(),
	}
	m.aliveNode(&s2, nil, false)

	state = m.nodeMap[nodeName]
	if state.State != stateAlive {
		t.Fatalf("should still be alive")
	}
	if !bytes.Equal(state.Meta, []byte("foo")) {
		t.Fatalf("meta should be updated")
	}
	if !bytes.Equal(state.Addr, []byte{127, 0, 0, 2}) {
		t.Fatalf("address should be updated")
	}
	if state.Port != 9000 {
		t.Fatalf("port should be updated")
	}
}

func TestMemberList_SuspectNode_NoNode(t *testing.T) {
	m := GetMemberlist(t, nil)
	defer m.Shutdown()

	s := suspect{Node: "test", Incarnation: 1}
	m.suspectNode(&s)
	if len(m.nodes) != 0 {
		t.Fatalf("don't expect nodes")
	}
}

func TestMemberList_SuspectNode(t *testing.T) {
	m := GetMemberlist(t, func(c *Config) {
		c.ProbeInterval = time.Millisecond
		c.SuspicionMult = 1
	})
	defer m.Shutdown()

	a := alive{Node: "test", Addr: []byte{127, 0, 0, 1}, Incarnation: 1, Vsn: m.config.BuildVsnArray()}
	m.aliveNode(&a, nil, false)

	m.changeNode("test", func(state *nodeState) {
		state.StateChange = state.StateChange.Add(-time.Hour)
	})

	s := suspect{Node: "test", Incarnation: 1}
	m.suspectNode(&s)

	if m.getNodeState("test") != stateSuspect {
		t.Fatalf("Bad state")
	}

	change := m.getNodeStateChange("test")
	if time.Now().Sub(change) > time.Second {
		t.Fatalf("bad change delta")
	}

	// Check a broad cast is queued
	if m.broadcasts.NumQueued() != 1 {
		t.Fatalf("expected only one queued message")
	}

	// Check its a suspect message
	if messageType(m.broadcasts.orderedView(true)[0].b.Message()[0]) != suspectMsg {
		t.Fatalf("expected queued suspect msg")
	}

	// Wait for the timeout
	time.Sleep(10 * time.Millisecond)

	if m.getNodeState("test") != stateDead {
		t.Fatalf("Bad state")
	}

	newChange := m.getNodeStateChange("test")
	if time.Now().Sub(newChange) > time.Second {
		t.Fatalf("bad change delta")
	}
	if !newChange.After(change) {
		t.Fatalf("should increment time")
	}

	// Check a broad cast is queued
	if m.broadcasts.NumQueued() != 1 {
		t.Fatalf("expected only one queued message")
	}

	// Check its a suspect message
	if messageType(m.broadcasts.orderedView(true)[0].b.Message()[0]) != deadMsg {
		t.Fatalf("expected queued dead msg")
	}
}

func TestMemberList_SuspectNode_DoubleSuspect(t *testing.T) {
	m := GetMemberlist(t, nil)
	defer m.Shutdown()

	a := alive{Node: "test", Addr: []byte{127, 0, 0, 1}, Incarnation: 1, Vsn: m.config.BuildVsnArray()}
	m.aliveNode(&a, nil, false)

	state := m.nodeMap["test"]
	state.StateChange = state.StateChange.Add(-time.Hour)

	s := suspect{Node: "test", Incarnation: 1}
	m.suspectNode(&s)

	if state.State != stateSuspect {
		t.Fatalf("Bad state")
	}

	change := state.StateChange
	if time.Now().Sub(change) > time.Second {
		t.Fatalf("bad change delta")
	}

	// clear the broadcast queue
	m.broadcasts.Reset()

	// Suspect again
	m.suspectNode(&s)

	if state.StateChange != change {
		t.Fatalf("unexpected state change")
	}

	// Check a broad cast is queued
	if m.broadcasts.NumQueued() != 0 {
		t.Fatalf("expected only one queued message")
	}

}

func TestMemberList_SuspectNode_OldSuspect(t *testing.T) {
	m := GetMemberlist(t, nil)
	defer m.Shutdown()

	a := alive{Node: "test", Addr: []byte{127, 0, 0, 1}, Incarnation: 10, Vsn: m.config.BuildVsnArray()}
	m.aliveNode(&a, nil, false)

	state := m.nodeMap["test"]
	state.StateChange = state.StateChange.Add(-time.Hour)

	// Clear queue
	m.broadcasts.Reset()

	s := suspect{Node: "test", Incarnation: 1}
	m.suspectNode(&s)

	if state.State != stateAlive {
		t.Fatalf("Bad state")
	}

	// Check a broad cast is queued
	if m.broadcasts.NumQueued() != 0 {
		t.Fatalf("expected only one queued message")
	}
}

func TestMemberList_SuspectNode_Refute(t *testing.T) {
	m := GetMemberlist(t, nil)
	defer m.Shutdown()

	a := alive{Node: m.config.Name, Addr: []byte{127, 0, 0, 1}, Incarnation: 1, Vsn: m.config.BuildVsnArray()}
	m.aliveNode(&a, nil, true)

	// Clear queue
	m.broadcasts.Reset()

	// Make sure health is in a good state
	if score := m.GetHealthScore(); score != 0 {
		t.Fatalf("bad: %d", score)
	}

	s := suspect{Node: m.config.Name, Incarnation: 1}
	m.suspectNode(&s)

	state := m.nodeMap[m.config.Name]
	if state.State != stateAlive {
		t.Fatalf("should still be alive")
	}

	// Check a broad cast is queued
	if m.broadcasts.NumQueued() != 1 {
		t.Fatalf("expected only one queued message")
	}

	// Should be alive mesg
	if messageType(m.broadcasts.orderedView(true)[0].b.Message()[0]) != aliveMsg {
		t.Fatalf("expected queued alive msg")
	}

	// Health should have been dinged
	if score := m.GetHealthScore(); score != 1 {
		t.Fatalf("bad: %d", score)
	}
}

func TestMemberList_DeadNode_NoNode(t *testing.T) {
	m := GetMemberlist(t, nil)
	defer m.Shutdown()

	d := dead{Node: "test", Incarnation: 1}
	m.deadNode(&d)
	if len(m.nodes) != 0 {
		t.Fatalf("don't expect nodes")
	}
}

func TestMemberList_DeadNode(t *testing.T) {
	ch := make(chan NodeEvent, 1)

	m := GetMemberlist(t, func(c *Config) {
		c.Events = &ChannelEventDelegate{ch}
	})
	defer m.Shutdown()

	a := alive{Node: "test", Addr: []byte{127, 0, 0, 1}, Incarnation: 1, Vsn: m.config.BuildVsnArray()}
	m.aliveNode(&a, nil, false)

	// Read the join event
	<-ch

	state := m.nodeMap["test"]
	state.StateChange = state.StateChange.Add(-time.Hour)

	d := dead{Node: "test", Incarnation: 1}
	m.deadNode(&d)

	if state.State != stateDead {
		t.Fatalf("Bad state")
	}

	change := state.StateChange
	if time.Now().Sub(change) > time.Second {
		t.Fatalf("bad change delta")
	}

	select {
	case leave := <-ch:
		if leave.Event != NodeLeave || leave.Node.Name != "test" {
			t.Fatalf("bad node name")
		}
	default:
		t.Fatalf("no leave message")
	}

	// Check a broad cast is queued
	if m.broadcasts.NumQueued() != 1 {
		t.Fatalf("expected only one queued message")
	}

	// Check its a suspect message
	if messageType(m.broadcasts.orderedView(true)[0].b.Message()[0]) != deadMsg {
		t.Fatalf("expected queued dead msg")
	}
}

func TestMemberList_DeadNode_Double(t *testing.T) {
	ch := make(chan NodeEvent, 1)
	m := GetMemberlist(t, nil)
	defer m.Shutdown()

	a := alive{Node: "test", Addr: []byte{127, 0, 0, 1}, Incarnation: 1, Vsn: m.config.BuildVsnArray()}
	m.aliveNode(&a, nil, false)

	state := m.nodeMap["test"]
	state.StateChange = state.StateChange.Add(-time.Hour)

	d := dead{Node: "test", Incarnation: 1}
	m.deadNode(&d)

	// Clear queue
	m.broadcasts.Reset()

	// Notify after the first dead
	m.config.Events = &ChannelEventDelegate{ch}

	// Should do nothing
	d.Incarnation = 2
	m.deadNode(&d)

	select {
	case <-ch:
		t.Fatalf("should not get leave")
	default:
	}

	// Check a broad cast is queued
	if m.broadcasts.NumQueued() != 0 {
		t.Fatalf("expected only one queued message")
	}
}

func TestMemberList_DeadNode_OldDead(t *testing.T) {
	m := GetMemberlist(t, nil)
	defer m.Shutdown()

	a := alive{Node: "test", Addr: []byte{127, 0, 0, 1}, Incarnation: 10, Vsn: m.config.BuildVsnArray()}
	m.aliveNode(&a, nil, false)

	state := m.nodeMap["test"]
	state.StateChange = state.StateChange.Add(-time.Hour)

	d := dead{Node: "test", Incarnation: 1}
	m.deadNode(&d)

	if state.State != stateAlive {
		t.Fatalf("Bad state")
	}
}

func TestMemberList_DeadNode_AliveReplay(t *testing.T) {
	m := GetMemberlist(t, nil)
	defer m.Shutdown()

	a := alive{Node: "test", Addr: []byte{127, 0, 0, 1}, Incarnation: 10, Vsn: m.config.BuildVsnArray()}
	m.aliveNode(&a, nil, false)

	d := dead{Node: "test", Incarnation: 10}
	m.deadNode(&d)

	// Replay alive at same incarnation
	m.aliveNode(&a, nil, false)

	// Should remain dead
	state, ok := m.nodeMap["test"]
	if ok && state.State != stateDead {
		t.Fatalf("Bad state")
	}
}

func TestMemberList_DeadNode_Refute(t *testing.T) {
	m := GetMemberlist(t, nil)
	defer m.Shutdown()

	a := alive{Node: m.config.Name, Addr: []byte{127, 0, 0, 1}, Incarnation: 1, Vsn: m.config.BuildVsnArray()}
	m.aliveNode(&a, nil, true)

	// Clear queue
	m.broadcasts.Reset()

	// Make sure health is in a good state
	if score := m.GetHealthScore(); score != 0 {
		t.Fatalf("bad: %d", score)
	}

	d := dead{Node: m.config.Name, Incarnation: 1}
	m.deadNode(&d)

	state := m.nodeMap[m.config.Name]
	if state.State != stateAlive {
		t.Fatalf("should still be alive")
	}

	// Check a broad cast is queued
	if m.broadcasts.NumQueued() != 1 {
		t.Fatalf("expected only one queued message")
	}

	// Should be alive mesg
	if messageType(m.broadcasts.orderedView(true)[0].b.Message()[0]) != aliveMsg {
		t.Fatalf("expected queued alive msg")
	}

	// We should have been dinged
	if score := m.GetHealthScore(); score != 1 {
		t.Fatalf("bad: %d", score)
	}
}

func TestMemberList_MergeState(t *testing.T) {
	m := GetMemberlist(t, nil)
	defer m.Shutdown()

	a1 := alive{Node: "test1", Addr: []byte{127, 0, 0, 1}, Incarnation: 1, Vsn: m.config.BuildVsnArray()}
	m.aliveNode(&a1, nil, false)
	a2 := alive{Node: "test2", Addr: []byte{127, 0, 0, 2}, Incarnation: 1, Vsn: m.config.BuildVsnArray()}
	m.aliveNode(&a2, nil, false)
	a3 := alive{Node: "test3", Addr: []byte{127, 0, 0, 3}, Incarnation: 1, Vsn: m.config.BuildVsnArray()}
	m.aliveNode(&a3, nil, false)

	s := suspect{Node: "test1", Incarnation: 1}
	m.suspectNode(&s)

	remote := []pushNodeState{
		pushNodeState{
			Name:        "test1",
			Addr:        []byte{127, 0, 0, 1},
			Incarnation: 2,
			State:       stateAlive,
		},
		pushNodeState{
			Name:        "test2",
			Addr:        []byte{127, 0, 0, 2},
			Incarnation: 1,
			State:       stateSuspect,
		},
		pushNodeState{
			Name:        "test3",
			Addr:        []byte{127, 0, 0, 3},
			Incarnation: 1,
			State:       stateDead,
		},
		pushNodeState{
			Name:        "test4",
			Addr:        []byte{127, 0, 0, 4},
			Incarnation: 2,
			State:       stateAlive,
		},
	}

	// Listen for changes
	eventCh := make(chan NodeEvent, 1)
	m.config.Events = &ChannelEventDelegate{eventCh}

	// Merge remote state
	m.mergeState(remote)

	// Check the states
	state := m.nodeMap["test1"]
	if state.State != stateAlive || state.Incarnation != 2 {
		t.Fatalf("Bad state %v", state)
	}

	state = m.nodeMap["test2"]
	if state.State != stateSuspect || state.Incarnation != 1 {
		t.Fatalf("Bad state %v", state)
	}

	state = m.nodeMap["test3"]
	if state.State != stateSuspect {
		t.Fatalf("Bad state %v", state)
	}

	state = m.nodeMap["test4"]
	if state.State != stateAlive || state.Incarnation != 2 {
		t.Fatalf("Bad state %v", state)
	}

	// Check the channels
	select {
	case e := <-eventCh:
		if e.Event != NodeJoin || e.Node.Name != "test4" {
			t.Fatalf("bad node %v", e)
		}
	default:
		t.Fatalf("Expect join")
	}

	select {
	case e := <-eventCh:
		t.Fatalf("Unexpect event: %v", e)
	default:
	}
}

func TestMemberlist_Gossip(t *testing.T) {
	ch := make(chan NodeEvent, 3)

	addr1 := getBindAddr()
	addr2 := getBindAddr()
	addr3 := getBindAddr()
	ip1 := []byte(addr1)
	ip2 := []byte(addr2)
	ip3 := []byte(addr3)

	m1 := HostMemberlist(addr1.String(), t, func(c *Config) {
		// Set the gossip interval fast enough to get a reasonable test,
		// but slow enough to avoid "sendto: operation not permitted"
		c.GossipInterval = 10 * time.Millisecond
	})
	defer m1.Shutdown()

	bindPort := m1.config.BindPort

	m2 := HostMemberlist(addr2.String(), t, func(c *Config) {
		c.BindPort = bindPort
		c.Events = &ChannelEventDelegate{ch}
		// Set the gossip interval fast enough to get a reasonable test,
		// but slow enough to avoid "sendto: operation not permitted"
		c.GossipInterval = 10 * time.Millisecond
	})
	defer m2.Shutdown()

	m3 := HostMemberlist(addr2.String(), t, func(c *Config) {
	})
	defer m3.Shutdown()

	a1 := alive{Node: addr1.String(), Addr: ip1, Port: uint16(bindPort), Incarnation: 1, Vsn: m1.config.BuildVsnArray()}
	m1.aliveNode(&a1, nil, true)
	a2 := alive{Node: addr2.String(), Addr: ip2, Port: uint16(bindPort), Incarnation: 1, Vsn: m2.config.BuildVsnArray()}
	m1.aliveNode(&a2, nil, false)
	a3 := alive{Node: addr3.String(), Addr: ip3, Port: uint16(bindPort), Incarnation: 1, Vsn: m3.config.BuildVsnArray()}
	m1.aliveNode(&a3, nil, false)

	// Gossip should send all this to m2. Retry a few times because it's UDP and
	// timing and stuff makes this flaky without.
	retry(t, 15, 250*time.Millisecond, func(failf func(string, ...interface{})) {
		m1.gossip()

		time.Sleep(3 * time.Millisecond)

		if len(ch) < 3 {
			failf("expected 3 messages from gossip but only got %d", len(ch))
		}
	})
}

func retry(t *testing.T, n int, w time.Duration, fn func(func(string, ...interface{}))) {
	t.Helper()
	for try := 1; try <= n; try++ {
		failed := false
		failFormat := ""
		failArgs := []interface{}{}
		failf := func(format string, args ...interface{}) {
			failed = true
			failFormat = format
			failArgs = args
		}

		fn(failf)

		if !failed {
			return
		}
		if try == n {
			t.Fatalf(failFormat, failArgs...)
		}
		time.Sleep(w)
	}
}

func TestMemberlist_GossipToDead(t *testing.T) {
	ch := make(chan NodeEvent, 2)

	addr1 := getBindAddr()
	addr2 := getBindAddr()
	ip1 := []byte(addr1)
	ip2 := []byte(addr2)

	m1 := HostMemberlist(addr1.String(), t, func(c *Config) {
		c.GossipInterval = time.Millisecond
		c.GossipToTheDeadTime = 100 * time.Millisecond
	})
	defer m1.Shutdown()

	bindPort := m1.config.BindPort

	m2 := HostMemberlist(addr2.String(), t, func(c *Config) {
		c.BindPort = bindPort
		c.Events = &ChannelEventDelegate{ch}
	})

	defer m2.Shutdown()

	a1 := alive{Node: addr1.String(), Addr: ip1, Port: uint16(bindPort), Incarnation: 1, Vsn: m1.config.BuildVsnArray()}
	m1.aliveNode(&a1, nil, true)
	a2 := alive{Node: addr2.String(), Addr: ip2, Port: uint16(bindPort), Incarnation: 1, Vsn: m2.config.BuildVsnArray()}
	m1.aliveNode(&a2, nil, false)

	// Shouldn't send anything to m2 here, node has been dead for 2x the GossipToTheDeadTime
	m1.nodeMap[addr2.String()].State = stateDead
	m1.nodeMap[addr2.String()].StateChange = time.Now().Add(-200 * time.Millisecond)
	m1.gossip()

	select {
	case <-ch:
		t.Fatalf("shouldn't get gossip")
	case <-time.After(50 * time.Millisecond):
	}

	// Should gossip to m2 because its state has changed within GossipToTheDeadTime
	m1.nodeMap[addr2.String()].StateChange = time.Now().Add(-20 * time.Millisecond)

	retry(t, 5, 10*time.Millisecond, func(failf func(string, ...interface{})) {
		m1.gossip()

		time.Sleep(3 * time.Millisecond)

		if len(ch) < 2 {
			failf("expected 2 messages from gossip")
		}
	})
}

func TestMemberlist_FailedRemote(t *testing.T) {
	type test struct {
		name     string
		err      error
		expected bool
	}
	tests := []test{
		{"nil error", nil, false},
		{"normal error", fmt.Errorf(""), false},
		{"net.OpError for file", &net.OpError{Net: "file"}, false},
		{"net.OpError for udp", &net.OpError{Net: "udp"}, false},
		{"net.OpError for udp4", &net.OpError{Net: "udp4"}, false},
		{"net.OpError for udp6", &net.OpError{Net: "udp6"}, false},
		{"net.OpError for tcp", &net.OpError{Net: "tcp"}, false},
		{"net.OpError for tcp4", &net.OpError{Net: "tcp4"}, false},
		{"net.OpError for tcp6", &net.OpError{Net: "tcp6"}, false},
		{"net.OpError for tcp with dial", &net.OpError{Net: "tcp", Op: "dial"}, true},
		{"net.OpError for tcp with write", &net.OpError{Net: "tcp", Op: "write"}, true},
		{"net.OpError for tcp with read", &net.OpError{Net: "tcp", Op: "read"}, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := failedRemote(test.err)
			if actual != test.expected {
				t.Fatalf("expected %t, got %t", test.expected, actual)
			}
		})
	}
}

func TestMemberlist_PushPull(t *testing.T) {
	addr1 := getBindAddr()
	addr2 := getBindAddr()
	ip1 := []byte(addr1)
	ip2 := []byte(addr2)

	ch := make(chan NodeEvent, 3)

	m1 := HostMemberlist(addr1.String(), t, func(c *Config) {
		c.GossipInterval = 10 * time.Second
		c.PushPullInterval = time.Millisecond
	})
	defer m1.Shutdown()

	bindPort := m1.config.BindPort

	m2 := HostMemberlist(addr2.String(), t, func(c *Config) {
		c.BindPort = bindPort
		c.GossipInterval = 10 * time.Second
		c.Events = &ChannelEventDelegate{ch}
	})
	defer m2.Shutdown()

	a1 := alive{Node: addr1.String(), Addr: ip1, Port: uint16(bindPort), Incarnation: 1, Vsn: m1.config.BuildVsnArray()}
	m1.aliveNode(&a1, nil, true)
	a2 := alive{Node: addr2.String(), Addr: ip2, Port: uint16(bindPort), Incarnation: 1, Vsn: m2.config.BuildVsnArray()}
	m1.aliveNode(&a2, nil, false)

	// Gossip should send all this to m2. It's UDP though so retry a few times
	retry(t, 5, 10*time.Millisecond, func(failf func(string, ...interface{})) {
		m1.pushPull()

		time.Sleep(3 * time.Millisecond)

		if len(ch) < 2 {
			failf("expected 2 messages from pushPull")
		}
	})
}

func TestVerifyProtocol(t *testing.T) {
	cases := []struct {
		Anodes   [][3]uint8
		Bnodes   [][3]uint8
		expected bool
	}{
		// Both running identical everything
		{
			Anodes: [][3]uint8{
				{0, 0, 0},
			},
			Bnodes: [][3]uint8{
				{0, 0, 0},
			},
			expected: true,
		},

		// One can understand newer, but speaking same protocol
		{
			Anodes: [][3]uint8{
				{0, 0, 0},
			},
			Bnodes: [][3]uint8{
				{0, 1, 0},
			},
			expected: true,
		},

		// One is speaking outside the range
		{
			Anodes: [][3]uint8{
				{0, 0, 0},
			},
			Bnodes: [][3]uint8{
				{1, 1, 1},
			},
			expected: false,
		},

		// Transitively outside the range
		{
			Anodes: [][3]uint8{
				{0, 1, 0},
				{0, 2, 1},
			},
			Bnodes: [][3]uint8{
				{1, 3, 1},
			},
			expected: false,
		},

		// Multi-node
		{
			Anodes: [][3]uint8{
				{0, 3, 2},
				{0, 2, 0},
			},
			Bnodes: [][3]uint8{
				{0, 2, 1},
				{0, 5, 0},
			},
			expected: true,
		},
	}

	for _, tc := range cases {
		aCore := make([][6]uint8, len(tc.Anodes))
		aApp := make([][6]uint8, len(tc.Anodes))
		for i, n := range tc.Anodes {
			aCore[i] = [6]uint8{n[0], n[1], n[2], 0, 0, 0}
			aApp[i] = [6]uint8{0, 0, 0, n[0], n[1], n[2]}
		}

		bCore := make([][6]uint8, len(tc.Bnodes))
		bApp := make([][6]uint8, len(tc.Bnodes))
		for i, n := range tc.Bnodes {
			bCore[i] = [6]uint8{n[0], n[1], n[2], 0, 0, 0}
			bApp[i] = [6]uint8{0, 0, 0, n[0], n[1], n[2]}
		}

		// Test core protocol verification
		testVerifyProtocolSingle(t, aCore, bCore, tc.expected)
		testVerifyProtocolSingle(t, bCore, aCore, tc.expected)

		//  Test app protocol verification
		testVerifyProtocolSingle(t, aApp, bApp, tc.expected)
		testVerifyProtocolSingle(t, bApp, aApp, tc.expected)
	}
}

func testVerifyProtocolSingle(t *testing.T, A [][6]uint8, B [][6]uint8, expect bool) {
	m := GetMemberlist(t, nil)
	defer m.Shutdown()

	m.nodes = make([]*nodeState, len(A))
	for i, n := range A {
		m.nodes[i] = &nodeState{
			Node: Node{
				PMin: n[0],
				PMax: n[1],
				PCur: n[2],
				DMin: n[3],
				DMax: n[4],
				DCur: n[5],
			},
		}
	}

	remote := make([]pushNodeState, len(B))
	for i, n := range B {
		remote[i] = pushNodeState{
			Name: fmt.Sprintf("node %d", i),
			Vsn:  []uint8{n[0], n[1], n[2], n[3], n[4], n[5]},
		}
	}

	err := m.verifyProtocol(remote)
	if (err == nil) != expect {
		t.Fatalf("bad:\nA: %v\nB: %v\nErr: %s", A, B, err)
	}
}
