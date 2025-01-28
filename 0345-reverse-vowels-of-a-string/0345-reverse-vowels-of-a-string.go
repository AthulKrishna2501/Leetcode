func reverseVowels(s string) string {
     i,j:=0,len(s)-1
    slice:=[]byte(s)
   
    for i<j{
        if isVowel(slice[i]) && isVowel(slice[j]){
            slice[i],slice[j]=slice[j],slice[i]
            i++
            j--
            continue
        }

        if !isVowel(slice[i]){
            i++
        }

        if !isVowel(slice[j]){
            j--
        }
    }

    return string(slice)
}


func isVowel(str byte)bool{
    if str=='a' || str=='e' || str=='i'||str=='o'|| str=='u'{
        return true
    }

    if str=='A' || str=='E' || str=='I' || str=='O'|| str=='U'{
        return true
    }

    return false
}