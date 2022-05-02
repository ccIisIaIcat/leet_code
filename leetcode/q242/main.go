package main

import "fmt"

func isAnagram(s string, t string) bool {
    map_1 := make(map[int32]int)
    map_2 := make(map[int32]int)
    for _,v := range(s){
        map_1[v] ++
    }
    for _,v := range(t){
        map_2[v] ++
    }
    if len(s) != len(t){
        return false
    }else{
        for _,v:= range(t){
            if map_1[v] != map_2[v]{
                return false
            }
        }
    }
    return true

}

func main()  {
	fmt.Println(3)
	
}