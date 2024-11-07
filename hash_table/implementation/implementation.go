package hashtableimplementation

import "fmt"

// KeyValue represents a key-value pair stored in our hash table
type KeyValue struct {
	Key   string
	Value interface{}
}

// HashTable represents our hash table structure
type HashTable struct {
	// We'll use an array of slices to handle collisions through chaining
	Buckets    [][]KeyValue
	BucketSize int
}

// NewHashTable creates and initializes a new HashTable
func NewHashTable(size int) *HashTable {
	// Initialize the hash table with empty buckets
	return &HashTable{
		Buckets:    make([][]KeyValue, size),
		BucketSize: size,
	}
}

// hash generates a simple hash code for a string key
func (ht *HashTable) hash(key string) int {
	// Simple hash function: sum of ASCII values of characters
	sum := 0
	for _, ch := range key {
		sum += int(ch)
	}
	// Use modulo to ensure the hash falls within our bucket size
	return sum % ht.BucketSize
}

// Put adds or updates a key-value pair in the hash table
func (ht *HashTable) Put(key string, value interface{}) {
	// Calculate which bucket this key belongs in
	index := ht.hash(key)

	// Check if the key already exists in the bucket
	for i, kv := range ht.Buckets[index] {
		if kv.Key == key {
			// Key found, update its value
			ht.Buckets[index][i].Value = value
			return
		}
	}

	// Key not found, append new key-value pair to bucket
	ht.Buckets[index] = append(ht.Buckets[index], KeyValue{key, value})
}

// Get retrieves a value from the hash table by its key
func (ht *HashTable) Get(key string) (interface{}, bool) {
	// Calculate which bucket to look in
	index := ht.hash(key)

	// Search through the bucket for our key
	for _, kv := range ht.Buckets[index] {
		if kv.Key == key {
			return kv.Value, true
		}
	}

	// Key not found
	return nil, false
}

// Delete removes a key-value pair from the hash table
func (ht *HashTable) Delete(key string) bool {
	// Calculate which bucket to look in
	index := ht.hash(key)

	// Search for the key in the bucket
	for i, kv := range ht.Buckets[index] {
		if kv.Key == key {
			// Remove the element by slicing it out
			ht.Buckets[index] = append(ht.Buckets[index][:i], ht.Buckets[index][i+1:]...)
			return true
		}
	}

	return false
}

func main() {
	// Create a new hash table with 10 buckets
	ht := NewHashTable(10)

	// Example usage
	ht.Put("name", "John")
	ht.Put("age", 30)
	ht.Put("city", "New York")

	// Retrieve and print values
	if name, found := ht.Get("name"); found {
		fmt.Printf("Name: %v\n", name)
	}

	if age, found := ht.Get("age"); found {
		fmt.Printf("Age: %v\n", age)
	}

	// Delete a key and try to retrieve it
	ht.Delete("city")
	if _, found := ht.Get("city"); !found {
		fmt.Println("City was deleted successfully")
	}
}
