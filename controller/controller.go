package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"training-golang-assignment-3/model"

	"github.com/labstack/echo/v4"
)

func UpdateData(c echo.Context) error {

	data := model.WaterWind{}

	if err := c.Bind(&data); err != nil { //ngebind object echo context yg masuk ke objek
		return err
	}

	fmt.Printf("Request masuk: %v", data)

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return err
	}

	writeToOrCreateFile("C:/Users/Dell/go/src/training-golang-assignment-3/output/output.json", jsonData)

	return c.JSON(http.StatusOK, data)
}

// writeToOrCreateFile appends data to a file with the given filename or creates the file if it doesn't exist
func writeToOrCreateFile(filename string, data []byte) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	// file, err := os.WriteFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
