package main

type item struct {
	c string
	v float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	edges := make(map[string][]*item, 16)
	for i := 0; i < len(values); i++ {
		c1 := equations[i][0]
		c2 := equations[i][1]

		if edges[c1] == nil {
			edges[c1] = make([]*item, 0, 16)
		}
		edges[c1] = append(edges[c1], &item{c: c2, v: values[i]})

		if edges[c2] == nil {
			edges[c2] = make([]*item, 0, 16)
		}
		edges[c2] = append(edges[c2], &item{c: c1, v: float64(1) / values[i]})
	}

	res := make([]float64, len(queries))
	for i := 0; i < len(res); i++ {
		res[i] = -1.0
	}

	for i, q := range queries {
		pathSet := make(map[string]struct{}, 16)
		dfs(q[0], q[1], 1.0, edges, res, i, pathSet)
	}

	return res
}

func dfs(b, e string, val float64, edges map[string][]*item, res []float64, i int, pathSet map[string]struct{}) {
	if _, ok := pathSet[b]; ok {
		return
	}
	pathSet[b] = struct{}{}
	_, ok := edges[b]
	if b == e && ok {
		res[i] = val
		return
	}
	for _, n := range edges[b] {
		dfs(n.c, e, val*n.v, edges, res, i, pathSet)
	}
}
