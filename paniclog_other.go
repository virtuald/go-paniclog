// Log the panic to the log file - for oses which can't do this

// +build !windows,!darwin,!dragonfly,!freebsd,!linux,!nacl,!netbsd,!openbsd

package log

import (
	"errors"
	"os"
)

func redirectStderr(f *os.File) error {
	return errors.New("Can't redirect stderr to file")
}