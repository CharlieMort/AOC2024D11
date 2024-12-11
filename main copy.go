package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1() {
	fmt.Println("AOC Day 9")
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
	//fmt.Println(nums)
	for i := 0; i < 25; i++ {
		fmt.Println("Iter", i)
		newNums := make([]int, 0)
		for _, num := range nums {
			numStr := fmt.Sprintf("%d", num)
			if num == 0 {
				newNums = append(newNums, 1)
			} else if len(numStr)%2 == 0 {
				leftHalf, _ := strconv.ParseInt(strings.Join(strings.Split(numStr, "")[:len(numStr)/2], ""), 10, 64)
				rightHalf, _ := strconv.ParseInt(strings.Join(strings.Split(numStr, "")[len(numStr)/2:], ""), 10, 64)
				newNums = append(newNums, int(leftHalf), int(rightHalf))
			} else {
				newNums = append(newNums, num*2024)
			}
		}
		nums = newNums
	}
	//fmt.Println(nums)
	fmt.Println("Count:", len(nums))
}
