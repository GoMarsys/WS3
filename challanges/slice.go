package main

import "fmt"

func main() {

	a := []int{9, 8, 7, 6, 5}
	b := a[0:2]
	b[1] = 2

	fmt.Println(a)

	if a[1] == 2 {
		panic("oh no!")
	}

}
