package file

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func GetJsonContent(filepath string) (string, bool) {
	fileAdapter, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	defer fileAdapter.Close()

	byteValue, _ := ioutil.ReadAll(fileAdapter)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	fmt.Println(result)

	return "test", false
}
