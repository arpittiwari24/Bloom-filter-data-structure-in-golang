package main

import (
	"fmt"
	"hash/fnv"
)

type BloomFilter struct {
	bitArray         []bool
	numHashFunctions int
}

func NewBloomFilter(size int, numHashFunctions int) *BloomFilter {
	return &BloomFilter{
		bitArray:         make([]bool, size),
		numHashFunctions: numHashFunctions,
	}
}

func (bf *BloomFilter) Add(value string) {
	for i := 0; i < bf.numHashFunctions; i++ {
		hash := fnv.New32()
		hash.Write([]byte(value))
		index := int(hash.Sum32()) % len(bf.bitArray)
		bf.bitArray[index] = true
	}
}

func (bf *BloomFilter) Contains(value string) bool {
	for i := 0; i < bf.numHashFunctions; i++ {
		hash := fnv.New32()
		hash.Write([]byte(value))
		index := int(hash.Sum32()) % len(bf.bitArray)
		if !bf.bitArray[index] {
			return false
		}
	}
	return true
}

func main() {
	bf := NewBloomFilter(100, 3)
	bf.Add("hello")
	bf.Add("world")
	fmt.Println(bf.Contains("hello")) // Output: true
	fmt.Println(bf.Contains("world")) // Output: true
	fmt.Println(bf.Contains("food"))  // Output: false
}
