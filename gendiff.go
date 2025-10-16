package code

import (
	"code/formatters"
	"code/pkg/encoding"
	"fmt"
	"strings"
)

func GenDiff(filepath1, filepath2, format string) (string, error) {
	var err error
	var body1, body2 map[string]any

	inputFormat := "json"
	if strings.HasSuffix(filepath1, ".yml") || strings.HasSuffix(filepath1, ".yaml") {
		inputFormat = "yaml"
	}

	if inputFormat == "yaml" {
		body1, err = encoding.EncodeYaml(filepath1)
		if err != nil {
			fmt.Println(err)
		}
		body2, err = encoding.EncodeYaml(filepath2)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		body1, err = encoding.EncodeJson(filepath1)
		if err != nil {
			fmt.Println(err)
		}
		body2, err = encoding.EncodeJson(filepath2)
		if err != nil {
			fmt.Println(err)
		}
	}

	result := formatters.GetFomatter(body1, body2, format)

	return result, nil
}
