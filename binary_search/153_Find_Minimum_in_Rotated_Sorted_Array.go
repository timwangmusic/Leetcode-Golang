package binary_search

func findMin(nums []int) int {
	start, end := 0, len(nums)-1
	for start < end {
		mid := start + (end - start) / 2
		if nums[mid] > nums[end] {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return nums[start]
}
