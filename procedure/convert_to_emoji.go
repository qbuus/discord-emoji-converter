package procedure

import (
	"emoji-converter/logger"
	"emoji-converter/utils"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/nfnt/resize"
)

func ConvertToEmoji() {
	var path string // path to file

	title := "Enter path to image: "
	description := "Leave empty to return to menu"
	placeholder := "C:\\Users\\you\\Pictures\\emoji.png"

	err := huh.NewInput().
		Title(title).
		Description(description).
		Placeholder(placeholder).
		Prompt("?").
		Value(&path).
		Run()
	if err != nil {
		logger.ErrorF("Error with input: %s", err.Error())
		return
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		logger.Error("File does not exist. Returning to menu")
		return
	}

	if !utils.IsSupportedFormat(path) {
		logger.Error("File format not supported. Returning to menu")
		return
	}

	logger.Info("✅ File exists and format is valid. Ready for conversion.")

	file,err := os.Open(path)
	if err != nil {
		logger.ErrorF("Error opening file: %s", err.Error())
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		logger.ErrorF("Error decoding image: %s", err.Error())
		return
	}

	resizedImg := resize.Thumbnail(128,128, img, resize.Lanczos3)

	home,err := os.UserHomeDir()
	if err != nil {
		logger.ErrorF("Error getting home directory: %s", err.Error())
		return
	}

	pathToSave := filepath.Join(home, "Downloads")

	if _, err := os.Stat(pathToSave); os.IsNotExist(err) {
		logger.Error("Downloads folder not found, saving to current directory instead")
		pathToSave = "."
	}

	base := filepath.Base(path)
	ext := filepath.Ext(base)
	name := strings.TrimSuffix(base, ext)

	outFileName := filepath.Join(pathToSave, name+"_emoji.png")

	outFile, err := os.Create(outFileName)
	if err != nil {
		logger.ErrorF("Error creating file: %s", err.Error())
		return
	}
	defer outFile.Close()

	err = png.Encode(outFile, resizedImg)
	if err != nil {
		logger.ErrorF("Failed to save resized image: %s", err.Error())
		return
	}

	logger.Info(fmt.Sprintf("✅ Emoji saved as: %s", outFileName))
}