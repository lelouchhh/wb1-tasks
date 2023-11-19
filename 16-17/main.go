package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{4, 2, 9, 1, 5, 6, 8, 3, 7}

	sort.Ints(arr)
	fmt.Println("Отсортированный массив:", arr)

	target := 5

	index := sort.SearchInts(arr, target)

	if index < len(arr) && arr[index] == target {
		fmt.Printf("Элемент %d найден в позиции %d\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден в массиве\n", target)
	}
}
