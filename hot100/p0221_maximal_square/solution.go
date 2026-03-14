package p0221_maximal_square

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	maxSide := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				minVal := dp[i-1][j]
				if dp[i][j-1] < minVal {
					minVal = dp[i][j-1]
				}
				if dp[i-1][j-1] < minVal {
					minVal = dp[i-1][j-1]
				}
				dp[i][j] = minVal + 1
				if dp[i][j] > maxSide {
					maxSide = dp[i][j]
				}
			}
		}
	}
	return maxSide * maxSide
}
