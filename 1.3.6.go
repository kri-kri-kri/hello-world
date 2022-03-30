package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var number int
	fmt.Scanf("%d", &number)
	str := strconv.Itoa(number)
	counter := make(map[int]int)
	for i := 0; i <= 9; i++ {
		num := strconv.Itoa(i)
		if strings.Count(str, num) != 0 {
			counter[i] = strings.Count(str, num)
		}
	}
	printCounter(counter)
}

func printCounter(counter map[int]int) {
	digits := make([]int, 0)
	for d := range counter {
		digits = append(digits, d)
	}
	sort.Ints(digits)
	for idx, digit := range digits {
		fmt.Printf("%d:%d", digit, counter[digit])
		if idx < len(digits)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Print("\n")
}
