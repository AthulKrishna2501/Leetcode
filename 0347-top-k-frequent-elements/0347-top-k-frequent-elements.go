type MaxHeap struct {
	arr []node
}
type node struct{
    value int
    freq int
}
func topKFrequent(nums []int, k int) []int {
    h:=&MaxHeap{}
    count:=make(map[int]int)
    for _,val:=range nums{
        count[val]++
    }
    for key,val:=range count{
        h.arr=append(h.arr,node{key,val})
    }
    h.build(h.arr)
    var topk[]int
    for k!=0{
        topk=append(topk,h.pop())
        k--
    }


    return topk
    
}
func (h *MaxHeap) pop() int {
	max := h.arr[0].value
	h.arr[0] = h.arr[len(h.arr)-1] 
	h.arr = h.arr[:len(h.arr)-1]   
	h.heapifyDown(0)

	return max
}
func (h *MaxHeap) build(arr []node) {
    for i:=len(h.arr)/2-1;i>=0;i--{
        h.heapifyDown(i)
    }
}
func (h *MaxHeap) heapifyDown(index int) {
	largest := index
	left := 2*index + 1
	right := 2*index + 2

	if left < len(h.arr) && h.arr[left].freq > h.arr[largest].freq  {
		largest = left
	}
	if right < len(h.arr) && h.arr[right].freq > h.arr[largest].freq  {
		largest = right
	}

	if largest != index {
		h.arr[index], h.arr[largest] = h.arr[largest], h.arr[index]
		h.heapifyDown(largest)
	}
}