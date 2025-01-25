# bloom
Implements a Bloom Filter in Go

## Standard Bloom Filter
Simplest Bloom filter that uses a Murmur3 hash function to get the hash index to place the data in. 
Calculate the number of 

$
k = \frac{m}{n} ln 2
$