package main

func numTrees(n int) int {
	if n == 0 {
		return 1
	} else {
		a := 0
		for i := 0; i < n; i++ {
			a += numTrees(i) * numTrees(n-i-1)
		}
		return a
	}
}

func main() {

}
