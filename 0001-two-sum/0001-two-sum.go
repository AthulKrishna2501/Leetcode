func twoSum(nums []int, target int) []int {
  seen:=make(map[int]int)

  for i, num:= range nums{
    diff:=target-num
    if index,found:=seen[diff];found{
        return []int{index,i}
    }

    seen[num]=i
  }

  return nil


}