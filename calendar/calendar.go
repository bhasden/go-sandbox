/*
January 1, 1900, was a Monday.
April, June, September, and November each have 30 days.
January, March, May, July, August, October, and December each have 31 days.
On years evenly divisible by 4 February has 29 days except on a century unless the year is divisible by 400.
On all other years February has 28 days.

Without using any date/time libraries, figure out how many Sundays fell on the first day of the month
during the twentieth century (January 1, 1901, to December 31, 2000).
*/

// Example calendar.go

package calendar

var firstDayToFirstSundaysLeapYears = [...]int{3, 2, 1, 2, 2, 1, 1}
var firstDayToFirstSundaysRegularYears = [...]int{2, 2, 2, 1, 3, 1, 1}

var wrappingLogic = [...]int{0, 1, 2, 3, 4, 5, 6, 0, 1}

func CountSundays() int {
	// Jan 1 1901 was tuesday
	dayOfWeek := 2
	sundayCount := 0
	var sundays int
	for year := 1; year < 101; year++ {
		sundays, dayOfWeek = numberOfFirstSundays(dayOfWeek, year)
		sundayCount += sundays
	}
	return sundayCount
}

func numberOfFirstSundays(firstDay int, year int) (firstSundays int, nextFirstDay int) {
	if isLeap(year) {
		return firstDayToFirstSundaysLeapYears[firstDay], wrappingLogic[firstDay+2]
	}
	return firstDayToFirstSundaysRegularYears[firstDay], wrappingLogic[firstDay+1]
}

func isLeap(y int) bool {
	if y%4 == 0 {
		if y%100 == 0 {
			return y%400 == 0
		}
		return true
	}
	return false
}
