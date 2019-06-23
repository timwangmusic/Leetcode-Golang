package sliding_window

func maxSatisfied(customers []int, grumpy []int, X int) int {
	last := 0
	for i:=0; i<X; i++{
		last += customers[i]
	}
	for i:=X; i<len(customers); i++{
		if grumpy[i] == 0{
			last += customers[i]
		}
	}
	res := last
	for i:=1; i+X <= len(customers); i++{
		if grumpy[i-1] == 1{
			last -= customers[i-1]
		}
		if grumpy[i+X-1] == 1{
			last += customers[i+X-1]
		}
		if last > res{
			res = last
		}
	}
	return res
}
