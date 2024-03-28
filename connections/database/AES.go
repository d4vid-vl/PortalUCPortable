package database

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

// ? Inicializamos el nonce y AESKey que vamos a usar
var globalNonce []byte
var globalAESKey []byte

// ? Nombre de los archivos para almacenar el nonce y la key
const nonceFileName = "connections/database/json/nonce.txt"
const AESKeyFileName = "connections/database/json/AESKey.txt"

// * AES Key
// Generar una clave aleatoria segura
func generateKey(size int) ([]byte, error) {
	key := make([]byte, size)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// Genera una key y la almacena en un archivo
func generateAndStoreAESKey(size int) error {
	key, err := generateKey(size)
	if err != nil {
		fmt.Println("Error al generar la llave AES: ", err)
	}

	// Almacenar la key en un archivo
	err = os.WriteFile(AESKeyFileName, key, 0644)
	if err != nil {
		return err
	}

	globalAESKey = key
	return nil
}

// Cargar llave del archivo
func loadKeyFromFile() error {
	key, err := os.ReadFile(AESKeyFileName)
	if err != nil {
		return err
	}

	globalAESKey = key
	return nil
}

// * Nonce
// Generar un nonce global
func generateNonce(size int) ([]byte, error) {
	nonce := make([]byte, size)
	_, err := io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}
	return nonce, nil
}

// Generar un nonce y almacenarlo en un archivo
func generateAndStoreNonce(size int) error {
	nonce, err := generateNonce(size)
	if err != nil {
		return err
	}

	// Almacenar el nonce en un archivo
	err = os.WriteFile(nonceFileName, nonce, 0644)
	if err != nil {
		return err
	}

	globalNonce = nonce
	return nil
}

// Cargar el nonce desde el archivo
func loadNonceFromFile() error {
	// Leer el nonce desde el archivo
	nonce, err := os.ReadFile(nonceFileName)
	if err != nil {
		return err
	}

	globalNonce = nonce
	return nil
}

// * Inicializar el nonce y llave al empezar el programa
func init() {
	// Intentar cargar el nonce desde el archivo
	err := loadNonceFromFile()
	if err != nil {
		// Si falla, generar un nuevo nonce y guardarla
		err := generateAndStoreNonce(12)
		if err != nil {
			fmt.Println("Error al inicializar el nonce: ", err)
		}
	}
	errkey := loadKeyFromFile()
	if errkey != nil {
		// Si falla, genera una nueva key y la guarda
		err := generateAndStoreAESKey(32)
		if err != nil {
			fmt.Println("Error al inicializar la llave AES: ", err)
		}
	}
}

// ! Encripta la información del usuario
func encryption(input string) string {
	// Se crea un cipher block para la encriptación
	c, err := aes.NewCipher(globalAESKey)
	if err != nil {
		fmt.Println("Error al crear el cipher: ", err)
		return ""
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println("Error al crear el GCM: ", err)
		return ""
	}

	gcmbyte := gcm.Seal(nil, globalNonce, []byte(input), nil)
	iEncrypt := base64.StdEncoding.EncodeToString(gcmbyte)
	// Encriptamos con gcm.Seal y devolvemos
	return iEncrypt
}

// ! Desencripta la información del usuario
func decryption(input string) string {
	// Se crea un cipher block para la encriptación
	c, err := aes.NewCipher(globalAESKey)
	if err != nil {
		fmt.Println("Error al crear el cipher: ", err)
		return ""
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println("Error al crear el GCM: ", err)
	}

	// Desencriptamos el base64 devuelto por la encriptación
	text, _ := base64.StdEncoding.DecodeString(input)

	iDecrypt, err := gcm.Open(nil, globalNonce, text, nil)
	if err != nil {
		fmt.Println("Error al desencriptar: ", err)
	}

	return string(iDecrypt)
}
