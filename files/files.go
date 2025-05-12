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

}
