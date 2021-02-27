package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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

func main() {

	// version := DST
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

	// Add Ingredients to the pot
	i1 := "Jerky"
	i2 := "Jerky"
	i3 := "Egg"
	i4 := ""
	crockPot := map[string]IngredientDetails{}
	crockPot[i1] = ingredientsData[i1]
	crockPot[i2] = ingredientsData[i2]
	crockPot[i3] = ingredientsData[i3]
	crockPot[i4] = ingredientsData[i4]

	// Initialize the baseline values
	titles := "[" + i1 + ", " + i2 + ", " + i3 + ", " + i4 + "]"
	meatVal := crockPot[i1].MEAT + crockPot[i2].MEAT + crockPot[i3].MEAT + crockPot[i4].MEAT
	fishVal := crockPot[i1].FISH + crockPot[i2].FISH + crockPot[i3].FISH + crockPot[i4].FISH
	eggVal := crockPot[i1].EGG + crockPot[i2].EGG + crockPot[i3].EGG + crockPot[i4].EGG
	fruitVal := crockPot[i1].FRUIT + crockPot[i2].FRUIT + crockPot[i3].FRUIT + crockPot[i4].FRUIT
	vegVal := crockPot[i1].VEGETABLE + crockPot[i2].VEGETABLE + crockPot[i3].VEGETABLE + crockPot[i4].VEGETABLE
	sweetVal := crockPot[i1].SWEETENER + crockPot[i2].SWEETENER + crockPot[i3].SWEETENER + crockPot[i4].SWEETENER
	monVal := crockPot[i1].MONSTER + crockPot[i2].MONSTER + crockPot[i3].MONSTER + crockPot[i4].MONSTER
	dairyVal := crockPot[i1].DAIRY + crockPot[i2].DAIRY + crockPot[i3].DAIRY + crockPot[i4].DAIRY
	bugVal := crockPot[i1].BUG + crockPot[i2].BUG + crockPot[i3].BUG + crockPot[i4].BUG
	inedVal := crockPot[i1].INEDIBLE + crockPot[i2].INEDIBLE + crockPot[i3].INEDIBLE + crockPot[i4].INEDIBLE
	miscVal := crockPot[i1].MISC + crockPot[i2].MISC + crockPot[i3].MISC + crockPot[i4].MISC
	healVal := crockPot[i1].HEALTH + crockPot[i2].HEALTH + crockPot[i3].HEALTH + crockPot[i4].HEALTH
	hungVal := crockPot[i1].HUNGER + crockPot[i2].HUNGER + crockPot[i3].HUNGER + crockPot[i4].HUNGER
	saniVal := crockPot[i1].SANITY + crockPot[i2].SANITY + crockPot[i3].SANITY + crockPot[i4].SANITY

	// counts type of ingredients
	// count Meat
	crockSlotsMea := []float64{crockPot[i1].MEAT, crockPot[i2].MEAT, crockPot[i3].MEAT, crockPot[i4].MEAT}
	meatCount := 0 // initializing an int outside of the loop so we can see it inside and outside of the "for"
	for _, v := range crockSlotsMea {
		if v > 0 {
			meatCount++
		}
	}

	// count Fish
	crockSlotsFis := []float64{crockPot[i1].FISH, crockPot[i2].FISH, crockPot[i3].FISH, crockPot[i4].FISH}
	fishCount := 0 // initializing an int outside of the loop so we can see it inside and outside of the "for"
	for _, v := range crockSlotsFis {
		if v > 0 {
			fishCount++
		}
	}
	// count egg
	crockSlotsEgg := []float64{crockPot[i1].EGG, crockPot[i2].EGG, crockPot[i3].EGG, crockPot[i4].EGG}
	eggCount := 0 // initializing an int outside of the loop so we can see it inside and outside of the "for"
	for _, v := range crockSlotsEgg {
		if v > 0 {
			eggCount++
		}
	}
	// count fruit
	crockSlotsFru := []float64{crockPot[i1].FRUIT, crockPot[i2].FRUIT, crockPot[i3].FRUIT, crockPot[i4].FRUIT}
	fruitCount := 0 // initializing an int outside of the loop so we can see it inside and outside of the "for"
	for _, v := range crockSlotsFru {
		if v > 0 {
			fruitCount++
		}
	}
	// count vegetables
	crockSlotsVeg := []float64{crockPot[i1].VEGETABLE, crockPot[i2].VEGETABLE, crockPot[i3].VEGETABLE, crockPot[i4].VEGETABLE}
	vegeCount := 0 // initializing an int outside of the loop so we can see it inside and outside of the "for"
	for _, v := range crockSlotsVeg {
		if v > 0 {
			vegeCount++
		}
	}
	// count sweetener
	crockSlotsSwe := []float64{crockPot[i1].SWEETENER, crockPot[i2].SWEETENER, crockPot[i3].SWEETENER, crockPot[i4].SWEETENER}
	sweetenerCount := 0 // initializing an int outside of the loop so we can see it inside and outside of the "for"
	for _, v := range crockSlotsSwe {
		if v > 0 {
			sweetenerCount++
		}
	}
	// count monster food
	crockSlotsMon := []float64{crockPot[i1].MONSTER, crockPot[i2].MONSTER, crockPot[i3].MONSTER, crockPot[i4].MONSTER}
	monsterCount := 0 // initializing an int outside of the loop so we can see it inside and outside of the "for"
	for _, v := range crockSlotsMon {
		if v > 0 {
			monsterCount++
		}
	}
	// count monster food
	crockSlotsDai := []float64{crockPot[i1].DAIRY, crockPot[i2].DAIRY, crockPot[i3].DAIRY, crockPot[i4].DAIRY}
	dairyCount := 0 // initializing an int outside of the loop so we can see it inside and outside of the "for"
	for _, v := range crockSlotsDai {
		if v > 0 {
			dairyCount++
		}
	}
	// count bug
	crockSlotsBug := []float64{crockPot[i1].BUG, crockPot[i2].BUG, crockPot[i3].BUG, crockPot[i4].BUG}
	bugCount := 0 // initializing an int outside of the loop so we can see it inside and outside of the "for"
	for _, v := range crockSlotsBug {
		if v > 0 {
			bugCount++
		}
	}
	// count inedable
	crockSlotsIne := []float64{crockPot[i1].INEDIBLE, crockPot[i2].INEDIBLE, crockPot[i3].INEDIBLE, crockPot[i4].INEDIBLE}
	inedibleCount := 0 // initializing an int outside of the loop so we can see it inside and outside of the "for"
	for _, v := range crockSlotsIne {
		if v > 0 {
			inedibleCount++
		}
	}
	// count misc (moleworm, butterfly, etc.)
	crockSlotsMis := []float64{crockPot[i1].MISC, crockPot[i2].MISC, crockPot[i3].MISC, crockPot[i4].MISC}
	miscCount := 0 // initializing an int outside of the loop so we can see it inside and outside of the "for"
	for _, v := range crockSlotsMis {
		if v > 0 {
			miscCount++
		}
	}

	// titles := []string{}

	// Look through each ingredient and sets values
	// for _, ingre := range crockPot {

	// titles = append(titles, ingre.NAME)
	// meatVal = meatVal + ingre.MEAT
	// fishVal += ingre.FISH
	// eggVal += ingre.EGG
	// fruitVal += ingre.FRUIT
	// vegVal += ingre.VEGETABLE
	// sweetVal += ingre.SWEETENER
	// monVal += ingre.MONSTER
	// dairyVal += ingre.DAIRY
	// bugVal += ingre.BUG
	// inedVal += ingre.INEDIBLE
	// miscVal += ingre.MISC
	// }

	//	Create an array with all possible recipes

	//	create an negative If conditional for each recipe

	// AMBEROSIA
	if strings.Count(titles, "Collected Dust") < 1 {
		delete(recipeData, "AMBEROSIA")
	}

	// ASPARAGUS SOUP
	if strings.Count(titles, "Asparagus") < 1 || vegVal < 1.5 || meatCount != 0 || inedibleCount != 0 {
		delete(recipeData, "ASPARAGUS SOUP")
	}

	// Bacon and Eggs
	if meatVal <= 1 || eggVal <= 1 || vegVal != 0 {
		delete(recipeData, "BACON AND EGGS")
	}

	// BANANA POP
	if strings.Count(titles, "Banana") < 1 || strings.Count(titles, "twigs") < 1 || meatCount != 0 || fishCount != 0 {
		delete(recipeData, "BANANA POP")
	}

	// Barnacle Linguine
	if strings.Count(titles, "Barnacles") != 2 || vegeCount != 2 {
		delete(recipeData, "BARNACLE LINGUINE")
	}

	// Barnacle Nigiri
	if strings.Count(titles, "Barnacles") < 1 || strings.Count(titles, "Kelp Fronds") < 1 || eggCount < 1 {
		delete(recipeData, "BARNACLE NIGIRI")
	}

	// BARNACLE PITA
	// fishval over 1 can cause stuffed fish head but idk the percentage
	if strings.Count(titles, "Barnacles") < 1 || strings.Count(titles, "Kelp Fronds") < 1 || eggCount < 1 {
		delete(recipeData, "BARNACLE PITA")
	}

	// BEEFY GREENS
	if strings.Count(titles, "Leafy Meat") < 1 || vegVal < 3 {
		delete(recipeData, "BEEFY GREENS")
	}

	// butter muffin
	if strings.Count(titles, "Butterfly Wings") < 1 || vegeCount < 1 || crockPot["Mandrake"].NAME != "" || meatCount != 0 {
		delete(recipeData, "BUTTER MUFFIN")
	}

	// CALIFORNIA ROLL
	if strings.Count(titles, "Kelp Fronds") < 2 || fishVal < 1 {
		delete(recipeData, "CALIFORNIA ROLL")
	}

	// ceviche
	if strings.Count(titles, "Ice") < 2 || fishVal < 2 || eggCount != 0 || inedibleCount != 0 {
		delete(recipeData, "CEVICHE")
	}

	// CREAMY POTATO PURÉE**
	// need to look at potato or roasted potato
	if strings.Count(titles, "Potato") < 2 || strings.Count(titles, "Garlic") < 1 || crockPot["Twigs"].NAME != "" || meatCount != 0 {
		delete(recipeData, "CREAMY POTATO PURÉE")
	}

	// dragonpie
	// need to look at dragon fruit or prepared dragon fruit
	if strings.Count(titles, "Dragon Fruit") < 1 || crockPot["Mandrake"].NAME != "" || meatCount != 0 {
		delete(recipeData, "DRAGONPIE")
	}

	// FANCY SPIRALLED TUBERS
	// need to look at potato or roasted potato
	if strings.Count(titles, "Potato") < 1 || strings.Count(titles, "Twigs") < 1 || inedibleCount-1 > 1 || meatCount != 0 {
		delete(recipeData, "FANCY SPIRALLED TUBERS")
	}

	// fish tacos
	// twig value - 1 twig 50% chance of  fish sticks
	// need to look at corn or popcorn
	if fishVal < 0.5 || (strings.Count(titles, "Corn") < 1 || strings.Count(titles, "Popcorn") < 1) {
		delete(recipeData, "FISH TACOS")
	}

	// fishsticks
	if fishVal < 0.25 || strings.Count(titles, "Twigs") != 1 || crockPot["Moleworm"].NAME != "" {
		delete(recipeData, "FISHSTICKS")
	}

	// fist full of jam
	if fruitVal < 0.5 || meatVal != 0 || vegVal != 0 || inedVal != 0 || crockPot["Dragon Fruit"].NAME != "" {
		delete(recipeData, "FIST FULL OF JAM")
	}

	// flower salad
	if strings.Count(titles, "Cactus Flower") < 1 || vegVal-.5 < 1.5 || meatVal != 0 || fruitVal != 0 || eggVal != 0 || sweetVal != 0 || crockPot["Twigs"].NAME != "" {
		delete(recipeData, "FLOWER SALAD")
	}

	// froggle bunwich
	// need to look at frog legs or cooked frog legs
	// makes kabob if only one stick
	if strings.Count(titles, "Frog Legs") < 1 || vegeCount < 1 || eggVal != 0 || sweetVal != 0 || crockPot["Mandrake"].NAME != "" {
		delete(recipeData, "FROGGLE BUNWICH")
	}

	// fruit medley
	// twigs is safest anything else 50% of fist full of jam
	if fruitVal < 3 || meatVal != 0 || vegVal != 0 || crockPot["Dragon Fruit"].NAME != "" {
		delete(recipeData, "FRUIT MEDLEY")
	}

	// guacamole
	// need to look at cactus flesh or stone fruit
	if strings.Count(titles, "Moleworm") < 1 || (strings.Count(titles, "Cactus Flesh") < 1 || strings.Count(titles, "Ripe Stone Fruit") < 1) || fruitVal != 0 {
		delete(recipeData, "GUACAMOLE")
	}

	// Honey Ham
	if strings.Count(titles, "Honey") < 1 || meatVal <= 1.5 || crockPot["Twigs"].NAME != "" || crockPot["Moleworm"].NAME != "" || crockPot["Mandrake"].NAME != "" || crockPot["Tallbird Egg"].NAME != "" {
		delete(recipeData, "HONEY HAM")
	}

	// Honey Nuggets
	if strings.Count(titles, "Honey") < 1 || meatVal > 1.5 || inedVal != 0 {
		delete(recipeData, "HONEY NUGGETS")
	}

	// ice cream
	if strings.Count(titles, "Ice") < 1 || dairyCount < 1 || sweetenerCount < 1 || meatVal != 0 || vegVal != 0 || eggVal != 0 || crockPot["Twigs"].NAME != "" {
		delete(recipeData, "ICE CREAM")
	}

	// jelly salad
	// look for cooked version
	if strings.Count(titles, "Leafy Meat") < 2 || sweetVal < 2 {
		delete(recipeData, "JELLY SALAD")
	}

	// jellybeans
	if strings.Count(titles, "Royal Jelly") < 1 || monVal != 0 || inedVal != 0 {
		delete(recipeData, "JELLYBEANS")
	}

	// kabobs
	if meatCount < 1 || strings.Count(titles, "Twigs") != 1 || crockPot["Moleworm"].NAME != "" || crockPot["Mandrake"].NAME != "" || fishVal != 0 {
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
	if meatVal >= 3 || meatVal == 0 || crockPot["Twigs"].NAME != "" {
		delete(recipeData, "MEATBALLS")
	}

	// meaty stew
	if meatVal < 3 || crockPot["Twigs"].NAME != "" || crockPot["Moleworm"].NAME != "" || crockPot["Honey"].NAME != "" || crockPot["Mandrake"].NAME != "" || crockPot["Tallbird Egg"].NAME != "" {
		delete(recipeData, "MEATY STEW")
	}

	// melonsicle
	if strings.Count(titles, "Watermelon") < 1 || strings.Count(titles, "Ice") < 1 || strings.Count(titles, "Twigs") < 1 || meatVal != 0 || vegVal != 0 || eggVal != 0 {
		delete(recipeData, "MELONSICLE")
	}

	// MILKMADE HAT
	if strings.Count(titles, "Nostrils") < 1 || strings.Count(titles, "Kelp Fronds") < 1 || dairyCount < 1 {
		delete(recipeData, "MILKMADE HAT")
	}

	// monster lasagna
	if monsterCount < 2 || crockPot["Twigs"].NAME != "" {
		delete(recipeData, "MONSTER LASAGNA")
	}

	// Mushy Cake
	if strings.Count(titles, "Moon Shroom") != 1 || strings.Count(titles, "Red Cap") != 1 || strings.Count(titles, "Blue Cap") != 1 || strings.Count(titles, "Green Cap") != 1 {
		delete(recipeData, "MUSHY CAKE")
	}

	// Pierogi
	if meatCount < 1 || eggCount < 1 || vegeCount < 1 || crockPot["Twigs"].NAME != "" || crockPot["Mandrake"].NAME != "" {
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
	if vegeCount < 1 || crockPot["Twigs"].NAME != "" || crockPot["Mandrake"].NAME != "" || crockPot["Butterfly Wings"].NAME != "" || crockPot["Dragon Fruit"].NAME != "" {
		delete(recipeData, "RATATOUILLE")
	}

	// SALSA FRESCA
	// look for cooked versions
	if strings.Count(titles, "Toma Root") < 1 || meatCount != 0 || inedibleCount != 0 || eggCount != 0 {
		delete(recipeData, "SALSA FRESCA")
	}

	// SEAFOOD GUMBO
	if strings.Count(titles, "Eel") < 1 || fishVal <= 2 {
		delete(recipeData, "SEAFOOD GUMBO")
	}

	//  Soothing tea
	// look for honey or comb
	if strings.Count(titles, "Forget-Me-Lots") < 1 || strings.Count(titles, "Honey") < 1 || strings.Count(titles, "Ice") < 1 || monsterCount != 0 || meatCount != 0 || fishCount != 0 || eggCount != 0 || inedibleCount != 0 || dairyCount != 0 {
		delete(recipeData, "SOOTHING TEA")
	}

	//  Spicy Chili
	if meatCount != 2 || vegeCount != 2 || meatVal < 1.5 || vegVal < 1.5 {
		delete(recipeData, "SPICY CHILI")
	}

	// STUFFED EGGPLANT
	if strings.Count(titles, "Eggplant") < 1 || vegeCount < 1 {
		delete(recipeData, "STUFFED EGGPLANT")
	}

	// stuffed fish heads
	if strings.Count(titles, "Barnacles") < 1 || fishVal-0.5 < 1 {
		delete(recipeData, "STUFFED FISH HEADS")
	}

	// STUFFED PEPPER POPPERS
	if strings.Count(titles, "Pepper") < 1 || meatVal > 1.5 || meatCount == 0 || crockPot["Twigs"].NAME != "" {
		delete(recipeData, "STUFFED PEPPER POPPERS")
	}

	// SURF 'N' TURF
	if meatVal < 2.5 || fishVal < 1.5 || crockPot["Ice"].NAME != "" {
		delete(recipeData, "SURF 'N' TURF")
	}

	// Taffy
	if sweetVal < 3 || meatCount != 0 {
		delete(recipeData, "TAFFY")
	}

	// trail mix
	// required berries must be uncooked
	if strings.Count(titles, "Roasted Birchnut") < 1 || strings.Count(titles, "Berries") < 1 || strings.Count(titles, "Roasted") > 3 || fruitCount-1 < 1 || meatCount != 0 || fishCount != 0 || eggCount != 0 || vegeCount != 0 || dairyCount != 0 {
		delete(recipeData, "TRAIL MIX")
	}

	// TURKEY DINNER
	// checkingg twice for veg or fruit
	if strings.Count(titles, "Drumstick") < 2 || meatVal-1 < 0.25 || (fruitVal < .5 || vegVal < .5) {
		delete(recipeData, "TURKEY DINNER")
	}

	// unagi
	if strings.Count(titles, "Eel") < 1 || (strings.Count(titles, "Lichen") < 1 || strings.Count(titles, "Kelp Fronds") < 1) {
		delete(recipeData, "UNAGI")
	}

	// VEGETABLE STINGER
	if (strings.Count(titles, "Toma Root") < 1 || strings.Count(titles, "Asparagus") < 1) || strings.Count(titles, "Ice") < 1 || vegVal-1 < 1.5 {
		delete(recipeData, "VEGETABLE STINGER")
	}

	// Veggie Burger
	if strings.Count(titles, "Leafy Meat") < 1 || strings.Count(titles, "Onion") < 1 || vegVal-1 < 1 {
		delete(recipeData, "VEGGIE BURGER")
	}

	// waffles
	if strings.Count(titles, "Butter") < 1 || strings.Count(titles, "Berries") < 1 || eggCount < 1 {
		delete(recipeData, "WAFFLES")
	}

	// wet goop
	// if everything is false

	// WOBSTER DINNER
	if strings.Count(titles, "Wobster") < 1 || strings.Count(titles, "Butter") < 1 || meatCount != 0 || fishCount != 0 || crockPot["Twigs"].NAME != "" {
		delete(recipeData, "WOBSTER DINNER")
	}

	// Output goes here
	// title := fmt.Sprintf("%s + %s + %s + %s:", titles[0], titles[1], titles[2], titles[3])

	fmt.Println(titles)

	//  this output only shows the vals that are greater than 0
	outVal := map[string]float64{"Meat:": meatVal, "Fish:": fishVal, "Egg:": eggVal, "Fruit:": fruitVal, "Vegetable:": vegVal, "Sweetener:": sweetVal, "Monster:": monVal, "Dairy:": dairyVal, "Bug:": bugVal, "Inedible:": inedVal, "Misc:": miscVal}
	for inName, existVal := range outVal {
		if existVal > 0 {
			fmt.Println(inName, existVal)
		}
	}
	// if eaten raw
	// fmt.Println("\n")
	fmt.Println("Ingredients")
	fmt.Println("Total Health", healVal)
	fmt.Println("Total Hunger", hungVal)
	fmt.Println("Total Sanity", saniVal)
	// prints raw exp
	fmt.Println(i1, "expires in", ingredientsData[i1].EXPIRE, "days")
	fmt.Println(i2, "expires in", ingredientsData[i2].EXPIRE, "days")
	fmt.Println(i3, "expires in", ingredientsData[i3].EXPIRE, "days")
	fmt.Println(i4, "expires in", ingredientsData[i4].EXPIRE, "days")

	// printed not organized
	// for t, existRec := range recipeData {
	// 	fmt.Println("\n")
	// 	fmt.Println(t, ":")
	// 	fmt.Println("Health:", existRec.HEALTH)
	// 	fmt.Println("Hunger:", existRec.HUNGER)
	// 	fmt.Println("Sanity:", existRec.SANITY)
	// 	fmt.Println("Ingredients:", existRec.INGREDIENTS)
	// 	fmt.Println("Exclude:", existRec.EXCLUDE)
	// 	fmt.Println("Notes:", existRec.NOTES)
	// 	fmt.Println("Expires:", existRec.EXPIRES)
	// 	fmt.Println("Cooktime:", existRec.COOKTIME)
	// 	fmt.Println("Priority:", existRec.PRIORITY)
	// }

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
				fmt.Println("\n")
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

	// *********************finding alternative values for cooked/murdered items*************************
	// is json, if ingredients are added, murder must be first, then  cook
	/*cdex := strings.Index(ingredientsData[i1].NOTES, "cook:")
	mdex := strings.Index(ingredientsData[i1].NOTES, "murder:")

	if mdex > -1 {
		fmt.Println(string(ingredientsData[i1].NOTES[mdex+8 : cdex-2]))
	}
	if cdex > -1 {
		fmt.Println(string(ingredientsData[i1].NOTES[cdex+6:]))
	}
	fmt.Println(mdex)
	fmt.Println(cdex)*/
	// *********************************************************************************************************************************************************
	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	for trueIn, trueInName := range ingredientsData {
		crockPot := map[string]IngredientDetails{}
		i1 := "Jerky"
		i2 := "Jerky"
		i3 := "Egg"
		i4 := trueIn
		crockPot[i1] = ingredientsData[i1]
		crockPot[i2] = ingredientsData[i2]
		crockPot[i3] = ingredientsData[i3]
		crockPot[i4] = ingredientsData[i4]

		// Initialize the baseline values
		titles := "[" + i1 + ", " + i2 + ", " + i3 + ", " + i4 + "]"
		meatVal := crockPot[i1].MEAT + crockPot[i2].MEAT + crockPot[i3].MEAT + crockPot[i4].MEAT
		// fishVal := crockPot[i1].FISH + crockPot[i2].FISH + crockPot[i3].FISH + crockPot[i4].FISH
		eggVal := crockPot[i1].EGG + crockPot[i2].EGG + crockPot[i3].EGG + crockPot[i4].EGG
		// fruitVal := crockPot[i1].FRUIT + crockPot[i2].FRUIT + crockPot[i3].FRUIT + crockPot[i4].FRUIT
		vegVal := crockPot[i1].VEGETABLE + crockPot[i2].VEGETABLE + crockPot[i3].VEGETABLE + crockPot[i4].VEGETABLE
		// sweetVal := crockPot[i1].SWEETENER + crockPot[i2].SWEETENER + crockPot[i3].SWEETENER + crockPot[i4].SWEETENER
		// monVal := crockPot[i1].MONSTER + crockPot[i2].MONSTER + crockPot[i3].MONSTER + crockPot[i4].MONSTER
		// dairyVal := crockPot[i1].DAIRY + crockPot[i2].DAIRY + crockPot[i3].DAIRY + crockPot[i4].DAIRY
		// bugVal := crockPot[i1].BUG + crockPot[i2].BUG + crockPot[i3].BUG + crockPot[i4].BUG
		// inedVal := crockPot[i1].INEDIBLE + crockPot[i2].INEDIBLE + crockPot[i3].INEDIBLE + crockPot[i4].INEDIBLE
		// miscVal := crockPot[i1].MISC + crockPot[i2].MISC + crockPot[i3].MISC + crockPot[i4].MISC
		// healVal := crockPot[i1].HEALTH + crockPot[i2].HEALTH + crockPot[i3].HEALTH + crockPot[i4].HEALTH
		// hungVal := crockPot[i1].HUNGER + crockPot[i2].HUNGER + crockPot[i3].HUNGER + crockPot[i4].HUNGER
		// saniVal := crockPot[i1].SANITY + crockPot[i2].SANITY + crockPot[i3].SANITY + crockPot[i4].SANITY

		// counts type of ingredients
		// count Meat
		crockSlotsMea := []float64{crockPot[i1].MEAT, crockPot[i2].MEAT, crockPot[i3].MEAT, crockPot[i4].MEAT}
		meatCount := 0 // initializing an int outside of the loop so we can see it inside and outside of the "for"
		for _, v := range crockSlotsMea {
			if v > 0 {
				meatCount++
			}
		}

		// count Fish
		crockSlotsFis := []float64{crockPot[i1].FISH, crockPot[i2].FISH, crockPot[i3].FISH, crockPot[i4].FISH}
		fishCount := 0 // initializing an int outside of the loop so we can see it inside and outside of the "for"
		for _, v := range crockSlotsFis {
			if v > 0 {
				fishCount++
			}
		}
		// count egg
		crockSlotsEgg := []float64{crockPot[i1].EGG, crockPot[i2].EGG, crockPot[i3].EGG, crockPot[i4].EGG}
		eggCount := 0 // initializing an int outside of the loop so we can see it inside and outside of the "for"
		for _, v := range crockSlotsEgg {
			if v > 0 {
				eggCount++
			}
		}
		// count fruit
		crockSlotsFru := []float64{crockPot[i1].FRUIT, crockPot[i2].FRUIT, crockPot[i3].FRUIT, crockPot[i4].FRUIT}
		fruitCount := 0 // initializing an int outside of the loop so we can see it inside and outside of the "for"
		for _, v := range crockSlotsFru {
			if v > 0 {
				fruitCount++
			}
		}
		// count vegetables
		crockSlotsVeg := []float64{crockPot[i1].VEGETABLE, crockPot[i2].VEGETABLE, crockPot[i3].VEGETABLE, crockPot[i4].VEGETABLE}
		vegeCount := 0 // initializing an int outside of the loop so we can see it inside and outside of the "for"
		for _, v := range crockSlotsVeg {
			if v > 0 {
				vegeCount++
			}
		}
		// count sweetener
		crockSlotsSwe := []float64{crockPot[i1].SWEETENER, crockPot[i2].SWEETENER, crockPot[i3].SWEETENER, crockPot[i4].SWEETENER}
		sweetenerCount := 0 // initializing an int outside of the loop so we can see it inside and outside of the "for"
		for _, v := range crockSlotsSwe {
			if v > 0 {
				sweetenerCount++
			}
		}
		// count monster food
		crockSlotsMon := []float64{crockPot[i1].MONSTER, crockPot[i2].MONSTER, crockPot[i3].MONSTER, crockPot[i4].MONSTER}
		monsterCount := 0 // initializing an int outside of the loop so we can see it inside and outside of the "for"
		for _, v := range crockSlotsMon {
			if v > 0 {
				monsterCount++
			}
		}
		// count monster food
		crockSlotsDai := []float64{crockPot[i1].DAIRY, crockPot[i2].DAIRY, crockPot[i3].DAIRY, crockPot[i4].DAIRY}
		dairyCount := 0 // initializing an int outside of the loop so we can see it inside and outside of the "for"
		for _, v := range crockSlotsDai {
			if v > 0 {
				dairyCount++
			}
		}
		// count bug
		crockSlotsBug := []float64{crockPot[i1].BUG, crockPot[i2].BUG, crockPot[i3].BUG, crockPot[i4].BUG}
		bugCount := 0 // initializing an int outside of the loop so we can see it inside and outside of the "for"
		for _, v := range crockSlotsBug {
			if v > 0 {
				bugCount++
			}
		}
		// count inedable
		crockSlotsIne := []float64{crockPot[i1].INEDIBLE, crockPot[i2].INEDIBLE, crockPot[i3].INEDIBLE, crockPot[i4].INEDIBLE}
		inedibleCount := 0 // initializing an int outside of the loop so we can see it inside and outside of the "for"
		for _, v := range crockSlotsIne {
			if v > 0 {
				inedibleCount++
			}
		}
		// count misc (moleworm, butterfly, etc.)
		crockSlotsMis := []float64{crockPot[i1].MISC, crockPot[i2].MISC, crockPot[i3].MISC, crockPot[i4].MISC}
		miscCount := 0 // initializing an int outside of the loop so we can see it inside and outside of the "for"
		for _, v := range crockSlotsMis {
			if v > 0 {
				miscCount++
			}
		}

		// Bacon and Eggs
		if meatVal <= 1 || eggVal <= 1 || vegVal != 0 {
		} else {
			fmt.Println(titles)
			fmt.Println(trueInName.NAME)
			fmt.Println("true")
		}
	}
	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// end of gigantic for loop
	// potential to call second ffunction one for tru one for false

	// ****standard output****
	// fmt.Println("FOR TESTINGS Meat: ", meatVal)
	// fmt.Println("FOR TESTINGS Fish: ", fishVal)
	// fmt.Println("FOR TESTINGS Egg: ", eggVal)
	// fmt.Println("FOR TESTINGS Fruit: ", fruitVal)
	// fmt.Println("FOR TESTINGS Vegetable: ", vegVal)
	// fmt.Println("FOR TESTINGS Sweetener: ", sweetVal)
	// fmt.Println("FOR TESTINGS Monster: ", monVal)
	// fmt.Println("FOR TESTINGS Dairy: ", dairyVal)
	// fmt.Println("FOR TESTINGS Bug: ", bugVal)
	// fmt.Println("FOR TESTINGS Inedible: ", inedVal)
	// fmt.Println("FOR TESTINGS Misc: ", miscVal)
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

}
