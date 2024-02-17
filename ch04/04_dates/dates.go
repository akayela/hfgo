package dates

const DaysInWeek = 7

func WeekToDays(week int) int {
	return week * DaysInWeek
}

func DaysToWeek(days int) float64 {
	return float64(days) / float64(DaysInWeek)
}
