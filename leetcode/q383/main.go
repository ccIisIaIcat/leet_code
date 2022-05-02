package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	search_map := make(map[int32]int)
	for _,v := range(magazine){
		search_map[v] ++
	}
	for _,v_2 := range(ransomNote){
		search_map[v_2] --
		if search_map[v_2] == -1{
			return false
		}
	}
	return true

}


func main()  {
	a := "soi"
	fmt.Printf("%T",a[0])
	for _,v := range(a){
		fmt.Printf("%T\n",v)
		fmt.Println(v)
	}
	
}

