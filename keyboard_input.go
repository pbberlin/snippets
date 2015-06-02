package main

import (
	"fmt"
	"io"
	"os"
	"syscall"
)

import "bufio"

var StdInDuringTest *os.File

func main() {

	{
		stat, err := os.Stdin.Stat()
		fmt.Printf("%#v, %#v, %#v\n", os.Stdin, stat, err)
	}
	{
		StdInDuringTest = os.NewFile(uintptr(syscall.Stdin), "/dev/tty")
		stat, err := StdInDuringTest.Stat()
		fmt.Printf("%#v, %#v, %#v\n", StdInDuringTest, stat, err)
	}
	{
		StdInDuringTest = os.NewFile(uintptr(syscall.Stdin), "/dev/stdin")
		stat, err := StdInDuringTest.Stat()
		fmt.Printf("%#v, %#v, %#v\n", StdInDuringTest, stat, err)
	}

	{
		StdInDuringTest = os.NewFile(0, "/dev/stdin2")
		stat, err := StdInDuringTest.Stat()
		fmt.Printf("%#v, %#v, %#v\n", StdInDuringTest, stat, err)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("tell! ")
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Print("EOF")
			} else {
				break
			}
		}
		fmt.Printf("u said %q \n", line)
		if line == "q" {
			break
		}
	}
}
