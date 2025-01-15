package typeHelper

import (
	"fmt"
	"strconv"
)

func StringToInt(input string, defaultValue int) int {
	result, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("Error converting string to int: %v. Returning default value: %d\n", err, defaultValue)
		return defaultValue
	}
	return result
}
