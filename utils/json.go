package utils

import (
	"encoding/json"
	"fmt"
)

func Dumps(jsonMap map[string]any) (string, error) {
	some, err := json.Marshal(jsonMap)
	if err != nil {
		fmt.Println(err)
	}
	jsonStr := string(some)
	return jsonStr, err

}

func Loads(jsonStr string) (map[string]any, error) {
	var jsonMap map[string]any
	err := json.Unmarshal([]byte(jsonStr), &jsonMap)
	return jsonMap, err
}
