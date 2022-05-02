package main

import "fmt"


func firstUniqChar(s string) int {
	search_map := make(map[string]int)
	for _,v := range(s){
		search_map[string(v)] ++
	}
	for i:=0;i<len(s);i++{
		if search_map[string(s[i])] == 1{
			return i
		}
	}
	return -1

}

func main()  {
	fmt.Println("hello world")
	s := "loveleetcode"
	fmt.Println(firstUniqChar(s))
}