package goexercises

import (
	"bytes"
	"goexercises"
	"io"
	"strings"
	"testing"
	"time"
)

func TestMyStringReader_Read_SmallBufferChunks(t *testing.T) {
	r := goexercises.NewMyStringReader("hello")
	buf := make([]byte, 2)

	// 1st chunk
	n, err := r.Read(buf)
	if n != 2 || err != nil {
		t.Fatalf("first read: n=%d err=%v; want n=2 err=nil", n, err)
	}
	if string(buf) != "he" {
		t.Fatalf("first read bytes=%q; want %q", string(buf), "he")
	}

	// 2nd chunk
	n, err = r.Read(buf)
	if n != 2 || err != nil {
		t.Fatalf("second read: n=%d err=%v; want n=2 err=nil", n, err)
	}
	if string(buf) != "ll" {
		t.Fatalf("second read bytes=%q; want %q", string(buf), "ll")
	}

	// last chunk should hit EOF
	n, err = r.Read(buf)
	if n != 1 || err != io.EOF {
		t.Fatalf("third read: n=%d err=%v; want n=1 err=EOF", n, err)
	}
	if string(buf[:n]) != "o" {
		t.Fatalf("third read bytes=%q; want %q", string(buf[:n]), "o")
	}
}

func TestMyStringReader_Read_ExactBuffer(t *testing.T) {
	r := goexercises.NewMyStringReader("world")
	buf := make([]byte, 5)

	n, err := r.Read(buf)
	if n != 5 || err != io.EOF {
		t.Fatalf("read: n=%d err=%v; want n=5 err=EOF", n, err)
	}
	if string(buf) != "world" {
		t.Fatalf("read bytes=%q; want %q", string(buf), "world")
	}
}

func TestMyStringReader_Read_EmptyString(t *testing.T) {
	r := goexercises.NewMyStringReader("")
	buf := make([]byte, 1)

	n, err := r.Read(buf)
	if n != 0 || err != io.EOF {
		t.Fatalf("read empty: n=%d err=%v; want n=0 err=EOF", n, err)
	}
}

func TestConsoleWriterWithTimestamp_Write(t *testing.T) {
	var out bytes.Buffer
	w := goexercises.NewConsoleWriterWithTimestamp(&out)

	const prefixLen = len("[2021-01-01 12:00:00] ")

	input := []byte("abc123")
	n, err := w.Write(input)
	if err != nil {
		t.Fatalf("Write returned error: %v", err)
	}
	if n != len(input)+prefixLen {
		t.Fatalf("Write n=%d; want %d", n, len(input)+prefixLen)
	}

	timestamp := out.Bytes()[:prefixLen]
	if _, err := time.Parse(time.DateTime, strings.TrimSpace(string(timestamp))); err != nil {
		t.Fatalf("timestamp %q is not in valid format", string(timestamp))
	}

	content := out.Bytes()[prefixLen:]
	if string(content) != string(input) {
		t.Fatalf("content %q is not equal to input %q", string(content), string(input))
	}
}
