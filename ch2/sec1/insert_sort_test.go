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

func TestAddBinaryIntsCase1(t *testing.T) {
    a := []int{1}
    b := []int{1}
    expected := []int{1,0}

    actual, err := AddBinaryIntegers(a,b)

    if err != nil {
        t.Fatalf("Failed: %v", err)
    }

    if !slices.Equal(actual, expected) {
        t.Fatalf("Failed: (%v,%v)", actual, expected)
    }
}

func TestAddBinaryIntsCase2(t *testing.T) {
    a := []int{0,1,1,1,0,0,1,0,1}
    b := []int{1,0,0,0,0,1,0,1,1}
    expected := []int{0,1,1,1,1,1,0,0,0,0}

    actual, err := AddBinaryIntegers(a,b)

    if err != nil {
        t.Fatalf("Failed: %v", err)
    }

    if !slices.Equal(actual, expected) {
        t.Fatalf("Failed: (%v,%v)", actual, expected)
    }
}

func TestAddBinaryIntsCase3(t *testing.T) {
    a := []int{1,1,1,1,1}
    b := []int{1,0,0,0}
    expected := []int{}

    actual, err := AddBinaryIntegers(a,b)

    if err == nil {
        t.Fatalf("Failed: Expected error")
    }

    if !slices.Equal(actual, expected) {
        t.Fatalf("Failed: (%v,%v)", actual, expected)
    }
}
