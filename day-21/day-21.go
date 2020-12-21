package main

import (
	"regexp"
	"sort"
	"strings"
)

func puzzle1(foodList []string) int {
	allergens := getAllergens(foodList)
	ingredients := getIngredients(foodList)
	options := getAllergenOptions(allergens, ingredients)

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

func puzzle2(foodList []string) string {
	allergens := getAllergens(foodList)
	ingredients := getIngredients(foodList)
	options := getAllergenOptions(allergens, ingredients)

	finished := false
	for !finished {
		for i, a := range options {
			for j, b := range options {
				if i == j {
					continue
				}
				if len(a) != 1 {
					continue
				}
				options[j] = remove(b, a[0])
			}
		}
		finished = allMatched(options)
	}

	allergenList := []string{}
	for a := range options {
		if !contains(allergenList, a) {
			allergenList = append(allergenList, a)
		}
	}
	sort.Strings(allergenList)

	dangerousIngredients := []string{}
	for _, a := range allergenList {
		dangerousIngredients = append(dangerousIngredients, options[a][0])
	}
	return strings.Join(dangerousIngredients, ",")
}

func getAllergenOptions(allergens [][]string, ingredients [][]string) map[string][]string {
	options := make(map[string][]string)
	for _, a := range allergens {
		for _, allergen := range a {
			options[allergen] = []string{}
		}
	}

	for i, a := range allergens {
		for _, allergen := range a {
			if len(options[allergen]) == 0 {
				options[allergen] = ingredients[i]
			} else {
				options[allergen] = intersect(options[allergen], ingredients[i])
			}
		}
	}
	return options
}

func getAllergens(foodList []string) (result [][]string) {
	r := regexp.MustCompile(`contains (.+)\)`)
	for _, item := range foodList {
		allergenString := r.FindStringSubmatch(item)[1]
		allergens := strings.Split(allergenString, ", ")
		result = append(result, allergens)
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

func allMatched(options map[string][]string) bool {
	for _, o := range options {
		if len(o) > 1 {
			return false
		}
	}
	return true
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
