package helpers

import (
	"assignment-3/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
)

func GetDataJson() models.Status {
	jsonFile, err := os.Open("./data/data.json")
	jsonByte, _ := ioutil.ReadAll(jsonFile)

	if err != nil {
		panic(err)
	}

	var dataWeather models.Status

	errMarshal := json.Unmarshal(jsonByte, &dataWeather)

	if errMarshal != nil {
		fmt.Println(errMarshal.Error())
	}

	return dataWeather
}

func UpdateJSONData() {
	water, wind := rand.Intn(100), rand.Intn(100)

	jsonString := fmt.Sprintf(`
	{
		"status": {
			"water": %v,
			"wind": %v
		}
	}`, water, wind)

	data, err := os.Create("./data/data.json")

	if err != nil {
		log.Fatal("Error in json update: ", err.Error())

		return
	}

	_, err = data.Write([]byte(jsonString))

	if err != nil {
		log.Println("Error in json update: ", err.Error())

		return
	}
}

func GetWaterStatus(water uint) string {
	switch {
	case water <= 5:
		return "Safe"
	case water <= 8:
		return "Alert"
	default:
		return "Danger"
	}
}

func GetWindStatus(wind uint) string {
	switch {
	case wind <= 6:
		return "Safe"
	case wind <= 15:
		return "Alert"
	default:
		return "Danger"
	}
}
