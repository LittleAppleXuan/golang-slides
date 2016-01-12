package main

import "fmt"

func main() {

	arrayI := [...]int{1, 2, 3}
	for i := 0; i < len(arrayI); i++ {
		fmt.Println(arrayI[i])
	}
	sliceI := []int{4, 5} //slice
	for i, _ := range arrayI {
		fmt.Println(i)
	}
	fmt.Println("After range []arr", arrayI)

	for i, v := range &arrayI {
		v *= 100
		fmt.Println(i, v)

	}
	fmt.Println("After range *[]arr", arrayI)
	for i, v := range sliceI {
		v *= 10
		fmt.Println(i, v)
	}
	fmt.Println("After range []slice", sliceI)

}
