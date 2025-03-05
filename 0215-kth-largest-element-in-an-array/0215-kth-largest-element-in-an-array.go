func findKthLargest(nums []int, k int) int {
    h:=&MaxHeap{}
    h.build(nums)
    result:=0
    for i:=0;i<k;i++{
        result=h.pop()
    }

    return result
}

type MaxHeap struct {
	arr []int
}


func (h *MaxHeap) pop() int {
	max := h.arr[0]
	h.arr[0] = h.arr[len(h.arr)-1] 
	h.arr = h.arr[:len(h.arr)-1]   
	h.heapifyDown(0)

	return max
}
func (h *MaxHeap) build(arr []int) {
    h.arr=arr
    for i:=len(h.arr)/2-1;i>=0;i--{
        h.heapifyDown(i)
    }
}
func (h *MaxHeap) heapifyDown(index int) {
	largest := index
	left := 2*index + 1
	right := 2*index + 2

	if left < len(h.arr) && h.arr[left] > h.arr[largest]  {
		largest = left
	}
	if right < len(h.arr) && h.arr[right] > h.arr[largest]  {
		largest = right
	}

	if largest != index {
		h.arr[index], h.arr[largest] = h.arr[largest], h.arr[index]
		h.heapifyDown(largest)
	}
}