package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"weatherapp/dto"
)

func main() {
	http.Handle("/weather", weatherHandler())
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })
	fmt.Println("Server is running and listening at port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func weatherHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//read the zipcode from query parameter
		zipCode := r.URL.Query().Get("zipcode")

		//build weather api request.
		url := "http://api.openweathermap.org/data/2.5/weather?zip=" + zipCode + ",us&appid=259f649c647791e124a3b150b8ae6cad&units=imperial" //hardcoded to serve only US zipcodes.

		//making http request
		response, err := http.Get(url)

		if err != nil {
			fmt.Printf("received bad response from weather api : %s", err)
			http.Error(w, err.Error(), 500)
			return
		}

		defer response.Body.Close()

		// read response
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("unable to read response body : %s", err)
			http.Error(w, err.Error(), 500)
			return
		}

		// parse response
		s, err := parseResponse(contents)

		if err != nil {
			fmt.Printf("unable to parse response json : %s", err)
			http.Error(w, err.Error(), 500)
			return
		}

		// write response
		fmt.Fprintf(w, "Current weather in zipcode %s is %f degree fahrenheit", zipCode, s.Main.Temp)
	})
}

// parseResponse function parses the JSON respose into weatherResponse struct
func parseResponse(body []byte) (*dto.WeatherResponse, error) {
	s := new(dto.WeatherResponse)
	err := json.Unmarshal(body, &s)
	if err != nil {
		return nil, err
	}
	return s, err

}
