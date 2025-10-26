package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
)

type countries_data struct {
	countries []country_data
}

type country_data struct {
	country string
	facts   []string
}

func main() {
	data, err := os.ReadFile("data/countries_data.json")
	if err != nil {
		log.Fatal("Error reading the file")
	}
	var dataSkeleton map[string][]string
	json.Unmarshal(data, &dataSkeleton)

	cdata := countries_data{}
	for key, value := range dataSkeleton {
		c := country_data{key, value}
		cdata.countries = append(cdata.countries, c)
	}
	cdata = shuffleData(cdata)
	sample := cdata.countries[0:5]
	countryName := sample[0].country
	facts := sample[0].facts
	for i := range len(sample[0].facts) {
		println(facts[i])
	}
	println("Try to guess the name of country based on the facts")
	println("You have 5 attempts")

	var attempedName string
	for _ = range 5 {
		fmt.Scanln(&attempedName)
		if attempedName == countryName {
			fmt.Println("Yayy, you guessed it correctly")
			return
		} else {
			fmt.Println("Try again, you can do it")
		}
	}
	fmt.Printf("The name of the country was %q\n", countryName)
}

func shuffleData(c countries_data) countries_data {
	rand.Shuffle(len(c.countries), func(i, j int) {
		c.countries[i], c.countries[j] = c.countries[j], c.countries[i]
	})
	return c
}
