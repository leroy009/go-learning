package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteDataToFile(fileName string, value float64) {
	dataText := fmt.Sprint(value)
	err := os.WriteFile(fileName, []byte(dataText), 0644)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func GetDataFromFile(fileName string) (float64, error){
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		//panic("Can't continue because file is not found.")
		return 100, errors.New("file not found")
	}

	balanceString := string(data)
	balance, err := strconv.ParseFloat(balanceString, 64)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		// fmt.Println(err)
		return 100, errors.New("failed to read data")
	}

	return balance, nil
}
