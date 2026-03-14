package p0323_number_of_connected_components

func countComponents(n int, edges [][]int) int {
	parent := make([]int, n)
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
	components := n
	for _, e := range edges {
		px, py := find(e[0]), find(e[1])
		if px != py {
			parent[px] = py
			components--
		}
	}
	return components
}
