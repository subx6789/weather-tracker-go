package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type ApiConfigData struct {
	OpenWeatherMapApiKey string `json:"openWeatherMapApiKey"`
}

type WeatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

func init() {
	// Loads values from .env into the system. If .env is not found, print a warning but continue.
	if err := godotenv.Load(); err != nil {
		log.Print("Warning: No .env file found")
	}
}

func main() {
	// Get the port from environment variables or use the default port 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not set
	}
	// Initialize the router
	r := mux.NewRouter()
	// Route for the root endpoint
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Weather Tracker Project in Golang\n")
	}).Methods(http.MethodGet)
	// Route for fetching weather data based on city
	r.HandleFunc("/weather/{city}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		city := params["city"] // Retrieval of the city from the route parameters
		// Query the weather API for the city
		data, err := query(city)
		if err != nil {
			// Send internal server error if the query fails
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Respond with JSON-encoded weather data
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(data)
	}).Methods(http.MethodGet)
	// Start the server
	fmt.Println("Server starting on port", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

// query fetches weather data for the specified city using OpenWeatherMap API
func query(city string) (WeatherData, error) {
	// Get the API key from the environment variables
	apiKey := os.Getenv("OPEN_WEATHER_MAP_API_KEY")
	if apiKey == "" {
		// Return error if API key is not set in environment variables
		return WeatherData{}, fmt.Errorf("missing OpenWeatherMap API key")
	}
	// Make an HTTP GET request to OpenWeatherMap API
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?appid=%s&q=%s", apiKey, city)
	resp, err := http.Get(url)
	if err != nil {
		// Return error if there was an issue with the API request
		return WeatherData{}, fmt.Errorf("failed to fetch weather data: %v", err)
	}
	defer resp.Body.Close()
	// Check if the response status code is OK (200)
	if resp.StatusCode != http.StatusOK {
		// Return an error if the status is not OK
		return WeatherData{}, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}
	// Decode the JSON response body into the WeatherData struct
	var d WeatherData
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		// Return error if there is an issue with decoding the response
		return WeatherData{}, fmt.Errorf("failed to decode weather data: %v", err)
	}
	// Return the decoded weather data
	return d, nil
}
