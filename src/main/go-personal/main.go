package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CityData struct {
	Averages map[string]float32
	Max      map[string]float32
	Min      map[string]float32
	Nb       map[string]int
}

func main() {

	filename := os.Args[1]

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := CityData{}

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}

		linePart := strings.Split(line, ";")
		city := linePart[0]
		temperatureString := linePart[1]

		temperature, err := strconv.ParseFloat(temperatureString, 32)
		if err != nil {
			log.Fatal(err)
		}

		updateData(&data, city, float32(temperature))
	}

	cities := make([]string, len(data.Nb), 0)

	for city := range data.Nb {
		cities = append(cities, city)
	}

	sort.Strings(cities)

	for _, city := range cities {
		fmt.Println(city, data.Min[city], data.Averages[city], data.Max[city])
	}

}

func updateData(data *CityData, city string, temperature float32) {
	if data.Averages[city] == 0 {
		data.Averages[city] = temperature
		data.Max[city] = temperature
		data.Min[city] = temperature
		data.Nb[city] = 1
	} else {
		data.Averages[city] = (data.Averages[city]*float32(data.Nb[city]) + temperature) / float32(data.Nb[city]+1)
		data.Nb[city]++
		if temperature > data.Max[city] {
			data.Max[city] = temperature
		}
		if temperature < data.Min[city] {
			data.Min[city] = temperature
		}
	}
}
