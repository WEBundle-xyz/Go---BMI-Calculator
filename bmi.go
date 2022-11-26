package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	fmt.Println("BMI Calculator")
	fmt.Println("----------------------")

	fmt.Println("Please enter your weight in Kg:  ")
	weightInput, _ := reader.ReadString('\n')
	fmt.Println("Please enter your height in m:  ")
	heightInput, _ := reader.ReadString('\n')

	weightInput = strings.Replace(weightInput, "\n", "", -1)
	heightInput = strings.Replace(heightInput, "\n", "", -1)

	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(weightInput, 64)

	bmi := weight / (height * height)
	fmt.Printf("Your BMI:  %.2f", bmi)
}
