package main

// required ingredient
func stringInSlice(a string, x []IngredientDetails) bool {
	for _, x := range x {
		if x.NAME == a {
			return true
		}
	}
	return false
}

// excluded ingredient
func stringNotInSlice(a string, x []IngredientDetails) bool {
	for _, x := range x {
		if x.NAME == a {
			return false
		}
	}
	return true
}

// count strings in cropost array
func countIngName(x string, crockPot []IngredientDetails) int {
	count := 0
	for _, cp := range crockPot {
		if cp.NAME == x {
			count++
		}
	}
	return count
}
