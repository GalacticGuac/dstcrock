package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type IngredientDetails struct {
	NAME      string  `json:"INGREDIENT"`
	MEAT      float64 `json:"MEAT"`
	FISH      float64 `json:"FISH"`
	EGG       float64 `json:"EGG"`
	FRUIT     float64 `json:"FRUIT"`
	VEGETABLE float64 `json:"VEGETABLE"`
	SWEETENER float64 `json:"SWEETENER"`
	MONSTER   float64 `json:"MONSTER"`
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
	crockPot := []IngredientDetails{}
	crockPot = append(crockPot, ingredientsData["Morsel"])
	crockPot = append(crockPot, ingredientsData["Egg"])
	crockPot = append(crockPot, ingredientsData["Egg"])
	crockPot = append(crockPot, ingredientsData["Monster Meat"])

	// Initialize the baseline values
	meatVal := 0.0
	fishVal := 0.0
	eggVal := 0.0
	fruitVal := 0.0
	vegVal := 0.0
	sweetVal := 0.0
	monVal := 0.0
	dairyVal := 0.0
	bugVal := 0.0
	inedVal := 0.0
	miscVal := 0.0

	titles := []string{}

	// Look through each ingredient
	for _, ingre := range crockPot {
		titles = append(titles, ingre.NAME)

		meatVal += ingre.MEAT

		fishVal += ingre.FISH

		eggVal += ingre.EGG

		fruitVal += ingre.FRUIT

		vegVal += ingre.VEGETABLE

		sweetVal += ingre.SWEETENER

		monVal += ingre.MONSTER

		dairyVal += ingre.DAIRY

		bugVal += ingre.BUG

		inedVal += ingre.INEDIBLE

		miscVal += ingre.MISC

	}

	// Output goes here
	title := fmt.Sprintf("%s + %s + %s + %s:", titles[0], titles[1], titles[2], titles[3])
	fmt.Println(title)

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
