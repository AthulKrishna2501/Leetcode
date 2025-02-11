func kidsWithCandies(candies []int, extraCandies int) []bool {
    m:=slices.Max(candies)
    result:=[]bool{}
    for _, c:= range candies{
        if c + extraCandies >= m{
            result = append(result,true)
        }else{
            result=append(result,false)
        }
    }

    return result

    
}