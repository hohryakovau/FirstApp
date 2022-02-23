package main

import "fmt"

func Solution(A []int, k int) []int {

	var test []int = A
	for i := 0; i < k; i++ {
		test = Rot(test)
		fmt.Println("Итерация ", i+1, " - ", test)
	}
	return test

}

func Rot(arr []int) []int {
	l := len(arr)
	fmt.Println("Длинна массива", l)
	var b []int = make([]int, l)
	for i := 0; i < l; i++ {
		if i == (l - 1) {
			b[0] = arr[i]
		} else {
			b[i+1] = arr[i]
		}
	}
	return b
}

func main() {

	ArrRot := []int{3, 8, 9, 7, 6}

	var N int
	N = 3
	ArrRot = Solution(ArrRot, N)
	fmt.Println("Результат", ArrRot)
}
