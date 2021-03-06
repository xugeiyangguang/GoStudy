package unitTestExcercise

func canCompleteCircuit(gas []int, cost []int) int {
	tank := 0
	for i := 0; i < len(gas); i++ {
		tank = gas[i]
		for j := 1; j <= len(gas); j++ {
			pos := (i + j) % len(gas)

			if pos != 0 {
				tank = tank - cost[pos-1]
			} else {
				tank = tank - cost[len(gas)-1]
			}
			if tank < 0 {
				break
			}
			if pos == i {
				return i
			}
			tank = tank + gas[pos]
		}
	}
	return -1
}
