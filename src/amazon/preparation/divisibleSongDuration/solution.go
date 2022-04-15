package divisibleSongDuration

import "sort"

func checkPermutations(time []int, j, k int) int {
	flags := make([]bool, len(time))
	left := time[j]
	right := time[k]
	flags[j] = true
	flags[k] = true
	jEnd := j
	kStart := k
	for i := j + 1; i < k; i++ {
		if time[i] == left {
			flags[i] = true
			jEnd = i
		}
	}
	for n := k - 1; n > j; n-- {
		if time[n] == right {
			flags[n] = true
			kStart = n
		}
	}
	if time[k] == time[j] {
		return (k - j) * (k - j + 1) / 2

	}
	return (jEnd - j + 1) * (k - kStart + 1)
}

func numPairsDivisibleBy60(time []int) int {
	reminders := make([]int, len(time), len(time))
	for i := 0; i < len(time); i++ {
		reminders[i] = time[i] % 60
	}
	sort.Ints(reminders)
	pairs := 0
	k := len(reminders) - 1
	for i := 0; i < k; {
		if (reminders[i]+reminders[k])%60 == 0 {
			pairs += checkPermutations(reminders, i, k)

			for ; i < k && i < len(reminders)-1 && reminders[i+1] == reminders[i]; i++ {
			}
			for ; i < k-1 && k > 0 && reminders[k-1] == reminders[k]; k-- {
			}
			i++
			k--
			continue
		}
		if reminders[i]+reminders[k] < 60 {
			i++
		} else {
			k--
		}
	}
	return pairs
}
