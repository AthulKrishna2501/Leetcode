func sortArray(nums []int) []int {
    if len(nums)<=1{
        return nums
    }
    mid:=len(nums)/2
    left:=sortArray(nums[:mid])
    right:=sortArray(nums[mid:])
    
    return mergeSort(left,right)
}

func mergeSort(left,right []int)[]int{
    result:=[]int{}
    i,j:=0,0
    for i<len(left) && j<len(right){
        if left[i]<=right[j]{
            result=append(result,left[i])
            i++
        }else{
            result=append(result,right[j])
            j++
        }
    }

    result=append(result,left[i:]...)
    result=append(result,right[j:]...)

    return result
    
}