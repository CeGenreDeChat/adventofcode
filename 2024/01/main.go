package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("inputs")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var left []int = make([]int, 1001)
	var right []int = make([]int, 1001)

	var nbOfRights map[int]int = make(map[int]int)

	index := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "   ")
		left[index], _ = strconv.Atoi(split[0])
		right[index], _ = strconv.Atoi(split[1])
		if _, ok := nbOfRights[right[index]]; !ok {
			nbOfRights[right[index]] = 1
		} else {
			nbOfRights[right[index]]++
		}
		index++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(left)
	sort.Ints(right)

	var total1, total2 int64 = 0, 0

	for i := 0; i < len(left); i++ {
		if addition := left[i] - right[i]; addition < 0 {
			total1 += int64(-addition)
		} else {
			total1 += int64(addition)
		}

		if _, ok := nbOfRights[left[i]]; ok {
			total2 += int64(left[i]) * int64(nbOfRights[left[i]])
		}
	}

	fmt.Printf("Total n°1 : %d\nTotal n°2 : %d\n", total1, total2)
}
