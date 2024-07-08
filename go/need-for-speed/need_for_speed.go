package speed

type Car struct {
	battery int
	batteryDrain int
	speed int
	distance int
}

func NewCar(speed, batteryDrain int) Car {
	return Car {
		battery: 100,
		batteryDrain: batteryDrain,
		speed: speed,
		distance: 0,
	}
}

type Track struct {
	distance int
}

func NewTrack(distance int) Track {
	return Track {
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if car.battery < car.batteryDrain {
		return car
	}
	return Car {
		battery: car.battery - car.batteryDrain,
		distance: car.distance + car.speed,
		batteryDrain: car.batteryDrain,
		speed: car.speed,
	}
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	return (car.battery / car.batteryDrain) * car.speed >= track.distance
}
