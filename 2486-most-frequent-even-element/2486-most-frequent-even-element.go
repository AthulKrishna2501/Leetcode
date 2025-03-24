func mostFrequentEven(nums []int) int {
	count := make(map[int]int)
    arr:=[]int{}

	for _, num := range nums {
		if num%2 == 0 {
			count[num]++
		}
	}

 

    maxCount:=math.MinInt

    for _,value:= range count{
        if value>maxCount{
            maxCount=value
        }
    }

    for key,value:= range count{
        if value==maxCount{
            arr=append(arr,key)
        }
    }

    if len(arr)==0{
        return -1
    }

    sort.Ints(arr)

    return arr[0]
}