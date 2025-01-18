func heightChecker(heights []int) int {
    expected:=[]int{}
    expected=append(expected,heights...)

    sort.Ints(expected)
    count:=0

    for i:=range heights{
        if heights[i]!=expected[i]{
           count++
        }
    }

    return count
   
}