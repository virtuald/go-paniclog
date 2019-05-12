// Log the panic under unix to the log file

// +build !windows,!solaris,!plan9

package paniclog

import (
	"errors"
	"os"

	"golang.org/x/sys/unix"
)

func redirectStderr(f *os.File) error {
	err := unix.Dup2(int(f.Fd()), int(os.Stderr.Fd()))
	if err != nil {
		return errors.New("Failed to redirect stderr to file: " + err.Error())
	}
	return nil
}
