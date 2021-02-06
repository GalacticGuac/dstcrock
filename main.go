package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Ingredient struct {
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

	jsonFile, err := os.Open("./ingredients.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	var ingredientsData []Ingredient
	err = json.Unmarshal(byteValue, &ingredientsData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ingredientsData[0].INGREDIENT)

}
