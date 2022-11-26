package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/stefg/bmi/info"
)

func main() {
	fmt.Println(info.MainTitle)
	fmt.Println(info.Separator)

	fmt.Println(info.WeightPromt)
	weightInput, _ := reader.ReadString('\n')
	fmt.Println(info.HeightPromt)
	heightInput, _ := reader.ReadString('\n')

	weightInput = strings.Replace(weightInput, "\n", "", -1)
	heightInput = strings.Replace(heightInput, "\n", "", -1)

	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(weightInput, 64)

	bmi := weight / (height * height)
	fmt.Printf("Your BMI:  %.2f", bmi)
}
