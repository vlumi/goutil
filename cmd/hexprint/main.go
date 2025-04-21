package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

var columns = 16
var bufferSize = columns * 64
var group = 8

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No file specified")
		return
	}
	err := iterate(os.Args[1], printRow)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
}

func printRow(offset int, buffer []byte) {
	fmt.Printf("%08x ", offset)
	for pos, b := range buffer {
		switch {
		case pos == 0:
		case pos%group == 0:
			fmt.Print(" ")
		}
		fmt.Printf("%02x ", b)
	}
	fmt.Println()
}

func iterate(fileName string, processRow func(int, []byte)) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	buffer := make([]byte, bufferSize)
	offset := 0
	for {
		count, err := f.Read(buffer)
		for start := 0; start < count; start += columns {
			end := min(start+columns, count)
			if start < end {
				processRow(offset+start, buffer[start:end])
			}
		}
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		offset += count
	}
}
