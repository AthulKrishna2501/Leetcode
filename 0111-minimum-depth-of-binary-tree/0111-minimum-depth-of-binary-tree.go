/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func minDepth(root *TreeNode) int {
    if root==nil{
        return 0
    }



        left:=minDepth(root.Left)

        right:=minDepth(root.Right) 

        if left==0{
            left=right
        }

        if right==0{
            right=left
        }




	
	return min(left, right)+1

}

func min(a, b int) int {
	if a < b {
		return a
	}
    

	return b
}