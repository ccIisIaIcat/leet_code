package main

func maximumWealth(accounts [][]int) int {
	max_num := 0
	for i := 0; i < len(accounts); i++ {
		temp := 0
		for j := 0; j < len(accounts[i]); j++ {
			temp += accounts[i][j]
		}
		if temp > max_num {
			max_num = temp
		}
	}
	return max_num
}

func main() {

}
