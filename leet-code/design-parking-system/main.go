package main

type ParkingSystem struct {
	small, medium, big int
	status             []int
}

const (
	BIG    = 1
	MEDIUM = 2
	SMALL  = 3
)

func Constructor(big int, medium int, small int) ParkingSystem {
	status := []int{0, 0, 0}
	return ParkingSystem{
		small:  small,
		medium: medium,
		big:    big,
		status: status,
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case BIG:
		if this.status[0] == this.big {
			return false
		}
		this.status[0]++
		return true
	case MEDIUM:
		if this.status[1] == this.medium {
			return false
		}
		this.status[1]++
		return true
	case SMALL:
		if this.status[2] == this.small {
			return false
		}
		this.status[2]++
		return true
	}
	panic("impossible")
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
