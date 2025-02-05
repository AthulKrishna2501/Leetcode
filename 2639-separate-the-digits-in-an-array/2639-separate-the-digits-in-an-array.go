func separateDigits(nums []int) []int {
    str:=""
    result:=[]int{}
    for _,num:= range nums{
        toStr:=strconv.Itoa(num)
        str+=toStr
    }

    for _,letter:= range str{
        num,_:=strconv.Atoi(string(letter))
        result=append(result,num)
    }
    

    return result
}