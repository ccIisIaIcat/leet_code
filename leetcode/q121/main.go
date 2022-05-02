package main

import "fmt"


func maxProfit(prices []int) int {
	size := len(prices)
	if size == 1{
		return 0
	}
	max_gap := make(map[int]int,0)
	max_gap[size-2]= max(prices[size-1]-prices[size-2],0)
	max_num := max_gap[size-2]
	for i:=size-3;i>=0;i--{
		max_gap[i] = max(prices[i+1]-prices[i]+max_gap[i+1],0)
		if max_gap[i]>=max_num{
			max_num = max_gap[i]
		}
	}
	return max_num

}

func max(a int,b int)  int{
	if a>=b{
		return a
	}else{
		return b
	}
}

func main()  {
	test := []int{7,1,5,3,6,4}
	fmt.Println(maxProfit(test))
	
}