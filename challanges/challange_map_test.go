package challange_map

import "testing"

func TestMapCreation(t *testing.T) {
	obj := validMapCreation()
	if obj["one"] != 1 {
		t.Errorf("Expected to return map object has a key: 'one' with value: 1\n")
	}
}
