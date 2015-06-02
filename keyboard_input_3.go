package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("tell! ")
		line, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				fmt.Print("EOF")
			} else {
				fmt.Print("other err ", err)
				break
			}
		}
		fmt.Printf("u said %q\n", line)
		if line == 'q' {
			break
		}
	}

}
