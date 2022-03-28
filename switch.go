package main

import (
	"fmt"
)

func main() {
    var code string
    fmt.Scan(&code)
    var lang string
    switch code {
    case "en":
        lang := "English"
    case "fr":
        lang := "French"
    case "ru":
        lang := "Russian"
    case "rus":
        lang := "Russian"
}
    fmt.Println(lang)
}