package B202403

import (
	"bufio"
	"fmt"
	"github/CeGenreDeChat/adventofcode/internal/algorithms"
	"log"
	"os"
	"strconv"
)

func Body_2024_03(graph *algorithms.Graph) {
	file, err := os.Open("resources/2024/03.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var total int64 = 0
	var do bool = true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		characters := []rune(scanner.Text())
		nb_l, nb_r := "", ""
		nl_s, nr_s := 0, 0
		m, u, l, q, p, nl, v, nr := false, false, false, false, false, false, false, false
		for index := 0; index < len(characters); index++ {

			if len(characters)-index > 3 && characters[index] == 'd' && characters[index+1] == 'o' && characters[index+2] == '(' && characters[index+3] == ')' {
				do = true
				index += 3
				continue
			}

			if len(characters)-index > 6 && characters[index] == 'd' && characters[index+1] == 'o' && characters[index+2] == 'n' && characters[index+3] == '\'' && characters[index+4] == 't' && characters[index+5] == '(' && characters[index+6] == ')' {
				do = false
				index += 6
				continue
			}

			if characters[index] == 'm' {
				m = true
			} else if m && characters[index] == 'u' {
				u = true
			} else if u && characters[index] == 'l' {
				l = true
			} else if l && characters[index] == '(' {
				q = true
			} else if q && !v && characters[index] >= '0' && characters[index] <= '9' && nl_s < 3 {
				nl = true
				nb_l = nb_l + string(characters[index])
				nl_s++
			} else if nl && characters[index] == ',' {
				v = true
			} else if v && characters[index] >= '0' && characters[index] <= '9' && nr_s < 3 {
				nr = true
				nb_r = nb_r + string(characters[index])
				nr_s++
			} else if nr && characters[index] == ')' {
				p = true
			} else {
				m, u, l, q, p, nl, v, nr, p = false, false, false, false, false, false, false, false, false
				nl_s, nr_s = 0, 0
				nb_l, nb_r = "", ""
			}

			if p {
				if do {
					left, _ := strconv.Atoi(nb_l)
					right, _ := strconv.Atoi(nb_r)
					total += int64(left) * int64(right)
				}
				nb_l, nb_r = "", ""
				m, u, l, q, p, nl, v, nr, p = false, false, false, false, false, false, false, false, false
				nl_s, nr_s = 0, 0
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total n°1 : %d\n", total)
}
