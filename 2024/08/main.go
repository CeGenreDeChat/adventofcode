package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Position struct {
	x int
	y int
}

const SIZE = 50

const DEBUG = false
const HARMONIC = true

func main() {
	file, err := os.Open("inputs")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var _map [SIZE][SIZE]rune
	var i int = 0
	var antennasPos []Position
	var a map[rune][]Position = make(map[rune][]Position, SIZE*SIZE)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for j := 0; j < len(_map[i]); j++ {
			_map[i][j] = rune(scanner.Text()[j])
			if DEBUG {
				fmt.Printf("%c", _map[i][j])
			}
			if _map[i][j] != '.' {
				a[_map[i][j]] = append(a[_map[i][j]], Position{x: i, y: j})
				antennasPos = append(antennasPos, Position{x: i, y: j})
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

	var antipoles map[Position]string = make(map[Position]string, SIZE*SIZE)

	for k, v := range a {
		foundAntipoles(k, &_map, antipoles, v)
		for _, aa := range v {
			if _, ok := antipoles[aa]; !ok {
				antipoles[aa] = string(k)
			} else {
				antipoles[aa] += string(k)
			}
		}
	}

	fmt.Printf("Nombre d'antennes : %v\n", len(antennasPos))
	fmt.Printf("Nombre d'antipoles : %v\n", len(antipoles))
}

func foundAntipoles(key rune, m *[SIZE][SIZE]rune, antipoles map[Position]string, antennas []Position) {
	var sz int
	if HARMONIC {
		sz = SIZE
	} else {
		sz = 2
	}
	for i := 0; i < len(antennas); i++ {
		for j := i + 1; j < len(antennas); j++ {
			p := Position{x: antennas[i].x - antennas[j].x, y: antennas[i].y - antennas[j].y}
			for k := 1; k < sz; k++ {
				a1 := Position{x: antennas[i].x + p.x*k, y: antennas[i].y + p.y*k}
				if a1.x < SIZE && a1.y < SIZE && a1.x >= 0 && a1.y >= 0 && m[a1.x][a1.y] != '#' && m[a1.x][a1.y] != key {
					if m[a1.x][a1.y] == '.' {
						m[a1.x][a1.y] = '#'
						if DEBUG {
							printMap(m)
						}
					}
					if _, ok := antipoles[a1]; !ok {
						antipoles[a1] = string(key)
					} else {
						antipoles[a1] += string(key)
					}
				}
			}

			for k := 1; k < sz; k++ {
				a2 := Position{x: antennas[j].x - p.x*k, y: antennas[j].y - p.y*k}
				if a2.x < SIZE && a2.y < SIZE && a2.x >= 0 && a2.y >= 0 && m[a2.x][a2.y] != '#' && m[a2.x][a2.y] != key {
					if m[a2.x][a2.y] == '.' {
						m[a2.x][a2.y] = '#'
						if DEBUG {
							printMap(m)
						}
					}
					if _, ok := antipoles[a2]; !ok {
						antipoles[a2] = string(key)
					} else {
						antipoles[a2] += string(key)
					}
				}
			}
		}
	}
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
