package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Point struct {
	x int
	y int
}

type Node struct {
	Value int
	Pos   Point
	Next  []*Node
}

var directions = []Point{
	{0, 1}, {1, 0}, {0, -1}, {-1, 0},
}

const SIZE = 50

const DEBUG = true

func main() {
	file, err := os.Open("inputs")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var matrix [][]int = make([][]int, SIZE)

	var startPos []Point = nil

	i := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matrix[i] = make([]int, SIZE)
		for j, c := range scanner.Text() {
			matrix[i][j] = int(c - '0')
			if matrix[i][j] == 0 {
				startPos = append(startPos, Point{x: i, y: j})
			}
		}
		i++
	}

	paths := findPaths(matrix)

	if DEBUG {
		printPaths(paths)
	}

	lll := 0

	var prev Point

	var exist map[Point]bool = nil

	for i := 0; i < len(paths); i++ {
		root := paths[i][0]
		if prev == root {
			break
		}
		prev = root
		exist = make(map[Point]bool)
		for j := i; j < len(paths); j++ {
			if paths[j][0] == root {
				if _, ok := exist[paths[j][9]]; !ok {
					exist[paths[j][9]] = true
					lll++
				}
			} else {
				i += j - i - 1
				break
			}
		}
	}

	fmt.Printf("Taille des chemins : %d\n", lll)
}

func printTree(node *Node, level int) {
	if node == nil {
		return
	}
	fmt.Printf("%sNode(%d) at (%d, %d)\n", strings.Repeat("  ", level), node.Value, node.Pos.x, node.Pos.y)
	for _, child := range node.Next {
		printTree(child, level+1)
	}
}

func isValid(x, y, rows, cols int) bool {
	return x >= 0 && y >= 0 && x < rows && y < cols
}

// buildTree crée l'arbre à partir d'un point initial
func findPaths(matrix [][]int) [][]Point {
	rows := len(matrix)
	cols := len(matrix[0])

	var paths [][]Point
	var currentPath []Point

	// Fonction récursive DFS
	var dfs func(x, y int)
	dfs = func(x, y int) {
		currentPath = append(currentPath, Point{x, y})

		if matrix[x][y] == 9 {
			// Si on atteint 9, enregistrer le chemin
			pathCopy := make([]Point, len(currentPath))
			copy(pathCopy, currentPath)
			paths = append(paths, pathCopy)
		} else {
			// Explorer les voisins valant la valeur actuelle + 1
			for _, dir := range directions {
				nx, ny := x+dir.x, y+dir.y
				if isValid(nx, ny, rows, cols) && matrix[nx][ny] == matrix[x][y]+1 {
					dfs(nx, ny)
				}
			}
		}

		// Backtracking : retirer le dernier point
		currentPath = currentPath[:len(currentPath)-1]
	}

	// Lancer DFS à partir de tous les points valant 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 0 {
				dfs(i, j)
			}
		}
	}

	return paths
}

func printPaths(paths [][]Point) {
	if len(paths) == 0 {
		fmt.Println("Aucun chemin trouvé.")
		return
	}
	for i, path := range paths {
		fmt.Printf("Chemin %d : ", i)
		for _, p := range path {
			fmt.Printf("(%d,%d) -> ", p.x, p.y)
		}
		fmt.Println("FIN")
	}
}
