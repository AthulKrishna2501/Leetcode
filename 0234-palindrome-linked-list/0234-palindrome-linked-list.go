/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    checkArr:=[]int{}

    current:=head
    for current!=nil{
        checkArr=append(checkArr,current.Val)
        current=current.Next
    }


    return checkPalindrome(checkArr)

}

func checkPalindrome(arr []int)bool{
    i,j:=0,len(arr)-1
    for i<=j{
        if arr[i]!=arr[j]{
            return false
        }

        i++
        j--
    }

    return true
}