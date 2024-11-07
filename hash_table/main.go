package main

import "fmt"

type KeyValue struct {
	Key   string
	Value interface{}
}

type HashTable struct {
	Buckets    [][]KeyValue
	BucketSize int
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		Buckets:    make([][]KeyValue, size),
		BucketSize: size,
	}
}

func (ht *HashTable) Hash(key string) int {
	sum := 0
	for _, char := range key {
		sum += int(char)
	}

	return sum % ht.BucketSize
}

func (ht *HashTable) Put(key string, value interface{}) {
	index := ht.Hash(key)

	for i, kv := range ht.Buckets[index] {
		if kv.Key == key {
			ht.Buckets[index][i].Value = value
			return
		}
	}

	ht.Buckets[index] = append(ht.Buckets[index], KeyValue{Key: key, Value: value})
}

func (ht *HashTable) Get(key string) (interface{}, bool) {
	index := ht.Hash(key)

	for _, kv := range ht.Buckets[index] {
		if kv.Key == key {
			return kv.Value, true
		}
	}

	return nil, false
}

func (ht *HashTable) Delete(key string) bool {
	index := ht.Hash(key)

	for i, kv := range ht.Buckets[index] {
		if kv.Key == key {
			ht.Buckets[index] = append(ht.Buckets[index][:i], ht.Buckets[index][i+1:]...)
			return true
		}
	}

	return false
}

func main() {
	ht := NewHashTable(5)
	ht.Put("name", "fawas")
	ht.Put("age", 25)
	if name, found := ht.Get("name"); found {
		fmt.Println(name)
	}
	if age, found := ht.Get("age"); found {
		fmt.Println(age)
	}
	ht.Delete("fawas")
	if name, found := ht.Get("name"); found {
		fmt.Println(name)
	}

}
