package printer

import "fmt"

func Runner()  {
	Printer()
	
}

func Printer() {

	// Print function
	fmt.Print("This is a ")
	fmt.Print("single line ")
	fmt.Print("with Print.\n")

	// Printf function
	name := "John"
	age := 30
	fmt.Printf("My name is %s and I am %d years old.\n", name, age)

	// Println function
	fmt.Println("This is a")
	fmt.Println("multi-line")
	fmt.Println("with Println.")

	// Sprint function
	sprintResult := fmt.Sprint("Sprint ", "concatenates ", "strings.")
	fmt.Println(sprintResult)

	// Sprintf function
	sprintfResult := fmt.Sprintf("My name is %s and I am %d years old.", name, age)
	fmt.Println(sprintfResult)
}