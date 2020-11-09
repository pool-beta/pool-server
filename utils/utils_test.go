package utils

import (
	"testing"

	. "github.com/pool-beta/pool-server/types"
)

func TestSliceFindAndRemove(t *testing.T) {
	invalid := UserID(420)

	testSlice := []UserID{
		0,
		1,
		59,
		3,
		4,
		5,
		6,
		7,
	}
	
	// Test Find
	for i, e := range testSlice {
		findIndex, ok := Find(testSlice, e)
		if !ok {
			t.Errorf("Expected but not found -- index: %v; element: %v\n", i, e)
		}
		if (i != findIndex) {
			t.Errorf("Found but conflicting index: -- expected index: %v; actual index: %v; element: %v\n", i, findIndex, e)
		}
	}

	// Try Finding Invalid
	_, ok := Find(testSlice, invalid)
	if ok {
		t.Errorf("Expected Invalid but found -- invalid element: %v\n", invalid)
	}

	// Test Remove
	for i, _ := range testSlice {
		returnSlice := Remove(testSlice, i)
		expectedSlice := append(testSlice[:i], testSlice[i+1:]...)
		if !isSliceEqual(returnSlice, expectedSlice) {
			t.Errorf("Invalid results from Remove -- expected slice: %v; actual slice: %v\n", expectedSlice, returnSlice)
		}
	}
}

func isSliceEqual(array1 []UserID, array2 []UserID) bool {
	if len(array1) != len(array2) {
		return false
	}
	
	for i, e := range array1 {
		if e != array2[i] {
			return false
		}
	}
	return true
}