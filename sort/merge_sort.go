package sort

func mergeSort(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	n := len(nums)
	left := mergeSort(nums[:n/2])
	right := mergeSort(nums[n/2:])
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	n1, n2 := len(left), len(right)
	res := make([]int, n1+n2)
	var p1, p2, p int // pointers to positions in left, right and result

	for p1 < n1 && p2 < n2 {
		if left[p1] <= right[p2] {
			res[p] = left[p1]
			p1++
		} else {
			res[p] = right[p2]
			p2++
		}
		p++
	}

	for p1 < n1 {
		res[p] = left[p1]
		p++
		p1++
	}

	for p2 < n2 {
		res[p] = right[p2]
		p++
		p2++
	}

	return res
}
