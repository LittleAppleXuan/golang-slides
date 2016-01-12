package main

import (
	"fmt"
)

type dualError struct {
	Num     int
	problem string
}

func (e dualError) Error() string {
	return fmt.Sprintf("输入不正确，因为\"%d\"不是偶数", e.Num)

}
func requireDual(n int) (int, error) {
	if n%2 == 1 {
		return -1, dualError{Num: n}

	}
	return n, nil

}
func main() {
	if result, err := requireDual(7); err != nil {
		fmt.Println("错误：", err)

	} else {
		fmt.Println("正确：", result)

	}

}
