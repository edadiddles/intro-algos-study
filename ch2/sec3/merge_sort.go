package section3


func MergeSort(arr *[]int) {
    r_MergeSort(arr, 0, len(*arr)) 
}

func r_MergeSort(arr *[]int, p int, r int) {
    if p >= r-1 {
        return
    }
    q := (p+r)/2

    r_MergeSort(arr, p, q)
    r_MergeSort(arr, q, r)
    
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
