package main

// required ingredient
func compilePossibleRecipes(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) []*RecipeDetails {
	possibleRecipes := []*RecipeDetails{}

	baconAndEggs := baconAndEggsPossible(crockPot, attributeVals, attributeCounts)
	if baconAndEggs != nil {
		// if not using pointers:
		// if baconAndEggs.Name != "" {
		possibleRecipes = append(possibleRecipes, baconAndEggs)
	}

	// for _, recipe := range possibleRecipes {
	// 	fmt.Println(recipe.NAME)
	// }
	return possibleRecipes

}

func amberosiaPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// AMBEROSIA
	amberosia := recipeData["AMBEROSIA"]
	if countIngName("Collected Dust", crockPot) < 1 {
		return nil
	}
	return &amberosia
}

func baconAndEggsPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	baconAndEggs := recipeData["BACON AND EGGS"]

	if attributeVals.vegVal != 0 {
		return nil
	}
	if attributeVals.eggVal <= 1 {
		baconAndEggs.UNMETCONDITIONS = append(baconAndEggs.UNMETCONDITIONS, "ev")
	}
	if attributeVals.meatVal <= 1 {
		baconAndEggs.UNMETCONDITIONS = append(baconAndEggs.UNMETCONDITIONS, "mv")
	}

	return &baconAndEggs
}
