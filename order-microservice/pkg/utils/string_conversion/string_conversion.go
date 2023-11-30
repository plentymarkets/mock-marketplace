package string_conversion

import "strconv"

func IntToString(int int) string {
	return strconv.Itoa(int)
}

func StringToInt(string string) (int, error) {
	return strconv.Atoi(string)
}
