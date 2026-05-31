package main

import (
	"cmp"
	"fmt"
)

func BinarySearch[T cmp.Ordered](xs []T, target T) (int, bool) {
	if len(xs) == 0 {
		return 0, false
	}
	if len(xs) == 1 {
		if xs[0] == target {
			return 0, true
		}
		return 0, false
	}

	lo := 0
	hi := len(xs) - 1

	for lo <= hi {
		mid := lo + (hi-lo)/2
		if target == xs[mid] {
			return mid, true
		}
		if target < xs[mid] {
			hi = mid - 1
			continue
		}
		lo = mid + 1
	}

	return 0, false
}

func main() {
	idx, ok := BinarySearch(smallInts, 13)
	fmt.Println("small ints", idx, ok)
	idx, ok = BinarySearch(largeInts, 54)
	fmt.Println("large ints", idx, ok)
	idx, ok = BinarySearch(largeInts, 53)
	fmt.Println("large ints fail", idx, ok)
	idx, ok = BinarySearch(words, "kiwi")
	fmt.Println("strings", idx, ok)
	
	temp := []int{4, 5, 3, 9}
	idx, ok = BinarySearch(temp, 3)
	fmt.Println("temp", idx, ok)


}
