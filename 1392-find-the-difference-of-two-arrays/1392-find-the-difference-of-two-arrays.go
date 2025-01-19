func findDifference(nums1 []int, nums2 []int) [][]int {
    seen1:=make(map[int]bool)
    seen2:=make(map[int]bool)
    arr1:=[]int{}
    arr2:=[]int{}

    for _, num1:= range nums1{
        seen1[num1]=true
    }

    for _, num2:= range nums2{
        seen2[num2]=true
    }


    for value := range seen1{
        if !seen2[value]{
            arr1=append(arr1,value)
        }
    }

    for value:= range seen2{
        if !seen1[value]{
            arr2= append(arr2,value)
        }
    }

    return [][]int{arr1,arr2}
}