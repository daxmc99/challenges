package main

import (
	"fmt"
	"sort"
)

func main() {
	ans := []int{5, 3, 0, 2, 6, 2, 0, 7, 2, 5}
	fmt.Println(ans)

	fmt.Println(removeZeros([]int{0, 0, 0, 2, 1, 0}))
	fmt.Println(descendingSort(ans))
	fmt.Println("Desc: ", descendingSort([]int{5, 1, 3, 4, 2}))
	//fmt.Println(lengthCheck(7, ans))
	fmt.Println(frontElim(4, []int{5, 4, 3, 2, 1}))
	fmt.Println(popOffFirst([]int{2, 3, 4, 5}))
	//fmt.Println("Havel-Hakimi", havelHakimi(ans))
	fmt.Println([]int{4, 2, 0, 1, 5, 0}, havelHakimi([]int{4, 2, 0, 1, 5, 0}))
}

func removeZeros(arr []int) []int {
	i := 0
	for _, x := range arr {
		if x != 0 {
			//rewrite slice in place
			arr[i] = x
			i++
		}
	}

	return arr[:i]
}

// remove with bounds checking
func removeElem(i int, arr []int) []int {
	if i > len(arr) {
		i = i - 1
	}
	arr[i] = arr[len(arr)-1]
	return arr[:len(arr)-1]
}

func descendingSort(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	return arr
}

func lengthIsGreater(lenth int, arr []int) bool {
	return lenth > len(arr)
}

func frontElim(count int, arr []int) []int {
	arr = descendingSort(arr)
	for i := 0; i < count; i++ {
		arr[i] = arr[i] - 1
	}
	return arr
}

func popOffFirst(arr []int) []int {
	return append(arr[:0], arr[0+1:]...)
}

func havelHakimi(arr []int) bool {
	for {
		arr = removeZeros(arr)
		if len(arr) == 0 {
			return true
		}
		arr = descendingSort(arr)
		N := arr[0]
		//order matters
		arr = popOffFirst(arr)
		if N > len(arr) {
			return false
		}
		arr = frontElim(N, arr)
		//continue
	}
}
