/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    newNode:=&TreeNode{Val:val}
    if root==nil{
        root=newNode
    }

    if newNode.Val<root.Val{
        if root.Left==nil{
            root.Left=newNode
        }else{
            insertIntoBST(root.Left,val)
        }
    }else if newNode.Val>root.Val{
        if root.Right==nil{
            root.Right=newNode
        }else{
            insertIntoBST(root.Right,val)
        }
    }
    return root
}