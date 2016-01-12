package main

import "fmt"

func main() {
	var ss = []string{"We", "use", "goto", "instead", "of", "for"}
	//myGoto(ss)

	//tGoto(ss)
	if len(ss) == 6 {
		fmt.Println("yes")
	} else {
		goto Another
	}
Another:
	fmt.Println("another")
}
func tGoto(ss []string) {
	fmt.Println(ss[0] + " ")
	if len(ss) > 1 {
		tGoto(ss[1:])
	}
}

func myGoto(ss []string) {
	var i = 0
JustDoit:
	fmt.Print(ss[i] + " ")
	if i < len(ss)-1 {
		i++
		goto JustDoit
	}
}
