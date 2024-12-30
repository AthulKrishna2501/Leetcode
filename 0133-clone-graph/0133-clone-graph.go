/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    if node==nil{
        return nil
    }

    clones:=make(map[*Node]*Node)

    var dfs func(*Node)*Node
    dfs=func(n *Node)*Node{
        if clones,exists:=clones[n];exists{
            return clones
        }

        clone:=&Node{Val:n.Val}
        clones[n]=clone

        for _, neighbour:=range n.Neighbors{
            clone.Neighbors=append(clone.Neighbors,dfs(neighbour))
        }
        return clone
    }

    return dfs(node)
}
