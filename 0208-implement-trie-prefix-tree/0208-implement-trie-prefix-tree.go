type Trie struct {
    children map[rune]*Trie
    isEnd bool
}


func Constructor() Trie {
    return Trie{
        children:make(map[rune]*Trie),
        isEnd:false,
    }
}


func (this *Trie) Insert(word string)  {
    node :=this
    for _, str:= range word{
        if _, exists:=node.children[str];!exists{
            node.children[str]=&Trie{children:make(map[rune]*Trie)}
        }
        node = node.children[str]
    }

    node.isEnd=true
}


func (this *Trie) Search(word string) bool {
    node := this
    for _, str:= range word{
        if _, exists:=node.children[str];!exists{
            return false
        }

        node=node.children[str]
    }

    return node.isEnd
}


func (this *Trie) StartsWith(prefix string) bool {
    node := this
    for _, str:= range prefix{
        if _, exists:=node.children[str];!exists{
            return false
        }
        node=node.children[str]
    }

    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */