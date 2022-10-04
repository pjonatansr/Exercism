package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100.0)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(float64(productionRate) * successRate / 100.0) / 60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	singleCostCarsCount := carsCount % 10
	totalSingleCost := singleCostCarsCount * 10000
	groupCostCarsCount := (carsCount - singleCostCarsCount) / 10
	totalGroupCost :=  groupCostCarsCount * 95000

	return uint(totalSingleCost + totalGroupCost)
}
