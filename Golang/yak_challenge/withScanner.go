package main

import (
	"bufio"
	"fmt"
	"io"
)

func main() {

}

func withScanner(input io.ReadSeeker, start int64) error {
	fmt.Println("--SCANNER, start:", start)
	if _, err := input.Seek(start, 0); err != nil {
		return err
	}
	scanner := bufio.NewScanner(input)

	pos := start
	scanLines := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanLines(data, atEOF)
		pos += int64(advance)
		return
	}
	scanner.Split(scanLines)

	for scanner.Scan() {
		fmt.Printf("Pos: %d, Scanned: %s\n", pos, scanner.Text())
	}
	return scanner.Err()
}
