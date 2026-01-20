package main

/**
 *	求和
 */
func Sum[T Number](vals ...T) T {
	var res T
	for _, val := range vals {
		res += val
	}
	return res
}

// 找最大值
func getMax[T Number](vals ...T) T {
	var maxNum T = 0
	for _, val := range vals {
		if val > maxNum {
			maxNum = val
		}
	}
	return maxNum
}

// 找最小值
func getMin[T Number](vals ...T) T {
	var minNum T = 0
	for _, val := range vals {
		if val < minNum {
			minNum = val
		}
	}
	return minNum
}

type Number interface {
	~int | int64 | float32 | float64
}

type Integer int
