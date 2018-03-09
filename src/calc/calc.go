/**
 * go 实现的缩小版linux 计算器
 */
package main

import (
	"fmt"
	"os"
	"simplemath"
	"strconv"
)
// printf help info
var Usage = func() {
	fmt.Println("usage: calc command [arguments] ...")
	fmt.Println("\nThe commands are: \n\tadd\tAddtion of two values.\n\tsqrt\tSquare root of a non-negative value.")
}
// main
func main() {
	args := os.Args[1:]
	// printf help when input error 
	if args == nil || len(args) < 2 {
		Usage()
		return
	}
	switch args[0] {
		case "add":
			if len(args) != 3 {
				fmt.Println("USAGE: calc add <integer1><integer2>")
				return
			}
			// strconv包实现了基本数据类型和其字符串表示的相互转换。
			v1, err1 := strconv.Atoi(args[1])
			v2, err2 := strconv.Atoi(args[2])

			if err1 != nil || err2 != nil {
				fmt.Println("USAGE: calc add <integer1><integer1>")
				return
			}
			
			ret := simplemath.Add(v1, v2)
			fmt.Println("result: ", ret)
		case "sqrt":
			if len(args) != 2 {
				fmt.Println("USAGE: calc sqrt <integer>")
				return
			}
			v, err := strconv.Atoi(args[1])

			if err != nil {
				fmt.Println("USAGE: calc sqrt <integer>")
				return
			}

			ret := simplemath.Sqrt(v)

			fmt.Println("result: ", ret)
		default:
			Usage()
	}
}