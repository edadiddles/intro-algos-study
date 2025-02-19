package problems

func BubbleSort(arr *[]int) {
    for i:=0; i < len(*arr)-1; i++ {
        for j:=len(*arr)-1; j > i; j-- {
            if (*arr)[j] < (*arr)[j-1] {
                temp := (*arr)[j]
                (*arr)[j] = (*arr)[j-1]
                (*arr)[j-1] = temp
            }
        }
    }
}
