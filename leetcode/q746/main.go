package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	dp := make([]int , len(cost)+1)
	a1 := complete_dp(cost,&dp)
	return a1
	// dp_2 := make([]int , len(cost))
	// a2 := complete_dp(cost[1:],&dp_2)
	// if a1<=a2 {
	// 	return a1
	// }else{
	// 	return a2
	// }

}

func complete_dp(cost []int, dp *[]int) int{
	(*dp)[0] = 0
	(*dp)[1] = 0
	for i:=2;i<(len(*dp));i++{
		if ((*dp)[i-1]+cost[i-1])<=((*dp)[i-2]+cost[i-2]){
			(*dp)[i] = (*dp)[i-1]+cost[i-1]
		}else{
			(*dp)[i] = (*dp)[i-2]+cost[i-2]
		}

	}
	fmt.Println(*dp)
	return (*dp)[len(*dp)-1]

}


func main()  {
	test := []int{0,1,1,1}
	fmt.Println(minCostClimbingStairs(test))
	
}