package logger

import "fmt"

func Log(message string, error error) {
	if error != nil {
		fmt.Printf("%s: %s\n", message, error.Error())
		return
	}

	fmt.Println(message)
	return
}
