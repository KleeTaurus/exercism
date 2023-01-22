package pascal

func Triangle(n int) [][]int {
	triangle := make([][]int, 0)
	for i := 1; i < n+1; i++ {
		triangle = append(triangle, make([]int, i))

		for j := 0; j < i; j++ {
			if j == 0 || j == i-1 {
				triangle[i-1][j] = 1
			} else {
				triangle[i-1][j] = triangle[i-2][j-1] + triangle[i-2][j]
			}
		}
	}
	return triangle
}
