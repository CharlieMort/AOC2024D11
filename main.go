package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Blink(num int, blinkNum int) []int {
	if blinkNum > 75 {
		return []int{num}
	}
	numStr := fmt.Sprintf("%d", num)
	if num == 0 {
		return Blink(1, blinkNum+1)
	} else if len(numStr)%2 == 0 {
		leftHalf, _ := strconv.ParseInt(strings.Join(strings.Split(numStr, "")[:len(numStr)/2], ""), 10, 64)
		rightHalf, _ := strconv.ParseInt(strings.Join(strings.Split(numStr, "")[len(numStr)/2:], ""), 10, 64)
		out := make([]int, 0)
		out = append(out, Blink(int(leftHalf), blinkNum+1)...)
		out = append(out, Blink(int(rightHalf), blinkNum+1)...)
		return out
	} else {
		return Blink(num*2024, blinkNum+1)
	}
}

func main() {
	fmt.Println("AOC Day 11")
	dat, err := os.ReadFile("inputs.txt")
	if err != nil {
		log.Fatal(err)
	}
	arr := strings.Split(string(dat), " ")
	nums := make([]int, 0)
	for _, ele := range arr {
		num, _ := strconv.ParseInt(ele, 10, 64)
		nums = append(nums, int(num))
	}

	out := make([]int, 0)
	for _, ele := range nums {
		out = append(out, Blink(ele, 1)...)
	}

	fmt.Println("Count:", len(out))
}
