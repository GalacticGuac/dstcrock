package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type IngrediantList struct {
	Ingredient map[string]IngredientDetails
}

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

	jsonFile, err := os.Open("../lib/ingredients-reformat.json")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	ingredientsData := map[string]IngredientDetails{}
	err = json.Unmarshal(byteValue, &ingredientsData)
	if err != nil {
		fmt.Println(err)
	}

	// Set up fake Data/user input data
	ingredients := []IngredientDetails{}
	ingredients = append(ingredients, ingredientsData["Morsel"])
	ingredients = append(ingredients, ingredientsData["Egg"])
	ingredients = append(ingredients, ingredientsData["Egg"])
	ingredients = append(ingredients, ingredientsData["Monster Meat"])

	meatVal := 0.0
	vegVal := 0.0
	fruitVal := 0.0
	eggVal := 0.0

	titles := []string{}

	for _, ingre := range ingredients {
		titles = append(titles, ingre.NAME)

		meatVal += ingre.MEAT

		vegVal += ingre.VEGETABLE

		fruitVal += ingre.FRUIT

		eggVal += ingre.EGG

	}

	title := fmt.Sprintf("%s + %s + %s + %s:", titles[0], titles[1], titles[2], titles[3])
	fmt.Println(title)

	fmt.Println("Meat: ", meatVal)
	fmt.Println("Vegetable: ", vegVal)
	fmt.Println("Fruit: ", fruitVal)
	fmt.Println("Egg: ", eggVal)

}
