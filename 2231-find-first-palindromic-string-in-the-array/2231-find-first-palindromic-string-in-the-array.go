func firstPalindrome(words []string) string {
     for _,str:= range words{
        reverse:=Reverse(str)
        isTrue,word:=IsPalindrome(reverse,str)
        if isTrue{
          return word
        }
    }
    return ""
}

func Reverse(str string)string{
    runes:=[]rune(str)
    for i,j:=0,len(runes)-1;i<j;i,j=i+1,j-1{
        runes[i],runes[j]=runes[j],runes[i]
    }
    return string(runes)
}


func IsPalindrome(reverse,str string)(bool,string){
  if reverse==str{
    return true,str
  }
  return false,""
}


