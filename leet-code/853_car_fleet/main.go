package _53_car_fleet

import "sort"

type Car struct {
	Pos   int
	Speed int
}

func carFleet(target int, position []int, speed []int) int {
	cars := make([]*Car, len(position))
	for i, v := range position {
		cars[i] = &Car{Pos: v, Speed: speed[i]}
	}
	// O(nlogn)
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].Pos > cars[j].Pos
	})

	times := make([]float32, len(cars))
	for i, v := range cars {
		times[i] = float32((target - v.Pos)) / float32(v.Speed)
	}
	stack := []float32{times[0]}
	for i := 1; i < len(times); i++ {
		if stack[len(stack)-1] < times[i] {
			stack = append(stack, times[i])
		}
	}
	return len(stack)
}
