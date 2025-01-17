func heightChecker(heights []int) int {
    expected:=[]int{}
    expected=append(expected,heights...)

    sort.Ints(expected)
    result:=[]int{}
    for i:=range heights{
        if heights[i]!=expected[i]{
            result=append(result,i)
        }
    }
    count:=0
   for i:=0;i<len(result);i++{
    count++
   }

    return count
}