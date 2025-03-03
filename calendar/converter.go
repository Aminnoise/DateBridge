package calendar

type Converter struct{}

func (c Converter) Convert(date Calendar, target Calendar) Calendar {
	days := date.ToDays()
	return target.FromDays(days)
}
