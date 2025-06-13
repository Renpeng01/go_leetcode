package main

// 超时
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	has := 0
	for i := 0; i < n; i++ {
		has = 0
		m := i
		for j := 0; j < n; j++ {
			if has+gas[m]-cost[m] >= 0 {
				has = has + gas[m] - cost[m]
				m = (m + 1) % n
				if j == n-1 {
					return i
				}
				continue
			}
			break
		}
	}
	return -1
}

func canCompleteCircuit1(gas []int, cost []int) int {
	n := len(gas)
	has := 0
	for i := 0; i < n; {
		has = 0
		m := i
		for j := 0; j < n; j++ {
			if has+gas[m]-cost[m] >= 0 {
				has = has + gas[m] - cost[m]
				m = (m + 1) % n
				if j == n-1 {
					return i
				}
				continue
			}

			// x走不到y，则x和y中间的任意一点都走不到y，所有从y+1开始
			i = i + j + 1 // 如果i不到i，则从i + j + 1开始继续往后走
			break
		}

	}
	return -1
}
