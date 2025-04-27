package helper

func Or(value1 string, value2 string) string {
	if value1 == "" {
		return value2
	} else {
		return value1
	}
}