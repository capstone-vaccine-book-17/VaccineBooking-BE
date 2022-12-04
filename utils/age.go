package utils

import (
	"time"
)

func Age(input string) int {
	layout:= "2006-01-02"

	date, _ := time.Parse(layout, input)
	today := time.Now()

	age := today.Year() - date.Year() 
	return age

}
