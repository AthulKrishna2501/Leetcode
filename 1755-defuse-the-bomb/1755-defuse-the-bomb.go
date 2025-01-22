func decrypt(code []int, k int) []int {
    result:=make([]int,len(code))

    if k==0{
        return result
    }
     if k>0{
        for i:=0;i<len(code);i++{
            count:=0
            sum:=0
            for j:=i+1;count<k;count++{
                if j==len(code){
                    j=0
                }
                sum+=code[j]
                j++
            }

            result[i]=sum

        }
    }

    if k<0{
        j:=0
        for i:=0;i<len(code);i++{
            count:=k
            sum:=0
            if i==0{
                j=len(code)-1
            }else{
                j=i-1
            }

            for count!=0{
                count++
                sum+=code[j]
                j--

                if j<0{
                    j=len(code)-1
                }
            }
            result[i]=sum



        }
    }

    return result


   

}