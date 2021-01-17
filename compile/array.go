package compile

// outOfRange
// SSA 分析索引溢出
func outOfRange() int {
	arr := [3]int{1, 2, 3}
	i := 4
	elem := arr[i]
	return elem
}
