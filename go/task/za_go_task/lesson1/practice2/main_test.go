package main

import (
	"fmt"
	"testing"
)

func TestBubble(t *testing.T) {
	var arr = []int{5, -1, 0, 12, 3, 5}
	fmt.Printf("original arr:%v\n", arr)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}

	}

	fmt.Printf("after bubble sort:%v\n", arr)
}
