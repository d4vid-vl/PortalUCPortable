package database

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SaveUserData(user string, password string) {
	// Manejo de la información del Usuario
	eUser := encryption(user)
	ePassword := encryption(password)
	UserData := User{
		Username: eUser,
		Password: ePassword,
	}

	bytes, _ := json.MarshalIndent(UserData, "", "	")

	userDataFile := "connections/database/json/UserInfo.json"

	// Revisa si el archivo JSON está creado
	if _, err := os.Stat(userDataFile); err == nil { // En caso de estar creado
		checker := checkData(eUser, ePassword)
		if checker { // Si la información proporcionada por el usuario es la misma, no hará nada
			return
		} else { // En cambio si no es la misma, sobreescribirá la información anterior
			// Abre el archivo y borra los contenidos de este
			_, err := os.OpenFile(userDataFile, os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				fmt.Println("Error al abrir el JSON: ", err)
			}
			// Escribe la información en el JSON
			if err := os.WriteFile(userDataFile, bytes, 0644); err != nil {
				fmt.Println("Error al escribir en el JSON: ", err)
				return
			}
		}
	} else { // En caso de no estar creado
		// Crea el archivo
		file, err := os.Create(userDataFile)
		if err != nil {
			fmt.Println("Error al crear el JSON: ", err)
			return
		}
		defer file.Close()

		// Escribe la información en el JSON
		if err := os.WriteFile(userDataFile, bytes, 0644); err != nil {
			fmt.Println("Error al escribir en el JSON: ", err)
			return
		}
	}
}

// * Revisa si la información del usuario está ya presente en el JSON
func checkData(user string, password string) bool {
	UserDataFile := "connections/database/json/UserInfo.json"

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

	// Dependiendo de si el Username y Contraseña son iguales, devolverá true o false
	if UserData.Username == user && UserData.Password == password {
		return true
	} else {
		return false
	}
}
