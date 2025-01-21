func containsNearbyDuplicate(nums []int, k int) bool {
   indexMap := make(map[int]int) 
    
    for i, num := range nums {
        if prevIndex, found := indexMap[num]; found && i-prevIndex <= k {
            return true 
        }
        indexMap[num] = i 
    }
    return false 

}