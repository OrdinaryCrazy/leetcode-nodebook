func search(nums []int, target int) int {
    left := 0
    right := len(nums) - 1
    for left < right {
        mid := left + (right - left) / 2
        if target == nums[mid]{
            return mid
        } else if target < nums[mid] {
            right = mid
        } else {
            left = mid + 1
        }
    } 
	// case len(nums) == 1
    if target == nums[left] {
        return left
    }
    return -1
}