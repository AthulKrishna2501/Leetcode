func pivotArray(nums []int, pivot int) []int {
    smArr:=[]int{}
    larArr:=[]int{}
    result:=[]int{}
    equalArr:=[]int{}

    for _, num:= range nums{
        if num<pivot{
            smArr=append(smArr,num)
        }else if num>pivot{
            larArr=append(larArr,num)
        }else{
            equalArr=append(equalArr,num)
        }
    }

    result=append(result,smArr...)
    result=append(result,equalArr...)
    result=append(result,larArr...)
    return result
}