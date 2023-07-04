package utils_test

import (
	"api-book/internal/infrastructure/utils"
	"regexp"
	"testing"
)

func TestGetPort(t *testing.T) {
	port := utils.GetPort()

	_, err := regexp.MatchString(`:(\d{4})`, port)
	if err != nil {
		t.Error("No se recibio un puerto valido")
		t.Fail()
	} else {
		t.Log("Puerto encontrado!")
	}

}
