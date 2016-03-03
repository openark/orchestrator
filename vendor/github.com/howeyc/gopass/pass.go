package gopass

import (
	"errors"
	"fmt"
	"io"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

var defaultGetCh = func() (byte, error) {
	buf := make([]byte, 1)
	if n, err := os.Stdin.Read(buf); n == 0 || err != nil {
		if err != nil {
			return 0, err
		}
		return 0, io.EOF
	}
	return buf[0], nil
}

var (
	ErrInterrupted = errors.New("Interrupted")

	// Provide variable so that tests can provide a mock implementation.
	getch = defaultGetCh
)

// getPasswd returns the input read from terminal.
// If masked is true, typing will be matched by asterisks on the screen.
// Otherwise, typing will echo nothing.
func getPasswd(masked bool) ([]byte, error) {
	var err error
	var pass, bs, mask []byte
	if masked {
		bs = []byte("\b \b")
		mask = []byte("*")
	}

	if terminal.IsTerminal(int(os.Stdin.Fd())) {
		if oldState, err := terminal.MakeRaw(int(os.Stdin.Fd())); err != nil {
			return pass, err
		} else {
			defer terminal.Restore(int(os.Stdin.Fd()), oldState)
		}
	}

	for {
		if v, e := getch(); e != nil {
			err = e
			break
		} else if v == 127 || v == 8 {
			if l := len(pass); l > 0 {
				pass = pass[:l-1]
				fmt.Print(string(bs))
			}
		} else if v == 13 || v == 10 {
			break
		} else if v == 3 {
			err = ErrInterrupted
			break
		} else if v != 0 {
			pass = append(pass, v)
			fmt.Print(string(mask))
		}
	}
	fmt.Println()
	return pass, err
}

// GetPasswd returns the password read from the terminal without echoing input.
// The returned byte array does not include end-of-line characters.
func GetPasswd() ([]byte, error) {
	return getPasswd(false)
}

// GetPasswdMasked returns the password read from the terminal, echoing asterisks.
// The returned byte array does not include end-of-line characters.
func GetPasswdMasked() ([]byte, error) {
	return getPasswd(true)
}
