package B202401

import (
	"fmt"
	"github/CeGenreDeChat/adventofcode/internal/algorithms"
	"sort"
)

func Body_2024_01(graph *algorithms.Graph) {
	var nbOfRights map[int]int = graph.GetNumberOfElementInColumn(1)

	left := graph.GetListOfIntColumn(0)
	right := graph.GetListOfIntColumn(1)

	sort.Ints(left)
	sort.Ints(right)

	var total1, total2 int64 = 0, 0

	for i := 0; i < len(left); i++ {
		if addition := left[i] - right[i]; addition < 0 {
			total1 += int64(-addition)
		} else {
			total1 += int64(addition)
		}

		if _, ok := nbOfRights[left[i]]; ok {
			total2 += int64(left[i]) * int64(nbOfRights[left[i]])
		}
	}

	fmt.Printf("Total n°1 : %d\nTotal n°2 : %d\n", total1, total2)
}
