package test

import (
	"reflect"
	"siper92/mansion-assessment/internal/location"
	"testing"
)

func TestLocationIsInRange(t *testing.T) {
	t.Run("the_same_location", func(t *testing.T) {
		loc1 := location.Location{
			Name:      "Hatfield",
			PostCode:  "AL9 5JP",
			Latitude:  51.776142,
			Longitude: -0.222034,
		}

		loc2 := location.Location{
			Name:      "Hatfield",
			PostCode:  "AL9 5JP",
			Latitude:  51.776142,
			Longitude: -0.222034,
		}

		if !loc1.IsInRange(loc2, 30) {
			t.Errorf("Invalid range comparison, locations should be in range")
		}
	})

	t.Run("in_radius", func(t *testing.T) {
		loc1 := location.Location{
			Name:      "Hatfield",
			PostCode:  "AL9 5JP",
			Latitude:  51.776142,
			Longitude: -0.222034,
		}

		loc2 := location.Location{
			Name:      "St Albans",
			PostCode:  "AL4 8TJ",
			Latitude:  51.806262,
			Longitude: -0.290208,
		}

		if !loc1.IsInRange(loc2, 10) {
			t.Errorf("Invalid range comparison, locations should be in range")
		}
	})

	t.Run("barely_in_radius", func(t *testing.T) {
		loc1 := location.Location{
			Name:      "Hatfield",
			PostCode:  "AL9 5JP",
			Latitude:  51.776142,
			Longitude: -0.222034,
		}

		loc2 := location.Location{
			// (1 degree 111 km)
			Name:      "One degree - little bite less",
			PostCode:  "AL4 8TJ",
			Latitude:  52.776042,
			Longitude: -0.222034,
		}

		if !loc1.IsInRange(loc2, 111) {
			t.Errorf("Invalid range comparison, locations should be in range")
		}
	})

	t.Run("out_of_range_not_in_radius", func(t *testing.T) {
		loc1 := location.Location{
			Name:      "Hatfield",
			PostCode:  "AL9 5JP",
			Latitude:  51.776142,
			Longitude: -0.222034,
		}

		loc2 := location.Location{
			Name:      "St Albans",
			PostCode:  "AL4 8TJ",
			Latitude:  51.806262,
			Longitude: -0.290208,
		}

		if loc1.IsInRange(loc2, 1) {
			t.Errorf("Invalid range comparison, locations should NOT be in range")
		}
	})

	t.Run("barely_out_of_range_not_in_radius", func(t *testing.T) {
		loc1 := location.Location{
			Name:      "Hatfield",
			PostCode:  "AL9 5JP",
			Latitude:  51.776142,
			Longitude: -0.222034,
		}

		loc2 := location.Location{
			// (1 degree 111 km)
			Name:      "One degree + little bite more",
			PostCode:  "AL4 8TJ",
			Latitude:  52.776242,
			Longitude: -0.222034,
		}

		if loc1.IsInRange(loc2, 111) {
			t.Errorf("Invalid range comparison, locations should NOT be in range")
		}
	})
}

func TestMapToLocationStruct(t *testing.T) {
	mapData := map[string]interface{}{
		"latitude":  51.776142,
		"longitude": -0.222034,
		"name":      "Hatfield",
		"postcode":  "AL9 5JP",
	}

	loc := location.MapToLocationStruct(mapData)
	if loc.Latitude != 51.776142 ||
		loc.Longitude != -0.222034 ||
		loc.Name != "Hatfield" ||
		loc.PostCode != "AL9 5JP" {
		t.Errorf("Cannot convert the Map to Location Struct")
	}
}

func TestLocationStructToMap(t *testing.T) {
	mapData := map[string]interface{}{
		"latitude":  51.776142,
		"longitude": -0.222034,
		"name":      "Hatfield",
		"postcode":  "AL9 5JP",
	}

	locMap := location.LocationStructToMap(location.Location{
		Name:      "Hatfield",
		PostCode:  "AL9 5JP",
		Latitude:  51.776142,
		Longitude: -0.222034,
	})

	if !reflect.DeepEqual(mapData, locMap) {
		t.Errorf("Cannot convert the Location Struct to Map")
	}
}
