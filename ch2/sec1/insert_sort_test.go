package insertion_sort

import (
    "testing"
    "slices"
)

func TestSortCase1(t *testing.T) {
    arr := []int{5,2,4,6,1,3}
    expected := []int{1,2,3,4,5,6}

    actual := InsertionSort(arr)

    if !slices.Equal(actual, expected) {
        t.Fatalf("Failed: (%v,%v)", actual, expected)
    }
}

func TestSortCase1Dec(t *testing.T) {
    arr := []int{5,2,4,6,1,3}
    expected := []int{6,5,4,3,2,1}

    actual := InsertionSortDec(arr)

    if !slices.Equal(actual, expected) {
        t.Fatalf("Failed: (%v,%v)", actual, expected)
    }
}
