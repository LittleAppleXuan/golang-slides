package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var src, dst *os.File
	src, _ = os.Open("defer.go")
	defer func() {
		src.Close()
		dst.Close()
	}()

	dst, _ = os.OpenFile("defer2.go", os.O_WRONLY|os.O_CREATE, 0644)
	//_, err := io.Copy(dst, src)
	//fmt.Println(err)
	buf := make([]byte, 1024)
	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, _ := dst.Write(buf[0:nr])
			fmt.Println("write size:", nw)
		}
		if er == io.EOF {
			break
		}
	}

}
