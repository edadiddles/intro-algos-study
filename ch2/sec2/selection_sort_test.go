package section2

import (
    "testing"
)

func TestSelectionSortCase1(t *testing.T) {
    arr := []int{4,3,2,6,5,1}
    expected := []int{1,2,3,4,5,6}

    actual := SelectionSort(arr)

    if !slices.Equal(actual, expected) {
        t.Fatalf("Failed: (%v,%v)", actual, expected)
    }
}
