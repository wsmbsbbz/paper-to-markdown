package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		reader := bufio.NewReaderSize(os.Stdin, 10000)
		writer := bufio.NewWriter(os.Stdout)
		rawToOutput(reader, writer)
	} else {
		for _, fn := range os.Args[1:] {
			f, err := os.Open(fn)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			of, err := os.Create(fn + "-output")
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			reader := bufio.NewReader(f)
			writer := bufio.NewWriter(of)
			rawToOutput(reader, writer)
			f.Close()
			of.Close()
		}
	}
}

func rawToOutput(r *bufio.Reader, w *bufio.Writer) error {
	defer w.Flush()
	for true {
		s, err := r.ReadString('\n')
		if err == io.EOF {
			w.WriteString(s)
			break
		}
		if s != "\n" {
			s = s[:len(s)-1]
			if s[len(s)-1] == '-' {
				s = s[:len(s)-1]
			} else {
				s = s + " "
			}
		} else {
			w.WriteString(s)
		}
		w.WriteString(s)
	}
	return nil
}
