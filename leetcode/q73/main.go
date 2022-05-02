package main

import "fmt"

func setZeroes(matrix [][]int)  {
	a := len(matrix)
	b := len(matrix[0])
	var line []int
	var raw []int
	for i:=0;i<a;i++{
		for j:=0;j<b;j++{
			if matrix[i][j] == 0{
				line = append(line,i)
				raw = append(raw,j)
			}
		}
	}
	for _,l := range(line){
		for j:=0;j<b;j++{
			matrix[l][j] = 0
		}
	}
	for _,r := range(raw){
		for j:=0;j<a;j++{
			matrix[j][r] = 0
		}
	}

}

func main()  {
	fmt.Println("lalala")
	
}