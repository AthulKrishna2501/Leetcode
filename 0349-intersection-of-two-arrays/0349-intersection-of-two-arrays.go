func intersection(nums1 []int, nums2 []int) []int {
    result:=[]int{}
    hashMap:=make(map[int]bool)

    for i:=0;i<len(nums1);i++{
        for j:=0;j<len(nums2);j++{
            if nums1[i]==nums2[j]{
                if !hashMap[nums1[i]]{
                    hashMap[nums1[i]]=true
                    result=append(result,nums1[i])
                }
            }
        }
    }

    return result
}