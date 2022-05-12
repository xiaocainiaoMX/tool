package slice

// 数组倒序
func RSort(list interface{}) []interface{} {
	arr := ToSlice(list)
	lens := len(arr)
	rlist := []interface{}{}
	for i := lens - 1; i >= 0; i-- {
		rlist = append(rlist, arr[i])
	}
	return rlist
}

