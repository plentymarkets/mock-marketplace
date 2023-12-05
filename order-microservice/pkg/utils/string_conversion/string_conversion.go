package string_conversion

import "strconv"

func IntToString(int int) string {
	return strconv.Itoa(int)
}

func StringToInt(string string) (int, error) {
	integer, err := strconv.Atoi(string)
	return integer, err
}

func StringToUint(string string) (uint, error) {
	integer, err := strconv.Atoi(string)
	return uint(integer), err
}
