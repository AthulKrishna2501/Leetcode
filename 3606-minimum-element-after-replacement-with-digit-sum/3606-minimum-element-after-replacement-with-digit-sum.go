func minElement(nums []int) int {
    result:=make([]int,len(nums))
    for i,num:=range nums{
        result[i]=findSum(num)
    }

    smElement:=findSmall(result)
    return smElement
}


func findSum(num int)int{
    sum:=0
    for num>0{
        sum+=num%10
        num/=10
    }
    return sum
}

func findSmall(arr []int)int{
    min:=math.MaxInt
    for _, num:= range arr{
        if num<min{
            min=num
        }
    }
    return min
}