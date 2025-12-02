package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

// ReadInput reads file containing input information and parses each line of file to T.
func ReadInput[R any](file string, mapFunc func(string) (R, error)) ([]R, error) {
	fi, err := os.Lstat(file)
	if err != nil {
		return nil, err
	}
	if m := fi.Mode(); !m.IsRegular() {
		return nil, fmt.Errorf("want regular file, got: %v", m)
	}

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return readFile(f, mapFunc)
}

func readFile[R any](r io.Reader, mapFunc func(string) (R, error)) ([]R, error) {
	var result []R
	br := bufio.NewReader(r)
	for i := 0; ; i++ {
		l, _, err := br.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}

		r, err := mapFunc(string(l))
		if err != nil {
			return nil, fmt.Errorf("failed to map file string #%d to result type: %w", i, err)
		}
		result = append(result, r)
	}

	return result, nil
}
