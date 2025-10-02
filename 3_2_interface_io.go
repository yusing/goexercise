package goexercises

import "io"

/*
Go IO Interface: A Brief Introduction

The io package in Go provides interfaces for reading and writing data.
It's like the Stream interface in JavaScript, but with a completely different model.

There are two main interfaces:
- io.Reader: Read data from a source, implement Read(p []byte) (n int, err error)
- io.Writer: Write data to a destination, implement Write(p []byte) (n int, err error)

There are also some helper functions:
- io.Copy(dst io.Writer, src io.Reader) (n int64, err error)
- io.ReadAll(r io.Reader) ([]byte, error)
- io.WriteString(w io.Writer, s string) (n int, err error)
*/

type MyStringReader struct {
	s   string
	pos int
}

func NewMyStringReader(s string) MyStringReader {
	return MyStringReader{s: s, pos: 0}
}

func (r *MyStringReader) Read(p []byte) (n int, err error) {
	// TODO: implement
	// read the string into the buffer
	// return the number of bytes read and any error
	//
	// tips: you can use r.pos to track the position in the string
	// tips: use builtin `copy(dst, src) int`; return io.EOF if you reach the end of the string
	return 0, nil
}

type ConsoleWriterWithTimestamp struct {
	w io.Writer
}

func NewConsoleWriterWithTimestamp(w io.Writer) ConsoleWriterWithTimestamp {
	return ConsoleWriterWithTimestamp{w: w}
}

func (w *ConsoleWriterWithTimestamp) Write(p []byte) (n int, err error) {
	// TODO: implement
	// write the string to the console with a timestamp in format of "[2021-01-01 12:00:00] content"
	// return the number of bytes written and any error
	//
	// tips: use fmt.Fprintf to write to the writer with format string "%s"
	//       you can use time.Now() to get the current timestamp
	return 0, nil
}
