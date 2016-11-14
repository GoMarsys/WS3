package challange_slice

func slicingWithNewUnderlyingArray(input []int) []int {
	return append(make([]int, 0), input[1:4]...)
}
