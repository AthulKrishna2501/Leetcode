func maxArea(height []int) int {
    maxarea:=0

    left,right:=0,len(height)-1

    for left<right{
        area:=0
        if height[left]>height[right]{
            area=height[right]*(right-left)
            right--
        }else{
            area=height[left]*(right-left)
            left++
        }

        maxarea=max(maxarea,area)
    }

    return maxarea



}