/*
	composicion de un tlv:
		Largo: 2 caracteres que indican el largo del valor, este campo es importante puesto que indica cuantos caracteres leer a continuación.
		Tipo: El tipo tiene un largo de 3 caracteres donde el primer caracter indica el tipo de dato (A: Alfanumérico y N: Numérico) y dos caracteres para indicar el numero de campo Ejemplo: "01" o "15".
		Valor: Este es el valor del campo, el cual corresponde al valor del campo, su largo esta determinado por la porción Largo.

	Ejemplo: 11A05AB398765UJ102N2300
						 2  3          index + largo
	[11 A05 AB398765UJ1 ; 02 N23 00]

	Largo: 11 (indica cuantos caracteres leer despues para extraer el valor)
	Typo: A05 (A indica tipo de dato: A alfanumerico, N Numerico. 05 indica el numero de campo o la posicion del primer caracter del Valor)
	Valor: AB398765UJ1 (son los 11 caracteres que dice el largo, a partir de la posicion 05 de la cadena total)

	Largo: 02 (2 caracteres de largo)
	Typo: N23 (N de numerico, en la posicion 23)
	Valor: 00

*/

package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func getType(value string) string {
	if strings.ToLower(value) == "n" {
		return "Numérico"
	} else if strings.ToLower(value) == "a" {
		return "Alfanumérico"
	} else {
		return "Sin formato definido"
	}
}

func getStringFormat(posicion string, tipo string, largo string, valor string) string {
	return posicion + " de tipo " + tipo + " de largo " + largo + " y valor " + valor
}

func tlvDecoder(value string) (map[string]string, error) {
	if len(value) < 6 {
		return nil, errors.New("invalid string, should be at least 6 characters long ")
	}
	response := make(map[string]string)
	runeFromValue := []rune(value)
	i := 0
	key := 1

	for i < len(value) {
		largo := string(runeFromValue[i : i+2])
		_, largoStringError := strconv.Atoi(largo)
		if largoStringError != nil {
			return nil, errors.New("invalid string type for largo value, should be a number: " + largo)
		}

		tipo := getType(string(runeFromValue[i+2 : i+3]))
		if tipo == "Sin formato definido" {
			return nil, errors.New("invalid string type for type value, should be N or A, got: " + string(runeFromValue[i+2:i+3]))
		}

		posicion := string(runeFromValue[i+3 : i+5])
		_, posicionStringError := strconv.Atoi(posicion)
		if posicionStringError != nil {
			return nil, errors.New("invalid string type for posicion value, should be a number, got: " + posicion)
		}

		newIterator, _ := strconv.Atoi(largo)
		valor := string(runeFromValue[i+5 : i+5+newIterator])
		response["key"+strconv.Itoa(key)] = getStringFormat(posicion, tipo, largo, valor)
		i += i + 5 + newIterator
		key++
	}

	return response, nil
}

func main() {
	var tlvValue string
	fmt.Print("Ingrese un string con formato TLV: ")
	fmt.Scan(&tlvValue)
	response, error := tlvDecoder(tlvValue)
	if error != nil {
		fmt.Println("Error: ", error)
	} else {
		fmt.Println(response)
	}
}
