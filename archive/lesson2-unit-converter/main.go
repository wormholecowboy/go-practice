package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	Fahrenheit = iota
	Celsius
	Miles
	Kilometers
)

type val struct {
	value   float64
	valType int
}

func (v val) Convert() val {
	var out val
	if v.valType == Fahrenheit {
		out.valType = Celsius
		out.value = (v.value - 32) * (5.0 / 9.0)
		return out
	}
	if v.valType == Celsius {
		out.valType = Fahrenheit
		out.value = (v.value * 9.0 / 5.0) + 32
		return out
	}
	if v.valType == Miles {
		out.valType = Kilometers
		out.value = v.value * 1.60934
		return out
	}
	if v.valType == Kilometers {
		out.valType = Miles 
		out.value = v.value / 1.60934
		return out
	}
	return out
}

func unitLabel(unit int) string {
	switch unit {
	case Fahrenheit:
	  return "Fahrenheit"
	case Celsius: 
		return "Celsius"
	case Miles:
	  return "Miles"
	case Kilometers:
	  return "Kilometers"
	}
	return ""
	
}

func mapUnits(unit string) int {
	switch unit {
	case "F":
		return Fahrenheit
	case "C":
		return Celsius
	case "km":
		return Kilometers
	case "mi":
		return Miles
	default:
	  fmt.Println("Unknown unit")
	  os.Exit(1)
	} 
	return -1
}

func mapArgs(value string, valType string) val {
	num, err := strconv.ParseFloat(value, 64)
	if err != nil {
		fmt.Print("Boom")
		os.Exit(1)
	}
	unitType := mapUnits(valType)
	out := val{num, unitType}
	return out
}

func main() {
	if len(os.Args) < 3 {
		os.Exit(1)
	}

	numArg := os.Args[1]
	unitArg := os.Args[2]
	input := mapArgs(numArg, unitArg)

	thingy := input.Convert()
	

	fmt.Printf("Here is the value: %v, and the unit: %v\n", thingy.value, unitLabel(thingy.valType))

}
