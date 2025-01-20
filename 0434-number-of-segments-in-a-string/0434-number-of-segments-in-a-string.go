func countSegments(s string) int {
    split:=strings.Fields(s)
    count:=0
    for i:=0;i<len(split);i++{
        count++
    }

    return count
}