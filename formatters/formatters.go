package formatters

func GetFomatter(body1, bosy2 map[string]any, format string) string {
	switch format {
	case "stylish":
		return stylishFormatter(body1, bosy2)
	}
	return ""
}
