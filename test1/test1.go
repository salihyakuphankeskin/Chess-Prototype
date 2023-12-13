package test1

import (
	"fmt"
)
func Runner()  {
	Test1()	
}

func Test1() {
	// Create an 8x8 matrix
	var chessboard [8][8]string

	// Create a map to associate matrix block names with their indices
	blockIndices := make(map[string][2]int)

	// Initialize the matrix and map
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			blockName := fmt.Sprintf("%c%d", 'A'+j, i+1)
			blockIndices[blockName] = [2]int{i, j}
		}
	}

	// Function to set the value of a matrix block by its name
	setValue := func(name string, value string) {
		indices := blockIndices[name]
		chessboard[indices[0]][indices[1]] = value
	}

	// Example: Set the value of "B5" to "Pawn"
	setValue("B5", "Pawn")

	// Print the chessboard
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Printf("%-5s", chessboard[i][j])
		}
		fmt.Println()
	}
}
