package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func filter(predicate func(int) bool, iterable []int) []int {
	var evens []int
	for _, num := range iterable {
		if predicate(num) {
			evens = append(evens, num)
		}
	}
	return evens
}

func main() {
	src := readInput()
	res := filter(func(num int) bool { return num%2 == 0 }, src)
	fmt.Println(res)
}

func readInput() []int {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}
	return nums
}
