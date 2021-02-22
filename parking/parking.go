package parking

type ParkingSystem struct {
	slots map[int]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		slots: map[int]int{
			1: big,
			2: medium,
			3: small,
		},
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.slots[carType] <= 0 {
		return false
	}
	this.slots[carType]--
	return true
}
