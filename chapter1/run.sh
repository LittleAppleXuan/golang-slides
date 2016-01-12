$ cd $GOPATH/src/example
$ ls
main.go

$ cat main.go
package main

import "fmt"

func main() {
    fmt.Println("Hello world!")
}

$ go run main.go
Hello world!