package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type saveData struct {
	Title string
	Data  interface{}
}

func structToString(title string, data interface{}) string {
	save := saveData{Title: title, Data: data}
	jsonBytes, _ := json.Marshal(save)
	return string(jsonBytes)
}

func analizePathFile(filePath string) string {
	if !strings.Contains(filePath, "tmp/") {
		filePath = fmt.Sprintf("%s %s", "tmp/", filePath)
	}
	if !strings.Contains(filePath, ".log") {
		filePath = fmt.Sprintf("%s %s", filePath, ".log")
	}
	return filePath
}

func logSave(filePath string, info string) (string, error) {
	absFilePath, err := filepath.Abs(filePath)
	if err != nil {
		return "", err
	}

	file, err := os.OpenFile(absFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return "", err
	}
	defer file.Close()

	log.SetOutput(file)
	log.Println(info)

	return absFilePath, nil
}

func LogInfo(filePath string, info string) {
	path := analizePathFile(filePath)
	infoComplete := fmt.Sprintf("[ INFO ] %s", info)
	logSave(path, infoComplete)
}

func LogError(filePath string, info string) {
	path := analizePathFile(filePath)
	infoComplete := fmt.Sprintf("[ ERROR ] %s", info)
	logSave(path, infoComplete)
}

func LogWarning(filePath string, info string) {
	path := analizePathFile(filePath)
	infoComplete := fmt.Sprintf("[ WARNING ] %s", info)
	logSave(path, infoComplete)
}

func LogInfoData(filePath string, title string, data interface{}) {
	info := structToString(title, data)
	LogInfo(filePath, info)
}

func LogErrorData(filePath string, title string, data interface{}) {
	info := structToString(title, data)
	LogError(filePath, info)
}

func LogWarningData(filePath string, title string, data interface{}) {
	info := structToString(title, data)
	LogWarning(filePath, info)
}
