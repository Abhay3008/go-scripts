package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	file, _ := os.Create("./file.txt")
	io.WriteString(file, "Hello World!")
	fmt.Println("file has been created")
	file.Close()
	readfile, _ := os.ReadFile("./file.txt")
	fmt.Println("reading file: ", string(readfile))

}
