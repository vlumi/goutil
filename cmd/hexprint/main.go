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
	var rawOutput, hexOutput string
	for pos := range columns {
		switch {
		case pos == 0:
		case pos%group == 0:
			rawOutput += " "
			hexOutput += " "
		}
		if pos < len(buffer) {
			b := buffer[pos]
			if '!' <= b && b <= '~' {
				rawOutput += fmt.Sprintf("%c", b)
			} else {
				rawOutput += " "
			}
			hexOutput += fmt.Sprintf("%02x ", b)
		} else {
			rawOutput += " "
			hexOutput += "   "
		}
	}
	fmt.Printf("%08x ", offset)
	fmt.Print(rawOutput)
	fmt.Print(" ")
	fmt.Print(hexOutput)
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
