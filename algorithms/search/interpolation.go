package search

// The Interpolation Search is an improvement over Binary Search for instances, where the values in a sorted array are uniformly distributed.
// Binary Search always goes to middle element to check.
// On the other hand interpolation search may go to different locations according the value of needle being searched.
// Here is the source code of the Go program to search element in an integer array using Interpolation search algorithm.
// The output shows the position of element in array.

// Interpolation search algorythm
func Interpolation(haystack []int, needle int) int {
	min, max := haystack[0], haystack[len(haystack)-1]
	low, high := 0, len(haystack)-1

	for {
		if needle < min {
			return low
		}

		if needle > max {
			return high + 1
		}

		// make a guess of the location
		var guess int
		if high == low {
			guess = high
		} else {
			size := high - low
			offset := int(float64(size-1) * (float64(needle-min) / float64(max-min)))
			guess = low + offset
		}

		// maybe we found it?
		if haystack[guess] == needle {
			// scan backwards for start of value range
			for guess > 0 && haystack[guess-1] == needle {
				guess--
			}
			return guess
		}

		// if we guessed to high, guess lower or vice versa
		if haystack[guess] > needle {
			high = guess - 1
			max = haystack[high]
		} else {
			low = guess + 1
			min = haystack[low]
		}
	}
}
