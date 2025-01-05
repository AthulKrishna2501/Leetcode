func sortArrayByParity(nums []int) []int {
    even:=[]int{}
    odd:=[]int{}
    for _, num:= range nums{
        if num%2==0{
            even=append(even,num)
        }else{
            odd=append(odd,num)
        }
    }

    mergeArray:=append(even,odd...)
    return mergeArray
}