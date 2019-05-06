package array

// O(N^2) solution
func maxSumTwoNoOverlap(A []int, L int, M int) int {
	accumSum := make([]int, 0)
	accumSum = append(accumSum, 0)
	for _, num := range A{
		if len(accumSum) == 0{
			accumSum = append(accumSum, num)
		} else{
			accumSum = append(accumSum, accumSum[len(accumSum)-1]+num)
		}
	}
	res := 0
	start := 1
	for start + L <= len(accumSum){
		curRes := accumSum[start+L-1] - accumSum[start-1] + maxSum(&accumSum, start, start+L-1, M)
		if curRes > res{
			res = curRes
		}
		start++
	}
	return res
}

// find max excluding start to end
func maxSum(accumSum *[]int, start int, end int, M int) int{
	res := 0
	rightStart := start - 1
	for rightStart - M >= 0{
		curSum := (*accumSum)[rightStart] - (*accumSum)[rightStart-M]
		if curSum > res{
			res = curSum
		}
		rightStart--
	}
	leftStart := end + 1
	for leftStart + M <= len(*accumSum){
		curSum := (*accumSum)[leftStart + M - 1] - (*accumSum)[leftStart-1]
		if curSum > res{
			res = curSum
		}
		leftStart++
	}
	return res
}
