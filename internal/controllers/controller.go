package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/faruqii/EAI/internal/response"
)

// SearchDestination is the controller for the search destination from consume airbnb api
func SearchDestination(location string, checkIn string, checkOut string, adults int, children int, infants int, currency string) {
	if children == 0 {
		children = 0
	}
	if infants == 0 {
		infants = 0
	}
	//default currency is USD
	if currency == "" {
		currency = "USD"
	}

	url := fmt.Sprintf("https://airbnb13.p.rapidapi.com/search-location?location=%s&checkin=%s&checkout=%s&adults=%d&children=%d&infants=%d&page=1&currency=%s", location, checkIn, checkOut, adults, children, infants, currency)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("X-RapidAPI-Key", os.Getenv("API_KEY"))
	req.Header.Add("X-RapidAPI-Host", "airbnb13.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	var response response.AirbnbResponse

	err = json.Unmarshal(body, &response)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("based on your search, we found this result")

	for _, result := range response.Results {
		fmt.Println("=======================================")
		fmt.Printf("Name: %s\n", result.Name)
		fmt.Printf("Address: %s\n", result.Address)
		fmt.Printf("City: %s\n", result.City)
		fmt.Printf("Neighborhood: %s\n", result.Neighborhood)
		fmt.Printf("Number of Bathrooms: %d\n", result.Bathrooms)
		fmt.Printf("Number of Beds: %d\n", result.Beds)
		fmt.Printf("Number of Reviews: %d\n", result.ReviewsCount)
		fmt.Printf("Rating: %.1f/5\n", result.Rating)
		fmt.Printf("Property Type: %s\n", result.Type)
		fmt.Printf("Currency: %s\n", result.Price.Currency)
		fmt.Printf("Total Price: %d\n", result.Price.Total)
		fmt.Println("Price Breakdown:")
		for _, item := range result.Price.PriceItems {
			fmt.Printf("- %s: %d\n", item.Title, item.Amount)
		}
		fmt.Println("Facilities:")
		for _, facility := range result.Facilities {
			fmt.Println("- ", facility)
		}
		fmt.Println("=======================================")
	}
}
