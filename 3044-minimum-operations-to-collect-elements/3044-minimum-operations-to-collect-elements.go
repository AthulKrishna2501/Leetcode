func minOperations(nums []int, k int) int {
    seen:=make(map[int]bool)
    count:=0
    var arr []int

    for i:=len(nums)-1;i>=0;i--{
        if nums[i]<=k{
            if !seen[nums[i]]{
                arr=append(arr,nums[i])
                seen[nums[i]]=true
                if len(arr)==k{
                    break
                }
            }
        }

        count++
    }
fmt.Println(count)
    count++
    return count
}