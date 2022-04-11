package robotInCircle

import (
	"fmt"
	"os"
)

func isRobotBounded(instructions string) bool {
	steps := []rune(instructions)
	x := 0
	y := 0
	dx := 0
	dy := 1
	for i := 0; i < len(steps); i++ {
		switch steps[i] {
		case 'G':
			x += dx
			y += dy
		case 'L':
			dx, dy = -dy, dx
		case 'R':
			dx, dy = dy, -dx
		default:
			_, _ = fmt.Fprintf(os.Stderr, "wrong instruction %q (pos. %d)", string(steps[i]), i)
			return false
		}
	}
	return (x == 0 && y == 0) || !(dx == 0 && dy == 1)
}
