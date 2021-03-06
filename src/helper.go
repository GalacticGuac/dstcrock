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
