package main

import (
	"bufio"
	"io"
	"strings"
)

// Exercise 7.1: Using the ideas from ByteCounter, implement counters for words and for lines.

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	count, err := scanCount(strings.NewReader(string(p)), bufio.ScanWords)
	if err != nil {
		return 0, err
	}
	*c += WordCounter(count)
	return int(*c), nil
}

type LineCounter int

func (lc *LineCounter) Write(p []byte) (int, error) {
	count, err := scanCount(strings.NewReader(string(p)), bufio.ScanLines)
	if err != nil {
		return 0, err
	}
	*lc += LineCounter(count)
	return int(*lc), nil
}

func scanCount(s io.Reader, fn bufio.SplitFunc) (int, error) {
	count := 0
	scanner := bufio.NewScanner(s)
	scanner.Split(fn)
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return count, nil
}

func count(counter io.Writer, s string) error {
	_, err := counter.Write([]byte(s))
	return err
}

func main() {

}
