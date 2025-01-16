func searchRange(nums []int, target int) []int {
    if len(nums)==0{
        return []int{-1,-1}
    }

    first,last:=-1,-1
    
    for i, num:= range nums{
        if num==target{
            if first==-1{
                first=i
            }
            last=i
        }
    }
    return []int{first,last}
}