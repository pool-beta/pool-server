package utils

import (
	. "github.com/pool-beta/pool-server/user/types"
)

// UserID Slice Utils

func FindAndRemove(array []UserID, element UserID) ([]UserID, bool) {
	i, exists := Find(array, element)
	if !exists {
		return nil, false
	}
	return Remove(array, i), true
}

// Returns an array with the i-th element removed
func Remove(array []UserID, i int) []UserID {
	return append(array[:i], array[i+1:]...)
}

// Returns the index of the element in the array
func Find(array []UserID, element UserID) (int, bool) {
	for i := 0; i < len(array); i++ {
		if array[i] == element {
			return i, true
		}
	}
	return -1, false
}