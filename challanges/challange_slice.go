// your task is to return the subslice, in a way,
// that it does not point to the original Underlying array

package challange_slice

func slicingWithNewUnderlyingArray(input []int) []int {
	return input[1:4]
}
