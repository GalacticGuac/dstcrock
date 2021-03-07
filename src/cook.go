package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type IngredientDetails struct {
	NAME      string  `json:"INGREDIENT"`
	MEAT      float64 `json:"MEAT"`
	FISH      float64 `json:"FISH"`
	EGG       float64 `json:"EGG"`
	FRUIT     float64 `json:"FRUIT"`
	VEGETABLE float64 `json:"VEGETABLE"`
	SWEETENER float64 `json:"SWEETENER"`
	MONSTER   float64 `json:"MONSTER FOOD"`
	DAIRY     float64 `json:"DAIRY"`
	BUG       float64 `json:"BUG"`
	INEDIBLE  float64 `json:"INEDIBLE"`
	MISC      float64 `json:"MISC"`
	HEALTH    float64 `json:"HEALTH"`
	HUNGER    float64 `json:"HUNGER"`
	SANITY    float64 `json:"SANITY"`
	EXPIRE    float64 `json:"EXPIRE"`
	NOTES     string  `json:"NOTES"`
}

type RecipeDetails struct {
	NAME            string   `json:"FOOD"`
	HEALTH          float64  `json:"HEALTH"`
	HUNGER          float64  `json:"HUNGER"`
	SANITY          float64  `json:"SANITY"`
	INGREDIENTS     string   `json:"INGREDIENTS"`
	EXCLUDE         string   `json:"EXCLUDE"`
	NOTES           string   `json:"NOTES"`
	EXPIRES         float64  `json:"EXPIRES"`
	COOKTIME        float64  `json:"COOKTIME"`
	PRIORITY        float64  `json:"PRIORITY"`
	UNMETCONDITIONS []string `json:"UNMETCONDITIONS"`
	// preferences?
}
type AttributeCounts struct {
	meatCount      float64 `json:"MEAT"`
	fishCount      float64 `json:"FISH"`
	eggCount       float64 `json:"EGG"`
	fruitCount     float64 `json:"FRUIT"`
	vegeCount      float64 `json:"VEGETABLE"`
	sweetenerCount float64 `json:"SWEETENER"`
	monsterCount   float64 `json:"MONSTER FOOD"`
	dairyCount     float64 `json:"DAIRY"`
	bugCount       float64 `json:"BUG"`
	inedibleCount  float64 `json:"INEDIBLE"`
	miscCount      float64 `json:"MISC"`
}
type AttributeVals struct {
	meatVal  float64 `json:"MEAT"`
	fishVal  float64 `json:"FISH"`
	eggVal   float64 `json:"EGG"`
	fruitVal float64 `json:"FRUIT"`
	vegVal   float64 `json:"VEGETABLE"`
	sweetVal float64 `json:"SWEETENER"`
	monVal   float64 `json:"MONSTER FOOD"`
	dairyVal float64 `json:"DAIRY"`
	bugVal   float64 `json:"BUG"`
	inedVal  float64 `json:"INEDIBLE"`
	miscVal  float64 `json:"MISC"`
	healVal  float64 `json:"HEALTH"`
	hungVal  float64 `json:"HUNGER"`
	saniVal  float64 `json:"SANITY"`
}

// func deleteFalseReipe(recipeData map[string]RecipeDetails, masterCounts AttributeCounts, masterVals AttributeVals, titles string) map[string]RecipeDetails
// count strings in cropost array

func masterCounts(crockPot []IngredientDetails) AttributeCounts {

	attributeCounts := AttributeCounts{}

	// count meat
	for i := 0; i < len(crockPot); i++ {
		if crockPot[i].MEAT > 0 {
			attributeCounts.meatCount++
		}
	}

	// count Fish
	for i := 0; i < len(crockPot); i++ {
		if crockPot[i].FISH > 0 {
			attributeCounts.fishCount++
		}
	}

	// count egg
	for i := 0; i < len(crockPot); i++ {
		if crockPot[i].EGG > 0 {
			attributeCounts.eggCount++
		}
	}

	// count fruit
	for i := 0; i < len(crockPot); i++ {
		if crockPot[i].FRUIT > 0 {
			attributeCounts.fruitCount++
		}
	}

	// count vegetables
	for i := 0; i < len(crockPot); i++ {
		if crockPot[i].VEGETABLE > 0 {
			attributeCounts.vegeCount++
		}
	}

	// count sweetner
	for i := 0; i < len(crockPot); i++ {
		if crockPot[i].SWEETENER > 0 {
			attributeCounts.sweetenerCount++
		}
	}

	// count monster
	for i := 0; i < len(crockPot); i++ {
		if crockPot[i].MONSTER > 0 {
			attributeCounts.monsterCount++
		}
	}

	// count dairy
	for i := 0; i < len(crockPot); i++ {
		if crockPot[i].DAIRY > 0 {
			attributeCounts.dairyCount++
		}
	}

	// count bug
	for i := 0; i < len(crockPot); i++ {
		if crockPot[i].BUG > 0 {
			attributeCounts.bugCount++
		}
	}

	// count inedable
	for i := 0; i < len(crockPot); i++ {
		if crockPot[i].INEDIBLE > 0 {
			attributeCounts.inedibleCount++
		}
	}

	// count misc (moleworm, butterfly, etc.)
	for i := 0; i < len(crockPot); i++ {
		if crockPot[i].MISC > 0 {
			attributeCounts.miscCount++
		}
	}

	return attributeCounts
}

func masterVals(crockPot []IngredientDetails) AttributeVals {

	attributeVals := AttributeVals{}

	for i := 0; i < len(crockPot); i++ {
		attributeVals.meatVal += crockPot[i].MEAT
	}

	for i := 0; i < len(crockPot); i++ {
		attributeVals.fishVal += crockPot[i].FISH
	}

	for i := 0; i < len(crockPot); i++ {
		attributeVals.eggVal += crockPot[i].EGG
	}

	for i := 0; i < len(crockPot); i++ {
		attributeVals.fruitVal += crockPot[i].FRUIT
	}

	for i := 0; i < len(crockPot); i++ {
		attributeVals.vegVal += crockPot[i].VEGETABLE
	}

	for i := 0; i < len(crockPot); i++ {
		attributeVals.sweetVal += crockPot[i].SWEETENER
	}

	for i := 0; i < len(crockPot); i++ {
		attributeVals.monVal += crockPot[i].MONSTER
	}

	for i := 0; i < len(crockPot); i++ {
		attributeVals.dairyVal += crockPot[i].DAIRY
	}

	for i := 0; i < len(crockPot); i++ {
		attributeVals.bugVal += crockPot[i].BUG
	}

	for i := 0; i < len(crockPot); i++ {
		attributeVals.inedVal += crockPot[i].INEDIBLE
	}

	for i := 0; i < len(crockPot); i++ {
		attributeVals.miscVal += crockPot[i].MISC
	}

	for i := 0; i < len(crockPot); i++ {
		attributeVals.healVal += crockPot[i].HEALTH
	}

	for i := 0; i < len(crockPot); i++ {
		attributeVals.hungVal += crockPot[i].HUNGER
	}

	for i := 0; i < len(crockPot); i++ {
		attributeVals.saniVal += crockPot[i].SANITY
	}

	return attributeVals
}

