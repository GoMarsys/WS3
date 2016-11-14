package challange_interface

import (
	"reflect"
	"testing"
)

func TestExpensiveReturnFix(t *testing.T) {

	result := expensiveReturn()

	if reflect.TypeOf(result).Kind() != reflect.Slice {
		t.Errorf("Invalid type returned!\n")
	}

}
