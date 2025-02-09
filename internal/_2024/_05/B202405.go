package B202405

import (
	"bufio"
	"fmt"
	"github/CeGenreDeChat/adventofcode/internal/algorithms"
	"log"
	"os"
	"strconv"
	"strings"
)

type rule struct {
	page   int
	before int
}

type update struct {
	pages []int
}

func Body_2024_05(graph *algorithms.Graph) {
	file, err := os.Open("resources/2024/05.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var rules []rule = make([]rule, 1200)
	var updates []update = make([]update, 300)

	var isPages map[int]string = make(map[int]string, 1200)
	var isBefores map[int]string = make(map[int]string, 1200)

	var isUpdate bool = false

	var maxRule, maxUpdate int = 0, 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			isUpdate = true
		} else if !isUpdate {
			split := strings.Split(scanner.Text(), "|")
			page, _ := strconv.Atoi(split[0])
			before, _ := strconv.Atoi(split[1])
			rules[maxRule].page = page
			rules[maxRule].before = before
			if _, ok := isPages[page]; !ok {
				isPages[page] = split[1]
			} else {
				isPages[page] += "," + split[1]
			}

			if _, ok := isBefores[before]; !ok {
				isBefores[before] = split[0]
			} else {
				isBefores[before] += "," + split[0]
			}
			maxRule++
		} else if isUpdate {
			split := strings.Split(scanner.Text(), ",")
			updates[maxUpdate].pages = make([]int, len(split))
			for index, s := range split {
				p, _ := strconv.Atoi(s)
				updates[maxUpdate].pages[index] = p
			}
			maxUpdate++
		}
	}

	//fmt.Printf("%v", isUpdates)

	var good, last bool = false, true
	var middles, total int = 0, 0

	for i := 0; i < len(updates); i++ {
		var middle int = 0
		if updates[i].pages == nil {
			break
		}
		//fmt.Printf("[")
	Line:
		for j := 0; j < len(updates[i].pages); j++ {
			middle = returnMiddle(updates[i].pages)
			// si updates[i].pages[j] est avant tout les suivants de la liste
		Row:
			for o := j + 1; o < len(updates[i].pages); o++ {
				last = false

				page := updates[i].pages[j]
				before := updates[i].pages[o]
				if _, ok := isPages[page]; ok {
					b := strings.Split(isPages[page], ",")

					for q := 0; q < len(b); q++ {
						B, _ := strconv.Atoi(b[q])
						if before == B {
							good = true
							continue Row
						}
					}
					break Row
				}
			}
			if good {
				//fmt.Printf("%d,", updates[i].pages[j])
				good = false
				last = true
			} else if last {
				//fmt.Printf("%d", updates[i].pages[j])
				middles += middle
				good = false
				last = true
			} else {
				//fmt.Printf("X...")
				total++
				break Line
			}
		}
		//fmt.Printf("]\n")
	}

	var middleCorrect int = 0

	for i := 0; i < len(updates); i++ {
		if updates[i].pages == nil {
			break
		}

		middleCorrect += returnMiddle(correctList(updates[i].pages, isPages))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Règles trouvées : %d\n", maxRule)
	fmt.Printf("Mise à jour trouvées : %d\n", maxUpdate)

	fmt.Printf("Règle fausses : %d\n", total)
	fmt.Printf("Sommes des milieux : %d\n", middles)
	fmt.Printf("Sommes des milieux corrigés : %d\n", middleCorrect-middles)
}

func returnMiddle(list []int) int {
	return list[len(list)/2]
}

func correctList(pages []int, isPage map[int]string) []int {
	l := len(pages)

	var re []int = make([]int, l)
	var tmp []int = make([]int, l)

	copy(re, pages)

Start:
	copy(tmp, re)
	for i := 0; i < l-1; i++ {
		page := re[i]
		before := re[i+1]
		if !strings.Contains(isPage[page], strconv.Itoa(before)) {
			re[i] = tmp[i+1]
			re[i+1] = tmp[i]
			goto Start
		}
	}

	return re
}
