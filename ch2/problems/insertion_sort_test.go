package problems

import (
    "testing"
    "slices"
)

func TestInsertionSortCase1(t *testing.T) {
    arr := []int{5,4,3,2,6,1}
    expected := []int{1,2,3,4,5,6}

    InsertionSort(&arr)

    if !slices.Equal(arr, expected) {
        t.Fatalf("Failed: (%v,%v)", arr, expected)
    }
}

func TestInsertionSortCase2(t *testing.T) {
    arr := []int{1}
    expected := []int{1}

    InsertionSort(&arr)

    if !slices.Equal(arr, expected) {
        t.Fatalf("Failed: (%v,%v)", arr, expected)
    }
}

func TestInsertionSortCase3(t *testing.T) {
    arr := []int{1,2,3,4,5,6,7}
    expected := []int{1,2,3,4,5,6,7}

    InsertionSort(&arr)

    if !slices.Equal(arr, expected) {
        t.Fatalf("Failed: (%v,%v)", arr, expected)
    }
}

func TestInsertionSortCase4(t *testing.T) {
    arr := []int{9,8,7,6,5,4,3,2,1}
    expected := []int{1,2,3,4,5,6,7,8,9}

    InsertionSort(&arr)

    if !slices.Equal(arr, expected) {
        t.Fatalf("Failed: (%v,%v)", arr, expected)
    }
}

func TestInsertionSortCase5(t *testing.T) {
    arr := []int{}
    expected := []int{}

    InsertionSort(&arr)

    if !slices.Equal(arr, expected) {
        t.Fatalf("Failed: (%v,%v)", arr, expected)
    }
}
