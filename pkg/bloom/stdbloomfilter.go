package bloom

import (
	"errors"
	"fmt"
	"math"

	"encoding/binary"

	"github.com/twmb/murmur3"
)

type StandardBloomFilter struct {
	Arr  []bool
	N    int     // Estimated total number of elements in the set
	M    int     // Length of bit array
	K    int     // Number of hash functions
	P    float64 // Desired false positive rate
	size int     // Number of elements in the data structure
}

// Get size of the bit array based on the desired number of elements
// and the desired false positive rate. We use the approximation for the FPR
// to calculate the size of the bit array.
func GetBitArraySize(N int, P float64) (int, error) {
	if N < 0 || (P <= 0 || P > 1) {
		return -1, errors.New("argument error while getting bit arr size, check N and P values")
	}

	m := -int((float64(N) * math.Log(P)) / (math.Pow(math.Ln2, 2)))
	return m, nil
}

// Get the number of hash functions based on the size of the Bloom filter and
// the bit array size. Use the same exponential relationship to derive this
func GetNumHashFunctions(N int, M int) (int, error) {
	if N <= 0 {
		return -1, errors.New("invalid size argument while creating bloom filter")
	}

	k := int((float64(M) * math.Log(2)) / float64(N))
	k = max(1, k)
	return k, nil
}

func NewStandardBloomFilter(N int, P float64) (*StandardBloomFilter, error) {
	if N <= 0 {
		return nil, errors.New("invalid size argument while creating bloom filter")
	}

	M, err := GetBitArraySize(N, P)
	if err != nil {
		return nil, err
	}

	K, err := GetNumHashFunctions(N, M)
	if err != nil {
		return nil, err
	}

	fmt.Printf("M = %d, K = %d\n", M, K)

	return &StandardBloomFilter{
		Arr:  make([]bool, M),
		N:    N,
		P:    P,
		M:    M,
		K:    K,
		size: 0,
	}, nil
}

func (b *StandardBloomFilter) GetHashIndex(value string, idx int) uint64 {
	numBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(numBytes, uint32(idx))

	hashFunc := murmur3.New64()
	hashFunc.Write([]byte(value))
	hashFunc.Write(numBytes)

	// Get the two 64-bit values
	hashValue := hashFunc.Sum64()

	return hashValue % uint64(b.M)
}

func (b *StandardBloomFilter) Add(ele interface{}) error {
	strRepr := ele.(string)

	if b.size >= b.N {
		return errors.New("bloom filter has reached its size limit")
	}

	for i := 0; i < b.K; i++ {
		// Convert index to bytes so that it can be used as input to Hash
		hashIndex := b.GetHashIndex(strRepr, i)
		b.Arr[hashIndex] = true
	}
	b.size++
	return nil
}

func (b *StandardBloomFilter) Contains(ele interface{}) bool {
	strRepr := ele.(string)
	contains := true

	for i := 0; i < b.K; i++ {
		// Convert index to bytes so that it can be
		// used to check the bit array index
		hashIndex := b.GetHashIndex(strRepr, i)
		contains = contains && b.Arr[hashIndex]
	}
	return contains
}

func (b *StandardBloomFilter) Clear() error {
	b.Arr = make([]bool, len(b.Arr))
	b.size = 0
	return nil
}
