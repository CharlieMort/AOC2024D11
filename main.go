package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var cache map[string]int

func Blink(num int, blinkNum int) int {
	val, ok := cache[fmt.Sprintf("%d/%d", num, blinkNum)]
	if ok {
		return val
	}
	if blinkNum == 0 {
		return 1
	}
	numStr := fmt.Sprintf("%d", num)
	if num == 0 {
		return Blink(1, blinkNum-1)
	} else if len(numStr)%2 == 0 {
		leftHalf, _ := strconv.ParseInt(strings.Join(strings.Split(numStr, "")[:len(numStr)/2], ""), 10, 64)
		rightHalf, _ := strconv.ParseInt(strings.Join(strings.Split(numStr, "")[len(numStr)/2:], ""), 10, 64)
		cache[fmt.Sprintf("%d/%d", num, blinkNum)] = Blink(int(leftHalf), blinkNum-1) + Blink(int(rightHalf), blinkNum-1)
		return Blink(int(leftHalf), blinkNum-1) + Blink(int(rightHalf), blinkNum-1)
	} else {
		return Blink(num*2024, blinkNum-1)
	}
}

func main() {
	fmt.Println("AOC Day ")
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

	cache = make(map[string]int)
	count := 0
	for _, ele := range nums {
		count += Blink(ele, 75)
	}
	fmt.Println("Count:", count)
}
