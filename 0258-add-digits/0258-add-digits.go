func addDigits(num int) int {
    for num >= 10 {
        sum := 0
        for num > 0 {
            sum += num % 10
            num /= 10
        }
        num = sum
    }
    return num


}