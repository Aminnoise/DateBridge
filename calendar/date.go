package calendar

type Calendar interface {
	ToDays() int
	FromDays(days int) Calendar
	Kabise(year int) bool
}
