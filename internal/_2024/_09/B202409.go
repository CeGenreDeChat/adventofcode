package B202409

import (
	"bufio"
	"fmt"
	"github/CeGenreDeChat/adventofcode/internal/algorithms"
	"log"
	"os"
	"strconv"
)

type Position struct {
	start int
	len   int
}

const SIZE = 50

const DEBUG = false
const MOVEBLOCK = true

func Body_2024_09(graph *algorithms.Graph) {
	file, err := os.Open("resources/2024/09.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var data []int = nil
	var open []string = nil
	var sorted []string = nil

	var isFile bool = true
	var id int = 0
	var idLen []int = nil
	var spaceLen []int = nil
	var sum int = 0
	var posFiles, posSpaces []Position = nil, nil

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = make([]int, len(scanner.Text()))
		if DEBUG {
			fmt.Printf("data   : %s\n", scanner.Text())
		}
		for index, c := range scanner.Text() {
			data[index] = int(c - '0')
		}
	}

	for _, d := range data {
		lenId := 0
		lenSpace := 0
		for i := 0; i < d; i++ {
			if isFile {
				open = append(open, strconv.Itoa(id))
				lenId++
			} else {
				open = append(open, ".")
				lenSpace++
			}
		}
		if isFile {
			idLen = append(idLen, lenId)
			id++
		} else {
			spaceLen = append(spaceLen, lenSpace)
		}
		isFile = !isFile
	}

	if DEBUG {
		fmt.Printf("open   : %s\n", open)
	}

	var isFileStart bool = true
	var isSpaceStart bool = true

	index_f := -1
	index_s := -1

	id = 0

	for i := 0; i < len(open); i++ {
		if open[i] != "." {
			isSpaceStart = true
			if isFileStart || i > 0 && (open[i] != open[i-1]) {
				posFiles = append(posFiles, Position{i, 1})
				index_f++
				isFileStart = false
			} else {
				posFiles[index_f].len++
			}
		} else {
			isFileStart = true
			if isSpaceStart {
				posSpaces = append(posSpaces, Position{i, 1})
				index_s++
				id++
				isSpaceStart = false
			} else {
				posSpaces[index_s].len++
			}
		}
	}

	sorted = make([]string, len(open))

	copy(sorted, open)

	ff := 0

	if MOVEBLOCK {
		indexSpace := 0
		for indexFile := len(posFiles) - ff - 1; indexFile >= 0; indexFile-- {
			spaces := calculSpace(sorted)
			for _, space := range spaces {
				if space.start > posFiles[indexFile].start {
					break
				}
				if ifSizeEnought(posFiles[indexFile], space) == 0 {
					moveFiletoSpace(sorted, posFiles[indexFile], space)
					indexSpace += space.len
				}
			}
			if DEBUG {
				fmt.Printf("sorted : %s\n", sorted)
			}
		}
	} else {
		u := 0
		for i := len(open) - 1; i > 0; i-- {
			if open[i] != "." {
				for j := u; j < len(sorted); j++ {
					if sorted[j] == "." {
						sorted[j] = open[i]
						sorted[i] = "."
						break
					} else if i == j {
						break
					}
				}
			}
		}
	}

	if DEBUG {
		fmt.Printf("sorted : %s\n", sorted)
	}

	for index, c := range sorted {
		if c == "." {
			continue
		} else {
			v, _ := strconv.Atoi(c)
			sum += index * v
		}
	}

	fmt.Printf("La somme de contrôle vaut %d\n", sum)
}

func ifSizeEnought(file, space Position) int {
	if file.len <= space.len {
		return 0
	}
	return -1
}

func moveFiletoSpace(sorted []string, file, space Position) {
	for i := 0; i < file.len; i++ {
		sorted[space.start+i] = sorted[file.start+i]
		sorted[file.start+i] = "."
	}
}

func calculSpace(sorted []string) []Position {
	var posSpaces []Position = nil
	var isSpaceStart bool = true
	index_s := -1
	for i := 0; i < len(sorted); i++ {
		if sorted[i] == "." {
			if isSpaceStart {
				posSpaces = append(posSpaces, Position{i, 1})
				index_s++
				isSpaceStart = false
			} else {
				posSpaces[index_s].len++
			}
		} else {
			isSpaceStart = true
		}
	}
	return posSpaces
}
