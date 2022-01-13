package arraySubsets

func Subsets(input []interface{}) [][][]interface{} {
	if len(input) == 0 {
		return nil
	}
	if len(input) == 1 {
		return [][][]interface{}{{input}}
	}
	return getArraySubsets(input)[1:]
}

func getArraySubsets(input []interface{}) [][][]interface{} {
	var solutions [][][]interface{}
	if len(input) == 0 {
		return nil
	}
	solutions = append(solutions, [][]interface{}{input})
	for div := 1; div < len(input); div++ {
		suffix := getArraySubsets(input[div:])
		for _, subSolution := range suffix {
			if len(subSolution) > 0 {
				solutions = append(solutions, append([][]interface{}{input[:div]}, subSolution...))
			}
		}
	}
	return solutions
}
