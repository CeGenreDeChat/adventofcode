package B202404

import (
	"bufio"
	"fmt"
	"github/CeGenreDeChat/adventofcode/internal/algorithms"
	"log"
	"os"
)

func Body_2024_04(graph *algorithms.Graph) {
	file, err := os.Open("resources/2024/04.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var matrix [140][140]byte

	masSize := 140

	index, total, xtotal := 0, 0, 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for jndex, c := range line {
			matrix[index][jndex] = byte(c)
		}
		index++
	}

	for i := 0; i < masSize; i++ {
		for j := 0; j < masSize; j++ {
			if i < masSize-3 && j < masSize-3 {
				//fmt.Printf("horiz: ")
				if isWord("XMAS", matrix[i][j], matrix[i][j+1], matrix[i][j+2], matrix[i][j+3]) {
					total++
				}
				//fmt.Printf("diago: ")
				if isWord("XMAS", matrix[i][j], matrix[i+1][j+1], matrix[i+2][j+2], matrix[i+3][j+3]) {
					total++
				}
				//fmt.Printf("rdiag: ")
				if isWord("XMAS", matrix[i][j+3], matrix[i+1][j+2], matrix[i+2][j+1], matrix[i+3][j]) {
					total++
				}
				//fmt.Printf("verti: ")
				if isWord("XMAS", matrix[i][j], matrix[i+1][j], matrix[i+2][j], matrix[i+3][j]) {
					total++
				}
				//fmt.Printf("--------------\n")
			} else if i < masSize-3 && j >= masSize-3 {
				//fmt.Printf("verti: ")
				if isWord("XMAS", matrix[i][j], matrix[i+1][j], matrix[i+2][j], matrix[i+3][j]) {
					total++
				}
				//fmt.Printf("--------------\n")
			} else if j < masSize-3 {
				//fmt.Printf("horiz: ")
				if isWord("XMAS", matrix[i][j], matrix[i][j+1], matrix[i][j+2], matrix[i][j+3]) {
					total++
				}
				//fmt.Printf("--------------\n")
			}

			//X-MAS
			if i < masSize-2 && j < masSize-2 {
				//fmt.Printf("diago:  ")
				if isXWord("MAS", matrix[i][j], matrix[i+1][j+1], matrix[i+2][j+2]) {
					//fmt.Printf("rdiago: ")
					if isXWord("MAS", matrix[i][j+2], matrix[i+1][j+1], matrix[i+2][j]) {
						xtotal++
					}
				}
				//fmt.Printf("--------------\n")
			}

		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("XMAS trouvé : %d\n", total)
	fmt.Printf("X-MAS trouvé : %d\n", xtotal)
}

func isWord(word string, a, b, c, d byte) bool {
	//fmt.Printf("%s %s %s %s", string(a), string(b), string(c), string(d))
	if word[0] == a && word[1] == b && word[2] == c && word[3] == d || word[0] == d && word[1] == c && word[2] == b && word[3] == a {
		//fmt.Println(" - find")
		return true
	}
	//fmt.Println("")
	return false
}

func isXWord(word string, a, b, c byte) bool {
	//fmt.Printf("%s %s %s", string(a), string(b), string(c))
	if word[0] == a && word[1] == b && word[2] == c || word[0] == c && word[1] == b && word[2] == a {
		//fmt.Println(" - find")
		return true
	}
	//fmt.Println("")
	return false
}
