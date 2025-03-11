func restoreString(s string, indices []int) string {
t := make([]byte, len(s))
    for i, str:= range s{
        t[indices[i]]=byte(str)
    }

    return string(t)

}