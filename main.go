package main

import "fmt"

func Solution(A []int, k int) []int {

	var res []int = A
	for i := 0; i < k; i++ {
		res = Rot(res)
		fmt.Println("Итерация ", i+1, " - ", res)
	}
	return res

}

func Rot(arr []int) []int {
	l := len(arr)
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
	N = 4
	ArrRot = Solution(ArrRot, N)
	fmt.Println("Результат", ArrRot)
}
