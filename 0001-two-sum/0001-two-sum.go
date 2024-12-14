func twoSum(nums []int, target int) []int {
    hashmap:=make(map[int]int)
   for i,num:= range nums{
    complement:=target-num
    if index,found:=hashmap[complement];found{
        return []int{index,i}
    }
    hashmap[num]=i
   }
   return nil
}