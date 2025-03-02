func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
    result:=[][]int{}
    i,j:=0,0

    for i<len(nums1)&& j<len(nums2){
        idx1,value1:=nums1[i][0],nums1[i][1]
        idx2,value2:=nums2[j][0],nums2[j][1]

            if idx1<idx2{
                result=append(result,nums1[i])
                i++
            }else if idx1>idx2{
                result=append(result,nums2[j])
                j++
            }else{
                result=append(result,[]int{idx1,value1+value2})
                i++
                j++
            }

    }

    if i<len(nums1){
        result=append(result,nums1[i:]...)
    }

    if j<len(nums2){
        result=append(result,nums2[j:]...)
    }

    return result
}