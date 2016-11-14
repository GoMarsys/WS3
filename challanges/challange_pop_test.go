package challange_pop

import "testing"

func TestPopWithAppend(t *testing.T) {

	actually := popWithAppend(initialValue())
	expected := []int{5, 10, 15}

	if !testEq(actually, expected) {
		t.Errorf("Expected to %v be eq with %v", actually, expected)
	}

}

func TestPopWithCopy(t *testing.T) {

	actually := popWithCopy(initialValue())
	expected := []int{5, 10, 15}

	if !testEq(actually, expected) {
		t.Errorf("Expected to %v be eq with %v", actually, expected)
	}

}

func TestPopWithSlicing(t *testing.T) {

	actually := popWithSlicing(initialValue())
	expected := []int{5, 10, 15}

	if !testEq(actually, expected) {
		t.Errorf("Expected to %v be eq with %v", actually, expected)
	}

}

func initialValue() []int {
	return []int{5, 10, 15, 20}
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
