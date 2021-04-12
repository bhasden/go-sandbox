package calendar

import "fmt"

func _() {
	regularYears := make(map[int]int)
	leapYears := make(map[int]int)
	for day := 0; day < 7; day++ {
		count, _ := CountFirstSundaysInYear(1937, day, 0)
		regularYears[day] = count
	}
	for day := 0; day < 7; day ++ {
		count, _ := CountFirstSundaysInYear(1996, day, 0)
		leapYears[day] = count
	}
	fmt.Println("pause")
}

// In production Go code we'd have this be in a separate pkg, not in main!
func CountSundaysOld() int {
	currDayOfWeek := 1
	sundaysCount := 0
	for y := 1900; y <= 2000; y++ {
		sundaysCount, currDayOfWeek = CountFirstSundaysInYear(y, currDayOfWeek, sundaysCount)
	}
	return sundaysCount
}

func CountFirstSundaysInYear(year int, currDayOfWeek int, sundaysCount int) (int, int) {
	for m := 0; m < 12; m++ {
		if currDayOfWeek%7 == 0 {
			sundaysCount++
		}
		currDayOfWeek = (currDayOfWeek + DaysInMonth(m, year)) % 7
	}
	return sundaysCount, currDayOfWeek
}

func DaysInMonth(m, y int) int {
	if m == 1 {
		if isLeap(y) {
			return 29
		}
		return 28
	}
	if m == 0 || m == 2 || m == 4 || m == 6 || m == 7 || m == 9 || m == 11 {
		return 31
	}
	return 30
}
