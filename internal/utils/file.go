package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func GetFileContent(filepath string) []byte {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	return byteValue
}

func FileExists(filepath string) bool {
	info, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func DirExists(filepath string) bool {
	info, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
}

func GetCacheFilePath(file string) string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%s/var/cache/%s.json", path, file)
}

func GetDataFilePath(file string) string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%s/data/%s", path, file)
}

func IsValidPostCode(code string) bool {
	return len(code) > 5
}
