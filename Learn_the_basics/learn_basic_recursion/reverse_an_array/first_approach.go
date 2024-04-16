package main

import "fmt"

func reverseArray(array []int, array_size int) []int {
	reverse_array := make([]int, array_size)
	for i := 0; i < array_size; i++ {
		reverse_array[i] = array[array_size-1-i]
	}
	return reverse_array
}

func main() {
	var size int
	fmt.Println("give the size of the input array : ")
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
