package main

import (
	"fmt"
	"strconv"
	"sideModules/blockcreate"
)

func main() {
	type chessBoardPlaces struct{
		horizontal string
		vertical int
	}

	chessPieces :=[6]string{
		"Pawn", 
		"Knight", 
		"Bishop", 
		"Rook", 
		"Queen", 
		"King",
	}

	for i := range chessPieces {
		fmt.Printf( strconv.Itoa(i)+"th " +chessPieces[i]+ "\n")
	}
	fmt.Print(chessPieces)

	// Create an empty 8x8 map
	grid := make(map[string]map[int]string)
	// Initialize the map
	for i := 'A'; i <= 'H'; i++ {
		columnLabel := string(i)
		grid[columnLabel] = make(map[int]string)

		for j := 1; j <= 8; j++ {
			grid[columnLabel][j] = fmt.Sprintf("%s%d", columnLabel, j)
		}
	}
	fmt.Println("")
	// Print the grid
	for i := 1; i <= 8; i++ {
		for j := 'A'; j <= 'H'; j++ {
			columnLabel := string(j)
			fmt.Printf("%s ", grid[columnLabel][i])
		}
		fmt.Println("")
	}	
		
	blockcreate.Runner()
}