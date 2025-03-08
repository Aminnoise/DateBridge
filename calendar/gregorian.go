package calendar

var miladiDaysInMonth = []int{0,
	31, 28, 31, 30, 31, 30,
	31, 31, 30, 31, 30, 31}

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
	days += (g.Year - 622) * 365
	days += g.Year / 4
	days -= g.Year / 100
	days += g.Year / 400
	days -= 622 / 4
	days += 622 / 100
	days -= 622 / 400

	days -= 80 // چون شمسی از 22 مارس شروع شده

	return days
}

func (g Gregorian) FromDays(days int) Calendar {
	year := 621       // شروع دوره 4 ساله
	days += 365 + 80  // 365روز برای یکسال عقب رفتن از مبدا و 80 روز برای از اول سال 622 تا مبدا
	D4 := days / 1461 // تعداد دوره های 4 ساله
	D100 := D4 / 25   // تعداد دوره های 100 ساله
	D400 := D100 / 4  // تعداد دوره های 400 ساله
	year += D4 * 4
	days += D100 - D400
	days = days % 1461
	for {
		daysInYear := 365
		if days < daysInYear {
			break
		}
		days -= daysInYear
		year++
	}

	month := 1
	for i := 1; i < 12; i++ {
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
