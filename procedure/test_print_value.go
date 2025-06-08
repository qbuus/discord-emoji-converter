package procedure

import (
	"emoji-converter/logger"
	"fmt"

	"github.com/charmbracelet/huh"
)

func TestPrintValue() {
	logger.Info("Test Print Value")

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
			logger.ErrorF("Error with input: %s", err.Error())
			return
		}

		if stringToPrint == "" {
			logger.Error("No input. Returning to menu")
			return
		}

		logger.Info("Entered Value: ", stringToPrint)
		fmt.Println("Entered Value: ", stringToPrint)
}