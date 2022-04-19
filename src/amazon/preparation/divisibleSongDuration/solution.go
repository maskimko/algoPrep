package divisibleSongDuration

func numPairsDivisibleBy60(time []int) int {
	counter := 0
	for i := 0; i < len(time)-1; i++ {
		for k := i + 1; k < len(time); k++ {
			if (time[i]+time[k])%60 == 0 {
				counter++
			}
		}
	}
	return counter
}
