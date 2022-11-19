package binary_search

/* 69. x 的平方根
给你一个非负整数 x ，计算并返回 x 的 算术平方根 。
由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。
注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。
*/

func MySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	left, right := 1, x
	middle := 0
	for left <= right {
		middle = left + (right-left)/2
		tempNum := middle * middle
		if tempNum == x {
			return middle
		} else if tempNum > x {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return right
}
