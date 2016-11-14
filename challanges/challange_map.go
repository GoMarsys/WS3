// $ go test challanges/challange_map*

// fix the code so it won't panic with "assignment to entry in nil map"
// and set int 1 under the string key "one"

package challange_map

func validMapCreation() map[string]int {
	var m map[string]int
	return m
}
