package main

import (
	"fmt"
	"os"
)

func main() {
	td := os.TempDir()

	fmt.Println("Hello, playground", td)

	file, err := os.Create(td + "xx.txt")
	fmt.Println(err)
	_ = file

	n, err := file.WriteString("blubb")
	fmt.Println(err)
	fmt.Println(n, "bytes written")

	file1, err := os.Open(td + "xx.txt")
	fmt.Println(err)
	b := make([]byte, 10)
	n1, err := file1.Read(b)
	fmt.Println(err)
	fmt.Println(n1, "bytes read")
	fmt.Println(string(b))

}
