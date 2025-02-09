package algorithms

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PriorityQueue []*Item

type Item struct {
	Value    *Node
	priority int
	index    int
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

type Graph struct {
	Values [][]Node
	queue  []Node
	maxX   int
	maxY   int
	start  *Node
	stop   *Node
}

type Node struct {
	x     int
	y     int
	Value rune
	Word  string
	cost  int
}

func (g *Graph) Init() {
	g.Values = [][]Node{}
	g.queue = []Node{}
}

func (g *Graph) GetStartPoint(symbol rune) {
	for x := range g.Values {
		for y := range g.Values[x] {
			if g.Values[x][y].Value == symbol {
				g.start = &g.Values[x][y]
				g.start.cost = 0
				break
			}
		}
	}
}

func (g *Graph) GetStopPoint(symbol rune) {
	for x := range g.Values {
		for y := range g.Values[x] {
			if g.Values[x][y].Value == symbol {
				g.stop = &g.Values[x][y]
				break
			}
		}
	}
}

func (g *Graph) PrintMap() {
	for x := range g.Values {
		for y := range g.Values[x] {
			fmt.Printf("%c", g.Values[x][y].Value)
		}
		fmt.Print("\n")
	}
}

func (g *Graph) PrintLine() {
	for x := range g.Values {
		fmt.Printf("%s\n", g.Values[x][0].Word)
	}
}

func (g *Graph) PrintColumn() {
	for x := range g.Values {
		for y := range g.Values[x] {
			fmt.Printf("%s ", g.Values[x][y].Word)
		}
		fmt.Print("\n")
	}
}

func (g *Graph) GetNeighbors(posX, posY int) [][]int {
	var neighbors [][]int
	var x, y int

	for _, dir := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		x = posX + dir[0]
		y = posY + dir[1]
		if 0 <= x && x < g.maxX && 0 <= y && y < g.maxY {
			neighbors = append(neighbors, []int{x, y})
		}
	}

	return neighbors
}

func (g *Graph) ReadLineFromFile(pathFile string) error {
	var line string
	var file *os.File
	var index int
	var err error

	if file, err = os.Open(pathFile); err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	index = 0

	for scanner.Scan() {
		if line = scanner.Text(); line == "" {
			break
		}

		g.Values = append(g.Values, []Node{})
		for _, word := range strings.Fields(line) {
			g.Values[index] = append(g.Values[index], Node{Word: word})
		}

		index += 1
	}

	return nil
}

func (g *Graph) ReadColumnFromFile(pathFile string) error {
	var line string
	var file *os.File
	var index int
	var err error

	if file, err = os.Open(pathFile); err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	index = 0

	for scanner.Scan() {
		if line = scanner.Text(); line == "" {
			break
		}

		g.Values = append(g.Values, []Node{})

		for _, word := range strings.Fields(line) {
			g.Values[index] = append(g.Values[index], Node{Word: word})
		}

		index += 1
	}

	return nil
}

func (g *Graph) ReadRuneFromFile(pathFile string) error {
	var line string
	var file *os.File
	var err error
	var x, y int

	if file, err = os.Open(pathFile); err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	x, y = 0, 0

	for scanner.Scan() {
		if line = scanner.Text(); line == "" {
			break
		} else {
			g.Values = append(g.Values, []Node{})
			y = 0
		}
		for _, symbol := range line {
			g.Values[x] = append(g.Values[x], Node{x: x, y: y, Value: symbol})
			y += 1
		}
		x += 1
	}

	if err = scanner.Err(); err != nil {
		return err
	}

	g.maxX = x
	g.maxY = y

	return nil
}

func (g *Graph) ReadCostFromFile(pathFile string) error {
	var line string
	var file *os.File
	var err error
	var x, y int

	if file, err = os.Open(pathFile); err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	x, y = 0, 0

	for scanner.Scan() {
		if line = scanner.Text(); line == "" {
			break
		} else {
			g.Values = append(g.Values, []Node{})
			y = 0
		}
		for _, symbol := range line {
			g.Values[x] = append(g.Values[x], Node{x: x, y: y, Value: symbol, cost: (int(symbol) - '0')})
			y += 1
		}
		x += 1
	}

	if err = scanner.Err(); err != nil {
		return err
	}

	g.maxX, g.maxY = x, y

	return nil
}

