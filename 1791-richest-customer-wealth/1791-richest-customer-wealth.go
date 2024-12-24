func maximumWealth(accounts [][]int) int {
    maxSum:=0
    for _, num:= range accounts{
        sum:=0
        for _, n:= range num{
            sum+=n
        }
        if sum>maxSum{
            maxSum=sum
        }
        sum=0
    }
    return maxSum
}