// write file
package main	

import (
	"fmt"
	"os"
	"io"
)

// func main
func main() {
	var (
		writeString = "hello world"
		// go run 时所用的当前目录为GOPath 指定的路径
		fileName = "./src/fileio/outprint.txt"
		f *os.File
		err error
	)
	// check file 
	if isExist(fileName) {
		f, err = os.OpenFile(fileName, os.O_APPEND, 0666)
		fmt.Println("file existed")
	} else {
		f, err = os.Create(fileName)
		fmt.Println("file no found")
	}
	n, err := io.WriteString(f, writeString)
	check(err)
	fmt.Printf("write %d byte ", n)
}

// check file exited
func isExist(file string) bool {
	exist := true
	if _, err := os.Stat(file); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

// check err
func check(err error) {
	if err != nil {
		panic(err)
	}
} 