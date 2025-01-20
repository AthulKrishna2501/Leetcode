func firstCompleteIndex(arr []int, mat [][]int) int {
      row:=len(mat)
   col:=len(mat[0])
   
   hashMap:=make(map[int][2]int)
   
   for i:=0;i<row;i++{
     for j:=0;j<col;j++{
       hashMap[mat[i][j]]=[2]int{i,j}
     }
   }
   
  rowCount:=make([]int,row)
  colCount:=make([]int,col)
  
  
  for i:=0;i<len(arr);i++{
    r, c := hashMap[arr[i]][0], hashMap[arr[i]][1]
    rowCount[r]++
    colCount[c]++
    
    if colCount[c]==row || rowCount[r]==col{
      return i
    }
  }
  return -1
}