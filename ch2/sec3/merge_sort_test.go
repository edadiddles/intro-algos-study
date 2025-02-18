package section3

import (
    "testing"
    "slices"
)

func TestMergeSortCase1(t *testing.T) {
    arr := []int{4,3,5,2,6,1}
    expected := []int{1,2,3,4,5,6}

    MergeSort(&arr)

    if !slices.Equal(arr, expected) {
        t.Fatalf("Failed: (%v,%v)", arr, expected)
    }
}


