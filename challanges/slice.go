package main

func main() {

	a := make([]int, 3)
	b := a[0:2]
	b[1] = 2

	for i := 0; i < 10; i++ {
		a = append(a, 5)
	}

	if a[1] != 5 {
		panic("oh no!")
	}

}
