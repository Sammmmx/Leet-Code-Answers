package main

import "strings"

func convert(s string, numRows int) string {
	var result strings.Builder

	if numRows == 1 {
		return s
	}

	rows := make([]strings.Builder, numRows)
	currentRow := 0
	goingDown := false

	for _, char := range s {
		rows[currentRow].WriteRune(char)

		if currentRow == 0 {
			goingDown = true
		} else if currentRow == numRows-1 {
			goingDown = false
		}

		if goingDown {
			currentRow++
		} else {
			currentRow--
		}
	}

	for _, row := range rows {
		result.WriteString(row.String())
	}
	return result.String()
}

func main() {
	result := convert("PAYPALISHIRING", 3)
	println(result)
}
