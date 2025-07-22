package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
)

func countLetter(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := make(map[string]int)
	for {
		n, err := r.Read(buf)
		for _, b := range buf[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}
}

func simpleCountLetter() error {
	s := "Hello World"
	sr := strings.NewReader(s)
	result, err := countLetter(sr)
	if err != nil {
		return err
	}
	fmt.Printf("result: %d\n", result)
	return nil
}

func buildGZipReader(fileName string) (*gzip.Reader, func(), error) {
	open, err := os.Open(fileName)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file %s: %w", fileName, err)
	}
	reader, err := gzip.NewReader(open)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating gzip reader for file %s: %w", fileName, err)
	}
	return reader, func() {
		reader.Close()
		open.Close()
	}, nil
}

func gzipCountLetters() error {
	reader, f, err := buildGZipReader("io/my_data.txt.gz")
	if err != nil {
		return fmt.Errorf("error building gzip reader: %w", err)
	}
	defer f()
	result, err := countLetter(reader)
	if err != nil {
		return fmt.Errorf("error counting letters in gzip file: %w", err)
	}
	fmt.Printf("result: %d\n", result)
	return nil
}

func main() {
	err := simpleCountLetter()
	if err != nil {
		slog.Error("error with simpleCountLetters", "msg", err)
	}

	err = gzipCountLetters()
	if err != nil {
		slog.Error("error with gzipCountLetters", "msg", err)
	}
}
