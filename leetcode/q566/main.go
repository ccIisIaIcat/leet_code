package main

import "fmt"


func matrixReshape(mat [][]int, r int, c int) [][]int {
	a := len(mat)
	b := len(mat[0])
	answer := make([][]int,r)
	if a*b != r*c{
		return mat
	}else{
		for i:=0;i<a;i++{
			for j:=0;j<b;j++{
				answer[((i*b+(j+1)-1))/c] = append(answer[((i*b+(j+1)-1))/c],mat[i][j])
			}
		}
	}
	return answer

}

func main()  {
	test := make([][]int,2)
	a := []int{1,2}
	b := []int{3,4}
	// test[0] = append(test[0],a...)
	// test[1] = append(test[1],b...)
	test[0] = a
	test[1] = b
	fmt.Println(matrixReshape(test,4,1))
}