/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    var dummy = new(ListNode)
    
    dummy.Next = head
    var point *ListNode
    point=dummy

    for point.Next!=nil && point.Next.Next!=nil{
        var swap1,swap2 *ListNode

        swap1=point.Next
        swap2=point.Next.Next

        swap1.Next=swap2.Next
        swap2.Next=swap1

        point.Next=swap2
        point=swap1
    }
    return dummy.Next
}