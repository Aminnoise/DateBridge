package calendar

import "fmt"

var jalaliDaysInMonth = []int{0, 31, 31, 31, 31, 31, 31, 30, 30, 30, 30, 30, 29}

type Jalali struct {
	Year  int
	Month int
	Day   int
}

func (j Jalali) ToDays() int {
	days := j.Day
	for i := 1; i < j.Month; i++ {
		days += jalaliDaysInMonth[i]
	}
	for i := 1; i < j.Year; i++ {
		days += 365
		if j.Kabise(i) {
			days++
		}
	}
	days-- // چون فاصله از یکم هست یکم باید کم بشه
	fmt.Println(days)
	return days
}

func (j Jalali) FromDays(days int) Calendar {
	year := 1
	for {
		daysInYear := 365
		if j.Kabise(year) && year != 1 {
			daysInYear = 366
		}
		if days < daysInYear {
			break
		}
		days -= daysInYear
		year++
	}

	month := 1
	for i := 1; i <= 12; i++ {
		daysInMonth := jalaliDaysInMonth[i]
		if i == 12 && j.Kabise(year) {
			daysInMonth = 30
		}
		if days < daysInMonth {
			break
		}
		days -= daysInMonth
		month++
	}

	return Jalali{Year: year, Month: month, Day: days}
}

func (j Jalali) Kabise(year int) bool {
	kabiseYears := []int{1, 5, 9, 13, 17, 22, 26, 30}
	k := year % 33
	for _, v := range kabiseYears {
		if k == v {
			return true
		}
	}
	return false
}
