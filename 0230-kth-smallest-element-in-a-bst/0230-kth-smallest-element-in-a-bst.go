/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    result:=0
    count:=0

    var inorder func(node *TreeNode)
    inorder=func (node *TreeNode){
        if node==nil{
            return
        }

        inorder(node.Left)
        count++
        if count==k{
            result=node.Val
        }
        inorder(node.Right)
    }
    inorder(root)
    return result
}


