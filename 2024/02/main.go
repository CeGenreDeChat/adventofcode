package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("inputs")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var nbOfSafeReports int = 0

	var lineIndex int = 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lineIndex++
		levels := strings.Split(scanner.Text(), " ")
		l := len(levels)
		var level = make([]int, l)
		for i := 0; i < l; i++ {
			level[i], _ = strconv.Atoi(levels[i])
		}

		if isLevelSafe(level) {
			fmt.Printf("%d : Levels %v is safe\n", lineIndex, levels)
			nbOfSafeReports++
		} else if isLevelSafeWithDeleteion(level) {
			fmt.Printf("%d : Levels %v is safe with deletion\n", lineIndex, levels)
			nbOfSafeReports++
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total n°1 : %d\n", nbOfSafeReports)
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
