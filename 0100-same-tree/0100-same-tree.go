/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p==nil && q==nil{
        return true
    }
    if p==nil || q==nil{
        return false
    }

    return (p.Val==q.Val)&& 
    isSameTree(q.Left,p.Left) && 
    isSameTree(p.Right,q.Right)
}