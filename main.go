package main

import "fmt"

var jalaliDaysInMonth = []int{0, 31, 31, 31, 31, 31, 31, 30, 30, 30, 30, 30, 29}
var miladiDaysInMonth = []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

type Date interface {
	Kabise(year int) bool
	Convert(days int) (int, int, int)
	Days(y, m, d int) (day int)
}

// -----------------------------------------------------------------------------

// ساختار تقویم شمسی
type jalali struct {
	day   int
	month int
	year  int
}

// به دست آوردن روزها از مبدا شمسی با زمان میلادی

func (j *jalali) Days(g Gregorian) (day int) {
	origin := 622 * 365
	origin += 150 //روزهای کبیسه
	origin += 59  // روزهای ماه ها
	origin += 22  // روز های اضافه
	//----------------------------
	now := g.year * 365
	now += g.year / 4
	now -= g.year / 100
	now += g.year / 400
	for i := 1; i < g.month; i++ {
		now += miladiDaysInMonth[i]
	}
	now += g.day
	//----------------------------
	day = now - origin
	return day + 1
}

// محاسیه سال های کبیسه شمسی

func (j *jalali) Kabise(year int) bool {
	kabiseYears := []int{1, 5, 9, 13, 17, 22, 26, 30}
	k := year % 33
	for _, v := range kabiseYears {
		if k == v {
			return true
		}
	}
	return false
}

// تبدیل روز ها به تاریخ شمسی

func (j *jalali) Convert(days int) (a jalali) {
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
			month = i
			break
		}
		days -= daysInMonth
	}

	a.year = year
	a.month = month
	a.day = days

	// 👇 اینجا همون مشکل اصلی بود
	return a
}

// ------------------------------------------------------------------------------

// ساختار تقویم میلادی

type Gregorian struct {
	day   int
	month int
	year  int
}

// به دست آوردن روزها از مبدا شمسی با زمان شمسی

func (g *Gregorian) Days(j jalali) (day int) {

	for i := 1; i < j.year; i++ {
		day += 356
		if j.Kabise(j.year) && i != 1 {
			day++
		}
	}
	for i := 1; i < j.month; i++ {
		day += jalaliDaysInMonth[i]
	}
	day += j.day
	return day
}

// محاسیه سال های کبیسه میلادی

func (g *Gregorian) Kabise(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			}
			return false
		}
		return true
	}
	return false
}

// تبدیل روز ها به تاریخ میلادی

func (g *Gregorian) Convert(days int) (a Gregorian) {

	year := 1
	for {
		daysInYear := 365
		if g.Kabise(year) && year != 1 {
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
		if i == 12 && g.Kabise(year) {
			daysInMonth = 30
		}
		if days < daysInMonth {
			month = i
			break
		}
		days -= daysInMonth
	}

	day := days

	a.day = day
	a.month = month
	a.year = year + 622

	return a
}

func main() {
	j := new(Gregorian)
	g := jalali{
		year:  145,
		month: 6,
		day:   9,
	}
	days := j.Days(g)
	fmt.Printf("jalali days: %d\n", days)
	fmt.Println(j.Convert(days))
}
