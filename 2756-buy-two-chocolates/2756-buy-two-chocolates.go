func buyChoco(prices []int, money int) int {
    sort.Ints(prices)
    first:=prices[0]
    second:=prices[1]
    sum:=first+second

    if sum>money{
        return money
    }

    return money-sum
}