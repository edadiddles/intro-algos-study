package section2

func SelectionSort(arr []int) []int {
    for i:=0; i < len(arr)-1; i++ {
        idx_min := i
        for j:=i+1; j < len(arr); j++ {
           if arr[j] < arr[idx_min] {
               idx_min = j
           }
       }
       temp := arr[i]
       arr[i] = arr[idx_min]
       arr[idx_min] = temp
   }

    return arr
}
