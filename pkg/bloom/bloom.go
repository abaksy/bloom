package bloom

type BloomFilter interface {
	// Add a new element into the Bloom Filter data structure
	Add(item interface{}) error

	// Return a Boolean value based on whether the input argument is
	// present within the Bloom Filter data structure or not
	//
	// NOTE: returning true indicates that the value "may be" present in the
	// data structure, i.e. the value may not exist even if the function returns true
	Contains(item interface{}) bool

	// Clear all the bits in the Bloom filter, i.e. set them to 0
	Clear() error

	// Returns the number of elements in the Bloom Filter
	Size() uint
}
