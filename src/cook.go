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
	NAME        string  `json:"FOOD"`
	HEALTH      float64 `json:"HEALTH"`
	HUNGER      float64 `json:"HUNGER"`
	SANITY      float64 `json:"SANITY"`
	INGREDIENTS string  `json:"INGREDIENTS"`
	EXCLUDE     string  `json:"EXCLUDE"`
	NOTES       string  `json:"NOTES"`
	EXPIRES     float64 `json:"EXPIRES"`
	COOKTIME    float64 `json:"COOKTIME"`
	PRIORITY    float64 `json:"PRIORITY"`
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

type CrockPot struct {
	i1 IngredientDetails
	i2 IngredientDetails
	i3 IngredientDetails
	i4 IngredientDetails
}

// func deleteFalseReipe(recipeData map[string]RecipeDetails, masterCounts AttributeCounts, masterVals AttributeVals, titles string) map[string]RecipeDetails

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
func deleteRecipes(recipeData map[string]RecipeDetails, attributeCounts AttributeCounts, attributeVals AttributeVals, titles string, crockPot map[string]IngredientDetails) map[string]RecipeDetails {

	// AMBEROSIA
	if strings.Count(titles, "Collected Dust") < 1 {
		delete(recipeData, "AMBEROSIA")
	}

	// ASPARAGUS SOUP
	if strings.Count(titles, "Asparagus") < 1 || attributeVals.vegVal < 1.5 || attributeCounts.meatCount != 0 || attributeCounts.inedibleCount != 0 {
		delete(recipeData, "ASPARAGUS SOUP")
	}

	// Bacon and Eggs
	if attributeVals.meatVal <= 1 || attributeVals.eggVal <= 1 || attributeVals.vegVal != 0 {
		delete(recipeData, "BACON AND EGGS")
	}

	// BANANA POP
	if strings.Count(titles, "Banana") < 1 || strings.Count(titles, "twigs") < 1 || attributeCounts.meatCount != 0 || attributeCounts.fishCount != 0 {
		delete(recipeData, "BANANA POP")
	}

	// Barnacle Linguine
	if strings.Count(titles, "Barnacles") != 2 || attributeCounts.vegeCount != 2 {
		delete(recipeData, "BARNACLE LINGUINE")
	}

	// Barnacle Nigiri
	if strings.Count(titles, "Barnacles") < 1 || strings.Count(titles, "Kelp Fronds") < 1 || attributeCounts.eggCount < 1 {
		delete(recipeData, "BARNACLE NIGIRI")
	}

	// BARNACLE PITA
	// attributeVals.fishval over 1 can cause stuffed fish head but idk the percentage
	if strings.Count(titles, "Barnacles") < 1 || strings.Count(titles, "Kelp Fronds") < 1 || attributeCounts.eggCount < 1 {
		delete(recipeData, "BARNACLE PITA")
	}

	// BEEFY GREENS
	if strings.Count(titles, "Leafy Meat") < 1 || attributeVals.vegVal < 3 {
		delete(recipeData, "BEEFY GREENS")
	}

	// butter muffin
	if strings.Count(titles, "Butterfly Wings") < 1 || attributeCounts.vegeCount < 1 || crockPot["Mandrake"].NAME != "" || attributeCounts.meatCount != 0 {
		delete(recipeData, "BUTTER MUFFIN")
	}

	// CALIFORNIA ROLL
	if strings.Count(titles, "Kelp Fronds") < 2 || attributeVals.fishVal < 1 {
		delete(recipeData, "CALIFORNIA ROLL")
	}

	// ceviche
	if strings.Count(titles, "Ice") < 2 || attributeVals.fishVal < 2 || attributeCounts.eggCount != 0 || attributeCounts.inedibleCount != 0 {
		delete(recipeData, "CEVICHE")
	}

	// CREAMY POTATO PURÉE**
	// need to look at potato or roasted potato
	if strings.Count(titles, "Potato") < 2 || strings.Count(titles, "Garlic") < 1 || crockPot["Twigs"].NAME != "" || attributeCounts.meatCount != 0 {
		delete(recipeData, "CREAMY POTATO PURÉE")
	}

	// dragonpie
	// need to look at dragon fruit or prepared dragon fruit
	if strings.Count(titles, "Dragon Fruit") < 1 || crockPot["Mandrake"].NAME != "" || attributeCounts.meatCount != 0 {
		delete(recipeData, "DRAGONPIE")
	}

	// FANCY SPIRALLED TUBERS
	// need to look at potato or roasted potato
	if strings.Count(titles, "Potato") < 1 || strings.Count(titles, "Twigs") < 1 || attributeCounts.inedibleCount-1 > 1 || attributeCounts.meatCount != 0 {
		delete(recipeData, "FANCY SPIRALLED TUBERS")
	}

	// fish tacos
	// twig value - 1 twig 50% chance of  fish sticks
	// need to look at corn or popcorn
	if attributeVals.fishVal < 0.5 || (strings.Count(titles, "Corn") < 1 || strings.Count(titles, "Popcorn") < 1) {
		delete(recipeData, "FISH TACOS")
	}

	// fishsticks
	if attributeVals.fishVal < 0.25 || strings.Count(titles, "Twigs") != 1 || crockPot["Moleworm"].NAME != "" {
		delete(recipeData, "FISHSTICKS")
	}

	// fist full of jam
	if attributeVals.fruitVal < 0.5 || attributeVals.meatVal != 0 || attributeVals.vegVal != 0 || attributeVals.inedVal != 0 || crockPot["Dragon Fruit"].NAME != "" {
		delete(recipeData, "FIST FULL OF JAM")
	}

	// flower salad
	if strings.Count(titles, "Cactus Flower") < 1 || attributeVals.vegVal-.5 < 1.5 || attributeVals.meatVal != 0 || attributeVals.fruitVal != 0 || attributeVals.eggVal != 0 || attributeVals.sweetVal != 0 || crockPot["Twigs"].NAME != "" {
		delete(recipeData, "FLOWER SALAD")
	}

	// froggle bunwich
	// need to look at frog legs or cooked frog legs
	// makes kabob if only one stick
	if strings.Count(titles, "Frog Legs") < 1 || attributeCounts.vegeCount < 1 || attributeVals.eggVal != 0 || attributeVals.sweetVal != 0 || crockPot["Mandrake"].NAME != "" {
		delete(recipeData, "FROGGLE BUNWICH")
	}

	// fruit medley
	// twigs is safest anything else 50% of fist full of jam
	if attributeVals.fruitVal < 3 || attributeVals.meatVal != 0 || attributeVals.vegVal != 0 || crockPot["Dragon Fruit"].NAME != "" {
		delete(recipeData, "FRUIT MEDLEY")
	}

	// guacamole
	// need to look at cactus flesh or stone fruit
	if strings.Count(titles, "Moleworm") < 1 || (strings.Count(titles, "Cactus Flesh") < 1 || strings.Count(titles, "Ripe Stone Fruit") < 1) || attributeVals.fruitVal != 0 {
		delete(recipeData, "GUACAMOLE")
	}

	// Honey Ham
	if strings.Count(titles, "Honey") < 1 || attributeVals.meatVal <= 1.5 || crockPot["Twigs"].NAME != "" || crockPot["Moleworm"].NAME != "" || crockPot["Mandrake"].NAME != "" || crockPot["Tallbird Egg"].NAME != "" {
		delete(recipeData, "HONEY HAM")
	}

	// Honey Nuggets
	if strings.Count(titles, "Honey") < 1 || attributeVals.meatVal > 1.5 || attributeVals.inedVal != 0 {
		delete(recipeData, "HONEY NUGGETS")
	}

	// ice cream
	if strings.Count(titles, "Ice") < 1 || attributeCounts.dairyCount < 1 || attributeCounts.sweetenerCount < 1 || attributeVals.meatVal != 0 || attributeVals.vegVal != 0 || attributeVals.eggVal != 0 || crockPot["Twigs"].NAME != "" {
		delete(recipeData, "ICE CREAM")
	}

	// jelly salad
	// look for cooked version
	if strings.Count(titles, "Leafy Meat") < 2 || attributeVals.sweetVal < 2 {
		delete(recipeData, "JELLY SALAD")
	}

	// jellybeans
	if strings.Count(titles, "Royal Jelly") < 1 || attributeVals.monVal != 0 || attributeVals.inedVal != 0 {
		delete(recipeData, "JELLYBEANS")
	}

	// kabobs
	if attributeCounts.meatCount < 1 || strings.Count(titles, "Twigs") != 1 || crockPot["Moleworm"].NAME != "" || crockPot["Mandrake"].NAME != "" || attributeVals.fishVal != 0 {
		delete(recipeData, "KABOBS")
	}

	// leafy meatloaf
	if strings.Count(titles, "Leafy Meat") < 2 {
		delete(recipeData, "LEAFY MEATLOAF")
	}

	// LOBSTER BISQUE
	if strings.Count(titles, "Wobster") < 1 || strings.Count(titles, "Ice") < 1 {
		delete(recipeData, "LOBSTER BISQUE")
	}

	// mandrake soup
	if strings.Count(titles, "Mandrake") < 1 {
		delete(recipeData, "MANDRAKE SOUP")
	}

	// Meatballs
	if attributeVals.meatVal >= 3 || attributeVals.meatVal == 0 || crockPot["Twigs"].NAME != "" {
		delete(recipeData, "MEATBALLS")
	}

	// meaty stew
	if attributeVals.meatVal < 3 || crockPot["Twigs"].NAME != "" || crockPot["Moleworm"].NAME != "" || crockPot["Honey"].NAME != "" || crockPot["Mandrake"].NAME != "" || crockPot["Tallbird Egg"].NAME != "" {
		delete(recipeData, "MEATY STEW")
	}

	// melonsicle
	if strings.Count(titles, "Watermelon") < 1 || strings.Count(titles, "Ice") < 1 || strings.Count(titles, "Twigs") < 1 || attributeVals.meatVal != 0 || attributeVals.vegVal != 0 || attributeVals.eggVal != 0 {
		delete(recipeData, "MELONSICLE")
	}

	// MILKMADE HAT
	if strings.Count(titles, "Nostrils") < 1 || strings.Count(titles, "Kelp Fronds") < 1 || attributeCounts.dairyCount < 1 {
		delete(recipeData, "MILKMADE HAT")
	}

	// monster lasagna
	if attributeCounts.monsterCount < 2 || crockPot["Twigs"].NAME != "" {
		delete(recipeData, "MONSTER LASAGNA")
	}

	// Mushy Cake
	if strings.Count(titles, "Moon Shroom") != 1 || strings.Count(titles, "Red Cap") != 1 || strings.Count(titles, "Blue Cap") != 1 || strings.Count(titles, "Green Cap") != 1 {
		delete(recipeData, "MUSHY CAKE")
	}

	// Pierogi
	if attributeCounts.meatCount < 1 || attributeCounts.eggCount < 1 || attributeCounts.vegeCount < 1 || crockPot["Twigs"].NAME != "" || crockPot["Mandrake"].NAME != "" {
		delete(recipeData, "PIEROGI")
	}

	//powdercake
	if strings.Count(titles, "Corn") < 1 || strings.Count(titles, "Honey") < 1 || strings.Count(titles, "Twigs") < 1 {
		delete(recipeData, "POWDERCAKE")
	}

	//Pumkin cookie
	// 3 honey or comb 50% chance of making taffy
	if strings.Count(titles, "Pumpkin") < 1 || strings.Count(titles, "Honey") <= 1 {
		delete(recipeData, "PUMPKIN COOKIE")
	}

	// RATATOUILLE
	if attributeCounts.vegeCount < 1 || crockPot["Twigs"].NAME != "" || crockPot["Mandrake"].NAME != "" || crockPot["Butterfly Wings"].NAME != "" || crockPot["Dragon Fruit"].NAME != "" {
		delete(recipeData, "RATATOUILLE")
	}

	// SALSA FRESCA
	// look for cooked versions
	if strings.Count(titles, "Toma Root") < 1 || attributeCounts.meatCount != 0 || attributeCounts.inedibleCount != 0 || attributeCounts.eggCount != 0 {
		delete(recipeData, "SALSA FRESCA")
	}

	// SEAFOOD GUMBO
	if strings.Count(titles, "Eel") < 1 || attributeVals.fishVal <= 2 {
		delete(recipeData, "SEAFOOD GUMBO")
	}

	//  Soothing tea
	// look for honey or comb
	if strings.Count(titles, "Forget-Me-Lots") < 1 || strings.Count(titles, "Honey") < 1 || strings.Count(titles, "Ice") < 1 || attributeCounts.monsterCount != 0 || attributeCounts.meatCount != 0 || attributeCounts.fishCount != 0 || attributeCounts.eggCount != 0 || attributeCounts.inedibleCount != 0 || attributeCounts.dairyCount != 0 {
		delete(recipeData, "SOOTHING TEA")
	}

	//  Spicy Chili
	if attributeCounts.meatCount != 2 || attributeCounts.vegeCount != 2 || attributeVals.meatVal < 1.5 || attributeVals.vegVal < 1.5 {
		delete(recipeData, "SPICY CHILI")
	}

	// STUFFED EGGPLANT
	if strings.Count(titles, "Eggplant") < 1 || attributeCounts.vegeCount < 1 {
		delete(recipeData, "STUFFED EGGPLANT")
	}

	// stuffed fish heads
	if strings.Count(titles, "Barnacles") < 1 || attributeVals.fishVal-0.5 < 1 {
		delete(recipeData, "STUFFED FISH HEADS")
	}

	// STUFFED PEPPER POPPERS
	if strings.Count(titles, "Pepper") < 1 || attributeVals.meatVal > 1.5 || attributeCounts.meatCount == 0 || crockPot["Twigs"].NAME != "" {
		delete(recipeData, "STUFFED PEPPER POPPERS")
	}

	// SURF 'N' TURF
	if attributeVals.meatVal < 2.5 || attributeVals.fishVal < 1.5 || crockPot["Ice"].NAME != "" {
		delete(recipeData, "SURF 'N' TURF")
	}

	// Taffy
	if attributeVals.sweetVal < 3 || attributeCounts.meatCount != 0 {
		delete(recipeData, "TAFFY")
	}

	// trail mix
	// required berries must be uncooked
	if strings.Count(titles, "Roasted Birchnut") < 1 || strings.Count(titles, "Berries") < 1 || strings.Count(titles, "Roasted") > 3 || attributeCounts.fruitCount-1 < 1 || attributeCounts.meatCount != 0 || attributeCounts.fishCount != 0 || attributeCounts.eggCount != 0 || attributeCounts.vegeCount != 0 || attributeCounts.dairyCount != 0 {
		delete(recipeData, "TRAIL MIX")
	}

	// TURKEY DINNER
	// checkingg twice for veg or fruit
	if strings.Count(titles, "Drumstick") < 2 || attributeVals.meatVal-1 < 0.25 || (attributeVals.fruitVal < .5 || attributeVals.vegVal < .5) {
		delete(recipeData, "TURKEY DINNER")
	}

	// unagi
	if strings.Count(titles, "Eel") < 1 || (strings.Count(titles, "Lichen") < 1 || strings.Count(titles, "Kelp Fronds") < 1) {
		delete(recipeData, "UNAGI")
	}

	// VEGETABLE STINGER
	if (strings.Count(titles, "Toma Root") < 1 || strings.Count(titles, "Asparagus") < 1) || strings.Count(titles, "Ice") < 1 || attributeVals.vegVal-1 < 1.5 {
		delete(recipeData, "VEGETABLE STINGER")
	}

	// Veggie Burger
	if strings.Count(titles, "Leafy Meat") < 1 || strings.Count(titles, "Onion") < 1 || attributeVals.vegVal-1 < 1 {
		delete(recipeData, "VEGGIE BURGER")
	}

	// waffles
	if strings.Count(titles, "Butter") < 1 || strings.Count(titles, "Berries") < 1 || attributeCounts.eggCount < 1 {
		delete(recipeData, "WAFFLES")
	}

	// wet goop
	// if everything is false

	// WOBSTER DINNER
	if strings.Count(titles, "Wobster") < 1 || strings.Count(titles, "Butter") < 1 || attributeCounts.meatCount != 0 || attributeCounts.fishCount != 0 || crockPot["Twigs"].NAME != "" {
		delete(recipeData, "WOBSTER DINNER")
	}
	return recipeData
}

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

