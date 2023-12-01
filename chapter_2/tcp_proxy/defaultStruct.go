package defaultStruct

type Reader interface {
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
}

type FooWriter struct{}

func (foo *FooWriter) Write(p []byte) (int, error) {
	return len(p), nil
}

// This is the generalised pattern to create a function that initialises a particular structure

