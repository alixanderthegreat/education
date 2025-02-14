package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"dozen":              12,
		"quarter_of_a_dozen": 12 / 4,
		"half_of_a_dozen":    12 / 2,
		"small_gross":        12 * 10,
		"gross":              12 * 12,
		"great_gross":        12 * 12 * 12,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	amnt, ok := units[unit]
	if !ok {
		return ok
	}
	bill[item] += amnt // addition assignment
	_, checker := bill[item]
	return checker
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	qty, ok := bill[item]
	if !ok {
		return ok
	}
	amnt, ok := units[unit]
	if !ok {
		return ok
	}
	if qty-amnt < 0 {
		return false
	} else if qty-amnt == 0 {
		delete(bill, item)
		return true
	} else {
		bill[item] -= units[unit]
		_, checker := bill[item]
		return checker
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	panic("Please implement the GetItem() function")
}
