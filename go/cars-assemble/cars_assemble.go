package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	successRatePercentage := successRate / 100
	return float64(productionRate) * successRatePercentage
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	groupNotTenOfCar := map[string]float64{
		"total": float64(carsCount % 10),
		"cost":  10000,
	}

	groupTenOfCar := map[string]float64{
		"total": float64(carsCount) - groupNotTenOfCar["total"],
		"cost":  95000,
	}

	totalCost := float64(0)
	totalCostGroupTenOfCar := float64(0)
	totalCostGroupNotTenOfCar := groupNotTenOfCar["total"] * groupNotTenOfCar["cost"]

	if groupTenOfCar["total"] > 0 {
		totalCostGroupTenOfCar = groupTenOfCar["total"] / 10 * groupTenOfCar["cost"]
	}

	totalCost = totalCostGroupTenOfCar + totalCostGroupNotTenOfCar

	return uint(totalCost)
}
