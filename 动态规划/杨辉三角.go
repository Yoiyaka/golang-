package main

import "fmt"

func generate(numRows int) [][]int {
	if numRows <= 0 {
		return nil
	}
	result := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		result[i] = make([]int, i+1)
		result[i][0], result[i][i] = 1, 1
		for j := 1; j < i; j++ {
			result[i][j] = result[i-1][j-1] + result[i-1][j]
		}
	}
	return result
}

func main() {
	var numsRows int
	fmt.Scanln(&numsRows)
	result := generate(numsRows)
	fmt.Println(result)
}
