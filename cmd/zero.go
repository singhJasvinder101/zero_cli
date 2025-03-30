package cmd

import (
	"fmt"
	"strconv"
)

func Add(a string, b string) string{	
	num1, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return "Invalid input"
	}
	num2, err := strconv.ParseFloat(b, 64)
	if err != nil {
		return "Invalid input"
	}
	result := num1 + num2
	return fmt.Sprintf("%f", result)
}

func Subtract(a, b string) string{
	num1, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return fmt.Sprintf("%f", num1)
	}
	num2, err := strconv.ParseFloat(b, 64)
	if err != nil {
		return fmt.Sprintf("%f", num2)
	}

	result := num1 - num2
	return fmt.Sprintf("%f", result)
}

