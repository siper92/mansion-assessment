package location

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"siper92/mansion-assessment/internal/cache"
	"siper92/mansion-assessment/internal/utils"

	PostCodesIo "github.com/DocHQ/go-postcodes"
)

type Location struct {
	Name      string  `json:"name"`
	PostCode  string  `json:"postcode"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

func (l Location) LoadFull() Location {
	if !utils.IsValidPostCode(l.PostCode) {
		log.Fatal(fmt.Sprintf("Invalid post code %s", l.PostCode))
	}

	locCache := cache.GetCacheContext(l.PostCode)
	if !locCache.IsValidCache() {
		post, err := PostCodesIo.Lookup(l.PostCode)
		if err != nil {
			log.Fatal(err)
		}

		if l.Name == "" {
			l.Name = post.Result.AdminDistrict
		}

		l.Latitude = post.Result.Latitude
		l.Longitude = post.Result.Longitude

		locCache = cache.Cache{
			CacheID:   cache.GetCacheId(l.PostCode),
			Data:      LocationStructToMap(l),
			ExpiresAt: cache.GetExpiresAfter(2 * 24),
		}
		locCache.SaveCacheContent()
	} else {
		l = MapToLocationStruct(locCache.Data)
	}

	return l
}

func (l Location) IsInRange(compLocation Location, radiusInKm int) bool {
	radiusInUnits := float64(1.0 / (111.0 / float64(radiusInKm)))

	return math.Pow((l.Latitude-compLocation.Latitude), 2)+
		math.Pow((l.Longitude-compLocation.Longitude), 2) <
		math.Pow(radiusInUnits, 2)
}

func MapToLocationStruct(v map[string]interface{}) Location {
	res := Location{}
	if jsonbody, err := json.Marshal(v); err != nil {
		log.Fatal(err)
	} else {
		if err := json.Unmarshal(jsonbody, &res); err != nil {
			log.Fatal(err)
		}
	}

	return res
}

func LocationStructToMap(v Location) map[string]interface{} {
	var inInterface map[string]interface{}
	tmpMarshal, _ := json.Marshal(v)
	json.Unmarshal(tmpMarshal, &inInterface)

	return inInterface
}

func GetStoredLocationsData() []Location {
	byteValue := utils.GetFileContent(utils.GetDataFilePath("locations.json"))

	var res []Location
	json.Unmarshal(byteValue, &res)

	return res
}
