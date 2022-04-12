package autoscalePolicyUtilizationCheck

func autoScale(utilization []int, instances int) int {
	var k = instances
	maxInstances := 2*10 ^ 8
	for i := 0; i < len(utilization); i++ {
		if utilization[i] < 25 && k > 1 {
			k = (k + 1) / 2
			i += 9
		}
		if utilization[i] > 60 && k < maxInstances {
			k *= 2
			if k > maxInstances {
				k = maxInstances
			}
			i += 9
		}
	}
	return k
}
