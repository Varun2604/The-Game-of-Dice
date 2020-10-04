package main

import "testing"

func TestEqualsIntArr(t *testing.T) {
	arr1 := []int{8, 2, 4, -1, 0, 1}
	arr2 := []int{8, 2, 4, -1, 0, 1}
	if !EqualsIntArr(arr1, arr2) {
		t.Error("EqualsIntArr malfunctioning - improper comparison of two equal arrays")
	}
	arr2 = append(arr2, 5)
	if EqualsIntArr(arr1, arr2) {
		t.Error("EqualsIntArr malfunctioning - improper comparison of two unequal length arrays")
	}
	arr2 = arr2[1:]
	if EqualsIntArr(arr1, arr2) {
		t.Error("EqualsIntArr malfunctioning - improper comparison of two unequal arrays")
	}
}
func TestSortIntDesc(t *testing.T) {
	arr := []int{8, 2, 4, -1, 0, 1}
	arr = SortIntDesc(arr)
	if !EqualsIntArr(arr, []int{8, 4, 2, 1, 0, -1}) {
		t.Error("SortIntDesc malfunctioning")
	}
}

func TestContainsStr(t *testing.T) {
	arr := []string{"a", "b", "c"}
	if !ContainsStr(arr, "a") {
		t.Error("ContainsStr malfunctioning - contains an ele but returns false")
	}
	if ContainsStr(arr, "d") {
		t.Error("ContainsStr malfunctioning - dose not contain an ele but returns true")
	}
}

func TestContainsInt(t *testing.T) {
	arr := []int{0, -1, 1, 8, 10000, -58}
	if !ContainsInt(arr, 1) {
		t.Error("ContainsInt malfunctioning - contains an ele but returns false")
	}
	if !ContainsInt(arr, -58) {
		t.Error("ContainsInt malfunctioning - contains an ele but returns false")
	}
	if !ContainsInt(arr, 0) {
		t.Error("ContainsInt malfunctioning - contains an ele but returns false")
	}
	if ContainsInt(arr, 20) {
		t.Error("ContainsStr malfunctioning - dose not contain an ele but returns true")
	}
}

func TestRoll(t *testing.T) {
	min := 1
	max := 6
	test := func(iteration int) {
		r := Roll(min, max)
		if r < 1 || r > 6 {
			t.Errorf("Roll malfunctions at %dth iteration", iteration)
		}
	}
	for i := 0; i < 100; i++ {
		test(i)
	}
	min = 6
	max = 5
	r := Roll(min, max)
	if r != -1 {
		t.Error("Roll malfunctions when min > max")
	}
	min = -1
	r = Roll(min, max)
	if r != -1 {
		t.Error("Roll malfunctions when min < 0")
	}
	min = 6
	max = -1
	r = Roll(min, max)
	if r != -1 {
		t.Error("Roll malfunctions when max < 0")
	}
}
