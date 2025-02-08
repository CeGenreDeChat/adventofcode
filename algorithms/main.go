package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	x     int
	y     int
	value rune
}

type Neighbor struct {
	x int
	y int
}

type Queue struct {
	items []Node
}

func (q *Queue) EnQueue(item *Node) {
	q.items = append(q.items, *item)
}

func (q *Queue) DeQueue() (*Node, error) {
	if len(q.items) == 0 {
		return nil, fmt.Errorf("La queue est vide")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return &item, nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func main() {
	var isStart, isStop bool
	var start, stop Node

	isStart, isStop = false, false

	nodes, maxX, maxY, _ := readMapFromFile("resources/map.txt")

	printMap(nodes)

Loop:
	for x := range *nodes {
		for y, node := range (*nodes)[x] {
			if node.value != '.' {
				if node.value == 'A' {
					start = Node{x: x, y: y, value: node.value}
					isStart = true
					if isStart && isStop {
						break Loop
					}
				} else if node.value == 'B' {
					stop = Node{x: x, y: y, value: node.value}
					isStop = true
					if isStart && isStop {
						break Loop
					}
				}
			}
		}
	}

	breadthFirstSearch(nodes, maxX, maxY, &start, &stop, 'M')

	fmt.Print("\n")

	printMap(nodes)

	os.Exit(0)
}

func readMapFromFile(pathFile string) (*[][]Node, int, int, error) {
	var array [][]Node
	var line string
	var file *os.File
	var err error
	var x, y int

	if file, err = os.Open(pathFile); err != nil {
		return nil, 0, 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	x, y = 0, 0

	for scanner.Scan() {
		if line = scanner.Text(); line == "" {
			break
		} else {
			array = append(array, []Node{})
			y = 0
		}
		for _, symbol := range line {
			array[x] = append(array[x], Node{x: x, y: y, value: symbol})
			y += 1
		}
		x += 1
	}

	if err = scanner.Err(); err != nil {
		return nil, 0, 0, err
	}

	return &array, x, y, nil
}

func printMap(nodes *[][]Node) {
	for _, x := range *nodes {
		for _, y := range x {
			fmt.Printf("%c", y.value)
		}
		fmt.Print("\n")
	}
}

func getNeighbors(node *Node, maxX, maxY int) []Node {
	var dirs = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	var neighbors []Node
	var x, y int

	for _, dir := range dirs {
		x = node.x + dir[0]
		y = node.y + dir[1]
		if 0 <= x && x < maxX && 0 <= y && y < maxY {
			neighbors = append(neighbors, Node{x: x, y: y, value: 0})
		}
	}

	return neighbors
}

func breadthFirstSearch(nodes *[][]Node, maxX, maxY int, start, stop *Node, except rune) []Node {
	var reached map[Node]Node
	var frontier Queue
	var current Node
	var path []Node

	frontier = Queue{}
	reached = make(map[Node]Node)

	reached[*start] = Node{}

	frontier.EnQueue(start)

	for !frontier.IsEmpty() {
		current, _ := frontier.DeQueue()
		for _, next := range getNeighbors(current, maxX, maxY) {
			next.value = (*nodes)[next.x][next.y].value
			if _, exist := reached[next]; !exist && next.value != except {
				frontier.EnQueue(&(*nodes)[next.x][next.y])
				reached[next] = *current
				(*nodes)[next.x][next.y].value = 'x'
			}
		}
	}

	current = (*nodes)[stop.x][stop.y]
	path = []Node{}

	for current.x != start.x || current.y != start.y {
		path = append(path, current)
		current = reached[current]
		(*nodes)[current.x][current.y].value = '.'
	}

	return path
}
