package main

import "fmt"

func main() {
	array := []int{2, 5, 6, 8, 3, 1, 4, 9, 10, 7}
	arrLength := len(array)
	sort(array, arrLength)
	print(array)
}

func sort(array []int, length int) []int {
	var temp, smallest int
	for i := 0; i < length-1; i++ {
		smallest = i
		for j := 0; j < length; j++ {
			if array[j] < array[smallest] {
				smallest = j
			}
		}

		temp = array[smallest]
		array[smallest] = array[i]
		array[i] = temp
	}

	return array
}

func print(v []int) {
	for _, val := range v {
		fmt.Println(val)
	}
}
