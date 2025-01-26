func minimumChairs(s string) int {
    count:=0
    maxCount:=0
    for _, str:= range s{
        if str=='E'{
            count++
        }else{
            count--
        }

        maxCount=max(maxCount,count)
    }

    return maxCount
}