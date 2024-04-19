package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{}
	units["quarter_of_a_dozen"] = 3
	units["half_of_a_dozen"] = 6
	units["dozen"] = 12
	units["small_gross"] = 120
	units["gross"] = 144
	units["great_gross"] = 1728
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, exists := units[unit]

	if !exists {
		return exists
	}
	bill[item] += units[unit]

	return exists
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, itemExists := bill[item]
	_, unitExists := units[unit]

	if !itemExists || !unitExists {
		return false
	}

	if newUnit := bill[item] - units[unit]; newUnit < 0 {
		return false
	} else if newUnit == 0 {
		delete(bill, item)
	} else if newUnit > 0 {
		bill[item] = newUnit
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	qtd, exists := bill[item]
	return qtd, exists
}
