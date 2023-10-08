package outputlog

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/adenhidayatuloh/assignment3-011/dto"
)

// func untuk melakukan http request pada go, dan akan langsung menampilkan response , status water dan status wind
func OutputLog() {

	for {

		apiURL := "http://localhost:8080/waterwind/"

		// Membuat request dengan metode PUT tanpa data
		request, err := http.NewRequest("PUT", apiURL, nil)
		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}

		// Melakukan HTTP PUT request
		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer response.Body.Close()

		ApiResponse := dto.Response{}
		err = json.NewDecoder(response.Body).Decode(&ApiResponse)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}

		indentedJSON, err := json.MarshalIndent(ApiResponse, "", "  ")
		if err != nil {
			fmt.Println("Error marshaling to JSON:", err)
			return
		}

		statusWater := getStatusWater(ApiResponse.Water)
		statusWind := getStatusWind(ApiResponse.Wind)

		// Mencetak data

		fmt.Println(string(indentedJSON))
		fmt.Println("status water : " + statusWater)
		fmt.Println("status wind : " + statusWind)
		fmt.Println(" ")

		// tunggu 15 detik
		time.Sleep(15 * time.Second)

	}
}

// func untuk mendapatkan status water
func getStatusWater(water int) string {

	if water > 8 {
		return "bahaya"
	}
	if water >= 6 && water <= 8 {
		return "siaga"
	}
	return "aman"
}

// func mendapatkan status wind
func getStatusWind(wind int) string {

	if wind > 15 {
		return "bahaya"
	}
	if wind >= 7 && wind <= 15 {
		return "siaga"
	}
	return "aman"
}
