package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func adjacentCombos(combo []byte) [][]byte {
	result := make([][]byte, 0)
	for n := 0; n < len(combo); n++ {
		comboUp := make([]byte, len(combo), len(combo))
		comboDown := make([]byte, len(combo), len(combo))
		copy(comboUp, combo)
		copy(comboDown, combo)
		comboUp[n] = (comboUp[n] + 1) % 10
		comboDown[n] = (10 + comboDown[n] - 1) % 10
		result = append(result, comboUp)
		result = append(result, comboDown)
	}
	return result
}

func toInt(combo []byte) int {
	val := 0
	for n := 0; n < len(combo); n++ {
		val = (val)*10 + int(combo[n])
	}
	return val
}

func fromInt(val int, length int) []byte {
	var combo []byte
	for val > 9 {
		digit := val % 10
		val /= 10
		combo = append([]byte{byte(digit)}, combo...)
	}
	combo = append([]byte{byte(val)}, combo...)
	for len(combo) < length {
		combo = append([]byte{0}, combo...)
	}
	return combo
}

func numStepsInt(combo []byte, trappedCombos map[int]struct{}) int {
	var queue [][]byte
	visited := make(map[int]struct{})
	queue = append(queue, []byte{0, 0, 0, 0})
	limit := 10*len(combo) + 1
	distance := 0
	for len(queue) > 0 {
		size := len(queue)
		distance++
		if distance > limit {
			break
		}
		for i := 0; i < size; i++ {
			c := queue[0]
			queue = queue[1:]
			if _, ok := visited[toInt(c)]; ok {
				continue
			}
			visited[toInt(c)] = struct{}{}
			for _, n := range adjacentCombos(c) {
				if _, ok := trappedCombos[toInt(n)]; ok {
					continue
				}
				if _, ok := visited[toInt(n)]; ok {
					continue
				}
				if bytes.Equal(n, combo) {
					return distance
				}
				queue = append(queue, n)
			}
		}
	}
	return -1
}

func numSteps(combo string, trappedCombos []string) int {
	// WRITE YOUR BRILLIANT CODE HERE
	comboInt, err := strconv.Atoi(combo)
	if err != nil {
		return -1
	}
	byteCombo := fromInt(comboInt, len(combo))
	tCombos := make(map[int]struct{})
	for tc := range trappedCombos {
		tComboInt, err := strconv.Atoi(trappedCombos[tc])
		if err != nil {
			return -1
		}
		tCombos[tComboInt] = struct{}{}
	}
	return numStepsInt(byteCombo, tCombos)
}

func splitWords(s string) []string {
	if s == "" {
		return []string{}
	}
	return strings.Split(s, " ")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	combo := scanner.Text()
	scanner.Scan()
	trappedCombos := splitWords(scanner.Text())
	res := numSteps(combo, trappedCombos)
	fmt.Println(res)
}
