package main

import "fmt"

func main() {
    fmt.Print("Enter fareignheight: ")
    var input float64
    fmt.Scanf("%f", &input)

    output := convertToCelsius(input)

    fmt.Println(output)
}

func convertToCelsius(temp float64) float64 {
	result := (temp - 32) * 5/9
	return result
}

func convertFeetToMetres(amt float64) float64 {
	const metresInFeet := 0.3048
	return amt * metresInFeet
}
