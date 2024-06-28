package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println(os.Args)

	readAndWriteFileContents(os.Args[1])

}

type customWriter struct{}

func (cw customWriter) Write(p []byte) (int, error) {
	fmt.Println("Calling my custom writer")
	fmt.Println(string(p))
	return len(p), nil
}

func readAndWriteFileContents(n string) {
	f, err := os.Open(n)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// io.Copy(os.Stdout, f)
	io.Copy(customWriter{}, f)

}