//	function that deletes false recipes (if any of the formulas are true, item is deleted, false=keep)
func deleteRecipes(recipeData map[string]RecipeDetails, attributeCounts AttributeCounts, attributeVals AttributeVals, crockPot []IngredientDetails) map[string]RecipeDetails {

	// AMBEROSIA
	if countIngName("Collected Dust", crockPot) < 1 {
		delete(recipeData, "AMBEROSIA")
	}

	// ASPARAGUS SOUP
	if countIngName("Asparagus", crockPot) < 1 || attributeVals.vegVal < 1.5 || attributeCounts.meatCount != 0 || attributeCounts.inedibleCount != 0 {
		delete(recipeData, "ASPARAGUS SOUP")
	}

	// Bacon and Eggs
	if attributeVals.meatVal <= 1 || attributeVals.eggVal <= 1 || attributeVals.vegVal != 0 {
		delete(recipeData, "BACON AND EGGS")
	}

	// BANANA POP
	if countIngName("Banana", crockPot) < 1 || countIngName("Twigs", crockPot) < 1 || attributeCounts.meatCount != 0 || attributeCounts.fishCount != 0 {
		delete(recipeData, "BANANA POP")
	}

	// Barnacle Linguine
	if countIngName("Barnacles", crockPot) != 2 || attributeCounts.vegeCount != 2 {
		delete(recipeData, "BARNACLE LINGUINE")
	}

	// Barnacle Nigiri
	if countIngName("Barnacles", crockPot) < 1 || countIngName("Kelp Fronds", crockPot) < 1 || attributeCounts.eggCount < 1 {
		delete(recipeData, "BARNACLE NIGIRI")
	}

	// BARNACLE PITA
	// attributeVals.fishval over 1 can cause stuffed fish head but idk the percentage
	if countIngName("Barnacles", crockPot) < 1 || countIngName("Kelp Fronds", crockPot) < 1 || attributeCounts.eggCount < 1 {
		delete(recipeData, "BARNACLE PITA")
	}

	// BEEFY GREENS
	if countIngName("Leafy Meat", crockPot) < 1 || attributeVals.vegVal < 3 {
		delete(recipeData, "BEEFY GREENS")
	}

	// butter muffin
	if countIngName("Butterfly Wings", crockPot) < 1 || attributeCounts.vegeCount < 1 || stringInSlice("Mandrake", crockPot) || attributeCounts.meatCount != 0 {
		delete(recipeData, "BUTTER MUFFIN")
	}

	// CALIFORNIA ROLL
	if countIngName("Kelp Fronds", crockPot) < 2 || attributeVals.fishVal < 1 {
		delete(recipeData, "CALIFORNIA ROLL")
	}

	// ceviche
	if countIngName("Ice", crockPot) < 2 || attributeVals.fishVal < 2 || attributeCounts.eggCount != 0 || attributeCounts.inedibleCount != 0 {
		delete(recipeData, "CEVICHE")
	}

	// CREAMY POTATO PURÉE**
	// need to look at potato or roasted potato
	if countIngName("Potato", crockPot) < 2 || countIngName("Garlic", crockPot) < 1 || stringInSlice("Twigs", crockPot) || attributeCounts.meatCount != 0 {
		delete(recipeData, "CREAMY POTATO PURÉE")
	}

	// dragonpie
	// need to look at dragon fruit or prepared dragon fruit
	if countIngName("Dragon Fruit", crockPot) < 1 || stringInSlice("Mandrake", crockPot) || attributeCounts.meatCount != 0 {
		delete(recipeData, "DRAGONPIE")
	}

	// FANCY SPIRALLED TUBERS
	// need to look at potato or roasted potato
	if countIngName("Potato", crockPot) < 1 || countIngName("Twigs", crockPot) < 1 || attributeCounts.inedibleCount-1 > 1 || attributeCounts.meatCount != 0 {
		delete(recipeData, "FANCY SPIRALLED TUBERS")
	}

	// fish tacos
	// twig value - 1 twig 50% chance of  fish sticks
	// need to look at corn or popcorn
	if attributeVals.fishVal < 0.5 || (countIngName("Corn", crockPot) < 1 || countIngName("Popcorn", crockPot) < 1) {
		delete(recipeData, "FISH TACOS")
	}

	// fishsticks
	if attributeVals.fishVal < 0.25 || countIngName("Twigs", crockPot) != 1 || stringInSlice("Moleworm", crockPot) {
		delete(recipeData, "FISHSTICKS")
	}

	// fist full of jam
	if attributeVals.fruitVal < 0.5 || attributeVals.meatVal != 0 || attributeVals.vegVal != 0 || attributeVals.inedVal != 0 || stringInSlice("Dragon Fruit", crockPot) {
		delete(recipeData, "FIST FULL OF JAM")
	}

	// flower salad
	if countIngName("Cactus Flower", crockPot) < 1 || attributeVals.vegVal-.5 < 1.5 || attributeVals.meatVal != 0 || attributeVals.fruitVal != 0 || attributeVals.eggVal != 0 || attributeVals.sweetVal != 0 || stringInSlice("Twigs", crockPot) {
		delete(recipeData, "FLOWER SALAD")
	}

	// froggle bunwich
	// need to look at frog legs or cooked frog legs
	// makes kabob if only one stick
	if countIngName("Frog Legs", crockPot) < 1 || attributeCounts.vegeCount < 1 || attributeVals.eggVal != 0 || attributeVals.sweetVal != 0 || stringInSlice("Mandrake", crockPot) {
		delete(recipeData, "FROGGLE BUNWICH")
	}

	// fruit medley
	// twigs is safest anything else 50% of fist full of jam
	if attributeVals.fruitVal < 3 || attributeVals.meatVal != 0 || attributeVals.vegVal != 0 || stringInSlice("Dragon Fruit", crockPot) {
		delete(recipeData, "FRUIT MEDLEY")
	}

	// guacamole
	// need to look at cactus flesh or stone fruit
	if countIngName("Moleworm", crockPot) < 1 || (countIngName("Cactus Flesh", crockPot) < 1 || countIngName("Ripe Stone Fruit", crockPot) < 1) || attributeVals.fruitVal != 0 {
		delete(recipeData, "GUACAMOLE")
	}

	// Honey Ham
	if countIngName("Honey", crockPot) < 1 || attributeVals.meatVal <= 1.5 || stringInSlice("Twigs", crockPot) || stringInSlice("Moleworm", crockPot) || stringInSlice("Mandrake", crockPot) || stringInSlice("Tallbird Egg", crockPot) {
		delete(recipeData, "HONEY HAM")
	}

	// Honey Nuggets
	if countIngName("Honey", crockPot) < 1 || attributeVals.meatVal > 1.5 || attributeVals.inedVal != 0 {
		delete(recipeData, "HONEY NUGGETS")
	}

	// ice cream
	if countIngName("Ice", crockPot) < 1 || attributeCounts.dairyCount < 1 || attributeCounts.sweetenerCount < 1 || attributeVals.meatVal != 0 || attributeVals.vegVal != 0 || attributeVals.eggVal != 0 || stringInSlice("Twigs", crockPot) {
		delete(recipeData, "ICE CREAM")
	}

	// jelly salad
	// look for cooked version
	if countIngName("Leafy Meat", crockPot) < 2 || attributeVals.sweetVal < 2 {
		delete(recipeData, "JELLY SALAD")
	}

	// jellybeans
	if countIngName("Royal Jelly", crockPot) < 1 || attributeVals.monVal != 0 || attributeVals.inedVal != 0 {
		delete(recipeData, "JELLYBEANS")
	}

	// kabobs
	if attributeCounts.meatCount < 1 || countIngName("Twigs", crockPot) != 1 || stringInSlice("Moleworm", crockPot) || stringInSlice("Mandrake", crockPot) || attributeVals.fishVal != 0 {
		delete(recipeData, "KABOBS")
	}

	// leafy meatloaf
	if countIngName("Leafy Meat", crockPot) < 2 {
		delete(recipeData, "LEAFY MEATLOAF")
	}

	// LOBSTER BISQUE
	if countIngName("Wobster", crockPot) < 1 || countIngName("Ice", crockPot) < 1 {
		delete(recipeData, "LOBSTER BISQUE")
	}

	// mandrake soup
	if countIngName("Mandrake", crockPot) < 1 {
		delete(recipeData, "MANDRAKE SOUP")
	}

	// Meatballs
	if attributeVals.meatVal >= 3 || attributeVals.meatVal == 0 || stringInSlice("Twigs", crockPot) {
		delete(recipeData, "MEATBALLS")
	}

	// meaty stew
	if attributeVals.meatVal < 3 || stringInSlice("Twigs", crockPot) || stringInSlice("Moleworm", crockPot) || stringInSlice("Honey", crockPot) || stringInSlice("Mandrake", crockPot) || stringInSlice("Tallbird Egg", crockPot) {
		delete(recipeData, "MEATY STEW")
	}

	// melonsicle
	if countIngName("Watermelon", crockPot) < 1 || countIngName("Ice", crockPot) < 1 || countIngName("Twigs", crockPot) < 1 || attributeVals.meatVal != 0 || attributeVals.vegVal != 0 || attributeVals.eggVal != 0 {
		delete(recipeData, "MELONSICLE")
	}

	// MILKMADE HAT
	if countIngName("Nostrils", crockPot) < 1 || countIngName("Kelp Fronds", crockPot) < 1 || attributeCounts.dairyCount < 1 {
		delete(recipeData, "MILKMADE HAT")
	}

	// monster lasagna
	if attributeCounts.monsterCount < 2 || stringInSlice("Twigs", crockPot) {
		delete(recipeData, "MONSTER LASAGNA")
	}

	// Mushy Cake
	if countIngName("Moon Shroom", crockPot) != 1 || countIngName("Red Cap", crockPot) != 1 || countIngName("Blue Cap", crockPot) != 1 || countIngName("Green Cap", crockPot) != 1 {
		delete(recipeData, "MUSHY CAKE")
	}

	// Pierogi
	if attributeCounts.meatCount < 1 || attributeCounts.eggCount < 1 || attributeCounts.vegeCount < 1 || stringInSlice("Twigs", crockPot) || stringInSlice("Mandrake", crockPot) {
		delete(recipeData, "PIEROGI")
	}

	//powdercake
	if countIngName("Corn", crockPot) < 1 || countIngName("Honey", crockPot) < 1 || countIngName("Twigs", crockPot) < 1 {
		delete(recipeData, "POWDERCAKE")
	}

	//Pumkin cookie
	// 3 honey or comb 50% chance of making taffy
	if countIngName("Pumpkin", crockPot) < 1 || countIngName("Honey", crockPot) <= 1 {
		delete(recipeData, "PUMPKIN COOKIE")
	}

	// RATATOUILLE
	if attributeCounts.vegeCount < 1 || stringInSlice("Twigs", crockPot) || stringInSlice("Mandrake", crockPot) || stringInSlice("Butterfly Wings", crockPot) || stringInSlice("Dragon Fruit", crockPot) {
		delete(recipeData, "RATATOUILLE")
	}

	// SALSA FRESCA
	// look for cooked versions
	if countIngName("Toma Root", crockPot) < 1 || attributeCounts.meatCount != 0 || attributeCounts.inedibleCount != 0 || attributeCounts.eggCount != 0 {
		delete(recipeData, "SALSA FRESCA")
	}

	// SEAFOOD GUMBO
	if countIngName("Eel", crockPot) < 1 || attributeVals.fishVal <= 2 {
		delete(recipeData, "SEAFOOD GUMBO")
	}

	//  Soothing tea
	// look for honey or comb
	if countIngName("Forget-Me-Lots", crockPot) < 1 || countIngName("Honey", crockPot) < 1 || countIngName("Ice", crockPot) < 1 || attributeCounts.monsterCount != 0 || attributeCounts.meatCount != 0 || attributeCounts.fishCount != 0 || attributeCounts.eggCount != 0 || attributeCounts.inedibleCount != 0 || attributeCounts.dairyCount != 0 {
		delete(recipeData, "SOOTHING TEA")
	}

	//  Spicy Chili
	if attributeCounts.meatCount != 2 || attributeCounts.vegeCount != 2 || attributeVals.meatVal < 1.5 || attributeVals.vegVal < 1.5 {
		delete(recipeData, "SPICY CHILI")
	}

	// STUFFED EGGPLANT
	if countIngName("Eggplant", crockPot) < 1 || attributeCounts.vegeCount < 1 {
		delete(recipeData, "STUFFED EGGPLANT")
	}

	// stuffed fish heads
	if countIngName("Barnacles", crockPot) < 1 || attributeVals.fishVal-0.5 < 1 {
		delete(recipeData, "STUFFED FISH HEADS")
	}

	// STUFFED PEPPER POPPERS
	if countIngName("Pepper", crockPot) < 1 || attributeVals.meatVal > 1.5 || attributeCounts.meatCount == 0 || stringInSlice("Twigs", crockPot) {
		delete(recipeData, "STUFFED PEPPER POPPERS")
	}

	// SURF 'N' TURF
	if attributeVals.meatVal < 2.5 || attributeVals.fishVal < 1.5 || stringInSlice("Ice", crockPot) {
		delete(recipeData, "SURF 'N' TURF")
	}

	// Taffy
	if attributeVals.sweetVal < 3 || attributeCounts.meatCount != 0 {
		delete(recipeData, "TAFFY")
	}

	// trail mix
	// required berries must be uncooked
	if countIngName("Roasted Birchnut", crockPot) < 1 || countIngName("Berries", crockPot) < 1 || countIngName("Roasted", crockPot) > 3 || attributeCounts.fruitCount-1 < 1 || attributeCounts.meatCount != 0 || attributeCounts.fishCount != 0 || attributeCounts.eggCount != 0 || attributeCounts.vegeCount != 0 || attributeCounts.dairyCount != 0 {
		delete(recipeData, "TRAIL MIX")
	}

	// TURKEY DINNER
	// checkingg twice for veg or fruit
	if countIngName("Drumstick", crockPot) < 2 || attributeVals.meatVal-1 < 0.25 || (attributeVals.fruitVal < .5 || attributeVals.vegVal < .5) {
		delete(recipeData, "TURKEY DINNER")
	}

	// unagi
	if countIngName("Eel", crockPot) < 1 || (countIngName("Lichen", crockPot) < 1 || countIngName("Kelp Fronds", crockPot) < 1) {
		delete(recipeData, "UNAGI")
	}

	// VEGETABLE STINGER
	if (countIngName("Toma Root", crockPot) < 1 || countIngName("Asparagus", crockPot) < 1) || countIngName("Ice", crockPot) < 1 || attributeVals.vegVal-1 < 1.5 {
		delete(recipeData, "VEGETABLE STINGER")
	}

	// Veggie Burger
	if countIngName("Leafy Meat", crockPot) < 1 || countIngName("Onion", crockPot) < 1 || attributeVals.vegVal-1 < 1 {
		delete(recipeData, "VEGGIE BURGER")
	}

	// waffles
	if countIngName("Butter", crockPot) < 1 || countIngName("Berries", crockPot) < 1 || attributeCounts.eggCount < 1 {
		delete(recipeData, "WAFFLES")
	}

	// wet goop
	// if everything is false

	// WOBSTER DINNER
	if countIngName("Wobster", crockPot) < 1 || countIngName("Butter", crockPot) < 1 || attributeCounts.meatCount != 0 || attributeCounts.fishCount != 0 || stringInSlice("Twigs", crockPot) {
		delete(recipeData, "WOBSTER DINNER")
	}
	return recipeData
}

