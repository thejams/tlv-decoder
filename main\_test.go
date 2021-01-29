package main

import "testing"

func TestStringFormat(t *testing.T) {
	response := getStringFormat("05", "Alfanumérico", "3", "123")
	if response != "05 de tipo Alfanumérico de largo 3 y valor 123" {
		t.Error("expected: 05 de tipo Alfanumérico de largo 3 y valor 123, but got:", response)
	}
}

func TestType(t *testing.T) {
	numericResponse := getType("N")
	if numericResponse != "Numérico" {
		t.Error("expected: Numérico, but got:", numericResponse)
	}
	alphaNumericResponse := getType("A")
	if alphaNumericResponse != "Alfanumérico" {
		t.Error("expected: Alfanumérico, but got:", alphaNumericResponse)
	}
	noResponse := getType("ASDA")
	if noResponse != "Sin formato definido" {
		t.Error("expected: Sin formato definido, but got:", noResponse)
	}
}

func TestTLVDecoder(t *testing.T) {
	response, error := tlvDecoder("11A05AB398765UJ102N2300")
	if error != nil {
		t.Error("expected: full map object response, but got error: ", error)
	}
	if len(response) != 2 {
		t.Error("expected: map length of 2, but got:", len(response))
	}
	if response["key1"] != "05 de tipo Alfanumérico de largo 11 y valor AB398765UJ1" {
		t.Error("expected: 05 de tipo Alfanumérico de largo 11 y valor AB398765UJ1 , but got:", response["key1"])
	}
	if response["key2"] != "23 de tipo Numérico de largo 02 y valor 00" {
		t.Error("expected: 23 de tipo Numérico de largo 02 y valor 00 , but got:", response["key2"])
	}
}

func TestTLVEmptyString(t *testing.T) {
	noResponse, error := tlvDecoder("")
	if error == nil {
		t.Error("expected: error response but got full map object response: ", noResponse)
	}
}

func TestTLVEInvalidStringValues(t *testing.T) {
	_, largoStringError := tlvDecoder("zzA05AB398765UJ102N2300")
	if largoStringError == nil {
		t.Error("expected: invalid string type for largo value, should be a number, got: ", largoStringError)
	}

	_, posicionStringError := tlvDecoder("11AzzAB398765UJ102N2300")
	if posicionStringError == nil {
		t.Error("expected: invalid string type for posicion value, should be a number, got: ", posicionStringError)
	}

	_, tipoStringError := tlvDecoder("11z05AB398765UJ102N2300")
	if tipoStringError == nil {
		t.Error("expected: invalid string type for posicion value, should be a number, got: ", tipoStringError)
	}
}
