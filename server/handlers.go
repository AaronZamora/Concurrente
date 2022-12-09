package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
		return
	}

	fmt.Fprintf(w, "Hello World!")

}

func getCountries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "apllication/json")
	json.NewEncoder(w).Encode(countries)

}

func addCountry(w http.ResponseWriter, r *http.Request) {

	Country := &country{}

	err := json.NewDecoder(r.Body).Decode(Country)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}

	countries = append(countries, Country)
	fmt.Fprintf(w, "Country added successfully")
}
