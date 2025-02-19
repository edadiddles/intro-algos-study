package problems

import (
    "testing"
    "slices"
)

func TestBubbleSortCase1(t *testing.T) {
    arr := []int{1,5,3,2,4,6}
    expected := []int{1,2,3,4,5,6}

    BubbleSort(&arr)

    if !slices.Equal(arr, expected) {
        t.Fatalf("Failed: (%v,%v)", arr, expected)
    }
}
