package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const NBOFSTEPS = 75

const DEBUG = false

type Separate struct {
	left  int
	right int
}

func main() {
	file, err := os.Open("inputs")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var stonesList map[int]int = make(map[int]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, c := range strings.Split(scanner.Text(), " ") {
			value, _ := strconv.Atoi(c)
			stonesList[value] = 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var sum int = 0

	if DEBUG {
		printMap(stonesList)
	}

	for i := 0; i < NBOFSTEPS; i++ {
		var new_stonesList map[int]int = make(map[int]int)
		for key, nb := range stonesList {
			for _, new_stone := range calcul(key) {
				new_stonesList[new_stone] += nb
			}
		}
		stonesList = new_stonesList
		if DEBUG {
			printMap(stonesList)
		}
	}

	for _, nb := range stonesList {
		sum += nb
	}

	fmt.Printf("Nombre de pierres : %d\n", sum)
}

func calcul(input int) []int {
	var output []int

	if input == 0 {
		output = append(output, 1)
	} else if len := intLengthMath(input); len%2 == 0 {
		left, right := splitNumber(input, len)
		output = append(output, left)
		output = append(output, right)
	} else {
		output = append(output, input*2024)
	}

	return output
}

func intLengthMath(n int) int {
	length := 0
	temp := n
	for temp > 0 {
		temp /= 10
		length++
	}
	return length
}

func splitNumber(n, s int) (int, int) {
	power := int(math.Pow(10, float64(s/2)))

	part1 := n / power
	part2 := n % power

	return part1, part2
}

func printMap(m map[int]int) {
	fmt.Printf("Pierres : ")
	for key := range m {
		fmt.Printf("%v ", key)
	}
	fmt.Printf("\n")
}
