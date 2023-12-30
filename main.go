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
var scanner = bufio.NewScanner(os.Stdin)
var chessBoard = make(map[string]map[string]string)
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
func main() {
	fmt.Print("\tHELLO WELCOME TO THE CHESS GAME IN GOLANG\n\n")

	//Creating a String x String => String Matrix
	chessBoard = chessBoardCreationLoop(chessBoard)
	userSideInput := sideSelection(chessBoard)
	printMap_String_String(userSideInput, chessBoard)
	gameLoop()



	fmt.Println("")
	fmt.Println("")	
	
	Empty(chessPieces,chessBoardCreationLoop,startingProgram, userSideInput)
}


func Empty(...any) any {

	return 0
}
func EmptyModule()  {
	blockcreate.Runner()
	test1.Runner()	
}
func printMap_String_String(myUserSideInput string,myMap map[string]map[string]string) {	
	defaultCallOuts:= func(verticalLines string,horizontalLines int){
		rowKey := verticalLines
		colKey := fmt.Sprintf("%d", horizontalLines)
		value := myMap[rowKey][colKey]
		fmt.Printf("%s%s: %s | ", rowKey, colKey, value)
	}

	if (myUserSideInput == "white" || 
		myUserSideInput == "w" || 
		myUserSideInput == "wh" ){
			fmt.Printf("WHITE SIDE \n")
			for  horizontalLines := 8; horizontalLines >= 1; horizontalLines--{
				for verticalLines := 'A'; verticalLines <= 'H'; verticalLines++  {
					defaultCallOuts(string(verticalLines), horizontalLines)
				}
				fmt.Printf("\n")
			}
	} else 
	if(	myUserSideInput == "black" || 
		myUserSideInput == "b" || 
		myUserSideInput == "bl" ){	
			fmt.Printf("BLACK SIDE \n")	
			for  horizontalLines := 1; horizontalLines <= 8; horizontalLines++{
				for verticalLines := 'H'; verticalLines >= 'A'; verticalLines--  {
					defaultCallOuts(string(verticalLines), horizontalLines)
				}
				fmt.Printf("\n")
			}
	}
}

func chessBoardCreationLoop(myMap map[string]map[string]string) map[string]map[string]string {
	for horizontal := 'A'; horizontal <= 'H'; horizontal++ {
		rowMatrix := string(horizontal)
		myMap[rowMatrix] = make(map[string]string)
		for vertical := 1; vertical <= 8; vertical++ {
			columnMatrix := fmt.Sprintf("%d", vertical)
			myMap[rowMatrix][columnMatrix] = "  "
		}
	}
	return myMap	
}

func startingProgram(){
}
func userInputInteraction(messageToGive string) string{
	fmt.Printf("%s ", messageToGive)
	scanner.Scan()
	userSideInput := scanner.Text()
	return userSideInput
}

func sideSelection(myMap map[string]map[string]string) string{
	for {
		messageForUser:= "Enter 'White' or 'Black': "
		userSideInput:= strings.ToLower(userInputInteraction(messageForUser))
		fmt.Print("\n")
		if err := chessBoardSetUp(userSideInput, myMap); err == nil {
			if  (
				userSideInput == "white"	|| 		userSideInput == "w" || 
				userSideInput == "wh"){
					return userSideInput
							
			}else if (
				userSideInput == "black" || userSideInput == "b" || 
				userSideInput == "bl"){
					return userSideInput
				}
			break
		} else {
			fmt.Println("Error:", err)
		}
	}
	return " "
}

func chessBoardSetUp(selectedSide string, myMap map[string]map[string]string)  error {
	if  !(
		selectedSide == "white"	|| 		selectedSide == "w" 	|| 
		selectedSide == "wh"	||		selectedSide == "black" || 		
		selectedSide == "b" 	|| 		selectedSide == "bl" 	){
		return errors.New("invalid input. Please enter either 'White' or 'Black'")
	}

	myMap["A"]["1"] = chessPieces["White Rook"]
		myMap["H"]["1"] = chessPieces["White Rook"]
		myMap["B"]["1"] = chessPieces["White Knight"]
		myMap["G"]["1"] = chessPieces["White Knight"]
		myMap["C"]["1"] = chessPieces["White Bishop"]
		myMap["F"]["1"] = chessPieces["White Bishop"]
		myMap["D"]["1"] = chessPieces["White Queen"]
		myMap["E"]["1"] = chessPieces["White King"]
		
		myMap["A"]["2"] = chessPieces["White Pawn"] 
		myMap["B"]["2"] = chessPieces["White Pawn"] 
		myMap["C"]["2"] = chessPieces["White Pawn"] 
		myMap["D"]["2"] = chessPieces["White Pawn"] 
		myMap["E"]["2"] = chessPieces["White Pawn"] 
		myMap["F"]["2"] = chessPieces["White Pawn"] 
		myMap["G"]["2"] = chessPieces["White Pawn"] 
		myMap["H"]["2"] = chessPieces["White Pawn"]

		myMap["A"]["8"] = chessPieces["Black Rook"]
		myMap["H"]["8"] = chessPieces["Black Rook"]
		myMap["B"]["8"] = chessPieces["Black Rook"]
		myMap["G"]["8"] = chessPieces["Black Knight"]
		myMap["C"]["8"] = chessPieces["Black Knight"]
		myMap["F"]["8"] = chessPieces["Black Bishop"]
		myMap["D"]["8"] = chessPieces["Black Queen"]
		myMap["E"]["8"] = chessPieces["Black King"]
		
		myMap["A"]["7"] = chessPieces["Black Pawn"] 
		myMap["B"]["7"] = chessPieces["Black Pawn"] 
		myMap["C"]["7"] = chessPieces["Black Pawn"] 
		myMap["D"]["7"] = chessPieces["Black Pawn"] 
		myMap["E"]["7"] = chessPieces["Black Pawn"] 
		myMap["F"]["7"] = chessPieces["Black Pawn"] 
		myMap["G"]["7"] = chessPieces["Black Pawn"] 
		myMap["H"]["7"] = chessPieces["Black Pawn"]
		return nil	
}

func playerMovementValidation()  string{
	messageToGive :="Your move: "
	userQuestion := userInputInteraction(messageToGive)
	for {
		if len(userQuestion) != 4  {
			continue							
		}	
		firstTwo:= userQuestion[:2]
		if _, ok := chessBoard[firstTwo]; !ok {
			continue
		}
		firstTwo= userQuestion[len(userQuestion)-2:]
		if _, ok := chessBoard[firstTwo]; !ok {
			continue
		}
		return userQuestion
	}
}
func gameLoop()  {
	playerOutput:= playerMovementValidation()	
	pieceLocation := playerOutput[:2]
	pieceNextLocation := playerOutput[len(playerOutput)-2:]

}