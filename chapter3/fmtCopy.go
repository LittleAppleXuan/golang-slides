package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	var fSrc, fDst *os.File
	var err error
	defer func() {
		fSrc.Close()
		fDst.Close()
	}()
	fSrc, err = os.Open("defer.go")
	if err != nil {
		fmt.Println("open failed")

	}

	fDst, err = os.Create("defer3.go")
	if err != nil {
		fmt.Println("create failed")

	}

	reader := bufio.NewReader(fSrc)

	for {
		line, err1 := reader.ReadString('\n')
		if err1 == io.EOF {
			break

		}
		if err1 != nil {
			fmt.Println("read failed")
		}

		fmt.Fprintf(fDst, "qiniu %s", line)

	}

}
