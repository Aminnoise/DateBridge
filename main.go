package main

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

func (j *jalali) Days(y, m, d int) (day int) {
	origin := 622 * 365
	origin += 150 //روزهای کبیسه
	origin += 59  // روزهای ماه ها
	origin += 22  // روز های اضافه
	//----------------------------
	now := y * 365
	now += y / 4
	now -= y / 100
	now += y / 400
	for i := 1; i < m; i++ {
		now += miladiDaysInMonth[i]
	}
	now += d
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

func (j *jalali) Convert(days int) (int, int, int) {
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

	// 👇 اینجا همون مشکل اصلی بود
	return year, month, days
}

// ------------------------------------------------------------------------------

// ساختار تقویم میلادی

type Gregorian struct {
	day   int
	month int
	year  int
}

// به دست آوردن روزها از مبدا شمسی با زمان شمسی

func (g *Gregorian) Days(y, m, d int) (day int) {
	return day
}

// محاسیه سال های کبیسه میلادی

func (g *Gregorian) Kabise(year int) bool {
	return false
}

// تبدیل روز ها به تاریخ میلادی

func (g *Gregorian) Convert(days int) (year, month, day int) {
	return year, month, day
}

func main() {

}
