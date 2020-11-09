package pool

import (
)

/* Stream Slice */

func FindAndRemove(array []Stream, element Stream) ([]Stream, bool) {
	i, exists := Find(array, element)
	if !exists {
		return nil, false
	}
	return Remove(array, i), true
}

// Returns an array with the i-th element removed
func Remove(array []Stream, i int) []Stream {
	return append(array[:i], array[i+1:]...)
}

// Returns the index of the element in the array
func Find(array []Stream, element Stream) (int, bool) {
	for i := 0; i < len(array); i++ {
		if array[i] == element {
			return i, true
		}
	}
	return -1, false
}
