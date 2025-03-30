package cmd

import (
	"fmt"
	"math"
	"strconv"
)

func Add(a string, b string) string {
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

func Subtract(a, b string) (result string) {
	num1, err := strconv.ParseFloat(a, 64)
	if err != nil {
		fmt.Println("Error: First value is invalid")
		return
	}
	num2, err := strconv.ParseFloat(b, 64)
	if err != nil {
		fmt.Println("Error: First value is invalid")
		return
	}

	result = fmt.Sprintf("%f", num1-num2)
	return
}

func Multiply(a, b string, isRoundUp bool) (result string) {
	num1, err := strconv.ParseFloat(a, 64)
	if err != nil {
		fmt.Println("Error: First value is invalid")
		return
	}
	num2, err := strconv.ParseFloat(b, 64)
	if err != nil {
		fmt.Println("Error: First value is invalid")
		return
	}
	if isRoundUp {
		num1 = float64(int(num1))
		num2 = float64(int(num2))
	}
	result = fmt.Sprintf("%f", num1*num2)
	return
}

func Divide(a, b string, isRoundUp bool) (result string, err error) {
	num1, err := strconv.ParseFloat(a, 64)
	if err != nil {
		err = fmt.Errorf("first value is not a number")
		return
	}
	num2, err := strconv.ParseFloat(b, 64)
	if err != nil {
		err = fmt.Errorf("second value is not a number")
		return
	}

	if num2 == 0 {
		err = fmt.Errorf("error: division by zero is not allowed ")
		return
	}

	if isRoundUp {
		num1 = math.Round(num1)
		num2 = math.Round(num2)
	}	

	result = fmt.Sprintf("%f", num1/num2)
	return
}