func blankRec(recipeData map[string]RecipeDetails, attributeCounts AttributeCounts, attributeVals AttributeVals, titles string, crockPot map[string]IngredientDetails) map[string]RecipeDetails {

	// AMBEROSIA
	if strings.Count(titles, "Collected Dust") < 1 {
	}

	// ASPARAGUS SOUP
	if strings.Count(titles, "Asparagus") < 1 || attributeVals.vegVal < 1.5 || attributeCounts.meatCount != 0 || attributeCounts.inedibleCount != 0 {
	}

	// Bacon and Eggs
	if attributeVals.meatVal <= 1 || attributeVals.eggVal <= 1 || attributeVals.vegVal != 0 {
	}

	// BANANA POP
	if strings.Count(titles, "Banana") < 1 || strings.Count(titles, "twigs") < 1 || attributeCounts.meatCount != 0 || attributeCounts.fishCount != 0 {
	}

	// Barnacle Linguine
	if strings.Count(titles, "Barnacles") != 2 || attributeCounts.vegeCount != 2 {
	}

	// Barnacle Nigiri
	if strings.Count(titles, "Barnacles") < 1 || strings.Count(titles, "Kelp Fronds") < 1 || attributeCounts.eggCount < 1 {
	}

	// BARNACLE PITA
	// attributeVals.fishval over 1 can cause stuffed fish head but idk the percentage
	if strings.Count(titles, "Barnacles") < 1 || strings.Count(titles, "Kelp Fronds") < 1 || attributeCounts.eggCount < 1 {
	}

	// BEEFY GREENS
	if strings.Count(titles, "Leafy Meat") < 1 || attributeVals.vegVal < 3 {
	}

	// butter muffin
	if strings.Count(titles, "Butterfly Wings") < 1 || attributeCounts.vegeCount < 1 || crockPot["Mandrake"].NAME != "" || attributeCounts.meatCount != 0 {
	}

	// CALIFORNIA ROLL
	if strings.Count(titles, "Kelp Fronds") < 2 || attributeVals.fishVal < 1 {
	}

	// ceviche
	if strings.Count(titles, "Ice") < 2 || attributeVals.fishVal < 2 || attributeCounts.eggCount != 0 || attributeCounts.inedibleCount != 0 {
	}

	// CREAMY POTATO PURÉE**
	// need to look at potato or roasted potato
	if strings.Count(titles, "Potato") < 2 || strings.Count(titles, "Garlic") < 1 || crockPot["Twigs"].NAME != "" || attributeCounts.meatCount != 0 {
	}

	// dragonpie
	// need to look at dragon fruit or prepared dragon fruit
	if strings.Count(titles, "Dragon Fruit") < 1 || crockPot["Mandrake"].NAME != "" || attributeCounts.meatCount != 0 {
	}

	// FANCY SPIRALLED TUBERS
	// need to look at potato or roasted potato
	if strings.Count(titles, "Potato") < 1 || strings.Count(titles, "Twigs") < 1 || attributeCounts.inedibleCount-1 > 1 || attributeCounts.meatCount != 0 {
	}

	// fish tacos
	// twig value - 1 twig 50% chance of  fish sticks
	// need to look at corn or popcorn
	if attributeVals.fishVal < 0.5 || (strings.Count(titles, "Corn") < 1 || strings.Count(titles, "Popcorn") < 1) {
	}

	// fishsticks
	if attributeVals.fishVal < 0.25 || strings.Count(titles, "Twigs") != 1 || crockPot["Moleworm"].NAME != "" {
	}

	// fist full of jam
	if attributeVals.fruitVal < 0.5 || attributeVals.meatVal != 0 || attributeVals.vegVal != 0 || attributeVals.inedVal != 0 || crockPot["Dragon Fruit"].NAME != "" {
	}

	// flower salad
	if strings.Count(titles, "Cactus Flower") < 1 || attributeVals.vegVal-.5 < 1.5 || attributeVals.meatVal != 0 || attributeVals.fruitVal != 0 || attributeVals.eggVal != 0 || attributeVals.sweetVal != 0 || crockPot["Twigs"].NAME != "" {
	}

	// froggle bunwich
	// need to look at frog legs or cooked frog legs
	// makes kabob if only one stick
	if strings.Count(titles, "Frog Legs") < 1 || attributeCounts.vegeCount < 1 || attributeVals.eggVal != 0 || attributeVals.sweetVal != 0 || crockPot["Mandrake"].NAME != "" {
	}

	// fruit medley
	// twigs is safest anything else 50% of fist full of jam
	if attributeVals.fruitVal < 3 || attributeVals.meatVal != 0 || attributeVals.vegVal != 0 || crockPot["Dragon Fruit"].NAME != "" {
	}

	// guacamole
	// need to look at cactus flesh or stone fruit
	if strings.Count(titles, "Moleworm") < 1 || (strings.Count(titles, "Cactus Flesh") < 1 || strings.Count(titles, "Ripe Stone Fruit") < 1) || attributeVals.fruitVal != 0 {
	}

	// Honey Ham
	if strings.Count(titles, "Honey") < 1 || attributeVals.meatVal <= 1.5 || crockPot["Twigs"].NAME != "" || crockPot["Moleworm"].NAME != "" || crockPot["Mandrake"].NAME != "" || crockPot["Tallbird Egg"].NAME != "" {
	}

	// Honey Nuggets
	if strings.Count(titles, "Honey") < 1 || attributeVals.meatVal > 1.5 || attributeVals.inedVal != 0 {
	}

	// ice cream
	if strings.Count(titles, "Ice") < 1 || attributeCounts.dairyCount < 1 || attributeCounts.sweetenerCount < 1 || attributeVals.meatVal != 0 || attributeVals.vegVal != 0 || attributeVals.eggVal != 0 || crockPot["Twigs"].NAME != "" {
	}

	// jelly salad
	// look for cooked version
	if strings.Count(titles, "Leafy Meat") < 2 || attributeVals.sweetVal < 2 {
	}

	// jellybeans
	if strings.Count(titles, "Royal Jelly") < 1 || attributeVals.monVal != 0 || attributeVals.inedVal != 0 {
	}

	// kabobs
	if attributeCounts.meatCount < 1 || strings.Count(titles, "Twigs") != 1 || crockPot["Moleworm"].NAME != "" || crockPot["Mandrake"].NAME != "" || attributeVals.fishVal != 0 {
	}

	// leafy meatloaf
	if strings.Count(titles, "Leafy Meat") < 2 {
	}

	// LOBSTER BISQUE
	if strings.Count(titles, "Wobster") < 1 || strings.Count(titles, "Ice") < 1 {
	}

	// mandrake soup
	if strings.Count(titles, "Mandrake") < 1 {
	}

	// Meatballs
	if attributeVals.meatVal >= 3 || attributeVals.meatVal == 0 || crockPot["Twigs"].NAME != "" {
	}

	// meaty stew
	if attributeVals.meatVal < 3 || crockPot["Twigs"].NAME != "" || crockPot["Moleworm"].NAME != "" || crockPot["Honey"].NAME != "" || crockPot["Mandrake"].NAME != "" || crockPot["Tallbird Egg"].NAME != "" {
	}

	// melonsicle
	if strings.Count(titles, "Watermelon") < 1 || strings.Count(titles, "Ice") < 1 || strings.Count(titles, "Twigs") < 1 || attributeVals.meatVal != 0 || attributeVals.vegVal != 0 || attributeVals.eggVal != 0 {
	}

	// MILKMADE HAT
	if strings.Count(titles, "Nostrils") < 1 || strings.Count(titles, "Kelp Fronds") < 1 || attributeCounts.dairyCount < 1 {
	}

	// monster lasagna
	if attributeCounts.monsterCount < 2 || crockPot["Twigs"].NAME != "" {
	}

	// Mushy Cake
	if strings.Count(titles, "Moon Shroom") != 1 || strings.Count(titles, "Red Cap") != 1 || strings.Count(titles, "Blue Cap") != 1 || strings.Count(titles, "Green Cap") != 1 {
	}

	// Pierogi
	if attributeCounts.meatCount < 1 || attributeCounts.eggCount < 1 || attributeCounts.vegeCount < 1 || crockPot["Twigs"].NAME != "" || crockPot["Mandrake"].NAME != "" {
	}

	//powdercake
	if strings.Count(titles, "Corn") < 1 || strings.Count(titles, "Honey") < 1 || strings.Count(titles, "Twigs") < 1 {
	}

	//Pumkin cookie
	// 3 honey or comb 50% chance of making taffy
	if strings.Count(titles, "Pumpkin") < 1 || strings.Count(titles, "Honey") <= 1 {
	}

	// RATATOUILLE
	if attributeCounts.vegeCount < 1 || crockPot["Twigs"].NAME != "" || crockPot["Mandrake"].NAME != "" || crockPot["Butterfly Wings"].NAME != "" || crockPot["Dragon Fruit"].NAME != "" {
	}

	// SALSA FRESCA
	// look for cooked versions
	if strings.Count(titles, "Toma Root") < 1 || attributeCounts.meatCount != 0 || attributeCounts.inedibleCount != 0 || attributeCounts.eggCount != 0 {
	}

	// SEAFOOD GUMBO
	if strings.Count(titles, "Eel") < 1 || attributeVals.fishVal <= 2 {
	}

	//  Soothing tea
	// look for honey or comb
	if strings.Count(titles, "Forget-Me-Lots") < 1 || strings.Count(titles, "Honey") < 1 || strings.Count(titles, "Ice") < 1 || attributeCounts.monsterCount != 0 || attributeCounts.meatCount != 0 || attributeCounts.fishCount != 0 || attributeCounts.eggCount != 0 || attributeCounts.inedibleCount != 0 || attributeCounts.dairyCount != 0 {
	}

	//  Spicy Chili
	if attributeCounts.meatCount != 2 || attributeCounts.vegeCount != 2 || attributeVals.meatVal < 1.5 || attributeVals.vegVal < 1.5 {
	}

	// STUFFED EGGPLANT
	if strings.Count(titles, "Eggplant") < 1 || attributeCounts.vegeCount < 1 {
	}

	// stuffed fish heads
	if strings.Count(titles, "Barnacles") < 1 || attributeVals.fishVal-0.5 < 1 {
	}

	// STUFFED PEPPER POPPERS
	if strings.Count(titles, "Pepper") < 1 || attributeVals.meatVal > 1.5 || attributeCounts.meatCount == 0 || crockPot["Twigs"].NAME != "" {
	}

	// SURF 'N' TURF
	if attributeVals.meatVal < 2.5 || attributeVals.fishVal < 1.5 || crockPot["Ice"].NAME != "" {
	}

	// Taffy
	if attributeVals.sweetVal < 3 || attributeCounts.meatCount != 0 {
	}

	// trail mix
	// required berries must be uncooked
	if strings.Count(titles, "Roasted Birchnut") < 1 || strings.Count(titles, "Berries") < 1 || strings.Count(titles, "Roasted") > 3 || attributeCounts.fruitCount-1 < 1 || attributeCounts.meatCount != 0 || attributeCounts.fishCount != 0 || attributeCounts.eggCount != 0 || attributeCounts.vegeCount != 0 || attributeCounts.dairyCount != 0 {
	}

	// TURKEY DINNER
	// checkingg twice for veg or fruit
	if strings.Count(titles, "Drumstick") < 2 || attributeVals.meatVal-1 < 0.25 || (attributeVals.fruitVal < .5 || attributeVals.vegVal < .5) {
	}

	// unagi
	if strings.Count(titles, "Eel") < 1 || (strings.Count(titles, "Lichen") < 1 || strings.Count(titles, "Kelp Fronds") < 1) {
	}

	// VEGETABLE STINGER
	if (strings.Count(titles, "Toma Root") < 1 || strings.Count(titles, "Asparagus") < 1) || strings.Count(titles, "Ice") < 1 || attributeVals.vegVal-1 < 1.5 {
	}

	// Veggie Burger
	if strings.Count(titles, "Leafy Meat") < 1 || strings.Count(titles, "Onion") < 1 || attributeVals.vegVal-1 < 1 {
	}

	// waffles
	if strings.Count(titles, "Butter") < 1 || strings.Count(titles, "Berries") < 1 || attributeCounts.eggCount < 1 {
	}

	// wet goop
	// if everything is false

	// WOBSTER DINNER
	if strings.Count(titles, "Wobster") < 1 || strings.Count(titles, "Butter") < 1 || attributeCounts.meatCount != 0 || attributeCounts.fishCount != 0 || crockPot["Twigs"].NAME != "" {
	}
	return recipeData
}

