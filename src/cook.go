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
	INGREDIENT string  `json:"INGREDIENT"`
	MEAT       float64 `json:"MEAT"`
	FISH       float64 `json:"FISH"`
	EGG        float64 `json:"EGG"`
	FRUIT      float64 `json:"FRUIT"`
	VEGETABLE  float64 `json:"VEGETABLE"`
	SWEETENER  float64 `json:"SWEETENER"`
	MONSTER    float64 `json:"MONSTER"`
	DAIRY      float64 `json:"DAIRY"`
	BUG        float64 `json:"BUG"`
	INEDIBLE   float64 `json:"INEDIBLE"`
	MISC       float64 `json:"MISC"`
	HEALTH     float64 `json:"HEALTH"`
	HUNGER     float64 `json:"HUNGER"`
	SANITY     float64 `json:"SANITY"`
	DAYS       float64 `json:"DAYS"`
	NOTES      string  `json:"NOTES"`
}

func main() {

	jsonFile, err := os.Open("../lib/ingredients.json")
	if err != nil {
		fmt.Println(err)
	}
	
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	var ingredientsData []IngredientDetails
	err = json.Unmarshal(byteValue, &ingredientsData)
	if err != nil {
		fmt.Println(err)
	}
	// 0 3 7 10

	ingredients := []IngredientDetails{}
	ingredients = append(ingredients, ingredientsData[0])
	ingredients = append(ingredients, ingredientsData[3])
	ingredients = append(ingredients, ingredientsData[7])
	ingredients = append(ingredients, ingredientsData[10])

	// fmt.Println(ingredients)
	meatVal := 0.0
	vegVal := 0.0
	fruitVal := 0.0
	eggVal := 0.0
	for _, ingre := range ingredients {
		meatVal += ingre.MEAT

		vegVal += ingre.VEGETABLE

		fruitVal += ingre.FRUIT

		eggVal += ingre.EGG

	}
	fmt.Println("meat", meatVal)
	fmt.Println("vegetable", vegVal)
	fmt.Println("fruit", fruitVal)
	fmt.Println("egg", eggVal)

}
