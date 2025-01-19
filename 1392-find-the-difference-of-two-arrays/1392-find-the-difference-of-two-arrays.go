func findDifference(nums1 []int, nums2 []int) [][]int {
    seen:=make(map[int]bool)
  arr1:=[]int{}
  arr2:=[]int{}


  for i:=0;i<len(nums1);i++{
      seen[nums1[i]]=true
}

for j:=0;j<len(nums2);j++{
  if seen[nums2[j]]{
    seen[nums2[j]]=false
  }
}

for key,value:= range seen{
  if value!=false{
    arr1=append(arr1,key)
  }
}

 for i:=0;i<len(nums2);i++{
      seen[nums2[i]]=true
}

for j:=0;j<len(nums1);j++{
  if seen[nums1[j]]{
    seen[nums1[j]]=false
  }
}

for key,value:= range seen{
  if value!=false{
    arr2=append(arr2,key)
  }
}

result:=[][]int{}
result=append(result,arr1,arr2)

return result

}