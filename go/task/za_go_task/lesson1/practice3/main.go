package main

import "fmt"

func main() {
	//	level := -1
	//	println("这是一个金字塔程序，请输入层数(2-20)：")
	//INPUT:
	//	fmt.Scan(&level)
	//	if level < 2 || level > 20 {
	//		println("这是一个金字塔程序，请输入层数(2-20)：")
	//		goto INPUT
	//	}
	//
	//	for i := 1; i <= level; i++ {
	//		for j := level; j > i; j-- {
	//			print(" ")
	//		}
	//		for k := 1; k <= i; k++ {
	//			print("A ")
	//		}
	//		println()
	//	}
	var words = []string{"LONGAN", "DIRECT", "SALES", "TECHNICAL", "TEAM"}
	var result = CartesianProduct(words)
	for _, s := range result {
		fmt.Println(s)
	}
	fmt.Println("size:", len(result))
}

func CartesianProduct(words []string) []string {
	var result []string
	length := len(words)
	if length == 0 {
		return result
	}

	// Calculate the Cartesian product of the words
	size := 1
	for _, word := range words {
		size *= len(word)
	}
	for i := 0; i < size; i++ {
		// Generate each individual result
		str := ""
		for j := 0; j < length; j++ {
			index := (i / length) % len(words[j])
			str += string(words[j][index])
		}
		result = append(result, str)
	}
	return result
}
