package paniclog

import "os"

// RedirectStderr to the file passed in, so that the output of any panics that
// occur will be sent to that file.
//
// Of course, anything else written to stderr will also be sent to that file,
// so don't do that unless that's your intent.
func RedirectStderr(f *os.File) error {
	return redirectStderr(f)
}
