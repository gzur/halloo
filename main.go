package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// var API_KEY string

func main() {
	http.HandleFunc("/hello", halloo)

	http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
		city := strings.SplitN(r.URL.Path, "/", 3)[2]
		data, err := query(city)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(data)
	})
	port := ":8080"
	log.Printf("Up and listening on port %s", port)
	http.ListenAndServe(port, nil)
}

func halloo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("halloo!"))
}

func query(city string) (weatherData, error) {
	api_key := "6bd299bd050124d5988440b62aa343a0"
	endpoint := "http://api.openweathermap.org/data/2.5/weather?q=%s&APPID=%s"
	url := fmt.Sprintf(endpoint, city, api_key)
	log.Printf("Querying %s", url)
	resp, err := http.Get(url + city)
	if err != nil {
		return weatherData{}, err
	}

	defer resp.Body.Close()

	var d weatherData

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}
	// log.Printf(json.NewDecoder(resp.Body).Decode(&d))
	return d, nil
}

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}
