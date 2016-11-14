// Pop the slice in the 3 most idiomatic way
//
// 1. with slicing
// 2. with append + ...
// 3. with copy
//
// initial actually := []int{5, 10, 15, 20}
// expected := []int{5, 10, 15}
//
package pop

func popWithSlicing(actually []int) []int {
	return actually[:len(actually)-1]
}

func popWithAppend(actually []int) []int {
	return append(make([]int, 0, cap(actually)), actually[:len(actually)-1]...)
}

func popWithCopy(actually []int) []int {
	lastIndex := len(actually) - 1
	returnSlice := make([]int, lastIndex)
	popedSlice := actually[:lastIndex]
	copy(returnSlice, popedSlice)
	return returnSlice
}
