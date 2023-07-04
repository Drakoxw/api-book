package utils

import (
	"testing"
)

func TestGenerateSalt(t *testing.T) {
	value := 10
	salt := GenerateSalt(value)
	if len(salt) != value {
		t.Errorf("Fallo al generar salt, Se ecperaba un tamaño: %d y es: %d", len(salt), value)
		t.Fail()
	} else {
		t.Log("Test success")
	}
}

func TestHashPassword(t *testing.T) {
	value := "passwordSecret"
	expect := "9c51811c36d923767891505d7a506155d23c2a9de9c88dde117a4d3880e2c9f3bc56da8354a43918b532eee918990961f5cac865154b40b0eed5861735206804"
	salt := GenerateSalt(13)
	passEnc := HashPassword(value, salt)
	if passEnc == expect {
		t.Log("Test success")
	} else {
		t.Error("\nNo hubo coincidencia en los datos!!! \n")
		t.Errorf("\nExpect : %s", expect)
		t.Errorf("\nPassEnc : %s", passEnc)
		t.Fail()
	}

}

func TestDoPasswordsMatch(t *testing.T) {
	password := "passwordSecret"
	hashedPass := "9c51811c36d923767891505d7a506155d23c2a9de9c88dde117a4d3880e2c9f3bc56da8354a43918b532eee918990961f5cac865154b40b0eed5861735206804"
	salt := GenerateSalt(13)
	match := DoPasswordsMatch(hashedPass, password, salt)
	if match {
		t.Log("Test success")
	} else {
		t.Error("\nNo hubo coincidencia en los datos!!! \n")
		t.Fail()
	}
}
