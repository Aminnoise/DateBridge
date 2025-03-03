package calendar

var miladiDaysInMonth = []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

type Gregorian struct {
	Year  int
	Month int
	Day   int
}

func (g Gregorian) ToDays() int {
	days := g.Day
	for i := 1; i < g.Month; i++ {
		daysMonth := miladiDaysInMonth[i]
		if i == 2 && g.Kabise(g.Year) {
			daysMonth = miladiDaysInMonth[i] + 1
		}
		days += daysMonth
	}
	for i := 622; i < g.Year; i++ {
		days += 365
		if g.Kabise(i) {
			days++
		}
	}
	days -= 80 // چون شمسی از 22 مارس شروع شده
	return days
}

func (g Gregorian) FromDays(days int) Calendar {
	year := 622
	days += 80
	for {
		daysInYear := 365
		if g.Kabise(year) {
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
		daysInMonth := miladiDaysInMonth[i]
		if i == 2 && g.Kabise(year) {
			daysInMonth = 29
		}
		if days < daysInMonth {
			break
		}
		days -= daysInMonth
		month++
	}

	return Gregorian{Year: year, Month: month, Day: days}
}

func (g Gregorian) Kabise(year int) bool {
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return true
	}
	return false
}
