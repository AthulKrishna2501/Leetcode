func targetIndices(nums []int, target int) []int {
	result := []int{}

	sort.Ints(nums)

    for i,num:= range nums{
        if num==target{
            result=append(result,i)
        }
    }

	return result
}