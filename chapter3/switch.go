package main

import (
	"fmt"
	"regexp"
)

var (
	reg0 = regexp.MustCompile(`^[0-9]+$`)
	reg1 = regexp.MustCompile(`^[0-9]+x$`)
)

func main() {
	var str string
	fmt.Scanf("%s", &str)
	switch {
	case reg0.MatchString(str):
		fmt.Println("case 0")
		fallthrough
	case reg1.MatchString(str):
		fmt.Println("case 1")
	default:
		fmt.Println("default")
	}
}
