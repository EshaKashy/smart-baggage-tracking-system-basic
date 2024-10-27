package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Luggage struct {
	ID       string `json:"id"`
	Owner    string `json:"owner"`
	Status   string `json:"status"`
	Location string `json:"location"`
}

var luggageData []Luggage

func loadLuggageData() {
	file, err := os.Open("data/luggage.json")
	if err != nil {
		log.Fatalf("Error opening luggage data file: %v", err)
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	json.Unmarshal(byteValue, &luggageData)
}

func main() {
	loadLuggageData()

	// Display the luggage data in the terminal
	for _, luggage := range luggageData {
		fmt.Printf("Luggage ID: %s\nOwner: %s\nStatus: %s\nLocation: %s\n\n", luggage.ID, luggage.Owner, luggage.Status, luggage.Location)
	}

	http.HandleFunc("/luggage", func(w http.ResponseWriter, r *http.Request) {
		// Enable CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		// Send luggage data as JSON
		json.NewEncoder(w).Encode(luggageData)
	})

	fmt.Println("Server is running on http://localhost:8080/luggage")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
