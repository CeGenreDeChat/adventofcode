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

func main() {
	file, err := os.Open("inputs")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var nbOfCorrect, sumOfCorrect int = 0, 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ":")

		fmt.Printf("split : %v\n", split)

		getValues := strings.Split(strings.TrimSpace(split[1]), " ")
		sizeValues := len(getValues)
		var values []int = make([]int, sizeValues)
		for i := 0; i < sizeValues; i++ {
			s := getValues[i]
			v, _ := strconv.Atoi(s)
			values[i] = v
		}
		//fmt.Printf("len(values) : %d\n", len(values))
		var nbOperations = int(math.Pow(2, float64(len(values)-1)))
		var nbOperands = nbOperations / 2

		//fmt.Printf("Nombre d'opération et d'opérandes : %d, %d\n", nbOperations, nbOperands)

		operations := make([][]string, 0)

		var generateOperations func([]string, int, int)
		generateOperations = func(current []string, index int, remaining int) {
			if remaining == 0 {
				operations = append(operations, append([]string(nil), current...))
				return
			}
			current[index] = "+"
			generateOperations(current, index+1, remaining-1)
			current[index] = "*"
			generateOperations(current, index+1, remaining-1)
			current[index] = "||"
			generateOperations(current, index+1, remaining-1)
		}

		generateOperations(make([]string, nbOperands), 0, len(values)-1)

		for _, operation := range operations {
			//fmt.Printf("Operation : %v\n", operation)
			var result int = values[0]
			for i := 0; i < len(values)-1; i++ {
				if operation[i] == "+" {
					result += values[i+1]
				} else if operation[i] == "*" {
					result *= values[i+1]
				} else {
					left := strconv.Itoa(result)
					right := strconv.Itoa(values[i+1])
					result, _ = strconv.Atoi(left + right)
				}
			}
			//fmt.Printf("Result : %d\n", result)
			if r, _ := strconv.Atoi(split[0]); r == result {
				nbOfCorrect++
				sumOfCorrect += result
				//fmt.Println("valide")
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Equation correctes : %d\nSommes des équations correctes : %d\n", nbOfCorrect, sumOfCorrect)
}
