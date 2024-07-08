package birdwatcher

func TotalBirdCount(birdsPerDay []int) int {
	var result int = 0

	for i := 0; i < len(birdsPerDay); i++ {
		result += birdsPerDay[i]
	}

	return result
}

func BirdsInWeek(birdsPerDay []int, week int) int {
	begin := (week -1) * 7
	end := (week * 7)

	if end > len(birdsPerDay) {
		return TotalBirdCount(birdsPerDay[begin:])
	} else {
		return TotalBirdCount(birdsPerDay[begin:end])
	}
}

func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i+= 1 {
		if i % 2 == 0 {
			birdsPerDay[i] += 1
		}
	}

	return birdsPerDay
}
