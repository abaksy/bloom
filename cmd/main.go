package main

import (
	"fmt"

	bloom "github.com/abaksy/bloom/pkg/bloom"
)

func main() {
	bf, err := bloom.NewStandardBloomFilter(10, 0.01)
	if err != nil {
		fmt.Println(err)
		return
	}

	bf.Add("hello123")
	fmt.Println("Added hello123!")
	bf.Add("hello456")
	fmt.Println("Added hello456!")

	ans := bf.Contains("hello123")
	if ans {
		fmt.Println("hello123 may be present!")
	} else {
		fmt.Println("hello123 not present")
	}

	ans = bf.Contains("hello789")
	if ans {
		fmt.Println("hello789 may be present!")
	} else {
		fmt.Println("hello789 not present")
	}
}
