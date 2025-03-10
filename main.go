package main

import (
	"fmt"
	"shamsi/calendar"
)

func main() {

	j := calendar.Jalali{Year: 534, Month: 6, Day: 23}
	g := calendar.Gregorian{}

	converter := calendar.Converter{}

	miladi := converter.Convert(j, g).(calendar.Gregorian)
	fmt.Printf("Gregorian: %d-%02d-%02d\n", miladi.Year, miladi.Month, miladi.Day)

	k := calendar.Jalali{}
	shamsi := converter.Convert(miladi, k).(calendar.Jalali)
	fmt.Printf("Jalali: %d-%02d-%02d\n", shamsi.Year, shamsi.Month, shamsi.Day)

}
