package questions

import "fmt"

/*
Write a caching system where the last n computed results are stored in a map for quick lookup.

Store the results of expensive computations so that if the same input comes again, the program doesn't recompute it — it simply retrieves the result from the cache.
*/

// Cache struct to hold the map and the order of keys.
type Cache struct {
	size  int
	store map[int]int
	order []int
}

// NewCache creates a new Cache with a given size.
func NewCache(size int) *Cache {
	return &Cache{
		size:  size,
		store: make(map[int]int),
		order: make([]int, 0, size),
	}
}

// Get retrieves the cached result or computes and stores it.
func (c *Cache) Get(input int) int {
	// Check if the result is in cache.
	if result, found := c.store[input]; found {
		fmt.Println("Cache Hit for input:", input)
		return result
	}

	// Simulate expensive computation: let's square the input.
	result := input * input

	// If the cache is full, remove the oldest entry.
	if len(c.order) >= c.size {
		oldest := c.order[0]
		delete(c.store, oldest)
		c.order = c.order[1:]
	}

	// Store the new result.
	c.store[input] = result
	c.order = append(c.order, input)

	fmt.Println("Cache Miss for input:", input, "Computed and Stored:", result)
	return result
}

func Cache_using_Map() {
	/*
		size = 3
		store = map[int]int[]
		order = []int len: 0, cap: 3, []
	*/
	cache := NewCache(3) // Create a cache with a max size of 3.

	// Test inputs.
	inputs := []int{2, 4, 2, 5, 4, 6, 2}

	for _, input := range inputs {
		result := cache.Get(input)
		fmt.Println("Result for", input, "is", result)
	}
}

/*
Initialization:




Create a Cache structure with:
			A map to store computed results (store).
			A slice to track the order of inserted inputs (order).
			A fixed size limit for the cache.




For Each Input:
		Cache Lookup:
				Check if the input exists in the map.
				If yes:
				→ Return the cached result. (Cache Hit)
		Cache Miss Handling:
				If no:
				Compute the result. (In this case: input * input).
		Cache Size Check:
				If the cache is full (len(order) >= size):
					Remove the oldest entry:
					Take the first element from the order slice.
					Delete it from the map.
					Remove it from the order slice.
		Cache Store:
				Store the new computed result in the map.
				Append the input to the order slice.
		Return the Result:
				Return the computed or cached result to the caller.
*/