func processPossible(x string) {
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

	conditions := []string{
		"mV",
		"mC",
		"twigs",
	}
	// for each possible recipe {
	for _, condition := range conditions {
		switch condition {
		case "mV":
			fmt.Println("Meat value too high")

		case "mv":
			fmt.Println("Meat value too low")

		default:
			fmt.Printf("Missing %s", condition)

		}
	}
	// }

	if strings.Contains(x, "mV") {
		fmt.Println("Meat value too high")
	}
	if strings.Contains(x, "mv") {
		fmt.Println("Meat value too low")
	}
	if strings.Contains(x, "mC") {
		fmt.Println("Meat count too high")
	}
	if strings.Contains(x, "mc") {
		fmt.Println("Meat count to low")
	}
	if strings.Contains(x, "hV") {
		fmt.Println("Fish value too high")
	}
	if strings.Contains(x, "hv") {
		fmt.Println("Fish value too low")
	}
	if strings.Contains(x, "hC") {
		fmt.Println("Fish count too high")
	}
	if strings.Contains(x, "hc") {
		fmt.Println("Fish count to low")
	}
	if strings.Contains(x, "eV") {
		fmt.Println("Egg value too high")
	}
	if strings.Contains(x, "ev") {
		fmt.Println("Egg value too low")
	}
	if strings.Contains(x, "eC") {
		fmt.Println("Egg count too high")
	}
	if strings.Contains(x, "ec") {
		fmt.Println("Egg count to low")
	}
	if strings.Contains(x, "fV") {
		fmt.Println("Fruit value too high")
	}
	if strings.Contains(x, "fv") {
		fmt.Println("Fruit value too low")
	}
	if strings.Contains(x, "fC") {
		fmt.Println("Fruit count too high")
	}
	if strings.Contains(x, "fc") {
		fmt.Println("Fruit count to low")
	}
	if strings.Contains(x, "gV") {
		fmt.Println("Vegetable value too high")
	}
	if strings.Contains(x, "gv") {
		fmt.Println("Vegetable value too low")
	}
	if strings.Contains(x, "gC") {
		fmt.Println("Vegetable count too high")
	}
	if strings.Contains(x, "gc") {
		fmt.Println("Vegetable count to low")
	}
	if strings.Contains(x, "sV") {
		fmt.Println("Sweetener value too high")
	}
	if strings.Contains(x, "sv") {
		fmt.Println("Sweetener value too low")
	}
	if strings.Contains(x, "sC") {
		fmt.Println("Sweetener count too high")
	}
	if strings.Contains(x, "sc") {
		fmt.Println("Sweetener count to low")
	}
	if strings.Contains(x, "tV") {
		fmt.Println("Monster value too high")
	}
	if strings.Contains(x, "tv") {
		fmt.Println("Monster value too low")
	}
	if strings.Contains(x, "tC") {
		fmt.Println("Monster count too high")
	}
	if strings.Contains(x, "tc") {
		fmt.Println("Monster count to low")
	}
	if strings.Contains(x, "dV") {
		fmt.Println("Dairy value too high")
	}
	if strings.Contains(x, "dv") {
		fmt.Println("Dairy value too low")
	}
	if strings.Contains(x, "dC") {
		fmt.Println("Dairy count too high")
	}
	if strings.Contains(x, "dc") {
		fmt.Println("Dairy count to low")
	}
	if strings.Contains(x, "bV") {
		fmt.Println("Bug value too high")
	}
	if strings.Contains(x, "bv") {
		fmt.Println("Bug value too low")
	}
	if strings.Contains(x, "bC") {
		fmt.Println("Bug count too high")
	}
	if strings.Contains(x, "bc") {
		fmt.Println("Bug count to low")
	}
	if strings.Contains(x, "iV") {
		fmt.Println("Inedible value too high")
	}
	if strings.Contains(x, "iv") {
		fmt.Println("Inedible value too low")
	}
	if strings.Contains(x, "iC") {
		fmt.Println("Inedible count too high")
	}
	if strings.Contains(x, "ic") {
		fmt.Println("Inedible count to low")
	}
	if strings.Contains(x, "lV") {
		fmt.Println("Misc value too high")
	}
	if strings.Contains(x, "lv") {
		fmt.Println("Misc value too low")
	}
	if strings.Contains(x, "lC") {
		fmt.Println("Misc count too high")
	}
	if strings.Contains(x, "lc") {
		fmt.Println("Misc count to low")
	}
	if strings.Contains(x, "oC") {
		fmt.Println("Required Ingredient count too high")
	}
	if strings.Contains(x, "oc") {
		fmt.Println("Required Ingredient count too low")
	}
	if strings.Contains(x, "x") {
		fmt.Println("Crock Has Exclusion")
	}
}
func testerRecipes(recipeData map[string]RecipeDetails, attributeCounts AttributeCounts, attributeVals AttributeVals, crockPot []IngredientDetails) map[string]RecipeDetails {

	// // AMBEROSIA
	// if countIngName("Collected Dust", crockPot) < 1 {
	// }

	// ASPARAGUS SOUP
	if countIngName("Asparagus", crockPot) < 1 {
	}
	if attributeVals.vegVal < 1.5 {
	}
	if attributeCounts.meatCount != 0 {
	}
	if attributeCounts.inedibleCount != 0 {
	}

	// Bacon and Eggs
	if attributeVals.meatVal <= 1 {
	}
	if attributeVals.eggVal <= 1 {
	}
	if attributeVals.vegVal != 0 {
	}

	// BANANA POP
	if countIngName("Banana", crockPot) < 1 {
	}
	if countIngName("Twigs", crockPot) < 1 {
	}
	if attributeCounts.meatCount != 0 {
	}
	if attributeCounts.fishCount != 0 {
	}

	// Barnacle Linguine
	if countIngName("Barnacles", crockPot) != 2 {
	}
	if attributeCounts.vegeCount != 2 {
	}

	// Barnacle Nigiri
	if countIngName("Barnacles", crockPot) < 1 {
	}
	if countIngName("Kelp Fronds", crockPot) < 1 {
	}
	if attributeCounts.eggCount < 1 {
	}

	// BARNACLE PITA
	if countIngName("Barnacles", crockPot) < 1 {
	}
	if countIngName("Kelp Fronds", crockPot) < 1 {
	}
	if attributeCounts.eggCount < 1 {
	}

	// BEEFY GREENS
	if countIngName("Leafy Meat", crockPot) < 1 {
	}
	if attributeVals.vegVal < 3 {
	}

	// butter muffin
	if countIngName("Butterfly Wings", crockPot) < 1 {
	}
	if attributeCounts.vegeCount < 1 {
	}
	if stringInSlice("Mandrake", crockPot) {
	}
	if attributeCounts.meatCount != 0 {
	}

	// CALIFORNIA ROLL
	if countIngName("Kelp Fronds", crockPot) < 2 {
	}
	if attributeVals.fishVal < 1 {
	}

	// ceviche
	if countIngName("Ice", crockPot) < 2 {
	}
	if attributeVals.fishVal < 2 {
	}
	if attributeCounts.eggCount != 0 {
	}
	if attributeCounts.inedibleCount != 0 {
	}

	// CREAMY POTATO PURÉE**
	// need to look at potato or roasted potato
	if countIngName("Potato", crockPot) < 2 {
	}
	if countIngName("Garlic", crockPot) < 1 {
	}
	if stringInSlice("Twigs", crockPot) {
	}
	if attributeCounts.meatCount != 0 {
	}

	// dragonpie
	// need to look at dragon fruit or prepared dragon fruit
	if countIngName("Dragon Fruit", crockPot) < 1 {
	}
	if stringInSlice("Mandrake", crockPot) {
	}
	if attributeCounts.meatCount != 0 {
	}

	// FANCY SPIRALLED TUBERS
	// need to look at potato or roasted potato
	if countIngName("Potato", crockPot) < 1 {
	}
	if countIngName("Twigs", crockPot) < 1 {
	}
	if attributeCounts.inedibleCount-1 > 1 {
	}
	if attributeCounts.meatCount != 0 {
	}

	// fish tacos
	// twig value - 1 twig 50% chance of  fish sticks
	// need to look at corn or popcorn
	if attributeVals.fishVal < 0.5 {
	}
	if countIngName("Corn", crockPot) < 1 && countIngName("Popcorn", crockPot) < 1 {
	}

	// fishsticks
	if attributeVals.fishVal < 0.25 {
	}
	if countIngName("Twigs", crockPot) != 1 {
	}
	if stringInSlice("Moleworm", crockPot) {
	}

	// fist full of jam
	if attributeVals.fruitVal < 0.5 {
	}
	if attributeVals.meatVal != 0 {
	}
	if attributeVals.vegVal != 0 {
	}
	if attributeVals.inedVal != 0 {
	}
	if stringInSlice("Dragon Fruit", crockPot) {
	}

	// flower salad
	if countIngName("Cactus Flower", crockPot) < 1 {
	}
	if attributeVals.vegVal-.5 < 1.5 {
	}
	if attributeVals.meatVal != 0 {
	}
	if attributeVals.fruitVal != 0 {
	}
	if attributeVals.eggVal != 0 {
	}
	if attributeVals.sweetVal != 0 {
	}
	if stringInSlice("Twigs", crockPot) {
	}

	// froggle bunwich
	// need to look at frog legs or cooked frog legs
	// makes kabob if only one stick
	if countIngName("Frog Legs", crockPot) < 1 {
	}
	if attributeCounts.vegeCount < 1 {
	}
	if attributeVals.eggVal != 0 {
	}
	if attributeVals.sweetVal != 0 {
	}
	if stringInSlice("Mandrake", crockPot) {
	}

	// fruit medley
	// twigs is safest anything else 50% of fist full of jam
	if attributeVals.fruitVal < 3 {
	}
	if attributeVals.meatVal != 0 {
	}
	if attributeVals.vegVal != 0 {
	}
	if stringInSlice("Dragon Fruit", crockPot) {
	}

	// guacamole
	// need to look at cactus flesh or stone fruit
	if countIngName("Moleworm", crockPot) < 1 {
	}
	if countIngName("Cactus Flesh", crockPot) < 1 && countIngName("Ripe Stone Fruit", crockPot) < 1 {
	}
	if attributeVals.fruitVal != 0 {
	}

	// Honey Ham
	if countIngName("Honey", crockPot) < 1 {
	}
	if attributeVals.meatVal <= 1.5 {
	}
	if stringInSlice("Twigs", crockPot) {
	}
	if stringInSlice("Moleworm", crockPot) {
	}
	if stringInSlice("Mandrake", crockPot) {
	}
	if stringInSlice("Tallbird Egg", crockPot) {
	}

	// Honey Nuggets
	if countIngName("Honey", crockPot) < 1 {
	}
	if attributeVals.meatVal > 1.5 {
	}
	if attributeVals.inedVal != 0 {
	}

	// ice cream
	if countIngName("Ice", crockPot) < 1 {
	}
	if attributeCounts.dairyCount < 1 {
	}
	if attributeCounts.sweetenerCount < 1 {
	}
	if attributeVals.meatVal != 0 {
	}
	if attributeVals.vegVal != 0 {
	}
	if attributeVals.eggVal != 0 {
	}
	if stringInSlice("Twigs", crockPot) {
	}

	// jelly salad
	// look for cooked version
	if countIngName("Leafy Meat", crockPot) < 2 {
	}
	if attributeVals.sweetVal < 2 {
	}

	// jellybeans
	if countIngName("Royal Jelly", crockPot) < 1 {
	}
	if attributeVals.monVal != 0 {
	}
	if attributeVals.inedVal != 0 {
	}

	// kabobs
	if attributeCounts.meatCount < 1 {
	}
	if countIngName("Twigs", crockPot) != 1 {
	}
	if stringInSlice("Moleworm", crockPot) {
	}
	if stringInSlice("Mandrake", crockPot) {
	}
	if attributeVals.fishVal != 0 {
	}

	// leafy meatloaf
	if countIngName("Leafy Meat", crockPot) < 2 {
	}

	// LOBSTER BISQUE
	if countIngName("Wobster", crockPot) < 1 {
	}
	if countIngName("Ice", crockPot) < 1 {
	}

	// mandrake soup
	if countIngName("Mandrake", crockPot) < 1 {
	}

	// Meatballs
	if attributeVals.meatVal >= 3 {
	}
	if attributeVals.meatVal == 0 {
	}
	if stringInSlice("Twigs", crockPot) {
	}

	// meaty stew
	if attributeVals.meatVal < 3 {
	}
	if stringInSlice("Twigs", crockPot) {
	}
	if stringInSlice("Moleworm", crockPot) {
	}
	if stringInSlice("Honey", crockPot) {
	}
	if stringInSlice("Mandrake", crockPot) {
	}
	if stringInSlice("Tallbird Egg", crockPot) {
	}

	// melonsicle
	if countIngName("Watermelon", crockPot) < 1 {
	}
	if countIngName("Ice", crockPot) < 1 {
	}
	if countIngName("Twigs", crockPot) < 1 {
	}
	if attributeVals.meatVal != 0 {
	}
	if attributeVals.vegVal != 0 {
	}
	if attributeVals.eggVal != 0 {
	}

	// MILKMADE HAT
	if countIngName("Nostrils", crockPot) < 1 {
	}
	if countIngName("Kelp Fronds", crockPot) < 1 {
	}
	if attributeCounts.dairyCount < 1 {
	}

	// monster lasagna
	if attributeCounts.monsterCount < 2 {
	}
	if stringInSlice("Twigs", crockPot) {
	}

	// Mushy Cake
	if countIngName("Moon Shroom", crockPot) != 1 {
	}
	if countIngName("Red Cap", crockPot) != 1 {
	}
	if countIngName("Blue Cap", crockPot) != 1 {
	}
	if countIngName("Green Cap", crockPot) != 1 {
	}

	// Pierogi
	if attributeCounts.meatCount < 1 {
	}
	if attributeCounts.eggCount < 1 {
	}
	if attributeCounts.vegeCount < 1 {
	}
	if stringInSlice("Twigs", crockPot) {
	}
	if stringInSlice("Mandrake", crockPot) {
	}

	//powdercake
	if countIngName("Corn", crockPot) < 1 {
	}
	if countIngName("Honey", crockPot) < 1 {
	}
	if countIngName("Twigs", crockPot) < 1 {
	}

	//Pumkin cookie
	// 3 honey or comb 50% chance of making taffy
	if countIngName("Pumpkin", crockPot) < 1 {
	}
	if countIngName("Honey", crockPot) <= 1 {
	}

	// RATATOUILLE
	if attributeCounts.vegeCount < 1 {
	}
	if stringInSlice("Twigs", crockPot) {
	}
	if stringInSlice("Mandrake", crockPot) {
	}
	if stringInSlice("Butterfly Wings", crockPot) {
	}
	if stringInSlice("Dragon Fruit", crockPot) {
	}

	// SALSA FRESCA
	// look for cooked versions
	if countIngName("Toma Root", crockPot) < 1 {
	}
	if attributeCounts.meatCount != 0 {
	}
	if attributeCounts.inedibleCount != 0 {
	}
	if attributeCounts.eggCount != 0 {
	}

	// SEAFOOD GUMBO
	if countIngName("Eel", crockPot) < 1 {
	}
	if attributeVals.fishVal <= 2 {
	}

	//  Soothing tea
	// look for honey or comb
	if countIngName("Forget-Me-Lots", crockPot) < 1 {
	}
	if countIngName("Honey", crockPot) < 1 {
	}
	if countIngName("Ice", crockPot) < 1 {
	}
	if attributeCounts.monsterCount != 0 {
	}
	if attributeCounts.meatCount != 0 {
	}
	if attributeCounts.fishCount != 0 {
	}
	if attributeCounts.eggCount != 0 {
	}
	if attributeCounts.inedibleCount != 0 {
	}
	if attributeCounts.dairyCount != 0 {
	}

	//  Spicy Chili
	if attributeCounts.meatCount != 2 {
	}
	if attributeCounts.vegeCount != 2 {
	}
	if attributeVals.meatVal < 1.5 {
	}
	if attributeVals.vegVal < 1.5 {
	}

	// STUFFED EGGPLANT
	if countIngName("Eggplant", crockPot) < 1 {
	}
	if attributeCounts.vegeCount < 1 {
	}

	// stuffed fish heads
	if countIngName("Barnacles", crockPot) < 1 {
	}
	if attributeVals.fishVal-0.5 < 1 {
	}

	// STUFFED PEPPER POPPERS
	if countIngName("Pepper", crockPot) < 1 {
	}
	if attributeVals.meatVal > 1.5 {
	}
	if attributeCounts.meatCount == 0 {
	}
	if stringInSlice("Twigs", crockPot) {
	}

	// SURF 'N' TURF
	if attributeVals.meatVal < 2.5 {
	}
	if attributeVals.fishVal < 1.5 {
	}
	if stringInSlice("Ice", crockPot) {
	}

	// Taffy
	if attributeVals.sweetVal < 3 {
	}
	if attributeCounts.meatCount != 0 {
	}

	// trail mix
	// required berries must be uncooked
	if countIngName("Roasted Birchnut", crockPot) < 1 {
	}
	if countIngName("Berries", crockPot) < 1 {
	}
	if countIngName("Roasted", crockPot) > 3 {
	}
	if attributeCounts.fruitCount-1 < 1 {
	}
	if attributeCounts.meatCount != 0 {
	}
	if attributeCounts.fishCount != 0 {
	}
	if attributeCounts.eggCount != 0 {
	}
	if attributeCounts.vegeCount != 0 {
	}
	if attributeCounts.dairyCount != 0 {
	}

	// TURKEY DINNER
	// checkingg twice for veg or fruit
	if countIngName("Drumstick", crockPot) < 2 {
	}
	if attributeVals.meatVal-1 < 0.25 {
	}
	if attributeVals.fruitVal < .5 && attributeVals.vegVal < .5 {
	}

	// unagi
	if countIngName("Eel", crockPot) < 1 {
	}
	if countIngName("Lichen", crockPot) < 1 && countIngName("Kelp Fronds", crockPot) < 1 {
	}

	// VEGETABLE STINGER
	if countIngName("Toma Root", crockPot) < 1 && countIngName("Asparagus", crockPot) < 1 {
	}
	if countIngName("Ice", crockPot) < 1 {
	}
	if attributeVals.vegVal-1 < 1.5 {
	}

	// Veggie Burger
	if countIngName("Leafy Meat", crockPot) < 1 {
	}
	if countIngName("Onion", crockPot) < 1 {
	}
	if attributeVals.vegVal-1 < 1 {
	}

	// waffles
	if countIngName("Butter", crockPot) < 1 {
	}
	if countIngName("Berries", crockPot) < 1 {
	}
	if attributeCounts.eggCount < 1 {
	}

	// wet goop
	// if everything is false

	// WOBSTER DINNER
	if countIngName("Wobster", crockPot) < 1 {
	}
	if countIngName("Butter", crockPot) < 1 {
	}
	if attributeCounts.meatCount != 0 {
	}
	if attributeCounts.fishCount != 0 {
	}
	if stringInSlice("Twigs", crockPot) {
	}
	return recipeData
}