func testerRec(recipeData map[string]RecipeDetails, attributeCounts AttributeCounts, attributeVals AttributeVals, titles string, crockPot map[string]IngredientDetails) map[string]RecipeDetails {

	// Bacon and Eggs
	if attributeVals.meatVal <= 1 || attributeVals.eggVal <= 1 || attributeVals.vegVal != 0 {
	}

	// Barnacle Nigiri
	if strings.Count(titles, "Barnacles") < 1 || strings.Count(titles, "Kelp Fronds") < 1 || attributeCounts.eggCount < 1 {
	}

	// Pierogi
	if attributeCounts.meatCount < 1 || attributeCounts.eggCount < 1 || attributeCounts.vegeCount < 1 || crockPot["Twigs"].NAME != "" || crockPot["Mandrake"].NAME != "" {
	}

	// WOBSTER DINNER
	if strings.Count(titles, "Wobster") < 1 || strings.Count(titles, "Butter") < 1 || attributeCounts.meatCount != 0 || attributeCounts.fishCount != 0 || crockPot["Twigs"].NAME != "" {
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

func main() {

	// baconAndEggs := []string{}
	// baconAndEggs = append(baconAndEggs, "poop")
	// baconAndEggs = append(baconAndEggs, "poopy")
	// baconAndEggs = append(baconAndEggs, "pooper")
	// for _, str := range baconAndEggs {
	// 	fmt.Println(str)
	// }

	// "amberosia":            0,
	// "asparagusSoup":        0,
	// "baconAndEggs":         0,
	// "bananaPop":            0,
	// "barnacleLinguine":     0,
	// "barnacleNigiri":       0,
	// "barnaclePita":         0,
	// "beefyGreens":          0,
	// "butterMuffin":         0,
	// "californiaRoll":       0,
	// "ceviche":              0,
	// "creamyPotatoPurée":    0,
	// "dragonpie":            0,
	// "fancySpiralledTubers": 0,
	// "fishTacos":            0,
	// "fishsticks":           0,
	// "fistFullOfJam":        0,
	// "flowerSalad":          0,
	// "froggleBunwich":       0,
	// "fruitMedley":          0,
	// "guacamole":            0,
	// "honeyHam":             0,
	// "honeyNuggets":         0,
	// "iceCream":             0,
	// "jellySalad":           0,
	// "jellybeans":           0,
	// "kabobs":               0,
	// "leafyMeatloaf":        0,
	// "lobsterBisque":        0,
	// "mandrakeSoup":         0,
	// "meatballs":            0,
	// "meatyStew":            0,
	// "melonsicle":           0,
	// "milkmadeHat":          0,
	// "monsterLasagna":       0,
	// "mushyCake":            0,
	// "pierogi":              0,
	// "powdercake":           0,
	// "pumpkinCookie":        0,
	// "ratatouille":          0,
	// "salsaFresca":          0,
	// "seafoodGumbo":         0,
	// "soothingTea":          0,
	// "spicyChili":           0,
	// "stuffedEggplant":      0,
	// "stuffedFishHeads":     0,
	// "stuffedPepperPoppers": 0,
	// "surfNturf":            0,
	// "taffy":                0,
	// "trailMix":             0,
	// "turkeyDinner":         0,
	// "unagi":                0,
	// "vegetableStinger":     0,
	// "veggieBurger":         0,
	// "waffles":              0,
	// "wetGoop":              0,
	// "wobsterDinner":        0,

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
	recipeData := map[string]RecipeDetails{}
	err = json.Unmarshal(byteValue2, &recipeData)
	if err != nil {
		fmt.Println(err)
	}

	possibleRecipeData := map[string]RecipeDetails{}

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
	i1 := "Egg"
	i2 := "Barnacles"
	i3 := "Egg"
	i4 := "Jerky"
	// Jerky
	// Egg
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

	// i1 := "Egg"
	// i2 := "Egg"
	// i3 := "Jerky"
	// i4 := "Jerky"
	// crockPot := map[string]IngredientDetails{}
	// crockPot[i1] = ingredientsData[i1]
	// crockPot[i2] = ingredientsData[i2]
	// crockPot[i3] = ingredientsData[i3]
	// crockPot[i4] = ingredientsData[i4]

	// titles is sting of ingredients
	// titles := "[" + i1 + ", " + i2 + ", " + i3 + ", " + i4 + "]"
	// Look through each ingredient and sets values
	// count Ingredient Valus

	attributeVals := masterVals(crockPot)

	// Count Ingredient types
	attributeCounts := masterCounts(crockPot)

	// if i4 != "" { star hurrr
	// function that deletes false recipes
	// deleteRecipes(recipeData, attributeCounts, attributeVals, titles, crockPot) ------------------------------------------------------

	// bacon eggs
	if attributeVals.meatVal > 1 && attributeVals.eggVal > 1 && attributeVals.vegVal == 0 {
		fmt.Println("ok bacon eggs")
	}

	// Pierogi
	if attributeCounts.meatCount >= 1 && attributeCounts.eggCount >= 1 && attributeCounts.vegeCount >= 1 && stringNotInSlice("Twigs", crockPot) && stringNotInSlice("Mandrake", crockPot) {
		fmt.Println("ok pierogi")
	}
	// // -----------------------------------------------------------------------------------------------------------------------------------

	// Output goes here

	//  this output only shows the vals that are greater than 0
	outVal := map[string]float64{"Meat:": attributeVals.meatVal, "Fish:": attributeVals.fishVal, "Egg:": attributeVals.eggVal, "Fruit:": attributeVals.fruitVal, "Vegetable:": attributeVals.vegVal, "Sweetener:": attributeVals.sweetVal, "Monster:": attributeVals.monVal, "Dairy:": attributeVals.dairyVal, "Bug:": attributeVals.bugVal, "Inedible:": attributeVals.inedVal, "Misc:": attributeVals.miscVal}
	for inName, existVal := range outVal {
		if existVal > 0 {
			fmt.Println(inName, "val", existVal)
		}
		// fmt.Println(inName, existVal)
	}

	//  this output only shows the counts that are greater than 0
	outCount := map[string]float64{"Meat:": attributeCounts.meatCount, "Fish:": attributeCounts.fishCount, "Egg:": attributeCounts.eggCount, "Fruit:": attributeCounts.fruitCount, "Vegetable:": attributeCounts.vegeCount, "Sweetener:": attributeCounts.sweetenerCount, "Monster:": attributeCounts.monsterCount, "Dairy:": attributeCounts.dairyCount, "Bug:": attributeCounts.bugCount, "Inedible:": attributeCounts.inedibleCount, "Misc:": attributeCounts.miscCount}
	for inName, existcount := range outCount {
		if existcount > 0 {
			fmt.Println(inName, "count", existcount)
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
	// orgValidRec(recipeData)
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
