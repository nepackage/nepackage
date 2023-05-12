package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"log"
	"os"
	"os/user"
	"strconv"
	"strings"

	"github.com/gobwas/glob"
)

var g glob.Glob

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func ConvertStringToStruct(input string) []string {
	convertedString := strings.Split(input, ",")
	return convertedString
}

func ConvertStringToUintStruct(input string) []int {
	convertedStringArray := strings.Split(input, ",")
	var convertedint = []int{}

	for _, i := range convertedStringArray {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		convertedint = append(convertedint, j)
	}
	return convertedint
}

func MatchValidator(mustValue string, toValidateValue string) bool {

	g = glob.MustCompile(mustValue)
	state := g.Match(toValidateValue)

	return state
}

func TmpFolderCreation(folderName string) (folderPath string, err error) {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	folderToCreate := path + "/" + folderName
	err = os.MkdirAll(folderToCreate, 0755)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return folderToCreate, nil
}

func GetCurrentUser() (string, error) {
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
		return "", err
	}

	username := user.Username
	return username, nil
}

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

func EncryptString(text, secret string) (string, error) {
	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText), nil
}

func DecryptString(text, secret string) (string, error) {
	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return "", err
	}
	cipherText := Decode(text)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}

func EnvExist(key string) bool {
	if _, ok := os.LookupEnv(key); ok {
		return true
	}
	return false
}
