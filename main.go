package main

import (
	"fmt"
)

func main() {
	var weight float64
	var height float64

	fmt.Print("Enter your weight (in kg): ")
	fmt.Scan(&weight)

	fmt.Print("Enter your height (in cm): ")
	fmt.Scan(&height)

	// Convert cm to meters for BMI calculation
	heightInMeters := height / 100
	bmi := weight / (heightInMeters * heightInMeters)

	fmt.Printf("\nYour BMI is: %.1f\n", bmi)
	fmt.Println("\nBMI Categories:")
	switch {
	case bmi < 18.5:
		targetWeight := 18.5 * heightInMeters * heightInMeters
		weightToGain := targetWeight - weight
		fmt.Println("You are underweight")
		fmt.Printf("You need to gain %.1f kg to reach normal weight\n", weightToGain)
	case bmi < 25:
		fmt.Println("You have a normal weight")
	case bmi < 30:
		targetWeight := 24.9 * heightInMeters * heightInMeters
		weightToLose := weight - targetWeight
		fmt.Println("You are overweight")
		fmt.Printf("You need to lose %.1f kg to reach normal weight\n", weightToLose)
	default:
		targetWeight := 24.9 * heightInMeters * heightInMeters
		weightToLose := weight - targetWeight
		fmt.Println("You are obese")
		fmt.Printf("You need to lose %.1f kg to reach normal weight\n", weightToLose)
	}
}
