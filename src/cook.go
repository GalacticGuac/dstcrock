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
	NAME     string  `json:"NAME"`
	HEALTH   float64 `json:"HEALTH"`
	HUNGER   float64 `json:"HUNGER"`
	SANITY   float64 `json:"SANITY"`
	NOTES    string  `json:"NOTES"`
	CRITERIA string  `json:"CRITERIA"`
	EXPIRES  float64 `json:"EXPIRES"`
	PRIORITY float64 `json:"PRIORITY"`
	// preferences?
}

func main() {
	// Find that file
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
	i1 := "Berries"
	i2 := "Barnacles"
	i3 := "Blue Cap"
	i4 := "Blue Cap"
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
	miscVal := crockPot[i1].MEAT + crockPot[i2].MEAT + crockPot[i3].MEAT + crockPot[i4].MEAT

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

	// Bacon and Eggs
	if meatVal <= 1 || eggVal <= 1 || vegVal > 0 {
		// remove bacon and eggs from possibles
		fmt.Println("remove bacon and eggs")
	} else {
		fmt.Println("keep bacon and eggs")
	}

	// Barnacle Linguine
	// *original* if crockPot["Barnacles"].NAME == "" || crockPot["Vegetable"].NAME == "" {}

	// testing if there are 2 vegetables in crockpot

	// first if is counting the occurence of "barnecle in "titles" (line 75)
	if strings.Count(titles, "Barnacles") != 2 {
		fmt.Println("remove barnacle liguine")
	} else {
		fmt.Println("keep barnacle liguine")
	}

	// second find vegetables tring to make "true" into 1 then adding to make 2. if != 2 then remove
	crockSlots := []float64{crockPot[i1].VEGETABLE, crockPot[i2].VEGETABLE, crockPot[i3].VEGETABLE, crockPot[i4].VEGETABLE}
	vegeCount := 0 // initializing an int outside of the loop so we can see it inside and outside of the "for"
	for _, v := range crockSlots {
		foundVege := false // initializing a boolean (true/false) variable outside the if statement, so we can alter it inside and check after

		if v > 0 {
			foundVege = true
			fmt.Println("true")
		}

		if foundVege { // if statements always read `if "true" do this`, so since it a boolean, this says the same thing as if foundVege == true
			vegeCount += 1 // Adding 1 to the number of veggies found.
		}
	}

	if vegeCount >= 2 {
		//remove recipe
	}

	// if i1, i2 := crockPot["Egg"], crockPot["BEgg"];  { // || _, ok := crockPot["Egg"]; ok {
	// 	//  remove barnacle recipe
	// 	fmt.Println("has no egg: ")
	// }

	// Meatballs
	if meatVal > 3 || meatVal == 0 || crockPot["Twigs"].NAME != "" {
		// remove meatball recipe from possibles
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
}