// for printing valid recipes
func orgValidRec(recipeData map[string]RecipeDetails) {
	// creating new map so remove recipes to keep integrity of original map recipeData
	fRecipeData := map[string]RecipeDetails{}
	for key, value := range recipeData {
		fRecipeData[key] = value
	}

	//  part of percentage printing
	var bigPval float64 = -100
	for _, existRec := range fRecipeData {
		if existRec.PRIORITY > bigPval {
			bigPval = existRec.PRIORITY
		}
	}
	// using new map to order priority by printing highest priority then removing and then continuing
	rdatacount := len(fRecipeData) /*this is the amount of items in the new map*/
	for rdatacount > 0 {           /*beginning of loop. the loop will occur how ever many items are in the map*/
		var pval float64 = -100 /*setting priority value to -100 (no priorty is that low) to find and set highest pvalue in list*/
		for _, existRec := range fRecipeData {
			if existRec.PRIORITY > pval {
				pval = existRec.PRIORITY
			}
		}
		var pcount float64 = 0 /* counts how many have same priority */
		for _, existRec := range fRecipeData {
			if existRec.PRIORITY == pval {
				pcount++
			}
		}
		// output
		/*then whatever priority matches pval then gets printed then is removed from list (will do multiple if more than one) */
		for t, existRec := range fRecipeData {
			if existRec.PRIORITY == pval {
				fmt.Println()
				if bigPval == existRec.PRIORITY && pcount == 1 {
					fmt.Println("%", 100) /*highest priorty*/
				} else if bigPval == existRec.PRIORITY && pcount != 1 {
					fmt.Println("%", 100/pcount) /*splits percentage if muli*/
				} else {
					fmt.Println("%", 0/pcount) /*0% for everything else*/
				}
				fmt.Println(t, ":")
				fmt.Println("Health:", existRec.HEALTH)
				fmt.Println("Hunger:", existRec.HUNGER)
				fmt.Println("Sanity:", existRec.SANITY)
				fmt.Println("Ingredients:", existRec.INGREDIENTS)
				fmt.Println("Exclude:", existRec.EXCLUDE)
				fmt.Println("Notes:", existRec.NOTES)
				fmt.Println("Expires:", existRec.EXPIRES)
				fmt.Println("Cooktime:", existRec.COOKTIME)
				fmt.Println("Priority:", existRec.PRIORITY)
				delete(fRecipeData, t)
			}
		}
		rdatacount-- /*record/s that match pval then are removed*/
	} /*end of for loop */
}

