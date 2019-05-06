package array

func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for i := 1; i <= numRows; i++{
		res[i-1] = make([]int, i)
		if i == 1{
			res[0][0] = 1
		}
	}
	for y := 0; y < numRows; y++{
		find(numRows-1, y, &res)
	}
	return res
}

func find(x, y int, grid *[][]int) int{
	Grid := *grid
	if Grid[x][y] > 0{
		return Grid[x][y]
	}
	if y == 0 || x == y{
		Grid[x][y] = 1
	} else{
		Grid[x][y] = find(x-1, y-1, grid) + find(x-1, y, grid)
	}
	return Grid[x][y]
}
