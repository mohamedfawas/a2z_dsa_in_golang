package main

import "fmt"

func main() {
	// Create a hash table to store student names and their grades
	studentGrades := make(map[string]int)

	// Insert key-value pairs
	studentGrades["Fawas"] = 99
	studentGrades["Alice"] = 90
	studentGrades["Bob"] = 85
	studentGrades["Charlie"] = 88

	// Return as a map
	fmt.Println(studentGrades)

	// the stored data will be returned in any order
	for key, value := range studentGrades {
		fmt.Printf("name : %v , grade : %v ", key, value)
		fmt.Println("")
	}

	// Retrieve a value
	fmt.Println("Alice's grade:", studentGrades["Alice"])

	// Update a value
	studentGrades["Alice"] = 92

	// Retrieve a value
	fmt.Println("Alice's grade after updation :", studentGrades["Alice"])

	// Delete a key value pair
	delete(studentGrades, "Bob")

	// Check if a key exists
	grade, exists := studentGrades["Bob"]
	if exists {
		fmt.Println("Bob's grade : ", grade)
	} else {
		fmt.Println("Bob's grade not found")
	}
}
