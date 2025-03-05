func searchMatrix(matrix [][]int, target int) bool {
    for _, num:= range matrix{
        left,right:=0,len(num)-1
        for left<=right{
            mid:=(left+right)/2 
            if num[mid]==target{
                return true
            }else if num[mid]<target{
                left=mid+1
            }else{
                right=mid-1
            }
        }
    }

    return false
}