func balancedStringSplit(s string) int{
    balance:=0
    count:=0

    for _, str:= range s{
        if str=='L'{
            balance++
        }else{
            balance--
        }

        if balance==0{
            count++
        }
    }

    return count
}