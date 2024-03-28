package database

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadUserData() [2]string {
	UserDataFile := "connections/database/json/UserInfo.json"

	// Leemos el JSON con la información del usuario
	content, err := os.ReadFile(UserDataFile)
	if err != nil {
		return [2]string{"", ""}
	}

	// Transferimos y desencriptamos la información que hay en este JSON para devolverla
	var UserData User
	if err := json.Unmarshal(content, &UserData); err != nil {
		fmt.Println(err)
	}

	dUser := decryption(UserData.Username)
	dPassword := decryption(UserData.Password)

	return [2]string{dUser, dPassword}
}
