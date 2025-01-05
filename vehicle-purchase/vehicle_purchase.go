package purchase

import "fmt"

var licenseRequiredVehicles = [2]string{"car", "truck"}

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	licenseRequired := false
	for i := 0; i < len(licenseRequiredVehicles); i++ {
		if kind == licenseRequiredVehicles[i] {
			licenseRequired = true
			break
		}
	}
	return licenseRequired
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	suffix := "is clearly the better choice."
	var betterChoice string

	if option1 < option2 {
		betterChoice = option1
	} else {
		betterChoice = option2
	}

	return fmt.Sprint(betterChoice, " ", suffix)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var remainingPercentage float64
	if age < 3 {
		remainingPercentage = 0.8
	} else if age < 10 {
		remainingPercentage = 0.7
	} else {
		remainingPercentage = 0.5
	}
	return originalPrice * remainingPercentage
}
