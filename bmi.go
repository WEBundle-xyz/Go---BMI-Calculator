package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	mainTitle   = "BMI Calculator"
	separator   = "-------------------"
	weightPromt = "Plese enter the weight in Kg:  "
	heightPromt = "Please enter the height in cm: "
)

func main() {
	fmt.Println(mainTitle)
	fmt.Println(separator)

	fmt.Println(weightPromt)
	weightInput, _ := reader.ReadString('\n')
	fmt.Println(heightPromt)
	heightInput, _ := reader.ReadString('\n')

	weightInput = strings.Replace(weightInput, "\n", "", -1)
	heightInput = strings.Replace(heightInput, "\n", "", -1)

	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(weightInput, 64)

	bmi := weight / (height * height)
	fmt.Printf("Your BMI:  %.2f", bmi)
}
