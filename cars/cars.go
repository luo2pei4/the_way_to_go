package main

import "fmt"

// Any 通用空指针
type Any interface{}

// Car 结构体
type Car struct {
	Model        string
	Manufacturer string
	BuildYear    int
}

// CarSlice Car类型的切片
type CarSlice []*Car

// Process CarSlice的方法
func (cs CarSlice) Process(f func(car *Car)) {

	for _, c := range cs {

		f(c)
	}
}

// FindAll CarSlice的方法
func (cs CarSlice) FindAll(f func(car *Car) bool) CarSlice {

	cars := make([]*Car, 0)

	cs.Process(func(c *Car) {

		if f(c) {
			cars = append(cars, c)
		}
	})

	return cars
}

// Map Process cars and create new data.
func (cs CarSlice) Map(f func(car *Car) Any) []Any {
	result := make([]Any, len(cs))
	ix := 0
	cs.Process(func(c *Car) {
		result[ix] = f(c)
		ix++
	})
	return result
}

// MakeSortedAppender 排序函数
func MakeSortedAppender(manufacturers []string) (func(car *Car), map[string]CarSlice) {
	// Prepare maps of sorted cars.
	sortedCars := make(map[string]CarSlice)

	for _, m := range manufacturers {
		sortedCars[m] = make([]*Car, 0)
	}
	sortedCars["Default"] = make([]*Car, 0)

	// Prepare appender function:
	appender := func(c *Car) {
		if _, ok := sortedCars[c.Manufacturer]; ok {
			sortedCars[c.Manufacturer] = append(sortedCars[c.Manufacturer], c)
		} else {
			sortedCars["Default"] = append(sortedCars["Default"], c)
		}
	}
	return appender, sortedCars
}

func main() {

	// make some cars:
	ford := &Car{"Fiesta", "Ford", 2008}
	bmw := &Car{"XL 450", "BMW", 2011}
	merc := &Car{"D600", "Mercedes", 2009}
	bmw2 := &Car{"X 800", "BMW", 2008}
	// query:
	allCars := CarSlice([]*Car{ford, bmw, merc, bmw2})

	// 调用FindAll方法时，传入了一个返回值为bool类型的匿名函数（闭包）
	allNewBMWs := allCars.FindAll(func(car *Car) bool {
		return (car.Manufacturer == "BMW") && (car.BuildYear > 2010)
	})
	fmt.Println("AllCars: ", allCars)
	fmt.Println("New BMWs: ", allNewBMWs)

	manufacturers := []string{"Ford", "Aston Martin", "Land Rover", "BMW", "Jaguar"}
	sortedAppender, sortedCars := MakeSortedAppender(manufacturers)
	allCars.Process(sortedAppender)
	fmt.Println("Map sortedCars: ", sortedCars)
	BMWCount := len(sortedCars["BMW"])
	fmt.Println("We have ", BMWCount, " BMWs")
}
