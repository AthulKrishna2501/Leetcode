/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    slice:=[]int{}
    for list1!=nil{
        slice=append(slice,list1.Val)
        list1=list1.Next
    }

    for list2!=nil{
        slice=append(slice,list2.Val)
        list2=list2.Next
    }
    sort.Ints(slice)
    newNode:=&ListNode{}
    current:=newNode
    for _, val:= range slice{
        current.Next=&ListNode{Val:val}
        current=current.Next
    }
    return newNode.Next
}