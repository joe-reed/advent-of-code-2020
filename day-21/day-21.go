package main

import (
	"regexp"
	"strings"
)

func puzzle1(foodList []string) int {
	allergens := getAllergens(foodList)
	ingredients := getIngredients(foodList)
	options := make(map[string][]string)
	for _, allergen := range allergens {
		options[allergen] = []string{}
	}

	for i, item := range foodList {
		for _, allergen := range allergens {
			if strings.Contains(item, allergen) {
				if len(options[allergen]) == 0 {
					options[allergen] = ingredients[i]
				} else {
					options[allergen] = intersect(options[allergen], ingredients[i])
				}
			}
		}
	}

	for i := range ingredients {
		for _, potentialAllergens := range options {
			for _, potentialAllergen := range potentialAllergens {
				ingredients[i] = remove(ingredients[i], potentialAllergen)
			}
		}
	}

	count := 0
	for _, v := range ingredients {
		count += len(v)
	}
	return count
}

func getAllergens(foodList []string) (result []string) {
	r := regexp.MustCompile(`contains (.+)\)`)
	for _, item := range foodList {
		allergenString := r.FindStringSubmatch(item)[1]
		allergens := strings.Split(allergenString, ", ")
		for _, allergen := range allergens {
			if !contains(result, allergen) {
				result = append(result, allergen)
			}
		}
	}
	return
}

func getIngredients(foodList []string) (result [][]string) {
	for _, item := range foodList {
		ingredientString := strings.Split(item, " (")[0]
		ingredients := strings.Split(ingredientString, " ")
		result = append(result, ingredients)
	}
	return
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func intersect(a, b []string) (result []string) {
	for _, v := range a {
		if contains(b, v) {
			result = append(result, v)
		}
	}
	return result
}

func remove(s []string, e string) (result []string) {
	for _, v := range s {
		if v != e {
			result = append(result, v)
		}
	}
	return
}
