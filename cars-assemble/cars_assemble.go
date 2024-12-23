package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	floatProductionRate := float64(productionRate)
	return (successRate / 100) * floatProductionRate
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	carsPerHour := CalculateWorkingCarsPerHour(productionRate, successRate)
	carsPerMinute := carsPerHour / 60
	return int(carsPerMinute)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	groupOfTenCost := 95_000
	individualCost := 10_000

	groupsOfTen := carsCount / 10
	individuals := carsCount % 10

	groupsOfTenCost := groupsOfTen * groupOfTenCost
	individualsCost := individualCost * individuals

	return uint(groupsOfTenCost + individualsCost)
}
