package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestRecursion4CartesianProduct(t *testing.T) {
	arr := []string{"LONGAN", "DIRECT", "SALES", "TECHNICAL", "TEAM"}
	var count int
	tempArr := make([]string, len(arr), len(arr))
	var level int

	cartesianProduct(arr, tempArr, level, &count)
	println("\ntotal:" + strconv.Itoa(count))
}

func cartesianProduct(arr []string, tempArr []string, level int, count *int) {
	for i := 0; i < len(arr[level]); i++ {
		tempArr[level] = string(arr[level][i])
		if level < len(arr)-1 {
			cartesianProduct(arr, tempArr, level+1, count)
		} else {
			*count++
			println(strconv.Itoa(*count) + ":" + strings.Join(tempArr, ""))
		}

	}
}
