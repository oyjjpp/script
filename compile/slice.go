package compile

func newSlice() []int {
	array := [3]int{1, 2, 3}
	slice := array[0:1]
	return slice
}
