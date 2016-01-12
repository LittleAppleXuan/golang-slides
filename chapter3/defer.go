package main

import "log"

func main() {

	for i := 0; i < 10; i++ {
		defer log.Println(i)
	}
	//	myDefer()
}

/* func myDefer() func() {
	defer log.Println("log")
	if a := 1; a == 1 {
		return func() { fmt.Println("1") }()
	}
	return func() { fmt.Println("end") }()
}*/
