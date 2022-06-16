package main

import "fmt"

func main() {
	var countryCapitalMap = make(map[string]string)

	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Kenya"] = "Kenya"
	countryCapitalMap["Tanzania"] = "Dodoma"
	countryCapitalMap["Nigeria"] = "Abuja"

	for country, capital := range countryCapitalMap {

		fmt.Println("Capital of ", country, " is ", capital)

	}

	delete(countryCapitalMap, "France")
	fmt.Println("Entry France deleted")
	fmt.Println("Updated map")
}