func (g *Graph) EnQueue(item *Node) {
	g.queue = append(g.queue, *item)
}

func (g *Graph) DeQueue() (Node, error) {
	if len(g.queue) == 0 {
		return Node{}, fmt.Errorf("la queue est vide")
	}
	item := g.queue[0]
	g.queue = g.queue[1:]
	return item, nil
}

func (g *Graph) LenQueue() int {
	return len(g.queue)
}

func (g *Graph) GetCost(x, y int) int {
	return g.Values[x][y].cost
}

func (g *Graph) BreadthFirstSearch(except rune) {
	var cameFrom map[Node]Node
	var current Node

	cameFrom = make(map[Node]Node)

	cameFrom[*g.start] = Node{}

	g.EnQueue(g.start)

	for g.LenQueue() > 0 {
		current, _ = g.DeQueue()

		if current.x == g.stop.x && current.y == g.stop.y {
			break
		}

		for _, pos := range g.GetNeighbors(current.x, current.y) {
			next := (g.Values)[pos[0]][pos[1]]
			if _, exist := cameFrom[next]; !exist && next.Value != except {
				g.EnQueue(&g.Values[pos[0]][pos[1]])
				cameFrom[next] = current
			}
		}
	}

	for current.x != g.start.x || current.y != g.start.y {
		current = (cameFrom[current])
		g.Values[current.x][current.y].Value = '.'
	}
}

func (g *Graph) Dijkstra(except int) {
	var current *Node
	pq := &PriorityQueue{}
	heap.Init(pq)

	cameFrom := make(map[*Node]*Node)
	costSoFar := make(map[Node]int)
	cameFrom[g.start] = &Node{}
	costSoFar[*g.start] = 0

	heap.Push(pq, &Item{Value: g.start, priority: 0})

	for pq.Len() > 0 {
		current = heap.Pop(pq).(*Item).Value

		if current.x == g.stop.x && current.y == g.stop.y {
			break
		}

		for _, pos := range g.GetNeighbors(current.x, current.y) {

			if g.Values[pos[0]][pos[1]].cost == except {
				continue
			}

			newCost := costSoFar[*current] + g.GetCost(pos[0], pos[1])

			if _, exist := costSoFar[g.Values[pos[0]][pos[1]]]; !exist || newCost < costSoFar[g.Values[pos[0]][pos[1]]] {
				costSoFar[g.Values[pos[0]][pos[1]]] = newCost
				heap.Push(pq, &Item{Value: &g.Values[pos[0]][pos[1]], priority: newCost})
				cameFrom[&g.Values[pos[0]][pos[1]]] = current
			}
		}
	}

	for current.x != g.start.x || current.y != g.start.y {
		current = cameFrom[current]
		g.Values[current.x][current.y].Value = '.'
	}
}

func (g *Graph) GetListOfIntColumn(nb int) []int {
	var list []int

	for x := range g.Values {
		value, _ := strconv.Atoi(g.Values[x][nb].Word)
		list = append(list, value)
	}

	return list
}

func (g *Graph) GetListOfIntLine(nb int) []int {
	var list []int

	for y := range g.Values[nb] {
		value, _ := strconv.Atoi(g.Values[nb][y].Word)
		list = append(list, value)
	}

	return list
}

func (g *Graph) GetNumberOfElementInColumn(nb int) map[int]int {
	var nbOfElement map[int]int = make(map[int]int)

	for x := range g.Values {
		value, _ := strconv.Atoi(g.Values[x][nb].Word)
		if _, ok := nbOfElement[value]; ok {
			nbOfElement[value] += 1
		} else {
			nbOfElement[value] = 1
		}
	}

	return nbOfElement
}
