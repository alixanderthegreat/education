package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) (need bool) {
	return (kind == "car" || kind == "truck")
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) (msg string) {
	msg = " is clearly the better choice."
	if option1 < option2 {
		return option1 + msg
	}
	return option2 + msg
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) (price float64) {
	if age < 3 {
		price = originalPrice * .8
	} else if age >= 10 {
		price = originalPrice * .5
	} else {
		price = originalPrice * .7
	}
	return price
}
