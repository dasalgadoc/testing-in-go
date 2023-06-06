package domain

type shift map[string][2]int

func generateShift() shift {
	return shift{
		"morning":   {6, 11},
		"afternoon": {12, 17},
		"night":     {18, 21},
	}
}

func isShiftRegistered(value string) bool {
	shift := generateShift()
	_, ok := shift[value]
	return ok
}

func getShiftFromHour(hour int) string {
	shift := generateShift()
	for key, value := range shift {
		if hour >= value[0] && hour <= value[1] {
			return key
		}
	}
	return ""
}
