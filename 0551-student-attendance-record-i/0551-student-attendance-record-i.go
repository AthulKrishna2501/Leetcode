func checkRecord(s string) bool {
    if strings.Count(s,"A")>=2 || strings.Contains(s,"LLL"){
        return false
    }

    return true



}