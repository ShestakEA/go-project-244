package code

import (
	"code/pkg/encoding"
	"fmt"
	"sort"
	"strings"
)

func GenDiff(filepath1, filepath2, format string) (string, error) {
	var result string

	if format == "json" {
		body1, err := encoding.EncodeJson(filepath1)
		if err != nil {
			fmt.Println(err)
		}
		body2, err := encoding.EncodeJson(filepath2)
		if err != nil {
			fmt.Println(err)
		}

		uniqueKeys := make([]string, 0, max(len(body1), len(body2)))
		for k := range body1 {
			uniqueKeys = append(uniqueKeys, k)
		}
		for k := range body2 {
			if _, ok := body1[k]; !ok {
				uniqueKeys = append(uniqueKeys, k)
			}
		}
		sort.Strings(uniqueKeys)
		s := make([]string, 0, len(uniqueKeys))
		for _, k := range uniqueKeys {
			val1, ok1 := body1[k]
			val2, ok2 := body2[k]
			if ok1 && ok2 {
				if val1 == val2 {
					s = append(s, fmt.Sprintf("  %s: %v", k, val1))
				} else {
					s = append(s, fmt.Sprintf("- %s: %v", k, val1))
					s = append(s, fmt.Sprintf("+ %s: %v", k, val2))
				}
			} else {
				if ok1 {
					s = append(s, fmt.Sprintf("- %s: %v", k, val1))
				}
				if ok2 {
					s = append(s, fmt.Sprintf("+ %s: %v", k, val2))
				}
			}
		}
		result = "{\n" + strings.Join(s, "\n") + "\n}"
	}

	return result, nil
}
