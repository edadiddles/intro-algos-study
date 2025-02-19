package problems

func ModMergeSort(arr *[]int, l int) {
    r_ModMergeSort(arr, l, 0, len(*arr))
}

func r_ModMergeSort(arr *[]int, l int, p int, r int) {
    if l >= r-p {
        a := (*arr)[p:r]
        InsertionSort(&a)
        return
    }

    q := (p+r)/2

    r_ModMergeSort(arr, l, p, q)
    r_ModMergeSort(arr, l, q, r)

    merge(arr, p, q, r)
}

func merge(arr *[]int, p int, q int, r int) {
    a := make([]int, 0)
    b := make([]int, 0)

    a = append(a, (*arr)[p:q]...)
    b = append(b, (*arr)[q:r]...)

    i,j := 0,0
    k := p
    for ; i < len(a) && j < len(b); {
        if a[i] < b[j] {
            (*arr)[k] = a[i]
            i++
        } else {
            (*arr)[k] = b[j]
            j++
        }
        k++
    }

    for ; i < len(a); i++ {
        (*arr)[k] = a[i]
        k++
    }

    for ; j < len(b); j++ {
        (*arr)[k] = b[j]
        k++
    }
}
