package main

import (
	"fmt"
	"strconv"
)

func squareIsWhite(coordinates string) bool {
	judge, _ := strconv.Atoi(string(coordinates[1]))
	judge = int(coordinates[0]-97) + judge
	fmt.Println(judge)
	if judge%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	squareIsWhite("a2")
}
