package jsonFileReader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func GetJSONFile(filePath string, key string) string {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	str, ok := result[key].(string)
	if ok {
		return str
	} else {
		return "Error"
	}
}