func readRecipeData() map[string]RecipeDetails {
	// Find that file recipes
	jsonFile2, err := os.Open("../lib/recipe.json")
	if err != nil {
		fmt.Println(err)
	}

	// Read that file
	byteValue2, err := ioutil.ReadAll(jsonFile2)
	if err != nil {
		fmt.Println(err)
	}

	// Fit the file
	innerRecipeData := map[string]RecipeDetails{}
	err = json.Unmarshal(byteValue2, &innerRecipeData)
	if err != nil {
		fmt.Println(err)
	}

	return innerRecipeData
}

var recipeData = readRecipeData()

func main() {

	// Find that file ingredients
	jsonFile, err := os.Open("../lib/ingredientsv2.json")
	if err != nil {
		fmt.Println(err)
	}

	// Read that file
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}

	// Fit the file
	ingredientsData := map[string]IngredientDetails{}
	err = json.Unmarshal(byteValue, &ingredientsData)
	if err != nil {
		fmt.Println(err)
	}

	// possibleRecipeData := map[string]RecipeDetails{}
	// notPossibleRecipeData := map[string]RecipeDetails{}

	// crockpot = 0
	// ingsearch = 1
	// recsearch = 2
	// mode = 0

	// deleting all ingredients that cant be added to crockpot
	for x, ingredata := range ingredientsData {
		if ingredata.NOTES == "cannot be added to crock pot" {
			delete(ingredientsData, x)
		}
	}

	// Add Ingredients to the pot
	i1 := "Jerky"
	i2 := "Jerky"
	i3 := "Jerky"
	i4 := "Jerky"
	// Jerky
	// Egg
	// Barnacles
	// Blue Cap
	crockPot := make([]IngredientDetails, 0)
	crockSlots := []string{i1, i2, i3, i4}
	for _, slot := range crockSlots {
		if slot != "" {
			crockPot = append(crockPot, ingredientsData[slot])
		}
	}

	// sorts crockpot by name
	sort.Slice(crockPot[:], func(i, j int) bool {
		return crockPot[i].NAME < crockPot[j].NAME
	})

	for _, crockPot := range crockPot {
		fmt.Print(crockPot.NAME, ", ")
	}
	fmt.Println()
	// Count Ingredient Vals
	attributeVals := masterVals(crockPot)
	// Count Ingredient types
	attributeCounts := masterCounts(crockPot)

	// if i4 != "" { star hurrr
	// function that deletes false recipes
	deleteRecipes(recipeData, attributeCounts, attributeVals, crockPot)
	// ------------------------------------------------------

	//possibleRecipes
	_ = compilePossibleRecipes(crockPot, attributeVals, attributeCounts)
	// TO PRINT:
	// Loop over possibleRecipes
	// Print recipe.Name
	// run processPossible(recipe)

	// *****
	// Pierogi
	if attributeCounts.meatCount >= 1 && attributeCounts.eggCount >= 1 && attributeCounts.vegeCount >= 1 && stringNotInSlice("Twigs", crockPot) && stringNotInSlice("Mandrake", crockPot) {
		fmt.Println("ok pierogi")
	}
	// // -----------------------------------------------------------------------------------------------------------------------------------

	// Output goes here

	//  this output only shows the vals that are greater than 0
	outVal := map[string]float64{"Meat value:": attributeVals.meatVal, "Fish value:": attributeVals.fishVal, "Egg value:": attributeVals.eggVal, "Fruit value:": attributeVals.fruitVal, "Vegetable value:": attributeVals.vegVal, "Sweetener value:": attributeVals.sweetVal, "Monster value:": attributeVals.monVal, "Dairy value:": attributeVals.dairyVal, "Bug value:": attributeVals.bugVal, "Inedible value:": attributeVals.inedVal, "Misc value:": attributeVals.miscVal}
	for inName, existVal := range outVal {
		if existVal > 0 {
			fmt.Println(inName, "value", existVal)
		}
		// fmt.Println(inName, existVal)
	}

	//  this output only shows the counts that are greater than 0
	outCount := map[string]float64{"Meat count:": attributeCounts.meatCount, "Fish count:": attributeCounts.fishCount, "Egg count:": attributeCounts.eggCount, "Fruit count:": attributeCounts.fruitCount, "Vegetable count:": attributeCounts.vegeCount, "Sweetener count:": attributeCounts.sweetenerCount, "Monster count:": attributeCounts.monsterCount, "Dairy count:": attributeCounts.dairyCount, "Bug count:": attributeCounts.bugCount, "Inedible count:": attributeCounts.inedibleCount, "Misc count:": attributeCounts.miscCount}
	for inName, existcount := range outCount {
		if existcount > 0 {
			fmt.Println(inName, existcount)
		}
		// fmt.Println(inName, existcount)
	}

	// if eaten raw
	fmt.Println()
	fmt.Println("Ingredients")
	fmt.Println("Total Health", attributeVals.healVal)
	fmt.Println("Total Hunger", attributeVals.hungVal)
	fmt.Println("Total Sanity", attributeVals.saniVal)

	// prints raw exp
	for i := 0; i < len(crockPot); i++ {
		fmt.Println(crockPot[i].NAME, "expires in", crockPot[i].EXPIRE, "days")
	}

	// creating new map so remove recipes to keep integrity of original map recipeData

	orgValidRec(recipeData)
	/**************************************************************************************************************start heerrrr

	} else { //////////////////////////////// START OF FIRST LOOP
		// BARNACLE NIGIRI
		barnacleNigiri := []string{}
		for ingreTrue := range ingredientsData {
			i4 := ingreTrue
			titles = "[" + i1 + ", " + i2 + ", " + i3 + ", " + i4 + "]"
			// Look through each ingredient and sets values
			// count Ingredient Valus
			attributeVals := masterVals(ingredientsData, i1, i2, i3, i4)
			// Count Ingredient types
			attributeCounts := masterCounts(ingredientsData, i1, i2, i3, i4)

			blankRec(recipeData, attributeCounts, attributeVals, titles, crockPot)
			// test
			if strings.Count(titles, "Barnacles") < 1 || strings.Count(titles, "Kelp Fronds") < 1 || attributeCounts.eggCount < 1 {
			} else {
				barnacleNigiri = append(barnacleNigiri, ingreTrue)
			}
		}
		if len(barnacleNigiri) != 0 {
			sort.Strings(barnacleNigiri)
			fmt.Println("BARNACLE NIGIRI")
			for _, ingre := range barnacleNigiri {
				fmt.Println(ingre)
			}
			fmt.Println()
		} // end of first loop

		baconAndEggs := []string{}
		for ingreTrue := range ingredientsData {
			i4 := ingreTrue
			titles = "[" + i1 + ", " + i2 + ", " + i3 + ", " + i4 + "]"
			// Look through each ingredient and sets values
			// count Ingredient Valus
			attributeVals := masterVals(ingredientsData, i1, i2, i3, i4)
			// Count Ingredient types
			attributeCounts := masterCounts(ingredientsData, i1, i2, i3, i4)

			blankRec(recipeData, attributeCounts, attributeVals, titles, crockPot)
			// test
			if attributeVals.meatVal <= 1 || attributeVals.eggVal <= 1 || attributeVals.vegVal != 0 {
			} else {
				baconAndEggs = append(baconAndEggs, ingreTrue)
			}
		}
		if len(baconAndEggs) != 0 {
			sort.Strings(baconAndEggs)
			fmt.Println("BACON AND EGGS")
			for _, ingre := range baconAndEggs {
				fmt.Println(ingre)
			}
			fmt.Println()
		} // end of first loop

		pierogi := []string{}
		for ingreTrue := range ingredientsData {
			i4 := ingreTrue
			titles = "[" + i1 + ", " + i2 + ", " + i3 + ", " + i4 + "]"
			// Look through each ingredient and sets values
			// count Ingredient Valus
			attributeVals := masterVals(ingredientsData, i1, i2, i3, i4)
			// Count Ingredient types
			attributeCounts := masterCounts(ingredientsData, i1, i2, i3, i4)

			blankRec(recipeData, attributeCounts, attributeVals, titles, crockPot)
			// test
			if attributeCounts.meatCount < 1 || attributeCounts.eggCount < 1 || attributeCounts.vegeCount < 1 || crockPot["Twigs"].NAME != "" || crockPot["Mandrake"].NAME != "" {
			} else {
				pierogi = append(pierogi, ingreTrue)
			}
		}
		if len(pierogi) != 0 {
			sort.Strings(pierogi)
			fmt.Println("PIEROGI")
			for _, ingre := range pierogi {
				fmt.Println(ingre)
			}
			fmt.Println()
		} // end of first loop
	}
	// ****standard output****
	// fmt.Println("FOR TESTINGS Meat: ", attributeVals.meatVal)
	// fmt.Println("FOR TESTINGS Fish: ", attributeVals.fishVal)
	// fmt.Println("FOR TESTINGS Egg: ", attributeVals.eggVal)
	// fmt.Println("FOR TESTINGS Fruit: ", attributeVals.fruitVal)
	// fmt.Println("FOR TESTINGS Vegetable: ", attributeVals.vegVal)
	// fmt.Println("FOR TESTINGS Sweetener: ", attributeVals.sweetVal)
	// fmt.Println("FOR TESTINGS Monster: ", attributeVals.monVal)
	// fmt.Println("FOR TESTINGS Dairy: ", attributeVals.dairyVal)
	// fmt.Println("FOR TESTINGS Bug: ", attributeVals.bugVal)
	// fmt.Println("FOR TESTINGS Inedible: ", attributeVals.inedVal)
	// fmt.Println("FOR TESTINGS Misc: ", attributeVals.miscVal)
	// ****count found types****
	// fmt.Println("FOR TESTINGS meat", meatCount)
	// fmt.Println("FOR TESTINGS fish", fishCount)
	// fmt.Println("FOR TESTINGS egg", eggCount)
	// fmt.Println("FOR TESTINGS fruit", fruitCount)
	// fmt.Println("FOR TESTINGS vege", vegeCount)
	// fmt.Println("FOR TESTINGS sweetener", sweetenerCount)
	// fmt.Println("FOR TESTINGS monster", monsterCount)
	// fmt.Println("FOR TESTINGS dairy", dairyCount)
	// fmt.Println("FOR TESTINGS bug", bugCount)
	// fmt.Println("FOR TESTINGS inedible", inedibleCount)
	// fmt.Println("FOR TESTINGS misc", miscCount)
	// printed not organized
	// for t, existRec := range recipeData {
	// fmt.Println("FOR TESTING", "\n")
	// 	fmt.Println("FOR TESTING", t , ":")
	// 	fmt.Println("FOR TESTING Health:", existRec.HEALTH)
	// 	fmt.Println("FOR TESTING Hunger:", existRec.HUNGER)
	// 	fmt.Println("FOR TESTING Sanity:", existRec.SANITY)
	// 	fmt.Println("FOR TESTING Ingredients:", existRec.INGREDIENTS)
	// 	fmt.Println("FOR TESTING Exclude:", existRec.EXCLUDE)
	// 	fmt.Println("FOR TESTING Notes:", existRec.NOTES)
	// 	fmt.Println("FOR TESTING Expires:", existRec.EXPIRES)
	// 	fmt.Println("FOR TESTING Cooktime:", existRec.COOKTIME)
	// 	fmt.Println("FOR TESTING Priority:", existRec.PRIORITY)
	*/ //end herrr******************************************************************************************
}
