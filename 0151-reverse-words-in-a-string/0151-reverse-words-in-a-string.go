func reverseWords(s string) string {
    split:=strings.Fields(s)
	for i,j:=0,len(split)-1;i<j;i,j=i+1,j-1{
	  split[i],split[j]=split[j],split[i]
	}
    
    return strings.Join(split," ")
     
}