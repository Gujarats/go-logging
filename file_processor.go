package logging

import (
	"bytes"
	"os"
	"path/filepath"
)

func SaveMessageToFile(message string) {
	// create a file in the current dir
	dir := currentDir()
	loggerFileName := getFileName()
	file := createLogFile(dir, loggerFileName)

	// append the message to a file
	_, err := file.WriteString(message + "\n")
	if err != nil {
		panic(err)
	}
}

// return absolute path in the current directory
func currentDir() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	dir := filepath.Dir(ex)

	return dir + "/"
}

// Get filename for file logger name
func getFileName() string {
	return "logger.log"
}

func removeExtension(filename string) string {
	dotIndex := bytes.Index([]byte(filename), []byte(`.`))
	return string(filename[:dotIndex])
}

func createLogFile(filePath, fileName string) *os.File {
	f, err := os.OpenFile(filePath+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		// error file is not created
		err = os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			panic(err)
		}
		createLogFile(filePath, fileName)
	}

	return f
}
