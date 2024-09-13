# 🌦️ Weather Tracker Project

This is a simple Weather Tracker project built in Golang using the OpenWeatherMap API and the Gorilla Mux router. The application fetches real-time weather data for a specified city and returns it as a JSON response.

## Features

- Fetch real-time weather information for any city using the OpenWeatherMap API.
- Display temperature in Kelvin.
- Error handling for missing API keys, incorrect city names, or API failures.
- Lightweight and efficient with a simple server setup.
- Organized with proper error handling and comments for future development.

## Prerequisites

- Go 1.16 or higher installed.
- An OpenWeatherMap API key (free version available at OpenWeatherMap).
- A .env file containing your OpenWeatherMap API key.

## Installation

- Clone the repository:

  ```bash
   git clone https://github.com/subx6789/weather-tracker-go.git
   cd weather-tracker-go
  ```

- Create a .env file in the root of the project and add your OpenWeatherMap API key:

  ```bash
   OPEN_WEATHER_MAP_API_KEY=your_api_key_here
  ```

- Install dependencies:

  ```bash
   go mod tidy
  ```

- Run the server:

  ```bash
   go run main.go
  ```

The server will start on port 8080 by default.

## API Endpoints

- GET `/`

Description: Root endpoint that welcomes the user.

Response:

```bash
 Welcome to Weather Tracker Project in Golang
```

- GET `/weather/{city}`

Description: Fetches weather information for the specified city.

Parameter city: The name of the city for which you want to fetch weather data.

Response: JSON object with the weather information of the city.

```bash
 {
     "name": "London",
     "main": {
         "temp": 289.67
     }
 }
```

## Project Structure

```bash
 ├── main.go         # Entry point of the application
 ├── .env            # Environment variables file (ignored by Git)
 ├── go.mod          # Go module file
 └── go.sum          # Go module dependencies file
```

## Built With

- **Golang** - The programming language used.
- **Gorilla Mux** - A powerful URL router and dispatcher for Golang.
- **OpenWeatherMap API** - API for accessing real-time weather data.

## Contributing

Feel free to contribute by submitting a pull request! Please ensure that your changes are well-documented.

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.

## Acknowledgements

- **OpenWeatherMap** for providing a free weather API.
- **Gorilla Mux** for the robust routing library.
- **godotenv** for easy environment variable management.

## Project Deployed with Render

- **Link:** [Weather-Tracker-Go](https://weather-tracker-go.onrender.com)
