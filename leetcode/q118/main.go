package main

import "fmt"

func generate(numRows int) [][]int {
	answer := make([][]int,numRows)
	for i:=0;i<numRows;i++{
		for j:=0;j<i+1;j++{
			// fmt.Println(i,"++++",j)
			if j==0 || j == i{
				answer[i] = append(answer[i],1)
			}else{
				answer[i] = append(answer[i],answer[i-1][j-1]+answer[i-1][j])
			}
		}
	}
	return answer

}

func main()  {
	fmt.Println(generate(4))
	
}