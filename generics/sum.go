package main

import "errors"

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
func getMax[T Number](vals ...T) (T, error) {
	if len(vals) == 0 {
		return vals[0], errors.New("no values given")
	}
	var maxNum T = vals[0]
	for i := 0; i < len(vals); i++ {
		if vals[i] > maxNum {
			maxNum = vals[i]
		}
	}
	return maxNum, nil
}

// 找最小值
func getMin[T Number](vals ...T) (T, error) {
	if len(vals) == 0 {
		return vals[0], errors.New("no values given")
	}
	var minNum T = vals[0]
	for i := 0; i < len(vals); i++ {
		if vals[i] < minNum {
			minNum = vals[i]
		}
	}
	return minNum, nil
}

// 在切片特定索引位置插入元素
func AddSlice[T any](slice []T, idx int, val T) ([]T, error) {
	if idx < 0 || idx > len(slice) {
		return nil, errors.New("下标出错")
	}

	res := make([]T, 0, len(slice)+1)
	// 添加原 slice中 0-idx索引元素数据到 res新切片中
	for i := 0; i < idx; i++ {
		res = append(res, slice[i])
	}
	// 使用 append 添加新元素，而不是直接索引赋值
	res = append(res, val)
	// 把原 slice里剩余元素都添加到 res中
	for i := idx; i < len(slice); i++ {
		res = append(res, slice[i])
	}
	return res, nil
}

type Number interface {
	~int | int64 | float32 | float64
}

type Integer int
