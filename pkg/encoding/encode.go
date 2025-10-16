package encoding

import (
	"encoding/json"
	"os"

	"gopkg.in/yaml.v3"
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

func EncodeYaml(filename string) (map[string]any, error) {
	body := make(map[string]any)
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(fileBytes, &body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
