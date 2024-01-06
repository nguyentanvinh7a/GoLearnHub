package main

import (
	"bytes"
	"fmt"
)

func main() {
	// var w Writer = &ConsoleWriter{}
	// w.Write([]byte("Hello Go!"))

	// ic := IntCounter(0)
	// var inc Increamenter = &ic
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(inc.Increamenter())
	// }

	// var wc WriterCloser = NewBufferWriterCloser()
	// wc.Write([]byte("Hello Go!"))
	// wc.Close()

	// var obj interface{} = NewBufferWriterCloser()
	// obj.(WriterCloser).Write([]byte("Hello Go!"))
	// obj.(WriterCloser).Close()

	// var obj interface{} = NewBufferWriterCloser()
	// wc, err := obj.(WriterCloser)
	// if err {
	// 	wc.Write([]byte("Hello Go!"))
	// 	wc.Close()
	// } else {
	// 	fmt.Println("Object is not a writercloser")
	// }

	var obj interface{} = "Hello"
	switch obj.(type) {
	case int:
		fmt.Println("This is an int")
	case string:
		fmt.Println("This is a string")
	default:
		fmt.Println("I don't know what this is")
	}
}

type Writer interface {
	Write([]byte) (int, error)
}

// type ConsoleWriter struct {}

// func (cw *ConsoleWriter) Write(data []byte) (int, error) {
// 	n, err := fmt.Println(string(data))
// 	return n, err
// }

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type BufferWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}
	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

func (bwc *BufferWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferWriterCloser() *BufferWriterCloser {
	return &BufferWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

type Increamenter interface {
	Increamenter() int
}

type IntCounter int

func (ic *IntCounter) Increamenter() int {
	*ic++
	return int(*ic)
}
