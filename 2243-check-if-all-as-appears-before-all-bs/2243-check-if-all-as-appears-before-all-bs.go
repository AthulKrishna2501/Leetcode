func checkString(s string) bool {
   countA:=0
    for _, str:= range s{
        if str=='a'{
            countA++
        }
    }
    fmt.Println(countA)


    for _, str:= range s{
        if str=='b'{
            if countA!=0{
                return false
            }
          }else{
                countA--
          }
    }

    return true
}