package procedure

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

func TestPrintValue() {
	fmt.Println("Test Print Value")

	var stringToPrint string
	title := "Enter a string to print: "
	description := "Leave empty to return to menu"

	err := huh.NewInput().
		Title(title).
		Description(description).
		Prompt("?").
		Value(&stringToPrint).
		Run()
	
		if err != nil {
			fmt.Println("Error with input", err.Error())
			return
		}

		if stringToPrint == "" {
			fmt.Println("No input. Returning to menu")
			return
		}

		fmt.Println("Entered Value: ", stringToPrint)
}