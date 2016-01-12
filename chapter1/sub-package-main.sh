$ cat main.go
package main

import (
    "fmt"
    "example/math"  // just "math" would not work, it would import std package math
)

func main() {
    fmt.Println("Hello world! My lucky number is", math.Mul2(privateHelperFunc()))
}

$ go run *.go
Hello world! My lucky number is 50