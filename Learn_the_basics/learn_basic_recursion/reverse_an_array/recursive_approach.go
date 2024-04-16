package main

import "fmt"

func reverseArray(array []int, start_index, end_index int) []int {
	if start_index >= end_index {
		return array
	}
	//swap
	var temp int = array[start_index]
	array[start_index] = array[end_index]
	array[end_index] = temp
	//recursive fn call
	reverseArray(array, start_index+1, end_index-1)
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
	rev_arr := reverseArray(arr, 0, size-1)
	fmt.Println("The reversed array is : ", rev_arr)

}
