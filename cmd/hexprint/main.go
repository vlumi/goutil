package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

var bufferSize = 512
var columns = 24
var group = 8

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No file specified")
		return
	}
	err := iterate(os.Args[1], process)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
}

func process(buffer []byte, offset int, count int) {
	for i, b := range buffer[:count] {
		pos := (i + offset)
		switch {
		case pos == 0:
		case pos%columns == 0:
			fmt.Println()
		case pos%group == 0:
			fmt.Print(" ")
		}
		fmt.Printf("%02x ", b)
	}
}

func iterate(fileName string, process func([]byte, int, int)) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	buffer := make([]byte, bufferSize)
	offset := 0
	for {
		count, err := f.Read(buffer)
		process(buffer, offset, count)
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		offset += count
	}
}
