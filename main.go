package main

import (
	B202401 "github/CeGenreDeChat/adventofcode/internal/_2024/_01"
	B202402 "github/CeGenreDeChat/adventofcode/internal/_2024/_02"
	B202403 "github/CeGenreDeChat/adventofcode/internal/_2024/_03"
	B202404 "github/CeGenreDeChat/adventofcode/internal/_2024/_04"
	B202405 "github/CeGenreDeChat/adventofcode/internal/_2024/_05"
	B202406 "github/CeGenreDeChat/adventofcode/internal/_2024/_06"
	B202407 "github/CeGenreDeChat/adventofcode/internal/_2024/_07"
	B202408 "github/CeGenreDeChat/adventofcode/internal/_2024/_08"
	B202409 "github/CeGenreDeChat/adventofcode/internal/_2024/_09"
	B202410 "github/CeGenreDeChat/adventofcode/internal/_2024/_10"
	B202411 "github/CeGenreDeChat/adventofcode/internal/_2024/_11"
	"github/CeGenreDeChat/adventofcode/internal/algorithms"
	"os"
)

func main() {
	var graph algorithms.Graph

	graph.Init()
	graph.ReadColumnFromFile("resources/2024/01.txt")
	B202401.Body_2024_01(&graph)

	graph.Init()
	graph.ReadLineFromFile("resources/2024/02.txt")
	B202402.Body_2024_02(&graph)

	B202403.Body_2024_03(&graph)

	B202404.Body_2024_04(&graph)

	B202405.Body_2024_05(&graph)

	B202406.Body_2024_06(&graph)

	B202407.Body_2024_07(&graph)

	B202408.Body_2024_08(&graph)

	B202409.Body_2024_09(&graph)

	B202410.Body_2024_10(&graph)

	B202411.Body_2024_11(&graph)

	os.Exit(0)
}
