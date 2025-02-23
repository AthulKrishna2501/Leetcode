/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    arr:=[]int{}

    current:=head
    for current!=nil{
        arr=append(arr,current.Val)
        current=current.Next
    }


    sort.Ints(arr)

    dummy:=&ListNode{}
    curr:=dummy

    for _, val:= range arr{
        curr.Next=&ListNode{Val:val}
        curr=curr.Next
    }

    return dummy.Next
}