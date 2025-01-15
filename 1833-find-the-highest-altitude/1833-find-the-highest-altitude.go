func largestAltitude(gain []int) int {
    result:=[]int{0}
    for i:=0;i<len(gain);i++{
        result=append(result,gain[i]+result[i])
    }

    max:=findMax(result)
    return max

}



func findMax(arr []int)int{
    max:=math.MinInt
    for _, num:= range arr{
        if num>max{
            max=num
        }
    }
    return max
}