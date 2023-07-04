package utils

import (
	"api-book/internal/domain/dtos"
	"encoding/json"
	"testing"
)

func TestCreateAwsResponse(t *testing.T) {
	expectedHeaders := map[string]string{
		"Content-Type":                 "application/json",
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Methods": "GET,HEAD,OPTIONS,POST",
	}

	response := CreateAwsResponse()

	if len(response.Headers) != len(expectedHeaders) {
		t.Errorf("Error al crear la respuesta de AWS, número incorrecto de headers. Esperado: %d, Obtenido: %d", len(expectedHeaders), len(response.Headers))
		t.Fail()
	}

	for key, value := range expectedHeaders {
		if response.Headers[key] != value {
			t.Errorf("Error al crear la respuesta de AWS, valor incorrecto para el header '%s'. Esperado: %s, Obtenido: %s", key, value, response.Headers[key])
			t.Fail()
		}
	}

	t.Log("Test success")
}

func TestErrorCreateAwsResponse(t *testing.T) {
	expectedHeaders := map[string]string{
		"Content-Type":                 "application",
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Methods": "GET,HEAD,OPTIONS,POST",
	}

	response := CreateAwsResponse()
	err := 0

	for key, value := range expectedHeaders {
		if response.Headers[key] != value {
			err++
		}
	}

	if err == 0 {
		t.Error("Se esperaba al menos un error")
		t.Fail()
	} else {
		t.Log("Test success")
	}
}

func TestCreateResponseApi(t *testing.T) {
	data := struct {
		ID   int
		Name string
	}{
		ID:   1,
		Name: "Ejemplo",
	}

	expectedStatus := "success"
	expectedMessage := "data found"

	response, err := CreateResponseApi(data)

	if err != nil {
		t.Errorf("Error al crear la respuesta de la API: %s", err.Error())
		t.Fail()
	}

	var responseDto dtos.ResponseDTO
	err = json.Unmarshal(response, &responseDto)
	if err != nil {
		t.Errorf("Error al deserializar la respuesta de la API: %s", err.Error())
		t.Fail()
	}

	if responseDto.Status != expectedStatus {
		t.Errorf("Error al crear la respuesta de la API, estado incorrecto. Esperado: %s, Obtenido: %s", expectedStatus, responseDto.Status)
		t.Fail()
	}

	if responseDto.Message != expectedMessage {
		t.Errorf("Error al crear la respuesta de la API, mensaje incorrecto. Esperado: %s, Obtenido: %s", expectedMessage, responseDto.Message)
		t.Fail()
	}

	if responseDto.Data == nil {
		t.Error("Error al crear la respuesta de la API, los datos están vacíos")
		t.Fail()
	}

	t.Log("Test success")
}

func TestBadResponse(t *testing.T) {
	expectedStatus := "error"
	expectedMessage := "Mensaje de error de prueba"
	responseStringExpected := `{"status":"error","message":"Mensaje de error de prueba","data":null}`

	response := BadResponse(expectedMessage)

	if response != responseStringExpected {
		t.Errorf("No se optuvo la respuesta esperada")
		t.Fail()
	}

	var responseDto dtos.ResponseDTO
	err := json.Unmarshal([]byte(response), &responseDto)
	if err != nil {
		t.Errorf("Error al deserializar la respuesta de error: %s", err.Error())
		t.Fail()
	}

	if responseDto.Status != expectedStatus {
		t.Errorf("Error al crear la respuesta de error, estado incorrecto. Esperado: %s, Obtenido: %s", expectedStatus, responseDto.Status)
		t.Fail()
	}

	if responseDto.Message != expectedMessage {
		t.Errorf("Error al crear la respuesta de error, mensaje incorrecto. Esperado: %s, Obtenido: %s", expectedMessage, responseDto.Message)
		t.Fail()
	}

	t.Log("Test success")
}
