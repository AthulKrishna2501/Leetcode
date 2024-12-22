func reverseWords(s string) string {
    result:=[]string{}
    split:=strings.Split(s," ")
    for i:=len(split)-1;i>=0;i--{
        if split[i]!=""{
            result=append(result,split[i])
        }
    }

    return strings.Join(result," ")
     
}