package B202406

import (
	"bufio"
	"fmt"
	"github/CeGenreDeChat/adventofcode/internal/algorithms"
	"log"
	"os"
	"strings"
)

type Position struct {
	x int
	y int
}

const SIZE = 130

const DEBUG = false
const TRACE = false

func Body_2024_06(graph *algorithms.Graph) {
	file, err := os.Open("resources/2024/06.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var _map [SIZE][SIZE]rune
	var nbPos int = SIZE * SIZE
	var nbPlaces int = 0
	var nbLoop int = 0
	var i int = 0
	var GuardPosInit Position

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for j := 0; j < len(_map[i]); j++ {
			_map[i][j] = rune(scanner.Text()[j])
			if DEBUG {
				fmt.Printf("%c", _map[i][j])
			}
			if _map[i][j] == '^' {
				GuardPosInit = Position{x: i, y: j}
			}
		}
		if DEBUG {
			fmt.Println("")
		}
		i++
	}

	if DEBUG {
		fmt.Println("")
	}

	nbMoves := calculPath(&_map, GuardPosInit, nbPos, nbPlaces, false, true)
	printMap(&_map)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var arrayOfX []Position = make([]Position, nbPos)

	for i := 0; i < len(_map); i++ {
		for j := 0; j < len(_map[i]); j++ {
			if _map[i][j] == 'X' {
				arrayOfX[nbPlaces].x = i
				arrayOfX[nbPlaces].y = j
				_map[i][j] = '.'
				nbPlaces++
			}
		}
	}

	if DEBUG {
		printMap(&_map)
	}

	fmt.Printf("Nombre de déplacement du gardien : %v\n", nbMoves)
	fmt.Printf("Nombre d'endroits visité par le gardien : %v\n", nbPlaces)

	for _, g := range arrayOfX {
		if g.x == GuardPosInit.x && g.y == GuardPosInit.y {
			continue
		}
		old := _map[g.x][g.y]
		_map[g.x][g.y] = 'O'
		if calculPath(&_map, GuardPosInit, nbPlaces, nbPlaces, true, false) == -1 {
			_map[g.x][g.y] = old
			nbLoop++
			continue
		}
		_map[g.x][g.y] = old
	}

	fmt.Printf("Nombre de chemin bouclant : %v\n", nbLoop)
}

func calculPath(_map *[SIZE][SIZE]rune, GuardPos Position, nbPos, nbPlaces int, look bool, trace bool) int {
	var knowPaths map[Position]string = make(map[Position]string, nbPlaces)
	var nbMoves int = 0
	for i := 0; i < nbPos; i++ {
		if !moveUp(_map, &GuardPos, &nbMoves, trace) {
			break
		}
		if look {
			if _, ok := knowPaths[GuardPos]; ok {
				if strings.Contains(knowPaths[GuardPos], "U") {
					return -1
				}
				knowPaths[GuardPos] += "U"
			} else {
				knowPaths[GuardPos] = "U"
			}
		}
		if DEBUG {
			printMap(_map)
		}
		if !moveRight(_map, &GuardPos, &nbMoves, trace) {
			break
		}
		if look {
			if _, ok := knowPaths[GuardPos]; ok {
				if strings.Contains(knowPaths[GuardPos], "R") {
					return -1
				}
				knowPaths[GuardPos] += "R"
			} else {
				knowPaths[GuardPos] = "R"
			}
		}
		if DEBUG {
			printMap(_map)
		}
		if !moveDown(_map, &GuardPos, &nbMoves, trace) {
			break
		}
		if look {
			if _, ok := knowPaths[GuardPos]; ok {
				if strings.Contains(knowPaths[GuardPos], "D") {
					return -1
				}
				knowPaths[GuardPos] += "D"
			} else {
				knowPaths[GuardPos] = "D"
			}
		}
		if DEBUG {
			printMap(_map)
		}
		if !moveLeft(_map, &GuardPos, &nbMoves, trace) {
			break
		}
		if look {
			if _, ok := knowPaths[GuardPos]; ok {
				if strings.Contains(knowPaths[GuardPos], "L") {
					return -1
				}
				knowPaths[GuardPos] += "L"
			} else {
				knowPaths[GuardPos] = "L"
			}
		}
		if DEBUG {
			printMap(_map)
		}
	}

	if DEBUG {
		printMap(_map)
	}
	return nbMoves
}

func moveUp(m *[SIZE][SIZE]rune, p *Position, M *int, trace bool) bool {
	m[p.x][p.y] = '.'
	*M += p.x
	for p.x >= 0 && m[p.x][p.y] != '#' && m[p.x][p.y] != 'O' {
		if trace {
			m[p.x][p.y] = 'X'
		}
		p.x--
	}
	p.x++
	*M -= p.x
	if p.x-1 < 0 {
		if trace {
			m[p.x][p.y] = 'X'
		}
	} else {
		m[p.x][p.y] = '>'
	}
	return p.x-1 >= 0
}

func moveRight(m *[SIZE][SIZE]rune, p *Position, M *int, trace bool) bool {
	m[p.x][p.y] = '.'
	*M -= p.y
	for p.y < len(m[p.y]) && m[p.x][p.y] != '#' && m[p.x][p.y] != 'O' {
		if trace {
			m[p.x][p.y] = 'X'
		}
		p.y++
	}
	p.y--
	*M += p.y
	if p.y+1 == len(m) {
		if trace {
			m[p.x][p.y] = 'X'
		}
	} else {
		m[p.x][p.y] = 'v'
	}
	return p.y+1 >= 0
}

func moveDown(m *[SIZE][SIZE]rune, p *Position, M *int, trace bool) bool {
	m[p.x][p.y] = '.'
	*M -= p.x
	for p.x < len(m) && m[p.x][p.y] != '#' && m[p.x][p.y] != 'O' {
		if trace {
			m[p.x][p.y] = 'X'
		}
		p.x++
	}
	p.x--
	*M += p.x
	if p.x+1 == len(m[p.x]) {
		if trace {
			m[p.x][p.y] = 'X'
		}
	} else {
		m[p.x][p.y] = '<'
	}
	return p.x+1 < len(m)
}

func moveLeft(m *[SIZE][SIZE]rune, p *Position, M *int, trace bool) bool {
	m[p.x][p.y] = '.'
	*M += p.y
	for p.y >= 0 && m[p.x][p.y] != '#' && m[p.x][p.y] != 'O' {
		if trace {
			m[p.x][p.y] = 'X'
		}
		p.y--
	}
	p.y++
	*M -= p.y
	if p.y-1 < 0 {
		if trace {
			m[p.x][p.y] = 'X'
		}
	} else {
		m[p.x][p.y] = '^'
	}
	return p.y-1 < len(m[p.x])
}

func printMap(m *[SIZE][SIZE]rune) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			fmt.Printf("%c", m[i][j])
		}
		fmt.Println("")
	}
	fmt.Println("")
}
