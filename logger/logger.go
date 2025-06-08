package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	pathOfFile *os.File
	InfoLogger *log.Logger
	ErrorLogger *log.Logger
	WarningLogger *log.Logger
)

func init() {
	year,month,day := time.Now().Date()

	err := os.MkdirAll("logger/logs", 0755)
	if err != nil {
		fmt.Println("[ERROR] creating logs directory:", err)
		return
	}

	file,err := filepath.Abs("logger/logs/" + fmt.Sprintf("%d-%d-%d.log", year, month, day))
	if err != nil {
		fmt.Println("[ERROR] getting abs file", err)
		return
	}

	path, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("[ERROR] opening file", err)
		return
	}

	pathOfFile = path
	InfoLogger = log.New(pathOfFile, "[INFO]: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(pathOfFile, "[ERROR]: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(pathOfFile, "[WARNING]: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Info(info ...any) {
	InfoLogger.Println(info...)
	fmt.Println(info...)
}

func InfoF(format string, info ...any) {
	InfoLogger.Printf(format, info...)
	fmt.Printf(format, info...)
}

func Warning(info ...any) {
	WarningLogger.Println(info...)
	fmt.Println(info...)
}

func WarningF(format string, info ...any) {
	WarningLogger.Printf(format, info...)
	fmt.Printf(format, info...)
}

func Error(info ...any) {
	ErrorLogger.Println(info...)
	fmt.Println(info...)
}

func ErrorF(format string, info ...any) {
	ErrorLogger.Printf(format, info...)
	fmt.Printf(format, info...)
}