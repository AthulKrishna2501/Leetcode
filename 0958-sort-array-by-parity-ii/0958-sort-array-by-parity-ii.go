func sortArrayByParityII(nums []int) []int {
    even:=[]int{}
    odd:=[]int{}
    result:=[]int{}

    for _, num:= range nums{
        if num%2==0{
            even=append(even,num)
        }else{
            odd=append(odd,num)
        }
    }


    for i:=0;i<len(even);i++{
        result=append(result,even[i])
        for j:=i;j<len(odd);j++{
            result=append(result,odd[i])
            break
        }
    }

    return result
}