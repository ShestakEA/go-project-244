package encoding

import (
	"encoding/json"
	"os"
)

func EncodeJson(filename string) (map[string]any, error) {
	body := make(map[string]any)
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(fileBytes, &body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
