// write file
package main	

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"io/ioutil"
	"strconv"
)
/**
 * 注意 fileName 使用的是相对路径, 根据测试发现该相对路径与执行 **go run **命令时所在的路径有关
 * 比如: 
 *    1 在**write.go**所在目录下执行 go run命令, 则会在该目录下创建outprint.txt
 *    2 在**src**同级目录下执行, 则会在src同级目录下创建outprint.txt
 */
var (
	// writeString = "hello world"
	fileName = "./outprint.txt"
	f *os.File
	err error
)
func main() {
	args := os.Args[1:]

	if  args == nil || 1 > len(args) {
		Usage()
	}
	// for _, s := range args {
	// 	fmt.Println(s)
	// }
	num, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("usage: command [args]<integer>")
		return
	}
	switch  num {
		case 1:
			writeByIO(fileName)
		case 2:
			writeByIOUtil(fileName)
		case 3: 
			writeByFile(fileName)
		case 4: 
			writeByBuf(fileName)
		default: 
			Usage()
	}
}
// help tip
func Usage() {
	fmt.Println("usage: command [args]<integer> ...")
	fmt.Println("\nThe args value are: \n\t1: write by IO\n\t2: write by ioutil\n\t3: write by File\n\t4: write by buffer.")
}
// check file exited
func isExist(file string) {
	exist := true
	if _, err := os.Stat(file); os.IsNotExist(err) {
		exist = false
	}
	if exist {
		f, err = os.OpenFile(file, os.O_APPEND, 0666)
		// f, err = os.Open(fileName)
		fmt.Println("file existed")
	} else {
		f, err = os.Create(file)
		fmt.Println("no found, created")
	}
	check(err)
	// return f
}
// check err
func check(err error) {
	if err != nil {
		panic(err)
	}
}
// 1 write by io.writeString
func writeByIO(fileName string) {
	isExist(fileName)
	n, err := io.WriteString(f, "hello golang by io")
	check(err)
	fmt.Printf("write %d byte ", n)
}
// 2 write by ioutil
func writeByIOUtil(fileName string) {
	d := []byte("hello golang by ioutil")
	err := ioutil.WriteFile(fileName, d, 0666)
	check(err)
}
// 3 write by file
func writeByFile(fileName string) {
	d := []byte("hello golang by File")
	isExist(fileName)
	n, err := f.Write(d)
	check(err)

	fmt.Printf("write byte %d.", n)
}
// 4 write by bufio
func writeByBuf(fileName string) {
	isExist(fileName)
	w := bufio.NewWriter(f)
	d, err := w.WriteString("hello golang by bufio")
	check(err)

	fmt.Printf("write %d byte.", d)
	w.Flush()
	f.Close()
}