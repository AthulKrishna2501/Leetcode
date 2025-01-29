func smallerNumbersThanCurrent(nums []int) []int {
     result:=[]int{}
    for i:=0;i<len(nums);i++{
        count:=0
        for j:=0;j<len(nums);j++{
            if j!=i && nums[j]<nums[i]{
                count++
            }
        }
        result=append(result,count)
    }
    return result
}