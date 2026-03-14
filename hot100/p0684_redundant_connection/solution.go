package p0684_redundant_connection

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	parent := make([]int, n+1)
	for i := range parent {
		parent[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	for _, e := range edges {
		px, py := find(e[0]), find(e[1])
		if px == py {
			return e
		}
		parent[px] = py
	}
	return nil
}
