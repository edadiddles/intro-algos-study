package problems

func InsertionSort(arr *[]int) {
    for i:=0; i < len(*arr)-1; i++ {
        for j:=i; j < len(*arr); j++ {
            if (*arr)[i] > (*arr)[j] {
                temp := (*arr)[i]
                (*arr)[i] = (*arr)[j]
                (*arr)[j] = temp
            }
        }
    }
}
