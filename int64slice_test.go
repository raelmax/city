package main

import "testing"

func TestInt64Slice_Len(t *testing.T) {
	var testSlice Int64Slice

	testSlice = append(testSlice, 1)
	if len1 := testSlice.Len(); len1 != 1 {
		t.Error("Expected 3, got ", len1)
	}

	testSlice = append(testSlice, 2)
	if len2 := testSlice.Len(); len2 != 2 {
		t.Error("Expected 3, got ", len2)
	}
}

func TestInt64Slice_Less(t *testing.T) {
	var testSlice Int64Slice

	testSlice = append(testSlice, 1)
	testSlice = append(testSlice, 2)
	testSlice = append(testSlice, 3)

	if less1 := testSlice.Less(1, 2); !less1 {
		t.Error("Expected true, got ", less1)
	}

	if less2 := testSlice.Less(2, 0); less2 {
		t.Error("Expected false, got ", less2)
	}
}

func TestInt64Slice_Swap(t *testing.T) {
	var testSlice Int64Slice

	testSlice = append(testSlice, 1)
	testSlice = append(testSlice, 2)
	testSlice = append(testSlice, 3)

	testSlice.Swap(0, 2)

	if !(testSlice[0] == 3 && testSlice[2] == 1) {
		t.Error("Expected 3 and 1, got ", testSlice[0], " and ", testSlice[2])
	}
}