package main

func minMutation(startGene string, endGene string, bank []string) int {

	q := make([]string, 0, 32)
	vis := make(map[string]struct{}, 32)
	q = append(q, startGene)
	vis[startGene] = struct{}{}

	res := 0
	for len(q) > 0 {

		n := len(q)
		for i := 0; i < n; i++ { // 注意这里不要写成 for len(q) > 0, 因为这里的q在for中会append
			g := q[0]
			q = q[1:]
			if g == endGene {
				return res
			}

			for _, v := range bank {
				_, ok := vis[v]
				if distance(g, v) == 1 && !ok {
					q = append(q, v)
					vis[v] = struct{}{}
				}
			}
		}
		res++
	}

	return -1

}

func distance(a, b string) int {

	cnt := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			cnt++
		}
	}
	return cnt
}
