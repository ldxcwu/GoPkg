package io_test

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

func TestIOCopy(t *testing.T) {
	reader := strings.NewReader("Hello World!\n")
	if _, err := io.Copy(os.Stdout, reader); err != nil {
		log.Fatal(err)
	}
}

func TestCopyBuffer(t *testing.T) {
	r1 := strings.NewReader("Reader 1\n")
	r2 := strings.NewReader("Reader 2\n")
	buf := make([]byte, 0)
	buf = append(buf, 'H')
	buf = append(buf, 'H')

	//In fact, strings.NewReader() has implemented WriterTo and ReaderFrom.
	//So it will not use the buf provided to copy.

	if _, err := io.CopyBuffer(os.Stdout, r1, buf); err != nil {
		log.Fatal(err)
	}

	if _, err := io.CopyBuffer(os.Stdout, r2, buf); err != nil {
		log.Fatal(err)
	}
}

func TestLimitReader(t *testing.T) {
	r := strings.NewReader("Something...\n")
	lr := io.LimitReader(r, 4)
	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}
}

func TestTeeReader(t *testing.T) {
	r := strings.NewReader("Something provided to be read...\n")
	tr := io.TeeReader(r, os.Stdout)
	if _, err := io.ReadAll(tr); err != nil {
		log.Fatal(err)
	}
}

func TestReaderAt(t *testing.T) {
	r := strings.NewReader("Something provided to be read...\n")
	b := make([]byte, 3)
	_, err := r.ReadAt(b, 1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("b: %v\n", b)
	_, err = r.ReadAt(b, 1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("b: %v\n", b)
	if _, err := io.CopyN(os.Stdout, r, 3); err != nil {
		panic(err)
	}
	_, err = r.ReadAt(b, 1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("b: %v\n", b)
	if _, err := io.Copy(os.Stdout, r); err != nil {
		panic(err)
	}
}

func TestReaderFrom(t *testing.T) {
	r := strings.NewReader("Something provided to be read...\n")
	os.Stdout.ReadFrom(r)

	r = strings.NewReader("Something provided to be read...\n")
	m := MyReaderFrom{os.Stdout}
	m.ReadFrom(r)

	r = strings.NewReader("Something provided to be read...\n")
	m = MyReaderFrom{&MyWriter{}}
	m.ReadFrom(r)
	fmt.Printf("m.w: %v\n", m.w)
}

// type ReaderFrom interface {
// 	ReadFrom(r Reader) (n int64, err error)
// }
type MyReaderFrom struct {
	w io.Writer
}

func (m *MyReaderFrom) ReadFrom(r io.Reader) (int64, error) {
	return io.Copy(m.w, r)
}

// type Writer interface {
// 	Write(p []byte) (n int, err error)
// }
type MyWriter struct {
	content []byte
}

func (m *MyWriter) Write(p []byte) (int, error) {
	m.content = p
	return len(p), nil
}
