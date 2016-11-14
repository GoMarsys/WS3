package challange_slice

import (
	"testing"
)

func TestSlicingReturnValueOK(t *testing.T) {

	input := []int{9, 8, 7, 6, 5}
	result := slicingWithNewUnderlyingArray(input)

	expected := []int{8, 7, 6}
	if !testEq(expected, result) {
		t.Errorf("expected array not sliced well (1..3), %v != %v\n", expected, result)
	}

}

func TestSlicingAndNewUnderlyingArrayIsDifferent(t *testing.T) {

	input := []int{0, 2, 4, 6, 8, 10}
	result := slicingWithNewUnderlyingArray(input)
	input[2] = 42

	expected := []int{2, 4, 6}
	if !testEq(result, expected) {
		t.Logf("expected array not sliced well (1..3), %v != %v\n", expected, result)
		t.Errorf("Underlying array is linked with the input!\n")
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
