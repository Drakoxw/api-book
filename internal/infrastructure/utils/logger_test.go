package utils

import (
	"testing"
)

type dog struct {
	Name   string
	IsGood bool
}

func TestStructToString(t *testing.T) {
	dog := dog{
		"Rex",
		false,
	}
	stringDog := structToString("titulo test", dog)
	if len(stringDog) > 0 {
		t.Log("Conversion satisfactoria!")
	} else {
		t.Error("No se ha podido convertir la estructura a cadena")
		t.Fail()
	}
}

func TestAnalizePathFile(t *testing.T) {
	result := analizePathFile("testing")
	expectedResult := "tmp/testing.log"
	if result == expectedResult {
		t.Logf("el resultado es satisfactorio")
	} else {
		t.Errorf("La ruta no cumple la estructura: %s", result)
		t.Fail()
	}
}

func TestLogSave(t *testing.T) {
	filepath := analizePathFile("testing")
	_, err := logSave(filepath, "info para guardar en el Log")
	if err != nil {
		t.Errorf("Hubo un error: %s", err.Error())
	} else {
		t.Log("El archivo fue generado correctamente.")
	}
}

func TestLoggerWidthData(t *testing.T) {
	dog := dog{
		"Rex",
		false,
	}
	LogWarningData("testing", "titulo", dog)
	LogInfoData("testing", "titulo", dog)
	LogErrorData("testing", "titulo", dog)
	t.Log("Se creo los logs sin problemas")
}

func TestLoggerWidthoutData(t *testing.T) {
	LogWarning("testing", "informacion en texto")
	LogInfo("testing", "informacion en texto")
	LogError("testing", "informacion en texto")
	t.Log("Se creo los logs sin problemas")
}
