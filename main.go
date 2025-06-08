package main

import (
	"emoji-converter/procedure"
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
)

func main() {
	for {
		fmt.Println("Welcome to Emoji Converter")

		var Option string

		form := huh.NewSelect[string]().
			Title("Select an option").Options(
				huh.NewOption("Test - Print value", "test_print_value"),
				huh.NewOption("Exit", "exit"),
			).Value(&Option)

		err := form.Run()
		if err != nil {
			fmt.Println("Error with selection", err.Error())
			continue
		}

		switch Option {
		case "test_print_value":
			procedure.TestPrintValue()	
		case "exit":
			fmt.Println("Goodbye")
			os.Exit(0)
		default:
			fmt.Println("Invalid option")
		} 
	}
}