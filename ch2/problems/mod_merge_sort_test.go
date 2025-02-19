package problems

import (
    "testing"
    "slices"
)

func TestModMergeSortCase1(t *testing.T) {
    arr := []int{4,3,5,2,1,6}
    expected := []int{1,2,3,4,5,6}

    ModMergeSort(&arr, 1)

    if !slices.Equal(arr, expected) {
        t.Fatalf("Failed: (%v,%v)", arr, expected)
    }
}

func TestModMergeSortCase2(t *testing.T) {
    arr := []int{4,3,5,2,1,6}
    expected := []int{1,2,3,4,5,6}

    ModMergeSort(&arr, 2)

    if !slices.Equal(arr, expected) {
        t.Fatalf("Failed: (%v,%v)", arr, expected)
    }
}

func TestModMergeSortCase3(t *testing.T) {
    arr := []int{4,3,5,2,1,6}
    expected := []int{1,2,3,4,5,6}

    ModMergeSort(&arr, 3)

    if !slices.Equal(arr, expected) {
        t.Fatalf("Failed: (%v,%v)", arr, expected)
    }
}

func TestModMergeSortCase4(t *testing.T) {
    arr := []int{4,3,5,2,1,6}
    expected := []int{1,2,3,4,5,6}

    ModMergeSort(&arr, 10)

    if !slices.Equal(arr, expected) {
        t.Fatalf("Failed: (%v,%v)", arr, expected)
    }
}
