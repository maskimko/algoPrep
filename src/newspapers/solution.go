package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func timeToRead(newspapers []int) int {
	sum := 0
	for i := range newspapers {
		sum += newspapers[i]
	}
	return sum
}

func newspapersSplit(newspapers []int, coworkers int) int {
	//Boundaries
	//max - sum all newspapers
	sumNewspapers := 0
	for _, n := range newspapers {
		sumNewspapers += n
	}
	max := sumNewspapers
	//min 1 item
	min := 1
	minTime := 1<<31 - 1
	midTime := -1
	for min <= max {
		mid := min + (max-min)/2
		midTime = time(newspapers, mid, coworkers)
		if midTime < minTime {
			minTime = midTime //Not needed if binsearch is correct
		}
		if mid == 1 {
			break //Just to stop the loop
		}
		preMidTime := midTime
		//lookup the previous
		for j := 0; preMidTime == midTime && mid > j; j++ {
			preMidTime = time(newspapers, mid-j, coworkers)
		}
		if preMidTime < minTime {
			minTime = preMidTime //Nod needed if binsearch is correct
		}
		if preMidTime == midTime {
			break
		}
		if midTime > preMidTime {
			max = mid - 1
		} else {
			min = mid + 1
		}
	}
	return midTime
}

func time(newspapers []int, amount int, coworkers int) int {
	times := make([]int, coworkers, coworkers)
	maxTime := 0
	qty := 0
	issue := 0

	for i := 0; i < len(newspapers); i += qty {
		if issue == coworkers-1 {
			rest := timeToRead(newspapers[i:])
			times[coworkers-1] = rest
			break
		}
		qty = assignNewspapers(newspapers[i:], amount)
		t := timeToRead(newspapers[i : i+qty])
		times[issue%coworkers] += t
		issue++
	}
	for i := range times {
		if times[i] > maxTime {
			maxTime = times[i]
		}
	}
	return maxTime
}

func assignNewspapers(newspapers []int, capacity int) int {
	i := 0
	for sum := 0; sum < capacity && i < len(newspapers); i++ {
		sum += newspapers[i]
	}
	return i
}

func arrayAtoi(arr []string) []int {
	res := []int{}
	for _, x := range arr {
		v, _ := strconv.Atoi(x)
		res = append(res, v)
	}
	return res
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
	newspapers := arrayAtoi(splitWords(scanner.Text()))
	scanner.Scan()
	coworkers, _ := strconv.Atoi(scanner.Text())
	res := newspapersSplit(newspapers, coworkers)
	fmt.Println(res)
}
