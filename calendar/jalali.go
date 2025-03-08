package calendar

var jalaliDaysInMonth = []int{0,
	31, 31, 31, 31, 31, 31,
	30, 30, 30, 30, 30, 29}

var jalaliDaySUMMonth = []int{0, // آرایه مجموع روزهای ماه‌های تقویم جلالی
	31, 62, 93, 124, 155, 186,
	216, 246, 276, 306, 336, 365}

var kabise33 = []int{
	0, 1, 1, 1, 1, 2, 2, 2, 2,
	3, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5,
	5, 6, 6, 6, 6, 7, 7, 7, 7, 8, 8, 8}

var kabiseYears = []int{1, 5, 9, 13, 17, 22, 26, 30}

type Jalali struct {
	Year  int
	Month int
	Day   int
}

func (j Jalali) ToDays() int {
	days := j.Day
	days += jalaliDaySUMMonth[j.Month-1]
	dorekamel := (j.Year - 1) / 33
	days += (j.Year - 1) * 365
	days += dorekamel * 8
	days += kabise33[j.Year-(dorekamel*33)]
	days-- // چون سال اول شمسی کبیسه نیس

	return days
}

func (j Jalali) FromDays(days int) Calendar {
	year := 1
	year += (days / 12053) * 33
	days = days % 12053
	for i := 1; i <= 8; i++ {
		if days < (365*kabiseYears[i])+(i*1) {
			days -= 366 * (i - 1)
			year += i - 1
			year += days / 365
			days = days % 365
			break
		}
	}

	month := 1
	for i := 1; i < 12; i++ {
		daysInMonth := jalaliDaysInMonth[i]
		/*	if i == 12 && j.Kabise(year) {
			daysInMonth = 30
		}*/
		if days < daysInMonth {
			break
		}
		days -= daysInMonth
		month++
	}

	return Jalali{Year: year, Month: month, Day: days}
}

func (j Jalali) Kabise(year int) bool {
	k := year % 33
	return binarySearch(k, 0, 7)
}

func binarySearch(x int, begin, end int) bool {
	if end < begin {
		return false
	}
	mid := (end + begin) / 2
	if kabiseYears[mid] == x {
		return true
	} else if kabiseYears[mid] > x {
		return binarySearch(x, begin, mid-1)
	} else {
		return binarySearch(x, mid+1, end)
	}
}
