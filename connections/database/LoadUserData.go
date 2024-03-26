package database

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadUserData() [2]string {
	UserDataFile := "connections/database/json/UserInfo.json"

	// Revisa que el archivo esté creado
	if _, err := os.Stat(UserDataFile); err == nil { // Si está creado, lee la información y la devuelve
		// Abre el archivo
		file, err := os.Open(UserDataFile)
		if err != nil {
			fmt.Println("Error abriendo el JSON: ", err)
		}
		defer file.Close()

		// Decodifica el JSON para su posterior análisis
		var UserData User
		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&UserData); err != nil {
			fmt.Println("Error decodificando el JSON: ", err)
		}

		// Desencripta el usuario y contraseña para ser devuelta y puesta en el texto
		dUser := decryption(UserData.Username)
		dPassword := decryption(UserData.Password)

		return [2]string{dUser, dPassword}
	} else { // En caso contrario, contará con espacios default
		return [2]string{"", ""}
	}
}
