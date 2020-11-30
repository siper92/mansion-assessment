package handler

import (
	"encoding/json"
	"net/http"

	"siper92/mansion-assessment/internal/location"
	"siper92/mansion-assessment/internal/utils"
)

var cachedLocations []location.Location

func InitCache() {
	for _, loc := range location.GetStoredLocationsData() {
		cachedLocations = append(cachedLocations, loc.LoadFull())
	}
}

type LocationFilter struct {
	PostCode string `json:"postcode"`
	Radius   int    `json:"radius"`
}

func FilterNearbyLocations(w http.ResponseWriter, r *http.Request) {
	var filter LocationFilter
	err := json.NewDecoder(r.Body).Decode(&filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !utils.IsValidPostCode(filter.PostCode) {
		http.Error(w, "Invalid post code format", http.StatusBadRequest)
		return
	}

	if filter.Radius < 1 {
		http.Error(w, "Invalid radius", http.StatusBadRequest)
		return
	}

	filterLocation := location.Location{PostCode: filter.PostCode}
	filterLocation = filterLocation.LoadFull()

	var res []location.Location

	if len(cachedLocations) < 1 {
		InitCache()
	}

	for _, loc := range cachedLocations {
		if filterLocation.IsInRange(loc, filter.Radius) {
			res = append(res, loc)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
