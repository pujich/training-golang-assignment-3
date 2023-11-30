package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"training-golang-assignment-3/controller"
	"training-golang-assignment-3/model"

	"math/rand"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	fmt.Println("Letsgo route")
	go routine()
	fmt.Println("sampe sini ga")
	e := echo.New()

	e.POST("/update", controller.UpdateData)

	e.Logger.Fatal(e.Start(":9000"))

	fmt.Println("Letsgo hit")

}

func routine() {
	//proses hit itself
	i := 0
	for {
		prosesHit()
		time.Sleep(time.Second * 15)
		i++
		if i > 5 {
			break
		}
	}

}

func prosesHit() {
	fmt.Println("start hit")
	// random seed

	requestData := model.WaterWind{
		Water: rand.Intn(20) + 1,
		Wind:  rand.Intn(20) + 1,
	}

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println("Error encoding request:", err)
		return
	}

	fmt.Printf("REQ: %v\n", requestData)

	res, err := http.NewRequest("POST", "http://localhost:9000/update", bytes.NewBuffer([]byte(requestBody)))
	if err != nil {
		log.Fatalln(err)
	}

	// Perform the request
	client := http.Client{}
	response, err := client.Do(res)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer response.Body.Close()

	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var responseStruct model.WaterWind
	// masukin ke struct
	err1 := json.Unmarshal([]byte(string(resBody)), &responseStruct)
	if err1 != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	fmt.Printf("RESPONSE: %v\n", responseStruct)
	fmt.Println("Done hit")

	printResult(responseStruct)
}

func printResult(data model.WaterWind) {
	var water int = data.Water
	var wind int = data.Wind

	statusWater := ""
	statusWind := ""

	if water <= 5 {
		statusWater = "aman"
	} else if (water > 5) && (water <= 8) {
		statusWater = "siaga"
	} else if water > 8 {
		statusWater = "bahaya"
	}

	if wind <= 6 {
		statusWind = "aman"
	} else if (wind > 6) && (wind <= 15) {
		statusWind = "siaga"
	} else if wind > 15 {
		statusWind = "bahaya"
	}

	// fmt.Printf("%v\n", data)
	fmt.Printf("water: %v, status water : %v\n", data.Water, statusWater)
	fmt.Printf("wind: %v, status wind : %v\n", data.Wind, statusWind)
}
