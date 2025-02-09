package B202402

import (
	"fmt"
	"github/CeGenreDeChat/adventofcode/internal/algorithms"
)

func Body_2024_02(graph *algorithms.Graph) {
	var nbOfSafeReports1, nbOfSafeReports2 int = 0, 0

	for x := range graph.Values {
		levels := graph.GetListOfIntLine(x)
		if isLevelSafe(levels) {
			nbOfSafeReports1 += 1
		} else if isLevelSafeWithDeleteion(levels) {
			nbOfSafeReports2 += 1
		}

	}

	fmt.Printf("Total n°1 : %d\n", nbOfSafeReports1)
	fmt.Printf("Total n°2 : %d\n", nbOfSafeReports1+nbOfSafeReports2)
}

func isLevelSafeWithDeleteion(level []int) bool {
	for i := 0; i < len(level); i++ {
		CopyLevel := make([]int, len(level))
		copy(CopyLevel, level)
		if i == len(level)-1 {
			CopyLevel = CopyLevel[:i]
		} else {
			CopyLevel = append(CopyLevel[:i], level[i+1:]...)
		}
		if isLevelSafe(CopyLevel) {
			return true
		}
	}
	return false
}

func isLevelSafe(level []int) bool {
	var up, down bool = false, false
	for i := 1; i < len(level); i++ {
		diff := level[i] - level[i-1]
		if diff > 0 {
			up = true
		} else if diff < 0 {
			down = true
		} else {
			return false
		}

		if up && down {
			return false
		}

		if diff > 3 || diff < -3 {
			return false
		}
	}
	return true
}
