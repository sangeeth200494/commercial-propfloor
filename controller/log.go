package controller

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	file, err := os.OpenFile("loggers/InfoLogs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 007)
	if err != nil {
		log.Fatal(err)
		//fmt.Println("info logger executed")
	}

	file1, err := os.OpenFile("loggers/WarningLogs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 007)
	if err != nil {
		log.Fatal(err)
		//fmt.Println("code not executed")
	}

	file2, err := os.OpenFile("loggers/ErrorLogs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 007)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO:", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file1, "WARNING ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file2, "ERROR ", log.Ldate|log.Ltime|log.Lshortfile)
}

func ReadFileInfo() {
	fmt.Printf("\n\nReading a file in loggers/InfoLogs\n")
	FileName := "InfoLogs.txt"

	// The ioutil package contains inbuilt methods like ReadFile that reads the filename and returns the contents.
	data, err := ioutil.ReadFile("loggers/InfoLogs.txt")
	if err != nil {
		ErrorLogger.Panicf("Failed Reading Data From InfoLogs: %s", err)
	}
	fmt.Printf("\nFile Name: %s", FileName)   // Prints the value as plain string
	fmt.Printf("\nSize: %d bytes", len(data)) // Base 10
	fmt.Printf("\nData: %s", data)
}

func ReadFileWarning() {
	fmt.Printf("\n\nReading a file in loggers/WarningLogs\n")
	FileName := "WarningLogs.txt"

	// The ioutil package contains inbuilt methods like ReadFile that reads the filename and returns the contents.
	data, err := ioutil.ReadFile("loggers/WarningLogs.txt")
	if err != nil {
		ErrorLogger.Panicf("Failed Reading Data From WarningLogs: %s", err)
	}
	fmt.Printf("\nFile Name: %s", FileName)   // Prints the value as plain string
	fmt.Printf("\nSize: %d bytes", len(data)) // Base 10
	fmt.Printf("\nData: %s", data)
}

func ReadFileError() {
	fmt.Printf("\n\nReading a file in loggers/ErrorLogs\n")
	FileName := "ErrorLogs.txt"

	// The ioutil package contains inbuilt methods like ReadFile that reads the filename and returns the contents.
	data, err := ioutil.ReadFile("loggers/ErrorLogs.txt")
	if err != nil {
		ErrorLogger.Panicf("Failed Reading Data From ErrorLogs: %s", err)
	}
	fmt.Printf("\nFile Name: %s", FileName)   // Prints the value as plain string
	fmt.Printf("\nSize: %d bytes", len(data)) // Base 10
	fmt.Printf("\nData: %s", data)
}
