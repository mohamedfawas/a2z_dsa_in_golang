package main

import "fmt"

func reverseArray(array []int, array_size int) []int {
	for i := 0; i < array_size/2; i++ {
		var temp int
		temp = array[i]
		array[i] = array[array_size-1-i]
		array[array_size-1-i] = temp
	}
	return array
}

func main() {
	fmt.Println("give the size of the array : ")
	var size int
	fmt.Scan(&size)

	var arr []int = make([]int, size)
	for i := range arr {
		fmt.Printf("give the element to add at the position %d : ", i)
		fmt.Scan(&arr[i])
	}

	fmt.Println(arr) // array creation done
	rev_arr := reverseArray(arr, size)
	fmt.Println(rev_arr)

}
