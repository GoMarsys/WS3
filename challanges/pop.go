// 1. your task is to alter the actually array content by slicing the slice
// 2. your task is to succeed the same with make + append

package main

func main() {

	actually := []int{5, 10, 15, 20}
	expected := []int{5, 10, 15}

	if !testEq(actually, expected) {
		panic("aww... not equal...")
	}

}

func testEq(a, b []int) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
