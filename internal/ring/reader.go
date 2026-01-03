package ring

import "io"

type Reader interface {
	io.ReadCloser
	io.ReaderAt
}
