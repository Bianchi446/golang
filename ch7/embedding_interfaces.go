/*
	It is possible to embbed an interface in an interface
*/

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Closer interface {
	Close() error

}

type ReadCloser interface {
	Reader
	Closer
}

