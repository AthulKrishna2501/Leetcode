/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
 return IsBST(root,math.MinInt,math.MaxInt)
}

func IsBST(node *TreeNode,min,max int)bool{
    if node==nil{
        return true
    }

    if node.Val <=min || node.Val>=max{
        return false
    }

    return IsBST(node.Left,min,node.Val) && IsBST(node.Right,node.Val,max)
}
