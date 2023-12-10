package main

import (
	"fmt"
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
	
	// Create an empty 8x8 map
	grid := make(map[string]map[int]string)
	// Initialize the map
	for i := 'A'; i <= 'H'; i++ {
		columnLabel := string(i)
		grid[columnLabel] = make(map[int]string)

		for j := 1; j <= 8; j++ {
			grid[columnLabel][j] = ""
		}
	}

	fmt.Println(grid["A"][3])
	fmt.Println(grid["B"][7])
	fmt.Println(grid["H"][5])
	
	// Print the grid
	for i := 1; i <= 8; i++ {
		for j := 'A'; j <= 'H'; j++ {
			columnLabel := string(j)
			fmt.Printf("%s ", grid[columnLabel][i])
		}
		fmt.Println("")
	}		
	fmt.Println(grid)

		

	
	blockcreate.Empty()
	Empty(chessPieces)
}

func Empty(...any) any {

	return 0
}
func seperate(){
	
	fmt.Println("--------------------")
}