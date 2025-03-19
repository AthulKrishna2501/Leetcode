/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    dummyNode:=&ListNode{Next:head}
    current:=dummyNode


    for current.Next!=nil{
        if current.Next.Val==val{
            current.Next=current.Next.Next
        }else{
        current=current.Next
        }
    }

    return dummyNode.Next
}