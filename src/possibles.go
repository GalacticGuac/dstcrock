package main

// required ingredient
func compilePossibleRecipes(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) []*RecipeDetails {
	possibleRecipes := []*RecipeDetails{}

	// baconAndEggs := baconAndEggsPossible(crockPot, attributeVals, attributeCounts)
	// if baconAndEggs != nil {
	// 	possibleRecipes = append(possibleRecipes, baconAndEggs)
	// }
	amberosia := amberosiaPossible(crockPot, attributeVals, attributeCounts)
	if amberosia != nil {
		possibleRecipes = append(possibleRecipes, amberosia)
	}
	asparagusSoup := asparagusSoupPossible(crockPot, attributeVals, attributeCounts)
	if asparagusSoup != nil {
		possibleRecipes = append(possibleRecipes, asparagusSoup)
	}
	baconAndEggs := baconAndEggsPossible(crockPot, attributeVals, attributeCounts)
	if baconAndEggs != nil {
		possibleRecipes = append(possibleRecipes, baconAndEggs)
	}
	bananaPop := bananaPopPossible(crockPot, attributeVals, attributeCounts)
	if bananaPop != nil {
		possibleRecipes = append(possibleRecipes, bananaPop)
	}
	barnacleLinguine := barnacleLinguinePossible(crockPot, attributeVals, attributeCounts)
	if barnacleLinguine != nil {
		possibleRecipes = append(possibleRecipes, barnacleLinguine)
	}
	barnacleNigiri := barnacleNigiriPossible(crockPot, attributeVals, attributeCounts)
	if barnacleNigiri != nil {
		possibleRecipes = append(possibleRecipes, barnacleNigiri)
	}
	barnaclePita := barnaclePitaPossible(crockPot, attributeVals, attributeCounts)
	if barnaclePita != nil {
		possibleRecipes = append(possibleRecipes, barnaclePita)
	}
	beefyGreens := beefyGreensPossible(crockPot, attributeVals, attributeCounts)
	if beefyGreens != nil {
		possibleRecipes = append(possibleRecipes, beefyGreens)
	}
	butterMuffin := butterMuffinPossible(crockPot, attributeVals, attributeCounts)
	if butterMuffin != nil {
		possibleRecipes = append(possibleRecipes, butterMuffin)
	}
	californiaRoll := californiaRollPossible(crockPot, attributeVals, attributeCounts)
	if californiaRoll != nil {
		possibleRecipes = append(possibleRecipes, californiaRoll)
	}
	ceviche := cevichePossible(crockPot, attributeVals, attributeCounts)
	if ceviche != nil {
		possibleRecipes = append(possibleRecipes, ceviche)
	}
	creamyPotatoPuree := creamyPotatoPureePossible(crockPot, attributeVals, attributeCounts)
	if creamyPotatoPuree != nil {
		possibleRecipes = append(possibleRecipes, creamyPotatoPuree)
	}
	dragonpie := dragonpiePossible(crockPot, attributeVals, attributeCounts)
	if dragonpie != nil {
		possibleRecipes = append(possibleRecipes, dragonpie)
	}
	fancySpiralledTubers := fancySpiralledTubersPossible(crockPot, attributeVals, attributeCounts)
	if fancySpiralledTubers != nil {
		possibleRecipes = append(possibleRecipes, fancySpiralledTubers)
	}
	fishTacos := fishTacosPossible(crockPot, attributeVals, attributeCounts)
	if fishTacos != nil {
		possibleRecipes = append(possibleRecipes, fishTacos)
	}
	fishsticks := fishsticksPossible(crockPot, attributeVals, attributeCounts)
	if fishsticks != nil {
		possibleRecipes = append(possibleRecipes, fishsticks)
	}
	fistFullOfJam := fistFullOfJamPossible(crockPot, attributeVals, attributeCounts)
	if fistFullOfJam != nil {
		possibleRecipes = append(possibleRecipes, fistFullOfJam)
	}
	flowerSalad := flowerSaladPossible(crockPot, attributeVals, attributeCounts)
	if flowerSalad != nil {
		possibleRecipes = append(possibleRecipes, flowerSalad)
	}
	froggleBunwich := froggleBunwichPossible(crockPot, attributeVals, attributeCounts)
	if froggleBunwich != nil {
		possibleRecipes = append(possibleRecipes, froggleBunwich)
	}
	fruitMedley := fruitMedleyPossible(crockPot, attributeVals, attributeCounts)
	if fruitMedley != nil {
		possibleRecipes = append(possibleRecipes, fruitMedley)
	}
	guacamole := guacamolePossible(crockPot, attributeVals, attributeCounts)
	if guacamole != nil {
		possibleRecipes = append(possibleRecipes, guacamole)
	}
	honeyHam := honeyHamPossible(crockPot, attributeVals, attributeCounts)
	if honeyHam != nil {
		possibleRecipes = append(possibleRecipes, honeyHam)
	}
	honeyNuggets := honeyNuggetsPossible(crockPot, attributeVals, attributeCounts)
	if honeyNuggets != nil {
		possibleRecipes = append(possibleRecipes, honeyNuggets)
	}
	iceCream := iceCreamPossible(crockPot, attributeVals, attributeCounts)
	if iceCream != nil {
		possibleRecipes = append(possibleRecipes, iceCream)
	}
	jellySalad := jellySaladPossible(crockPot, attributeVals, attributeCounts)
	if jellySalad != nil {
		possibleRecipes = append(possibleRecipes, jellySalad)
	}
	jellybeans := jellybeansPossible(crockPot, attributeVals, attributeCounts)
	if jellybeans != nil {
		possibleRecipes = append(possibleRecipes, jellybeans)
	}
	kabobs := kabobsPossible(crockPot, attributeVals, attributeCounts)
	if kabobs != nil {
		possibleRecipes = append(possibleRecipes, kabobs)
	}
	leafyMeatloaf := leafyMeatloafPossible(crockPot, attributeVals, attributeCounts)
	if leafyMeatloaf != nil {
		possibleRecipes = append(possibleRecipes, leafyMeatloaf)
	}
	lobsterBisque := lobsterBisquePossible(crockPot, attributeVals, attributeCounts)
	if lobsterBisque != nil {
		possibleRecipes = append(possibleRecipes, lobsterBisque)
	}
	mandrakeSoup := mandrakeSoupPossible(crockPot, attributeVals, attributeCounts)
	if mandrakeSoup != nil {
		possibleRecipes = append(possibleRecipes, mandrakeSoup)
	}
	meatballs := meatballsPossible(crockPot, attributeVals, attributeCounts)
	if meatballs != nil {
		possibleRecipes = append(possibleRecipes, meatballs)
	}
	meatyStew := meatyStewPossible(crockPot, attributeVals, attributeCounts)
	if meatyStew != nil {
		possibleRecipes = append(possibleRecipes, meatyStew)
	}
	melonsicle := melonsiclePossible(crockPot, attributeVals, attributeCounts)
	if melonsicle != nil {
		possibleRecipes = append(possibleRecipes, melonsicle)
	}
	milkmadeHat := milkmadeHatPossible(crockPot, attributeVals, attributeCounts)
	if milkmadeHat != nil {
		possibleRecipes = append(possibleRecipes, milkmadeHat)
	}
	monsterLasagna := monsterLasagnaPossible(crockPot, attributeVals, attributeCounts)
	if monsterLasagna != nil {
		possibleRecipes = append(possibleRecipes, monsterLasagna)
	}
	mushyCake := mushyCakePossible(crockPot, attributeVals, attributeCounts)
	if mushyCake != nil {
		possibleRecipes = append(possibleRecipes, mushyCake)
	}
	pierogi := pierogiPossible(crockPot, attributeVals, attributeCounts)
	if pierogi != nil {
		possibleRecipes = append(possibleRecipes, pierogi)
	}
	powdercake := powdercakePossible(crockPot, attributeVals, attributeCounts)
	if powdercake != nil {
		possibleRecipes = append(possibleRecipes, powdercake)
	}
	pumkinCookie := pumkinCookiePossible(crockPot, attributeVals, attributeCounts)
	if pumkinCookie != nil {
		possibleRecipes = append(possibleRecipes, pumkinCookie)
	}
	ratatouille := ratatouillePossible(crockPot, attributeVals, attributeCounts)
	if ratatouille != nil {
		possibleRecipes = append(possibleRecipes, ratatouille)
	}
	salsaFresca := salsaFrescaPossible(crockPot, attributeVals, attributeCounts)
	if salsaFresca != nil {
		possibleRecipes = append(possibleRecipes, salsaFresca)
	}
	seafoodGumbo := seafoodGumboPossible(crockPot, attributeVals, attributeCounts)
	if seafoodGumbo != nil {
		possibleRecipes = append(possibleRecipes, seafoodGumbo)
	}
	soothingTea := soothingTeaPossible(crockPot, attributeVals, attributeCounts)
	if soothingTea != nil {
		possibleRecipes = append(possibleRecipes, soothingTea)
	}
	spicyChili := spicyChiliPossible(crockPot, attributeVals, attributeCounts)
	if spicyChili != nil {
		possibleRecipes = append(possibleRecipes, spicyChili)
	}
	stuffedEggplant := stuffedEggplantPossible(crockPot, attributeVals, attributeCounts)
	if stuffedEggplant != nil {
		possibleRecipes = append(possibleRecipes, stuffedEggplant)
	}
	stuffedFishHeads := stuffedFishHeadsPossible(crockPot, attributeVals, attributeCounts)
	if stuffedFishHeads != nil {
		possibleRecipes = append(possibleRecipes, stuffedFishHeads)
	}
	stuffedPepperPoppers := stuffedPepperPoppersPossible(crockPot, attributeVals, attributeCounts)
	if stuffedPepperPoppers != nil {
		possibleRecipes = append(possibleRecipes, stuffedPepperPoppers)
	}
	surfNTurf := surfNTurfPossible(crockPot, attributeVals, attributeCounts)
	if surfNTurf != nil {
		possibleRecipes = append(possibleRecipes, surfNTurf)
	}
	trailMix := trailMixPossible(crockPot, attributeVals, attributeCounts)
	if trailMix != nil {
		possibleRecipes = append(possibleRecipes, trailMix)
	}
	turkeyDinner := turkeyDinnerPossible(crockPot, attributeVals, attributeCounts)
	if turkeyDinner != nil {
		possibleRecipes = append(possibleRecipes, turkeyDinner)
	}
	unagi := unagiPossible(crockPot, attributeVals, attributeCounts)
	if unagi != nil {
		possibleRecipes = append(possibleRecipes, unagi)
	}
	vegetableStinger := vegetableStingerPossible(crockPot, attributeVals, attributeCounts)
	if vegetableStinger != nil {
		possibleRecipes = append(possibleRecipes, vegetableStinger)
	}
	veggieBurger := veggieBurgerPossible(crockPot, attributeVals, attributeCounts)
	if veggieBurger != nil {
		possibleRecipes = append(possibleRecipes, veggieBurger)
	}
	waffles := wafflesPossible(crockPot, attributeVals, attributeCounts)
	if waffles != nil {
		possibleRecipes = append(possibleRecipes, waffles)
	}
	wetGoop := wetGoopPossible(crockPot, attributeVals, attributeCounts)
	if wetGoop != nil {
		possibleRecipes = append(possibleRecipes, wetGoop)
	}
	wobsterDinner := wobsterDinnerPossible(crockPot, attributeVals, attributeCounts)
	if wobsterDinner != nil {
		possibleRecipes = append(possibleRecipes, wobsterDinner)
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

func asparagusSoupPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// ASPARAGUS SOUP
	asparagusSoup := recipeData["ASPARAGUS SOUP"]
	if attributeCounts.meatCount != 0 {
		return nil
	}
	if attributeCounts.inedibleCount != 0 {
		return nil
	}
	if countIngName("Asparagus", crockPot) < 1 {
		asparagusSoup.UNMETCONDITIONS = append(asparagusSoup.UNMETCONDITIONS, "Asparagus")
	}

	if attributeVals.vegVal < 1.5 {
		asparagusSoup.UNMETCONDITIONS = append(asparagusSoup.UNMETCONDITIONS, "ev")
	}
	return &asparagusSoup
}

func baconAndEggsPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// BACON AND EGGS
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
func bananaPopPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// BANANA POP
	bananaPop := recipeData["BANANA POP"]
	if attributeCounts.meatCount != 0 {
		return nil
	}
	if attributeCounts.fishCount != 0 {
		return nil
	}
	if countIngName("Banana", crockPot) < 1 {
		bananaPop.UNMETCONDITIONS = append(bananaPop.UNMETCONDITIONS, "Banana")
	}
	if countIngName("Twigs", crockPot) < 1 {
		bananaPop.UNMETCONDITIONS = append(bananaPop.UNMETCONDITIONS, "Twigs")
	}

	return &bananaPop
}

func barnacleLinguinePossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// BARNACLE LINGUINE
	barnacleLinguine := recipeData["BARNACLE LINGUINE"]
	x := countIngName("Barnacles", crockPot)
	if x != 2 {
		switch x {
		case 0:
			barnacleLinguine.UNMETCONDITIONS = append(barnacleLinguine.UNMETCONDITIONS, "2 Barnacles")
		case 1:
			barnacleLinguine.UNMETCONDITIONS = append(barnacleLinguine.UNMETCONDITIONS, "1 Barnacles")
		default:
			return nil
		}
	}
	if attributeVals.vegVal < 2 {
		barnacleLinguine.UNMETCONDITIONS = append(barnacleLinguine.UNMETCONDITIONS, "ev")
	}
	return &barnacleLinguine
}

func barnacleNigiriPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// BARNACLE NIGIRI
	barnacleNigiri := recipeData["BARNACLE NIGIRI"]
	if countIngName("Barnacles", crockPot) < 1 {
		barnacleNigiri.UNMETCONDITIONS = append(barnacleNigiri.UNMETCONDITIONS, "Barnacles")
	}
	if countIngName("Kelp Fronds", crockPot) < 1 {
		barnacleNigiri.UNMETCONDITIONS = append(barnacleNigiri.UNMETCONDITIONS, "Kelp Fronds")
	}
	if attributeCounts.eggCount < 1 {
		barnacleNigiri.UNMETCONDITIONS = append(barnacleNigiri.UNMETCONDITIONS, "ev")
	}
	return &barnacleNigiri
}

func barnaclePitaPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// BARNACLE PITA
	barnaclePita := recipeData["BARNACLE PITA"]
	if countIngName("Barnacles", crockPot) < 1 {
		barnaclePita.UNMETCONDITIONS = append(barnaclePita.UNMETCONDITIONS, "Barnacles")
	}
	if attributeVals.vegVal < 0.5 {
		barnaclePita.UNMETCONDITIONS = append(barnaclePita.UNMETCONDITIONS, "ev")
	}
	return &barnaclePita
}

func beefyGreensPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// BEEFY GREENS********************************************************************************* higher val needed
	beefyGreens := recipeData["BEEFY GREENS"]
	if countIngName("Leafy Meat", crockPot) < 1 {
		beefyGreens.UNMETCONDITIONS = append(beefyGreens.UNMETCONDITIONS, "Leafy Meat")
	}

	if attributeVals.vegVal < 3 {
		beefyGreens.UNMETCONDITIONS = append(beefyGreens.UNMETCONDITIONS, "ev")
	}
	return &beefyGreens
}

func butterMuffinPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// BUTTER MUFFIN
	butterMuffin := recipeData["BUTTER MUFFIN"]
	if stringInSlice("Mandrake", crockPot) {
		return nil
	}
	if attributeCounts.meatCount != 0 {
		return nil
	}
	if countIngName("Butterfly Wings", crockPot) < 1 {
		butterMuffin.UNMETCONDITIONS = append(butterMuffin.UNMETCONDITIONS, "Butterfly Wings")
	}
	if attributeCounts.vegeCount < 1 {
		butterMuffin.UNMETCONDITIONS = append(butterMuffin.UNMETCONDITIONS, "ev")
	}
	return &butterMuffin
}

func californiaRollPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// CALIFORNIA ROLL
	californiaRoll := recipeData["CALIFORNIA ROLL"]
	x := countIngName("Kelp Fronds", crockPot)
	if x < 2 {
		switch x {
		case 0:
			californiaRoll.UNMETCONDITIONS = append(californiaRoll.UNMETCONDITIONS, "2 Kelp Fronds")
		case 1:
			californiaRoll.UNMETCONDITIONS = append(californiaRoll.UNMETCONDITIONS, "1 Kelp Fronds")
		}
	}
	if attributeVals.fishVal < 1 {
		californiaRoll.UNMETCONDITIONS = append(californiaRoll.UNMETCONDITIONS, "hv")
	}
	return &californiaRoll
}

func cevichePossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// CEVICHE
	ceviche := recipeData["CEVICHE"]
	if attributeCounts.eggCount != 0 {
		return nil
	}
	if attributeCounts.inedibleCount != 0 {
		return nil
	}
	x := countIngName("Ice", crockPot)
	if x < 2 {
		switch x {
		case 0:
			ceviche.UNMETCONDITIONS = append(ceviche.UNMETCONDITIONS, "2 Ice")
		case 1:
			ceviche.UNMETCONDITIONS = append(ceviche.UNMETCONDITIONS, "1 Ice")
		}
	}
	if attributeVals.fishVal < 2 {
		ceviche.UNMETCONDITIONS = append(ceviche.UNMETCONDITIONS, "hv")
	}
	return &ceviche
}

func creamyPotatoPureePossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// CREAMY POTATO PURÉE***************************************************************** uhhhh the word PURÉE is wierd help and count
	creamyPotatoPuree := recipeData["CREAMY POTATO PURÉE"]
	if attributeCounts.meatCount != 0 {
		return nil
	}
	if stringInSlice("Twigs", crockPot) {
		return nil
	}
	x := countIngName("Potato", crockPot)
	if x < 2 {
		switch x {
		case 0:
			creamyPotatoPuree.UNMETCONDITIONS = append(creamyPotatoPuree.UNMETCONDITIONS, "2 Potato")
		case 1:
			creamyPotatoPuree.UNMETCONDITIONS = append(creamyPotatoPuree.UNMETCONDITIONS, "Potato")
		}
	}
	if countIngName("Garlic", crockPot) < 1 {
		creamyPotatoPuree.UNMETCONDITIONS = append(creamyPotatoPuree.UNMETCONDITIONS, "Garlic")
	}
	return &creamyPotatoPuree
}

func dragonpiePossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// DRAGONPIE
	dragonpie := recipeData["DRAGONPIE"]
	if attributeCounts.meatCount != 0 {
		return nil
	}
	if stringInSlice("Mandrake", crockPot) {
		return nil
	}
	if countIngName("Dragon Fruit", crockPot) < 1 {
		dragonpie.UNMETCONDITIONS = append(dragonpie.UNMETCONDITIONS, "Dragon Fruit")
	}
	return &dragonpie
}

func fancySpiralledTubersPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// FANCY SPIRALLED TUBERS*****************************************************************
	fancySpiralledTubers := recipeData["FANCY SPIRALLED TUBERS"]
	if attributeCounts.meatCount != 0 {
		return nil
	}
	if countIngName("Potato", crockPot) < 1 {
		fancySpiralledTubers.UNMETCONDITIONS = append(fancySpiralledTubers.UNMETCONDITIONS, "Potato")
	}
	if countIngName("Twigs", crockPot) < 1 {
		fancySpiralledTubers.UNMETCONDITIONS = append(fancySpiralledTubers.UNMETCONDITIONS, "Twigs")
	}
	if attributeCounts.inedibleCount-1 > 1 {
		fancySpiralledTubers.UNMETCONDITIONS = append(fancySpiralledTubers.UNMETCONDITIONS, "iC")
	}
	return &fancySpiralledTubers
}

func fishTacosPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// fish tacos
	fishTacos := recipeData["FISH TACOS"]
	if attributeVals.fishVal < 0.5 {
		fishTacos.UNMETCONDITIONS = append(fishTacos.UNMETCONDITIONS, "hv")
	}
	if countIngName("Corn", crockPot) < 1 && countIngName("Popcorn", crockPot) < 1 {
		fishTacos.UNMETCONDITIONS = append(fishTacos.UNMETCONDITIONS, "Pop/Corn")
	}
	return &fishTacos
}

func fishsticksPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// FISHSTICKS
	fishsticks := recipeData["FISHSTICKS"]
	x := countIngName("Twigs", crockPot)
	if x != 1 {
		switch x {
		case 0:
			fishsticks.UNMETCONDITIONS = append(fishsticks.UNMETCONDITIONS, "Twigs")
		case 2, 3, 4:
			return nil
		}
		return nil
	}
	if stringInSlice("Moleworm", crockPot) {
		return nil
	}
	if attributeVals.fishVal < 0.25 {
		fishsticks.UNMETCONDITIONS = append(fishsticks.UNMETCONDITIONS, "hv")
	}
	return &fishsticks
}

func fistFullOfJamPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// FIST FULL OF JAM
	fistFullOfJam := recipeData["FIST FULL OF JAM"]
	if attributeVals.meatVal != 0 {
		return nil
	}
	if attributeVals.vegVal != 0 {
		return nil
	}
	if attributeVals.inedVal != 0 {
		return nil
	}
	if stringInSlice("Dragon Fruit", crockPot) {
		return nil
	}
	if attributeVals.fruitVal < 0.5 {
		fistFullOfJam.UNMETCONDITIONS = append(fistFullOfJam.UNMETCONDITIONS, "fv")
	}
	return &fistFullOfJam
}

func flowerSaladPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// FLOWER SALAD
	flowerSalad := recipeData["FLOWER SALAD"]
	if attributeVals.meatVal != 0 {
		return nil
	}
	if attributeVals.fruitVal != 0 {
		return nil
	}
	if attributeVals.eggVal != 0 {
		return nil
	}
	if attributeVals.sweetVal != 0 {
		return nil
	}
	if stringInSlice("Twigs", crockPot) {
		return nil
	}
	if countIngName("Cactus Flower", crockPot) < 1 {
		flowerSalad.UNMETCONDITIONS = append(flowerSalad.UNMETCONDITIONS, "Cactus Flower")
	}
	if attributeVals.vegVal-.5 < 1.5 {
		flowerSalad.UNMETCONDITIONS = append(flowerSalad.UNMETCONDITIONS, "gv")
	}
	return &flowerSalad
}

func froggleBunwichPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// FROGGLE BUNWICH
	froggleBunwich := recipeData["FROGGLE BUNWICH"]
	if attributeVals.eggVal != 0 {
		return nil
	}
	if attributeVals.sweetVal != 0 {
		return nil
	}
	if stringInSlice("Mandrake", crockPot) {
		return nil
	}
	if countIngName("Frog Legs", crockPot) < 1 {
		froggleBunwich.UNMETCONDITIONS = append(froggleBunwich.UNMETCONDITIONS, "Frog Legs")
	}
	if attributeCounts.vegeCount < 1 {
		froggleBunwich.UNMETCONDITIONS = append(froggleBunwich.UNMETCONDITIONS, "gv")
	}

	return &froggleBunwich
}

func fruitMedleyPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// FRUIT MEDLEY
	fruitMedley := recipeData["FRUIT MEDLEY"]
	if attributeVals.meatVal != 0 {
		return nil
	}
	if attributeVals.vegVal != 0 {
		return nil
	}
	if stringInSlice("Dragon Fruit", crockPot) {
		return nil
	}
	if attributeVals.fruitVal < 3 {
		fruitMedley.UNMETCONDITIONS = append(fruitMedley.UNMETCONDITIONS, "fv")
	}
	return &fruitMedley
}

func guacamolePossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// guacamole
	guacamole := recipeData["GUACAMOLE"]
	if attributeVals.fruitVal != 0 {
		return nil
	}
	if countIngName("Moleworm", crockPot) < 1 {
		guacamole.UNMETCONDITIONS = append(guacamole.UNMETCONDITIONS, "Moleworm")
	}
	if countIngName("Cactus Flesh", crockPot) < 1 && countIngName("Ripe Stone Fruit", crockPot) < 1 {
		guacamole.UNMETCONDITIONS = append(guacamole.UNMETCONDITIONS, "Cactus Flesh/Ripe Stone Fruit")
	}

	return &guacamole
}

func honeyHamPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// HONEY HAM
	honeyHam := recipeData["HONEY HAM"]
	if stringInSlice("Twigs", crockPot) {
		return nil
	}
	if stringInSlice("Moleworm", crockPot) {
		return nil
	}
	if stringInSlice("Mandrake", crockPot) {
		return nil
	}
	if stringInSlice("Tallbird Egg", crockPot) {
		return nil
	}
	if countIngName("Honey", crockPot) < 1 {
		honeyHam.UNMETCONDITIONS = append(honeyHam.UNMETCONDITIONS, "Honey")
	}
	if attributeVals.meatVal <= 1.5 {
		honeyHam.UNMETCONDITIONS = append(honeyHam.UNMETCONDITIONS, "mv")
	}
	return &honeyHam
}

func honeyNuggetsPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// honeyNuggets *************************************************************************************************
	honeyNuggets := recipeData["HONEY NUGGETS"]
	if attributeVals.inedVal != 0 {
		return nil
	}
	if countIngName("Honey", crockPot) < 1 {
		honeyNuggets.UNMETCONDITIONS = append(honeyNuggets.UNMETCONDITIONS, "Honey")
	}
	if attributeVals.meatVal > 1.5 {
		honeyNuggets.UNMETCONDITIONS = append(honeyNuggets.UNMETCONDITIONS, "mV")
	}
	return &honeyNuggets
}

func iceCreamPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// ice cream
	iceCream := recipeData["ICE CREAM"]
	if attributeVals.meatVal != 0 {
		return nil
	}
	if attributeVals.vegVal != 0 {
		return nil
	}
	if attributeVals.eggVal != 0 {
		return nil
	}
	if stringInSlice("Twigs", crockPot) {
		return nil
	}
	if countIngName("Ice", crockPot) < 1 {
		iceCream.UNMETCONDITIONS = append(iceCream.UNMETCONDITIONS, "Ice")
	}
	if attributeCounts.dairyCount < 1 {
		iceCream.UNMETCONDITIONS = append(iceCream.UNMETCONDITIONS, "dc")
	}
	if attributeCounts.sweetenerCount < 1 {
		iceCream.UNMETCONDITIONS = append(iceCream.UNMETCONDITIONS, "sc")
	}
	return &iceCream
}

func jellySaladPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// jellySalad
	jellySalad := recipeData["JELLY SALAD"]
	x := countIngName("Leafy Meat", crockPot)
	if x < 2 {
		switch x {
		case 0:
			jellySalad.UNMETCONDITIONS = append(jellySalad.UNMETCONDITIONS, "2 Leafy Meat")
		case 1:
			jellySalad.UNMETCONDITIONS = append(jellySalad.UNMETCONDITIONS, "Leafy Meat")
		}
	}
	if attributeVals.sweetVal < 2 {
		jellySalad.UNMETCONDITIONS = append(jellySalad.UNMETCONDITIONS, "sv")
	}
	return &jellySalad
}

func jellybeansPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// jellybeans
	jellybeans := recipeData["JELLYBEANS"]
	if attributeVals.monVal != 0 {
		return nil
	}
	if attributeVals.inedVal != 0 {
		return nil
	}
	if countIngName("Royal Jelly", crockPot) < 1 {
		jellybeans.UNMETCONDITIONS = append(jellybeans.UNMETCONDITIONS, "Royal Jelly")
	}
	return &jellybeans
}

func kabobsPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// KABOBS
	kabobs := recipeData["KABOBS"]
	x := countIngName("Twigs", crockPot)
	if x != 1 {
		switch x {
		case 0:
			kabobs.UNMETCONDITIONS = append(kabobs.UNMETCONDITIONS, "Twigs")
		case 2:
			return nil
		case 3:
			return nil
		}
	}
	if attributeVals.fishVal != 0 {
		return nil
	}
	if stringInSlice("Moleworm", crockPot) {
		return nil
	}
	if stringInSlice("Mandrake", crockPot) {
		return nil
	}
	if attributeCounts.meatCount < 1 {
		kabobs.UNMETCONDITIONS = append(kabobs.UNMETCONDITIONS, "mc")
	}
	return &kabobs
}

func leafyMeatloafPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// LEAFY MEATLOAF
	leafyMeatloaf := recipeData["LEAFY MEATLOAF"]
	x := countIngName("Leafy Meat", crockPot)
	if x < 2 {
		switch x {
		case 0:
			leafyMeatloaf.UNMETCONDITIONS = append(leafyMeatloaf.UNMETCONDITIONS, "2 Leafy Meat")
		case 1:
			leafyMeatloaf.UNMETCONDITIONS = append(leafyMeatloaf.UNMETCONDITIONS, "Leafy Meat")
		}
	}
	return &leafyMeatloaf
}

func lobsterBisquePossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// LOBSTER BISQUE
	lobsterBisque := recipeData["LOBSTER BISQUE"]
	if countIngName("Wobster", crockPot) < 1 {
		lobsterBisque.UNMETCONDITIONS = append(lobsterBisque.UNMETCONDITIONS, "Wobster")
	}
	if countIngName("Ice", crockPot) < 1 {
		lobsterBisque.UNMETCONDITIONS = append(lobsterBisque.UNMETCONDITIONS, "Ice")
	}
	return &lobsterBisque
}

func mandrakeSoupPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// MANDRAKE SOUP
	mandrakeSoup := recipeData["MANDRAKE SOUP"]
	if countIngName("Mandrake", crockPot) < 1 {
		mandrakeSoup.UNMETCONDITIONS = append(mandrakeSoup.UNMETCONDITIONS, "Mandrake")
	}
	return &mandrakeSoup
}

func meatballsPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// MEATBALLS
	meatballs := recipeData["MEATBALLS"]
	if stringInSlice("Twigs", crockPot) {
		return nil
	}
	x := attributeVals.meatVal
	switch x {
	case 0:
		meatballs.UNMETCONDITIONS = append(meatballs.UNMETCONDITIONS, "mv")
	case 3:
		return nil
	}
	return &meatballs
}

func meatyStewPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// MEATY STEW
	meatyStew := recipeData["MEATY STEW"]
	if stringInSlice("Twigs", crockPot) {
		return nil
	}
	if stringInSlice("Moleworm", crockPot) {
		return nil
	}
	if stringInSlice("Honey", crockPot) {
		return nil
	}
	if stringInSlice("Mandrake", crockPot) {
		return nil
	}
	if stringInSlice("Tallbird Egg", crockPot) {
		return nil
	}
	if attributeVals.meatVal < 3 {
		meatyStew.UNMETCONDITIONS = append(meatyStew.UNMETCONDITIONS, "mv")
	}
	return &meatyStew
}

func melonsiclePossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// MELONSICLE
	melonsicle := recipeData["MELONSICLE"]
	if attributeVals.meatVal != 0 {
		return nil
	}
	if attributeVals.vegVal != 0 {
		return nil
	}
	if attributeVals.eggVal != 0 {
		return nil
	}
	if countIngName("Watermelon", crockPot) < 1 {
		melonsicle.UNMETCONDITIONS = append(melonsicle.UNMETCONDITIONS, "Watermelon")
	}
	if countIngName("Ice", crockPot) < 1 {
		melonsicle.UNMETCONDITIONS = append(melonsicle.UNMETCONDITIONS, "Ice")
	}
	if countIngName("Twigs", crockPot) < 1 {
		melonsicle.UNMETCONDITIONS = append(melonsicle.UNMETCONDITIONS, "Twigs")
	}
	return &melonsicle
}

func milkmadeHatPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// MILKMADE HAT
	milkmadeHat := recipeData["MILKMADE HAT"]
	if countIngName("Nostrils", crockPot) < 1 {
		milkmadeHat.UNMETCONDITIONS = append(milkmadeHat.UNMETCONDITIONS, "Nostrils")
	}
	if countIngName("Kelp Fronds", crockPot) < 1 {
		milkmadeHat.UNMETCONDITIONS = append(milkmadeHat.UNMETCONDITIONS, "Kelp Fronds")
	}
	if attributeCounts.dairyCount < 1 {
		milkmadeHat.UNMETCONDITIONS = append(milkmadeHat.UNMETCONDITIONS, "dc")
	}
	return &milkmadeHat
}

func monsterLasagnaPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// MONSTER LASAGNA
	monsterLasagna := recipeData["MONSTER LASAGNA"]
	if attributeCounts.monsterCount < 2 {
		monsterLasagna.UNMETCONDITIONS = append(monsterLasagna.UNMETCONDITIONS, "tc")
	}
	if stringInSlice("Twigs", crockPot) {
		return nil
	}
	return &monsterLasagna
}

func mushyCakePossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// MUSHY CAKE
	mushyCake := recipeData["MUSHY CAKE"]
	w := countIngName("Moon Shroom", crockPot)
	x := countIngName("Red Cap", crockPot)
	y := countIngName("Blue Cap", crockPot)
	z := countIngName("Green Cap", crockPot)
	if w != 1 {
		switch w {
		case 0:
			mushyCake.UNMETCONDITIONS = append(mushyCake.UNMETCONDITIONS, "Moon Shroom")
		case 2, 3:
			return nil
		}
	}
	if x != 1 {
		switch x {
		case 0:
			mushyCake.UNMETCONDITIONS = append(mushyCake.UNMETCONDITIONS, "Red Cap")
		case 2, 3:
			return nil
		}
	}
	if y != 1 {
		switch y {
		case 0:
			mushyCake.UNMETCONDITIONS = append(mushyCake.UNMETCONDITIONS, "Blue Cap")
		case 2, 3:
			return nil
		}
	}
	if z != 1 {
		switch z {
		case 0:
			mushyCake.UNMETCONDITIONS = append(mushyCake.UNMETCONDITIONS, "Green Cap")
		case 2, 3:
			return nil
		}
	}
	return &mushyCake
}

func pierogiPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// PIEROGI
	pierogi := recipeData["PIEROGI"]
	if stringInSlice("Twigs", crockPot) {
		return nil
	}
	if stringInSlice("Mandrake", crockPot) {
		return nil
	}
	if attributeCounts.meatCount < 1 {
		pierogi.UNMETCONDITIONS = append(pierogi.UNMETCONDITIONS, "mc")
	}
	if attributeCounts.eggCount < 1 {
		pierogi.UNMETCONDITIONS = append(pierogi.UNMETCONDITIONS, "ec")
	}
	if attributeCounts.vegeCount < 1 {
		pierogi.UNMETCONDITIONS = append(pierogi.UNMETCONDITIONS, "gc")
	}
	return &pierogi
}

func powdercakePossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	//powdercake
	powdercake := recipeData["POWDERCAKE"]
	if countIngName("Corn", crockPot) < 1 {
		powdercake.UNMETCONDITIONS = append(powdercake.UNMETCONDITIONS, "Corn")
	}
	if countIngName("Honey", crockPot) < 1 {
		powdercake.UNMETCONDITIONS = append(powdercake.UNMETCONDITIONS, "Honey")
	}
	if countIngName("Twigs", crockPot) < 1 {
		powdercake.UNMETCONDITIONS = append(powdercake.UNMETCONDITIONS, "Twigs")
	}
	return &powdercake
}

func pumkinCookiePossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// PUMPKIN COOKIE
	pumkinCookie := recipeData["PUMPKIN COOKIE"]
	if countIngName("Pumpkin", crockPot) < 1 {
		pumkinCookie.UNMETCONDITIONS = append(pumkinCookie.UNMETCONDITIONS, "Pumpkin")
	}
	if countIngName("Honey", crockPot) <= 1 {
		pumkinCookie.UNMETCONDITIONS = append(pumkinCookie.UNMETCONDITIONS, "Honey")
	}
	return &pumkinCookie
}

func ratatouillePossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// ratatouille
	ratatouille := recipeData["RATATOUILLE"]
	if stringInSlice("Twigs", crockPot) {
		return nil
	}
	if stringInSlice("Mandrake", crockPot) {
		return nil
	}
	if stringInSlice("Butterfly Wings", crockPot) {
		return nil
	}
	if stringInSlice("Dragon Fruit", crockPot) {
		return nil
	}
	if attributeCounts.vegeCount < 1 {
		ratatouille.UNMETCONDITIONS = append(ratatouille.UNMETCONDITIONS, "vc")
	}
	return &ratatouille
}

func salsaFrescaPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// SALSA FRESCA
	salsaFresca := recipeData["SALSA FRESCA"]
	if countIngName("Toma Root", crockPot) < 1 {
		salsaFresca.UNMETCONDITIONS = append(salsaFresca.UNMETCONDITIONS, "Toma Root")
	}
	if attributeCounts.meatCount != 0 {
		return nil
	}
	if attributeCounts.inedibleCount != 0 {
		return nil
	}
	if attributeCounts.eggCount != 0 {
		return nil
	}
	return &salsaFresca
}

func seafoodGumboPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// SEAFOOD GUMBO
	seafoodGumbo := recipeData["SEAFOOD GUMBO"]
	if countIngName("Eel", crockPot) < 1 {
		seafoodGumbo.UNMETCONDITIONS = append(seafoodGumbo.UNMETCONDITIONS, "Eel")
	}
	if attributeVals.fishVal <= 2 {
		seafoodGumbo.UNMETCONDITIONS = append(seafoodGumbo.UNMETCONDITIONS, "hv")
	}
	return &seafoodGumbo
}

func soothingTeaPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	//  soothingTea
	soothingTea := recipeData["SOOTHING TEA"]
	if attributeCounts.monsterCount != 0 {
		return nil
	}
	if attributeCounts.meatCount != 0 {
		return nil
	}
	if attributeCounts.fishCount != 0 {
		return nil
	}
	if attributeCounts.eggCount != 0 {
		return nil
	}
	if attributeCounts.inedibleCount != 0 {
		return nil
	}
	if attributeCounts.dairyCount != 0 {
		return nil
	}
	if countIngName("Forget-Me-Lots", crockPot) < 1 {
		soothingTea.UNMETCONDITIONS = append(soothingTea.UNMETCONDITIONS, "Forget-Me-Lots")
	}
	if countIngName("Honey", crockPot) < 1 {
		soothingTea.UNMETCONDITIONS = append(soothingTea.UNMETCONDITIONS, "Honey")
	}
	if countIngName("Ice", crockPot) < 1 {
		soothingTea.UNMETCONDITIONS = append(soothingTea.UNMETCONDITIONS, "Ice")
	}
	return &soothingTea
}

func spicyChiliPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// SPICY CHILI
	spicyChili := recipeData["SPICY CHILI"]
	x := attributeCounts.meatCount
	y := attributeCounts.vegeCount
	if x != 2 {
		switch x {
		case 0, 1:
			spicyChili.UNMETCONDITIONS = append(spicyChili.UNMETCONDITIONS, "mc")
		case 3:
			return nil
		}
	}
	if y != 2 {
		switch x {
		case 0, 1:
			spicyChili.UNMETCONDITIONS = append(spicyChili.UNMETCONDITIONS, "gc")
		case 3:
			return nil
		}
	}
	if attributeVals.meatVal < 1.5 {
		spicyChili.UNMETCONDITIONS = append(spicyChili.UNMETCONDITIONS, "mv")
	}
	if attributeVals.vegVal < 1.5 {
		spicyChili.UNMETCONDITIONS = append(spicyChili.UNMETCONDITIONS, "gv")
	}
	return &spicyChili
}

func stuffedEggplantPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// STUFFED EGGPLANT
	stuffedEggplant := recipeData["STUFFED EGGPLANT"]
	if countIngName("Eggplant", crockPot) < 1 {
		stuffedEggplant.UNMETCONDITIONS = append(stuffedEggplant.UNMETCONDITIONS, "Eggplant")
	}
	if attributeCounts.vegeCount < 1 {
		stuffedEggplant.UNMETCONDITIONS = append(stuffedEggplant.UNMETCONDITIONS, "gv")
	}
	return &stuffedEggplant
}

func stuffedFishHeadsPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// stuffedFishHeads
	stuffedFishHeads := recipeData["STUFFED FISH HEADS"]
	if countIngName("Barnacles", crockPot) < 1 {
		stuffedFishHeads.UNMETCONDITIONS = append(stuffedFishHeads.UNMETCONDITIONS, "Barnacles")
	}
	if attributeVals.fishVal-0.5 < 1 {
		stuffedFishHeads.UNMETCONDITIONS = append(stuffedFishHeads.UNMETCONDITIONS, "hv")
	}
	return &stuffedFishHeads
}

func stuffedPepperPoppersPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// STUFFED PEPPER POPPERS
	stuffedPepperPoppers := recipeData["STUFFED PEPPER POPPERS"]
	if stringInSlice("Twigs", crockPot) {
		return nil
	}
	if countIngName("Pepper", crockPot) < 1 {
		stuffedPepperPoppers.UNMETCONDITIONS = append(stuffedPepperPoppers.UNMETCONDITIONS, "Pepper")
	}
	if attributeVals.meatVal > 1.5 {
		stuffedPepperPoppers.UNMETCONDITIONS = append(stuffedPepperPoppers.UNMETCONDITIONS, "mV")
	}
	return &stuffedPepperPoppers
}

func surfNTurfPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// SURF 'N' TURF
	surfNTurf := recipeData["SURF 'N' TURF"]
	if stringInSlice("Ice", crockPot) {
		return nil
	}
	if attributeVals.meatVal < 2.5 {
		surfNTurf.UNMETCONDITIONS = append(surfNTurf.UNMETCONDITIONS, "mv")
	}
	if attributeVals.fishVal < 1.5 {
		surfNTurf.UNMETCONDITIONS = append(surfNTurf.UNMETCONDITIONS, "hv")
	}

	return &surfNTurf
}

func trailMixPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// TRAIL MIX
	trailMix := recipeData["TRAIL MIX"]
	if attributeCounts.meatCount != 0 {
		return nil
	}
	if attributeCounts.fishCount != 0 {
		return nil
	}
	if attributeCounts.eggCount != 0 {
		return nil
	}
	if attributeCounts.vegeCount != 0 {
		return nil
	}
	if attributeCounts.dairyCount != 0 {
		return nil
	}
	if countIngName("Roasted", crockPot) > 3 {
		return nil
	}
	if countIngName("Roasted Birchnut", crockPot) < 1 {
		trailMix.UNMETCONDITIONS = append(trailMix.UNMETCONDITIONS, "Roasted Birchnut")
	}
	if countIngName("Berries", crockPot) < 1 {
		trailMix.UNMETCONDITIONS = append(trailMix.UNMETCONDITIONS, "Berries")
	}
	if attributeCounts.fruitCount-1 < 1 {
		trailMix.UNMETCONDITIONS = append(trailMix.UNMETCONDITIONS, "fc")
	}
	return &trailMix
}

func turkeyDinnerPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// TURKEY DINNER
	turkeyDinner := recipeData["TURKEY DINNER"]
	x := countIngName("Drumstick", crockPot)
	if x < 2 {
		switch x {
		case 0:
			turkeyDinner.UNMETCONDITIONS = append(turkeyDinner.UNMETCONDITIONS, "2 Drumstick")
		case 1:
			turkeyDinner.UNMETCONDITIONS = append(turkeyDinner.UNMETCONDITIONS, "Drumstick")
		}
	}
	if attributeVals.meatVal-1 < 0.25 {
		turkeyDinner.UNMETCONDITIONS = append(turkeyDinner.UNMETCONDITIONS, "mv")
	}
	if attributeVals.fruitVal < .5 && attributeVals.vegVal < .5 {
		turkeyDinner.UNMETCONDITIONS = append(turkeyDinner.UNMETCONDITIONS, "fv")
		turkeyDinner.UNMETCONDITIONS = append(turkeyDinner.UNMETCONDITIONS, "gv")
	}
	return &turkeyDinner
}

func unagiPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// unagi
	unagi := recipeData["UNAGI"]
	if countIngName("Eel", crockPot) < 1 {
		unagi.UNMETCONDITIONS = append(unagi.UNMETCONDITIONS, "Eel")
	}
	if countIngName("Lichen", crockPot) < 1 && countIngName("Kelp Fronds", crockPot) < 1 {
		unagi.UNMETCONDITIONS = append(unagi.UNMETCONDITIONS, "Lichen")
		unagi.UNMETCONDITIONS = append(unagi.UNMETCONDITIONS, "Kelp Fronds")
	}
	return &unagi
}

func vegetableStingerPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// VEGETABLE STINGER
	vegetableStinger := recipeData["VEGETABLE STINGER"]
	if countIngName("Toma Root", crockPot) < 1 && countIngName("Asparagus", crockPot) < 1 {
		vegetableStinger.UNMETCONDITIONS = append(vegetableStinger.UNMETCONDITIONS, "Toma Root")
		vegetableStinger.UNMETCONDITIONS = append(vegetableStinger.UNMETCONDITIONS, "Asparagus")
	}
	if countIngName("Ice", crockPot) < 1 {
		vegetableStinger.UNMETCONDITIONS = append(vegetableStinger.UNMETCONDITIONS, "Ice")
	}
	if attributeVals.vegVal-1 < 1.5 {
		vegetableStinger.UNMETCONDITIONS = append(vegetableStinger.UNMETCONDITIONS, "gv")
	}
	return &vegetableStinger
}

func veggieBurgerPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// VEGGIE BURGER
	veggieBurger := recipeData["VEGGIE BURGER"]
	if countIngName("Leafy Meat", crockPot) < 1 {
		veggieBurger.UNMETCONDITIONS = append(veggieBurger.UNMETCONDITIONS, "Leafy Meat")
	}
	if countIngName("Onion", crockPot) < 1 {
		veggieBurger.UNMETCONDITIONS = append(veggieBurger.UNMETCONDITIONS, "Onion")
	}
	if attributeVals.vegVal-1 < 1 {
		veggieBurger.UNMETCONDITIONS = append(veggieBurger.UNMETCONDITIONS, "gv")
	}
	return &veggieBurger
}

func wafflesPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// WAFFLES
	waffles := recipeData["WAFFLES"]
	if countIngName("Butter", crockPot) < 1 {
		waffles.UNMETCONDITIONS = append(waffles.UNMETCONDITIONS, "Butter")
	}
	if countIngName("Berries", crockPot) < 1 {
		waffles.UNMETCONDITIONS = append(waffles.UNMETCONDITIONS, "Berries")
	}
	if attributeCounts.eggCount < 1 {
		waffles.UNMETCONDITIONS = append(waffles.UNMETCONDITIONS, "ec")
	}
	return &waffles
}

func wetGoopPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// WET GOOP***********************************************************************************************
	wetGoop := recipeData["WET GOOP"]
	// if everything is false
	return &wetGoop
}

func wobsterDinnerPossible(crockPot []IngredientDetails, attributeVals AttributeVals, attributeCounts AttributeCounts) *RecipeDetails {
	// WOBSTER DINNER
	wobsterDinner := recipeData["WOBSTER DINNER"]
	if attributeCounts.meatCount != 0 {
		return nil
	}
	if attributeCounts.fishCount != 0 {
		return nil
	}
	if stringInSlice("Twigs", crockPot) {
		return nil
	}
	if countIngName("Wobster", crockPot) < 1 {
		wobsterDinner.UNMETCONDITIONS = append(wobsterDinner.UNMETCONDITIONS, "Wobster")
	}
	if countIngName("Butter", crockPot) < 1 {
		wobsterDinner.UNMETCONDITIONS = append(wobsterDinner.UNMETCONDITIONS, "Butter")
	}
	return &wobsterDinner
}

// MEATS = m
// FISHES= h
// EGGS= e
// FRUITS= f
// VEGETABLES= g
// SWEETENERS= s
// MONSTER FOODS= t
// DAIRY= d
// BUGS= b
// INEDIBLE= i
// MISC.= l
// INCLUDE= o
// EXCLUDE= x
// TOO MUCH VAL = V
// NOTE ENOUGH VAL = v
// TOO MUCH COUNT = C
// NOTE ENOUGH COUNT = v
