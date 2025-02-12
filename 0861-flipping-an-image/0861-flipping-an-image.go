func flipAndInvertImage(image [][]int) [][]int {
    reversedArr:=[][]int{}
    for i:=0;i<len(image);i++{
        rvr:=Reverse(image[i])
        reversedArr=append(reversedArr,rvr)
    }


    for i,arr:= range image{
        for j,num:= range arr{
            if num==1{
                reversedArr[i][j]=0
            }else{
                reversedArr[i][j]=1
            }
        }
    }

    return reversedArr
}



func Reverse(arr []int)[]int{
    for i,j:=0,len(arr)-1;i<j;i,j=i+1,j-1{
        arr[i],arr[j]=arr[j],arr[i]
    }

    return arr
}