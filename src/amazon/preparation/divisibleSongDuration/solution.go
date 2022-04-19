package divisibleSongDuration

import "sort"

func numPairsDivisibleBy60(time []int) int {
	reminders := make([]int, len(time), len(time))
	for i := 0; i < len(time); i++ {
		reminders[i] = time[i] % 60
	}
	sort.Ints(reminders)
	counter := 0
	i := 0
	k := len(reminders) - 1

	for k > i {
		if reminders[i] == 0 {
			zi := i
			for zi < len(reminders) && reminders[zi] == 0 {
				zi++
			}
			counter += (zi - i) * (zi - i - 1) / 2
			i += zi
			continue
		}

		if (reminders[k]+reminders[i])%60 == 0 {

			if reminders[k] == reminders[i] {
				counter += (k - i + 1) * (k - i) / 2
				return counter
			}

			iDup := 1
			for reminders[i+1] == reminders[i] && i < k {
				iDup++
				i++
			}
			kDup := 1
			for reminders[k-1] == reminders[k] && i < k {
				kDup++
				k--
			}
			counter += iDup * kDup
			k--
			i++

		}

		if (reminders[k] + reminders[i]) < 60 {
			i++
			continue
		}
		if (reminders[k] + reminders[i]) > 60 {
			k--
			continue
		}
	}

	return counter
}
