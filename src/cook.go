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
	DAYS      float64 `json:"DAYS"`
	NOTES     string  `json:"NOTES"`
}

type RecipeDetails struct {
	NAME        string  `json:"NAME"`
	HEALTH      float64 `json:"HEALTH"`
	HUNGER      float64 `json:"HUNGER"`
	SANITY      float64 `json:"SANITY"`
	INGREDIENTS string  `json:"INGREDIENTS"`
	EXCLUDE     string  `json:"EXCLUDE"`
	NOTES       string  `json:"NOTES"`
	EXPIRES     float64 `json:"EXPIRES"`
	COOKTIME    float64 `json:"COOK TIME"`
	PRIORITY    float64 `json:"PRIORITY"`
	// preferences?
}

func main() {

	// version := DST
	// Find that file ingredients
	jsonFile, err := os.Open("../lib/ingredients.json")
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

	// Add Ingredients to the pot
	i1 := "Barnacles"
	i2 := "Monster Jerky"
	i3 := "Egg"
	i4 := "Berries"
	crockPot := map[string]IngredientDetails{}
	crockPot[i1] = ingredientsData[i1]
	crockPot[i2] = ingredientsData[i2]
	crockPot[i3] = ingredientsData[i3]
	crockPot[i4] = ingredientsData[i4]

	// Initialize the baseline values *brian these were 0.0 before im messing with it to see if it works right
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
		fmt.Println("remove AMBEROSIA")
	} else {
		fmt.Println("keep AMBEROSIA")
	}

	// ASPARAGUS SOUP
	if strings.Count(titles, "Asparagus") < 1 || vegVal < 1.5 || meatCount != 0 || inedibleCount != 0 {
		fmt.Println("remove ASPARAGUS SOUP")
	} else {
		fmt.Println("keep ASPARAGUS SOUP")
	}

	// Bacon and Eggs
	if meatVal <= 1 || eggVal <= 1 || vegVal != 0 {
		fmt.Println("remove bacon and eggs")
	} else {
		fmt.Println("keep bacon and eggs")
	}

	// BANANA POP
	if strings.Count(titles, "Banana") < 1 || strings.Count(titles, "twigs") < 1 || meatCount != 0 || fishCount != 0 {
		fmt.Println("remove BANANA POP")
	} else {
		fmt.Println("keep BANANA POP")
	}

	// Barnacle Linguine
	if strings.Count(titles, "Barnacles") != 2 || vegeCount != 2 {
		fmt.Println("remove barnacle liguine")
	} else {
		fmt.Println("keep barnacle liguine")
	}

	// Barnacle Nigiri
	if strings.Count(titles, "Barnacles") < 1 || strings.Count(titles, "Kelp Fronds") < 1 || eggCount < 1 {
		fmt.Println("remove barnacle Nigiri")
	} else {
		fmt.Println("keep barnacle Nigiri")
	}

	// BARNACLE PITA
	// fishval over 1 can cause stuffed fish head but idk the percentage
	if strings.Count(titles, "Barnacles") < 1 || vegVal < 0.5 {
		fmt.Println("remove barnacle pita")
	} else {
		fmt.Println("keep barnacle pita")
	}

	// BEEFY GREENS
	if strings.Count(titles, "Leafy Meat") < 1 || vegVal < 3 {
		fmt.Println("remove beefy greens")
	} else {
		fmt.Println("keep beefy greens")
	}

	// butter muffin
	if strings.Count(titles, "Butterfly Wings") < 1 || vegeCount < 1 || crockPot["Mandrake"].NAME != "" || meatCount != 0 {
		fmt.Println("remove butter muffin")
	} else {
		fmt.Println("keep butter muffin")
	}

	// CALIFORNIA ROLL
	if strings.Count(titles, "Kelp Fronds") < 2 || fishVal < 1 {
		fmt.Println("remove CALIFORNIA ROLL")
	} else {
		fmt.Println("keep CALIFORNIA ROLL")
	}

	// ceviche
	if strings.Count(titles, "Ice") < 2 || fishVal < 2 || eggCount != 0 || inedibleCount != 0 {
		fmt.Println("remove CALIFORNIA ROLL")
	} else {
		fmt.Println("keep CALIFORNIA ROLL")
	}

	// CREAMY POTATO PURÉE**
	// need to look at potato or roasted potato
	if strings.Count(titles, "Potato") < 2 || strings.Count(titles, "Garlic") < 1 || crockPot["Twigs"].NAME != "" || meatCount != 0 {
		fmt.Println("remove CREAMY POTATO PURÉE")
	} else {
		fmt.Println("keep CREAMY POTATO PURÉE")
	}

	// dragonpie
	// need to look at dragon fruit or prepared dragon fruit
	if strings.Count(titles, "Dragon Fruit") < 1 || crockPot["Mandrake"].NAME != "" || meatCount != 0 {
		fmt.Println("remove dragon pie")
	} else {
		fmt.Println("keep dragon pie")
	}

	// FANCY SPIRALLED TUBERS
	// need to look at potato or roasted potato
	if strings.Count(titles, "Potato") < 1 || strings.Count(titles, "Twigs") < 1 || inedibleCount-1 > 1 || meatCount != 0 {
		fmt.Println("remove FANCY SPIRALLED TUBERS")
	} else {
		fmt.Println("keep FANCY SPIRALLED TUBERS")
	}

	// fish tacos
	// twig value - 1 twig 50% chance of  fish sticks
	// need to look at corn or popcorn
	if fishVal < 0.5 || (strings.Count(titles, "Corn") < 1 || strings.Count(titles, "Popcorn") < 1) {
		fmt.Println("remove fish tacos")
	} else {
		fmt.Println("keep fish tacos")
	}

	// fishsticks
	if fishVal < 0.25 || strings.Count(titles, "Twigs") != 1 || crockPot["Moleworm"].NAME != "" {
		fmt.Println("remove fishsticks")
	} else {
		fmt.Println("keep fishsticks")
	}

	// fist full of jam
	if fruitVal < 0.5 || meatVal != 0 || vegVal != 0 || inedVal != 0 || crockPot["Dragon Fruit"].NAME != "" {
		fmt.Println("remove fist full of jam")
	} else {
		fmt.Println("keep fist full of jam")
	}

	// flower salad
	if strings.Count(titles, "Cactus Flower") < 1 || vegVal-.5 < 1.5 || meatVal != 0 || fruitVal != 0 || eggVal != 0 || sweetVal != 0 || crockPot["Twigs"].NAME != "" {
		fmt.Println("remove flower salad")
	} else {
		fmt.Println("keep flower salad")
	}

	// froggle bunwich
	// need to look at frog legs or cooked frog legs
	// makes kabob if only one stick
	if strings.Count(titles, "Frog Legs") < 1 || vegeCount < 1 || eggVal != 0 || sweetVal != 0 || crockPot["Mandrake"].NAME != "" {
		fmt.Println("remove froggle bunwich")
	} else {
		fmt.Println("keep froggle bunwich")
	}

	// fruit medley
	// twigs is safest anything else 50% of fist full of jam
	// SIDNEY: We will definitely need to think about suggestions for the "best" version of the recipes
	if fruitVal < 3 || meatVal != 0 || vegVal != 0 || crockPot["Dragon Fruit"].NAME != "" {
		fmt.Println("remove fruit medley")
	} else {
		fmt.Println("keep fruit medley")
	}

	// guacamole
	// need to look at cactus flesh or stone fruit
	if strings.Count(titles, "Moleworm") < 1 || (strings.Count(titles, "Cactus Flesh") < 1 || strings.Count(titles, "Ripe Stone Fruit") < 1) || fruitVal != 0 {
		fmt.Println("remove guacamole")
	} else {
		fmt.Println("keep guacamole")
	}

	// Honey Ham
	if strings.Count(titles, "Honey") < 1 || meatVal <= 1.5 || crockPot["Twigs"].NAME != "" || crockPot["Moleworm"].NAME != "" || crockPot["Mandrake"].NAME != "" || crockPot["Tallbird Egg"].NAME != "" {
		fmt.Println("remove honey ham")
	} else {
		fmt.Println("keep honey ham")
	}

	// Honey Nuggets
	if strings.Count(titles, "Honey") < 1 || meatVal > 1.5 || inedVal != 0 {
		fmt.Println("remove honey nuggets")
	} else {
		fmt.Println("keep honey nuggets")
	}

	// ice cream
	if strings.Count(titles, "Ice") < 1 || dairyCount < 1 || sweetenerCount < 1 || meatVal != 0 || vegVal != 0 || eggVal != 0 || crockPot["Twigs"].NAME != "" {
		fmt.Println("remove ice cream")
	} else {
		fmt.Println("keep ice cream")
	}

	// jelly salad
	// look for cooked version
	if strings.Count(titles, "Leafy Meat") < 2 || sweetVal < 2 {
		fmt.Println("remove jelly salad")
	} else {
		fmt.Println("keep jelly salad")
	}

	// jellybeans
	if strings.Count(titles, "Royal Jelly") < 1 || monVal != 0 || inedVal != 0 {
		fmt.Println("remove jellybeans")
	} else {
		fmt.Println("keep jellybeans")
	}

	// kabobs
	if meatCount < 1 || strings.Count(titles, "Twigs") != 1 || crockPot["Moleworm"].NAME != "" || crockPot["Mandrake"].NAME != "" || fishVal != 0 {
		fmt.Println("remove kabobs")
	} else {
		fmt.Println("keep kabobs")
	}

	// leafy meatloaf
	if strings.Count(titles, "Leafy Meat") < 2 {
		fmt.Println("remove leafy meat loaf")
	} else {
		fmt.Println("keep leafy meat loaf")
	}

	// LOBSTER BISQUE
	if strings.Count(titles, "Wobster") < 1 || strings.Count(titles, "Ice") < 1 {
		fmt.Println("remove leafy meat loaf")
	} else {
		fmt.Println("keep leafy meat loaf")
	}

	// mandrake soup
	if strings.Count(titles, "Mandrake") < 1 {
		fmt.Println("remove mandrake soup")
	} else {
		fmt.Println("keep mandrake soup")
	}

	// Meatballs
	if meatVal >= 3 || meatVal == 0 || crockPot["Twigs"].NAME != "" {
		fmt.Println("remove meatballs")
	} else {
		fmt.Println("keep meatballs")
	}

	// meaty stew
	if meatVal < 3 || crockPot["Twigs"].NAME != "" || crockPot["Moleworm"].NAME != "" || crockPot["Honey"].NAME != "" || crockPot["Mandrake"].NAME != "" || crockPot["Tallbird Egg"].NAME != "" {
		fmt.Println("remove meaty stew")
	} else {
		fmt.Println("keep meaty stew")
	}

	// melonsicle
	if strings.Count(titles, "Watermelon") < 1 || strings.Count(titles, "Ice") < 1 || strings.Count(titles, "Twigs") < 1 || meatVal != 0 || vegVal != 0 || eggVal != 0 {
		fmt.Println("remove melonsicle")
	} else {
		fmt.Println("keep melonsicle")
	}

	// MILKMADE HAT
	if strings.Count(titles, "Nostrils") < 1 || strings.Count(titles, "Kelp Fronds") < 1 || dairyCount < 1 {
		fmt.Println("remove MILKMADE HAT")
	} else {
		fmt.Println("keep MILKMADE HAT")
	}

	// monster lasagna
	if monsterCount < 2 || crockPot["Twigs"].NAME != "" {
		fmt.Println("remove monster lasagna")
	} else {
		fmt.Println("keep monster lasagna")
	}

	// Mushy Cake
	if strings.Count(titles, "Moon Shroom") != 1 || strings.Count(titles, "Red Cap") != 1 || strings.Count(titles, "Blue Cap") != 1 || strings.Count(titles, "Green Cap") != 1 {
		fmt.Println("remove Mushy cake")
	} else {
		fmt.Println("keep Mushy cake")
	}

	// Pierogi
	if meatCount < 1 || eggCount < 1 || vegeCount < 1 || crockPot["Twigs"].NAME != "" || crockPot["Mandrake"].NAME != "" {
		fmt.Println("remove Pierogi")
	} else {
		fmt.Println("keep Pierogi")
	}

	//powdercake
	if strings.Count(titles, "Corn") < 1 || strings.Count(titles, "Honey") < 1 || strings.Count(titles, "Twigs") < 1 {
		fmt.Println("remove powdercake")
	} else {
		fmt.Println("keep powder cake")
	}

	//Pumkin cookie
	// 3 honey or comb 50% chance of making taffy
	if strings.Count(titles, "Pumpkin") < 1 || strings.Count(titles, "Honey") <= 1 {
		fmt.Println("remove Pumpkin cookie")
	} else {
		fmt.Println("keep Pumpkin cookie")
	}

	// RATATOUILLE
	if vegeCount < 1 || crockPot["Twigs"].NAME != "" || crockPot["Mandrake"].NAME != "" || crockPot["Butterfly Wings"].NAME != "" || crockPot["Dragon Fruit"].NAME != "" {
		fmt.Println("remove Ratatouille")
	} else {
		fmt.Println("keep Ratatouille")
	}

	// SALSA FRESCA
	// look for cooked versions
	if strings.Count(titles, "Toma Root") < 1 || meatCount != 0 || inedibleCount != 0 || eggCount != 0 {
		fmt.Println("remove Salsa Fresca")
	} else {
		fmt.Println("keep Salsa Fresca")
	}

	// SEAFOOD GUMBO
	if strings.Count(titles, "Eel") < 1 || fishVal <= 2 {
		fmt.Println("remove SEAFOOD GUMBO")
	} else {
		fmt.Println("keep SEAFOOD GUMBO")
	}

	//  Soothing tea
	// look for honey or comb
	if strings.Count(titles, "Forget-Me-Lots") < 1 || strings.Count(titles, "Honey") < 1 || strings.Count(titles, "Ice") < 1 || monsterCount != 0 || meatCount != 0 || fishCount != 0 || eggCount != 0 || inedibleCount != 0 || dairyCount != 0 {
		fmt.Println("remove Soothing tea")
	} else {
		fmt.Println("keep Soothing tea")
	}

	//  Spicy Chili
	if meatCount != 2 || vegeCount != 2 || meatVal < 1.5 || vegVal < 1.5 {
		fmt.Println("remove Spicy Chili")
	} else {
		fmt.Println("keep Spicy Chili")
	}

	// STUFFED EGGPLANT
	if strings.Count(titles, "Eggplant") < 1 || vegeCount < 1 {
		fmt.Println("remove STUFFED EGGPLANT")
	} else {
		fmt.Println("keep STUFFED EGGPLANT")
	}

	// stuffed fish heads
	if strings.Count(titles, "Barnacles") < 1 || fishVal-0.5 < 1 {
		fmt.Println("remove stuffed fish heads")
	} else {
		fmt.Println("keep stuffed fish heads")
	}

	// STUFFED PEPPER POPPERS
	if strings.Count(titles, "Pepper") < 1 || meatVal > 1.5 || meatCount == 0 || crockPot["Twigs"].NAME != "" {
		fmt.Println("remove STUFFED PEPPER POPPERS")
	} else {
		fmt.Println("keep STUFFED PEPPER POPPERS")
	}

	// SURF 'N' TURF
	if meatVal < 2.5 || fishVal < 1.5 || crockPot["Ice"].NAME != "" {
		fmt.Println("remove SURF 'N' TURF")
	} else {
		fmt.Println("keep SURF 'N' TURF")
	}

	// Taffy
	if sweetVal < 3 || meatCount != 0 {
		fmt.Println("remove Taffy")
	} else {
		fmt.Println("keep Taffy")
	}

	// trail mix
	// required berries must be uncooked
	if strings.Count(titles, "Roasted Birchnut") < 1 || strings.Count(titles, "Berries") < 1 || strings.Count(titles, "Roasted") > 3 || fruitCount-1 < 1 || meatCount != 0 || fishCount != 0 || eggCount != 0 || vegeCount != 0 || dairyCount != 0 {
		fmt.Println("remove trail mix")
	} else {
		fmt.Println("keep trail mix")
	}

	// TURKEY DINNER
	// checkingg twice for veg or fruit
	if strings.Count(titles, "Drumstick") < 2 || meatVal-1 < 0.25 || (fruitVal < .5 || vegVal < .5) {
		fmt.Println("remove TURKEY DINNER")
	} else {
		fmt.Println("keep TURKEY DINNER") //fmt.Println("remove TURKEY DINNER")
	}

	// unagi
	if strings.Count(titles, "Eel") < 1 || (strings.Count(titles, "Lichen") < 1 || strings.Count(titles, "Kelp Fronds") < 1) {
		fmt.Println("remove unagi")
	} else {
		fmt.Println("keep unagi")
	}

	// VEGETABLE STINGER
	if (strings.Count(titles, "Toma Root") < 1 || strings.Count(titles, "Asparagus") < 1) || strings.Count(titles, "Ice") < 1 || vegVal-1 < 1.5 {
		fmt.Println("remove VEGETABLE STINGER")
	} else {
		fmt.Println("keep VEGETABLE STINGER")
	}

	// Veggie Burger
	if strings.Count(titles, "Leafy Meat") < 1 || strings.Count(titles, "Onion") < 1 || vegVal-1 < 1 {
		fmt.Println("remove Veggie Burger")
	} else {
		fmt.Println("keep Veggie Burger")
	}

	// waffles
	if strings.Count(titles, "Butter") < 1 || strings.Count(titles, "Berries") < 1 || eggCount < 1 {
		fmt.Println("remove waffles")
	} else {
		fmt.Println("keep waffles")
	}

	// wet goop
	// if everything is false

	// WOBSTER DINNER
	if strings.Count(titles, "Wobster") < 1 || strings.Count(titles, "Butter") < 1 || meatCount != 0 || fishCount != 0 || crockPot["Twigs"].NAME != "" {
		fmt.Println("remove waffles")
	} else {
		fmt.Println("keep waffles")
	}

	// Output goes here
	// title := fmt.Sprintf("%s + %s + %s + %s:", titles[0], titles[1], titles[2], titles[3])

	fmt.Println(titles)

	fmt.Println("Meat: ", meatVal)
	fmt.Println("Fish: ", fishVal)
	fmt.Println("Egg: ", eggVal)
	fmt.Println("Fruit: ", fruitVal)
	fmt.Println("Vegetable: ", vegVal)
	fmt.Println("Sweetener: ", sweetVal)
	fmt.Println("Monster: ", monVal)
	fmt.Println("Dairy: ", dairyVal)
	fmt.Println("Bug: ", bugVal)
	fmt.Println("Inedible: ", inedVal)
	fmt.Println("Misc: ", miscVal)
	// count fouud types
	// fmt.Println("meat", meatCount)
	// fmt.Println("fish", fishCount)
	// fmt.Println("egg", eggCount)
	// fmt.Println("fruit", fruitCount)
	// fmt.Println("vege", vegeCount)
	// fmt.Println("sweetener", sweetenerCount)
	// fmt.Println("monster", monsterCount)
	// fmt.Println("dairy", dairyCount)
	// fmt.Println("bug", bugCount)
	// fmt.Println("inedible", inedibleCount)
	// fmt.Println("misc", miscCount)
}
