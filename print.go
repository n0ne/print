package print

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/fatih/color"
)

func Struct(strct interface{}) {
	argValue := reflect.ValueOf(strct)
	if argValue.Kind() != reflect.Struct {
		Error("Error: Argument is not a struct")
		return
	}

	jsonBytes, err := json.MarshalIndent(strct, "", "    ")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Convert the indented JSON bytes to a string and replace tabs with spaces
	indentedJSON := strings.ReplaceAll(string(jsonBytes), "\t", "    ")

	fmt.Println(indentedJSON)
}

func Structc(strct interface{}) {
	argValue := reflect.ValueOf(strct)
	if argValue.Kind() != reflect.Struct {
		Error("Error: Argument is not a struct")
		return
	}

	jsonBytes, err := json.MarshalIndent(strct, "", "    ")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Convert the indented JSON bytes to a string and replace tabs with spaces
	indentedJSON := strings.ReplaceAll(string(jsonBytes), "\t", "    ")

	// Create color functions
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	// Colorize the keys and values in JSON
	colorized := colorizeJSON(indentedJSON, red, green)

	// Print the colorized JSON
	fmt.Println(colorized)

}

func Slice[T any](s []T) {
	// Convert the slice values to strings
	strValues := make([]string, 0)
	for _, value := range s {
		strValues = append(strValues, fmt.Sprint(value))
	}

	typeOfSlice := fmt.Sprintf("%T", s)

	// Concatenate the values with a space separator and add brackets
	result := typeOfSlice + "{ " + strings.Join(strValues, ", ") + " }"

	// Print the concatenated string
	fmt.Println(result)
}

func colorizeJSON(input string, keyColorFunc, valueColorFunc func(a ...interface{}) string) string {
	// Split the input into lines
	lines := strings.Split(input, "\n")

	// Iterate through each line
	for i, line := range lines {
		// Find the index of the colon character
		colonIndex := strings.Index(line, ":")

		// If colon exists and it's not the last character
		if colonIndex > 0 && colonIndex < len(line)-1 {
			// Colorize the key before the colon and the value after the colon

			// Colorize the key using the keyColorFunc
			colorizedKey := keyColorFunc(line[:colonIndex]) + line[colonIndex:]

			// Find the colon index in the colorized key
			colonIndex = strings.Index(colorizedKey, ":")

			// Colorize the value using the valueColorFunc
			colorizedValue := colorizedKey[:colonIndex+1] + valueColorFunc(colorizedKey[colonIndex+1:])

			// Update the line with the colorized value
			lines[i] = colorizedValue
		}
	}

	// Join the lines back into a string and return
	return strings.Join(lines, "\n")
}

func Error(error string) {
	redColor := "\033[31m"
	resetColor := "\033[0m"

	if CheckString(error) {
		fmt.Println(redColor + error + resetColor)
	} else {
		fmt.Println("Argument is not a string")
	}
}
