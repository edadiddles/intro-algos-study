package insertion_sort

import (
    "errors"
)

func InsertionSort(arr []int) []int {
    for i:=1; i < len(arr); i++ {
        key := arr[i]
        j := i-1
        for ; j>=0 && arr[j] > key; j-- {
            arr[j+1] = arr[j]
        }
        arr[j+1] = key
    }

    return arr
}

func InsertionSortDec(arr []int) []int {
    for i:=1; i < len(arr); i++ {
        key := arr[i]
        j := i-1
        for ; j>=0 && arr[j] < key; j-- {
            arr[j+1] = arr[j]
        }
        arr[j+1] = key
    }

    return arr
}

func AddBinaryIntegers(a []int, b []int) ([]int, error) {
    if len(a) != len(b) {
        return []int{}, errors.New("Array's must be same length")
    }

    c := make([]int, len(a)+1)

    carry_val := 0
    for i:=len(a)-1; i >= 0; i-- {
        v := (a[i]+b[i]+carry_val)%2
        carry_val = (a[i]+b[i]+carry_val)/2
        c[i+1] = v
    }
    c[0] = carry_val

    return c, nil
}

