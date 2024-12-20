func twoSum(nums []int, target int) []int {
    hashMap:=make(map[int]int)
    for i,num:= range nums{
        sum:=target-num
        if index,found:= hashMap[sum];found{
            return []int{index,i}
        }
        hashMap[num]=i
    }
    return nil
}