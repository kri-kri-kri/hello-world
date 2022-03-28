package main

import (
	"fmt"
)

func main() {
	var source string
	var times int
	fmt.Scan(&source, &times)

	var result string
	var res string
	for i := 0; i < times; i++ {
		result += source
	}
	i := 1
	for i <= times {
    		res += source
    		i = i + 1
	}

	fmt.Println(result)
	fmt.Println(res)
}