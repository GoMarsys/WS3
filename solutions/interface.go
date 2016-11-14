package challange_interface

func expensiveReturn() []int {
	return []int{1, 2, 3}
}

func cheapReturn() interface{} {
	return 3
}
