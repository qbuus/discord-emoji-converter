package main

import (
	"emoji-converter/logger"
	"emoji-converter/procedure"
	"os"

	"github.com/charmbracelet/huh"
)

func main() {
	for {
		logger.Info("Welcome to Emoji Converter")

		var Option string

		form := huh.NewSelect[string]().
			Title("Select an option").Options(
				huh.NewOption("Test - Print value", "test_print_value"),
				huh.NewOption("Exit", "exit"),
			).Value(&Option)

		err := form.Run()
		if err != nil {
			logger.ErrorF("Error with selection: %s", err.Error())
			continue
		}

		switch Option {
		case "test_print_value":
			procedure.TestPrintValue()	
		case "exit":
			logger.Info("Goodbye")
			os.Exit(0)
		default:
			logger.Info("Invalid option")
		} 
	}
}