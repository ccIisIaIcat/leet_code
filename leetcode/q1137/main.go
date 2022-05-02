package main

import "fmt"


func tribonacci(n int) int {
    if n==0{
		return 0
	}else if n<=2{
		return 1
	}
	a,b,c,d := 0,1,1,2
	for i:=3;i<n;i++{
		a = b
		b = c 
		c = d
		d = a+b+c 
	}
	return d 
}

func main()  {
	fmt.Println(tribonacci(34))
	
}