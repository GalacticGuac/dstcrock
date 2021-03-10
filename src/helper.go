package main

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

// count strings in cropost array
func countIngName(x string, crockPot []IngredientDetails) int {
	count := 0
	for _, cp := range crockPot {
		if cp.NAME == x {
			count++
		}
	}
	return count
}

// if strings.Contains(x, "mV") {
// 	fmt.Println("Meat value too high")
// }
// if strings.Contains(x, "mv") {
// 	fmt.Println("Meat value too low")
// }
// if strings.Contains(x, "mC") {
// 	fmt.Println("Meat count too high")
// }
// if strings.Contains(x, "mc") {
// 	fmt.Println("Meat count to low")
// }
// if strings.Contains(x, "hV") {
// 	fmt.Println("Fish value too high")
// }
// if strings.Contains(x, "hv") {
// 	fmt.Println("Fish value too low")
// }
// if strings.Contains(x, "hC") {
// 	fmt.Println("Fish count too high")
// }
// if strings.Contains(x, "hc") {
// 	fmt.Println("Fish count to low")
// }
// if strings.Contains(x, "eV") {
// 	fmt.Println("Egg value too high")
// }
// if strings.Contains(x, "ev") {
// 	fmt.Println("Egg value too low")
// }
// if strings.Contains(x, "eC") {
// 	fmt.Println("Egg count too high")
// }
// if strings.Contains(x, "ec") {
// 	fmt.Println("Egg count to low")
// }
// if strings.Contains(x, "fV") {
// 	fmt.Println("Fruit value too high")
// }
// if strings.Contains(x, "fv") {
// 	fmt.Println("Fruit value too low")
// }
// if strings.Contains(x, "fC") {
// 	fmt.Println("Fruit count too high")
// }
// if strings.Contains(x, "fc") {
// 	fmt.Println("Fruit count to low")
// }
// if strings.Contains(x, "gV") {
// 	fmt.Println("Vegetable value too high")
// }
// if strings.Contains(x, "gv") {
// 	fmt.Println("Vegetable value too low")
// }
// if strings.Contains(x, "gC") {
// 	fmt.Println("Vegetable count too high")
// }
// if strings.Contains(x, "gc") {
// 	fmt.Println("Vegetable count to low")
// }
// if strings.Contains(x, "sV") {
// 	fmt.Println("Sweetener value too high")
// }
// if strings.Contains(x, "sv") {
// 	fmt.Println("Sweetener value too low")
// }
// if strings.Contains(x, "sC") {
// 	fmt.Println("Sweetener count too high")
// }
// if strings.Contains(x, "sc") {
// 	fmt.Println("Sweetener count to low")
// }
// if strings.Contains(x, "tV") {
// 	fmt.Println("Monster value too high")
// }
// if strings.Contains(x, "tv") {
// 	fmt.Println("Monster value too low")
// }
// if strings.Contains(x, "tC") {
// 	fmt.Println("Monster count too high")
// }
// if strings.Contains(x, "tc") {
// 	fmt.Println("Monster count to low")
// }
// if strings.Contains(x, "dV") {
// 	fmt.Println("Dairy value too high")
// }
// if strings.Contains(x, "dv") {
// 	fmt.Println("Dairy value too low")
// }
// if strings.Contains(x, "dC") {
// 	fmt.Println("Dairy count too high")
// }
// if strings.Contains(x, "dc") {
// 	fmt.Println("Dairy count to low")
// }
// if strings.Contains(x, "bV") {
// 	fmt.Println("Bug value too high")
// }
// if strings.Contains(x, "bv") {
// 	fmt.Println("Bug value too low")
// }
// if strings.Contains(x, "bC") {
// 	fmt.Println("Bug count too high")
// }
// if strings.Contains(x, "bc") {
// 	fmt.Println("Bug count to low")
// }
// if strings.Contains(x, "iV") {
// 	fmt.Println("Inedible value too high")
// }
// if strings.Contains(x, "iv") {
// 	fmt.Println("Inedible value too low")
// }
// if strings.Contains(x, "iC") {
// 	fmt.Println("Inedible count too high")
// }
// if strings.Contains(x, "ic") {
// 	fmt.Println("Inedible count to low")
// }
// if strings.Contains(x, "lV") {
// 	fmt.Println("Misc value too high")
// }
// if strings.Contains(x, "lv") {
// 	fmt.Println("Misc value too low")
// }
// if strings.Contains(x, "lC") {
// 	fmt.Println("Misc count too high")
// }
// if strings.Contains(x, "lc") {
// 	fmt.Println("Misc count to low")
// }
// if strings.Contains(x, "oC") {
// 	fmt.Println("Required Ingredient count too high")
// }
// if strings.Contains(x, "oc") {
// 	fmt.Println("Required Ingredient count too low")
// }
// if strings.Contains(x, "x") {
// 	fmt.Println("Crock Has Exclusion")
// }
// }
