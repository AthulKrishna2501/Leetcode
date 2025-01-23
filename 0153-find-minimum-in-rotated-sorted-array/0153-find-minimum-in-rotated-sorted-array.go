func findMin(nums []int) int {
    result:=nums[0]

    left,right:=0,len(nums)-1

    for left<=right{
        //check the array is aldready sorted
        if nums[left]<nums[right]{
            result=min(result,nums[left])
            break
        }

        mid:=(left+right)/2
        result=min(result,nums[mid])
        if nums[left]<=nums[mid]{
            left=mid+1
        }else{
            right=mid-1
        }
    }

    return result
}