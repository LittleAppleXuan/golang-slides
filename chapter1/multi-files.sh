$ cd $GOPATH/src/example
$ cat helper.go
package main

func privateHelperFunc() int {
    return 25
}

# 在另一个文件的函数中直接调用上面定义的函数
$ tail -3 main.go
func main() {
    fmt.Println("Hello world! My lucky number is", privateHelperFunc())
}