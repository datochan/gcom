package utils

/**
 * 生成指定的索引数组
 */
func GenerateIndex(start, step, max int) []int {
	var result []int

	if step > 0 {
		for idx:=start; idx < max; idx+=step { result = append(result, idx) }
		return result
	}

	for idx:=start; idx > max; idx+=step { result = append(result, idx) }

	return result
}

// 从 slice 中查找指定的数字
// 返回相应的索引值，找不到返回-1
func FindInIntegerSlice(n int, nAry []int) int {
	for i, e := range nAry {
		if e == n { return i }
	}

	return -1
}
