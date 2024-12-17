func countSeniors(details []string) int {
    var count int
    for _, word:= range details{
        splitWord:=word[11:13]
        age,_:=strconv.Atoi(splitWord)
        if age>60{
            count++
        }

    }
    return count

}