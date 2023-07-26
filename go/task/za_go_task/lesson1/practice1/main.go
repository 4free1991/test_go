package main

import (
	"fmt"
)

func main() {
	level := -1
	println("这是一个金字塔程序，请输入层数(2-20)：")
INPUT:
	fmt.Scan(&level)
	if level < 2 || level > 20 {
		println("这是一个金字塔程序，请输入层数(2-20)：")
		goto INPUT
	}

	for i := 1; i <= level; i++ {
		for j := level; j > i; j-- {
			print(" ")
		}
		for k := 1; k <= i; k++ {
			print("A ")
		}
		println()
	}
}
