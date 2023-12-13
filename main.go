package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sideModules/blockcreate"
	"sideModules/test1"
	"strings"
)
func main() {

	//Creating a String x String => String Matrix
	chessBoard := make(map[string]map[string]string)
	chessBoard = chessBoardCreationLoop(chessBoard)
	sideSelection(chessBoard)
	printMap_String_String(chessBoard)

	fmt.Println("")
	fmt.Println("")	
	
	Empty(chessPieces,chessBoardCreationLoop,startingProgram)
}

var chessPieces = map[string]string{
	"White Pawn"	: "WP",
	"White Knight"	: "WK", 
	"White Bishop"	: "WB", 
	"White Rook"	: "WR", 
	"White Queen"	: "WQ", 
	"White King"	: "WK",

	"Black Pawn"	: "BP", 
	"Black Knight"	: "BK", 
	"Black Bishop"	: "BB", 
	"Black Rook"	: "BR", 
	"Black Queen"	: "BQ", 
	"Black King"	: "BK",
}

func Empty(...any) any {

	return 0
}
func EmptyModule()  {
	blockcreate.Runner()
	test1.Runner()	
}
func printMap_String_String(myMap map[string]map[string]string) {

	for verticalLines := 'A'; verticalLines <= 'H'; verticalLines++ {
		rowKey := string(verticalLines)
		for horizontalLines := 1; horizontalLines <= 8; horizontalLines++ {
			colKey := fmt.Sprintf("%d", horizontalLines)
			value := myMap[rowKey][colKey]
			fmt.Printf("%s %s: %s | ", rowKey, colKey, value)
		}
		fmt.Printf("\n")
	}
}

func chessBoardCreationLoop(myMap map[string]map[string]string) map[string]map[string]string {
	
	// for loop A, B, C, D, E, F, G, H
	for horizontal := 'A'; horizontal <= 'H'; horizontal++ {
		//for loop takes chars as integer value
		// So, we convert them onto strings
		rowMatrix := string(horizontal)
		// Create a vertical Maps
		myMap[rowMatrix] = make(map[string]string)

		// for loop 1, 2, 3, 4, 5, 6, 7, 8
		for vertical := 1; vertical <= 8; vertical++ {
			// Convert integer into string
			columnMatrix := fmt.Sprintf("%d", vertical)
			// Register values into givin [x][y]
			myMap[rowMatrix][columnMatrix] = "  "
		}
	}
	return myMap	
}

var chessPiecesDupricated =[12]string{
	"White Pawn",
	"White Knight", 
	"White Bishop", 
	"White Rook", 
	"White Queen", 
	"White King",

	"Black Pawn", 
	"Black Knight", 
	"Black Bishop", 
	"Black Rook", 
	"Black Queen", 
	"Black King",
}

func startingProgram(){
	fmt.Println("HELLO WELCOME TO THE CHESS GAME IN GOLANG")	
}
func sideSelection(myMap map[string]map[string]string)  {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter 'White' or 'Black': ")
		scanner.Scan()
		userInput := scanner.Text()

		if err := chessBoardSetUp(userInput, myMap); err == nil {
			break
		} else {
			fmt.Println("Error:", err)
		}
	}

}

func chessBoardSetUp(selectedSide string, myMap map[string]map[string]string)  error {
	selectedSide = strings.ToLower(selectedSide)
	if  (
		selectedSide == "white"	|| 		selectedSide == "w" || 
		selectedSide == "wh"	|| 		selectedSide == "1"){
			myMap["A"]["1"]= chessPieces["White Rook"]
			myMap["H"]["1"]= chessPieces["White Rook"]
			myMap["B"]["1"]= chessPieces["White Knight"]
			myMap["G"]["1"]= chessPieces["White Knight"]
			myMap["C"]["1"]= chessPieces["White Bishop"]
			myMap["F"]["1"]= chessPieces["White Bishop"]
			myMap["D"]["1"]= chessPieces["White Queen"]
			myMap["E"]["1"]= chessPieces["White King"]
			
			myMap["A"]["2"]= chessPieces["White Pawn"] 
			myMap["B"]["2"]= chessPieces["White Pawn"] 
			myMap["C"]["2"]= chessPieces["White Pawn"] 
			myMap["D"]["2"]= chessPieces["White Pawn"] 
			myMap["E"]["2"]= chessPieces["White Pawn"] 
			myMap["F"]["2"]= chessPieces["White Pawn"] 
			myMap["G"]["2"]= chessPieces["White Pawn"] 
			myMap["H"]["2"]= chessPieces["White Pawn"]
			return nil	
					
	}else if (
		selectedSide == "black" || selectedSide == "b" || 
		selectedSide == "bl" 	|| selectedSide == "0"){	
			myMap["A"]["8"] = chessPieces["Black Rook"]
			myMap["H"]["8"] = chessPieces["Black Rook"]
			myMap["B"]["8"] = chessPieces["Black Rook"]
			myMap["G"]["8"] = chessPieces["Black Knight"]
			myMap["C"]["8"] = chessPieces["Black Knight"]
			myMap["F"]["8"] = chessPieces["Black Bishop"]
			myMap["D"]["8"] = chessPieces["Black Queen"]
			myMap["E"]["8"] = chessPieces["Black King"]
			
			myMap["A"]["7"]= chessPieces["Black Pawn"] 
			myMap["B"]["7"]= chessPieces["Black Pawn"] 
			myMap["C"]["7"]= chessPieces["Black Pawn"] 
			myMap["D"]["7"]= chessPieces["Black Pawn"] 
			myMap["E"]["7"]= chessPieces["Black Pawn"] 
			myMap["F"]["7"]= chessPieces["Black Pawn"] 
			myMap["G"]["7"]= chessPieces["Black Pawn"] 
			myMap["H"]["7"]= chessPieces["Black Pawn"]
			return nil
	}else {
		return errors.New("Invalid input. Please enter either 'White' or 'Black'")
	}
	
}
