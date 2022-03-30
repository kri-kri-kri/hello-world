package main

import (
	"fmt"
)

func main() {
	var text string
	var width int
	var res string
	fmt.Scanf("%s %d", &text, &width)
	bytes := []byte(text)
	if len(bytes) > width {
		res = string(bytes[:width]) + "..."
	} else {
		res = string(bytes)
	}
	fmt.Println(res)
}
