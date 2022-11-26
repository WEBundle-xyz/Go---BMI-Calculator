package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	fmt.Println("BMI Calculator")
	fmt.Println("----------------------")

	fmt.Println("Please enter your weight in Kg:  ")
	weightText, _ := reader.ReadString('\n')
	fmt.Println("Please enter your height in m:  ")
	heightText, _ := reader.ReadString('\n')

	fmt.Println(weightText)
	fmt.Println(heightText)
}

// this will give me the to the input text
// so that what it is
