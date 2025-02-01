func isArraySpecial(nums []int) bool {
    if len(nums)<=1{
        return true
    }

    for i:=0;i<len(nums)-1;i++{
        if nums[i]%2==0{
            if nums[i+1]%2==0{
                return false
            }
        }else if nums[i]%2!=0{
            if nums[i+1]%2!=0{
                return false
            }
        }
    }

    return true



}