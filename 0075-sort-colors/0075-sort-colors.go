func sortColors(nums []int)  {
    n:=len(nums)

    for i:=0;i<n-1;i++{
        isSwapped:=false
        for j:=0;j<n-i-1;j++{
            if nums[j]>nums[j+1]{
                nums[j],nums[j+1]=nums[j+1],nums[j]
                isSwapped=true
            }
        }
        if isSwapped==false{
            break
        }
    }


}