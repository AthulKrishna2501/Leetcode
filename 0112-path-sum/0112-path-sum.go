/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
    return isPathSum(root,0,targetSum)
}

func isPathSum(root *TreeNode,currentSum,targetSum int)bool{
    if root==nil{
        return false
    }

    currentSum+=root.Val

    if root.Left==nil && root.Right==nil{
        if currentSum==targetSum{
            return true
        }

        return false
    }

    left:=isPathSum(root.Left,currentSum,targetSum)
    right:=isPathSum(root.Right,currentSum,targetSum)

    return left||right
}