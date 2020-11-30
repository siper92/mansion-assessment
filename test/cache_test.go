package test

import (
	"siper92/mansion-assessment/internal/cache"
	"siper92/mansion-assessment/internal/location"
	"testing"
	"time"
)

func TestGetCacheId(t *testing.T) {
	t.Run("currect_cache_id_format", func(t *testing.T) {
		id := cache.GetCacheId("tEst tEsT")

		if id != "test_test" {
			t.Errorf("Invalid cache ID")
		}
	})

	t.Run("currect_cache_id_format_2", func(t *testing.T) {
		id := cache.GetCacheId(" tEst   ")

		if id != "_test___" {
			t.Errorf("Invalid cache ID")
		}
	})
}

func TestGetExpiresAfter(t *testing.T) {
	sec := time.Now().Unix()
	hours := 4

	timeAfter4hours := int64(hours*60*60) + sec
	// add one second to compansate for any delay or test execution time
	timeAfter4hours = timeAfter4hours + 1

	expireTime := cache.GetExpiresAfter(4)
	expireInSec := expireTime.Unix()
	if expireInSec > timeAfter4hours {
		t.Errorf("GetExpiresAfter returns wrong time")
	}
}

func TestCacheStruct(t *testing.T) {
	sec := time.Now().Unix()
	hours := 4

	timeAfter4hours := int64(hours*60*60) + sec
	// add one second to compansate for any delay or test execution time
	timeAfter4hours = timeAfter4hours + 1

	expireTime := cache.GetExpiresAfter(4)
	expireInSec := expireTime.Unix()
	if expireInSec > timeAfter4hours {
		t.Errorf("GetExpiresAfter returns wrong time")
	}
}

func TestCacheValidation(t *testing.T) {
	data := location.LocationStructToMap(
		location.Location{
			Name:      "St_Albans",
			PostCode:  "AL1 2RJ",
			Latitude:  51.741753,
			Longitude: -0.341337,
		})

	t.Run("cache_is_valid", func(t *testing.T) {
		entity := cache.Cache{
			CacheID:   cache.GetCacheId("test a"),
			Data:      data,
			ExpiresAt: cache.GetExpiresAfter(1)}

		if !entity.IsValidCache() {
			t.Errorf("Cache is invalid when it should be valid")
		}
	})

	t.Run("cache_is_invalid_when_missing_id", func(t *testing.T) {
		entity := cache.Cache{
			Data:      data,
			ExpiresAt: cache.GetExpiresAfter(2 * 24)}

		if entity.IsValidCache() {
			t.Errorf("Cache is valid with missing CacheID")
		}
	})

	t.Run("cache_is_invalid_when_expired", func(t *testing.T) {
		entity := cache.Cache{
			CacheID:   cache.GetCacheId("test a"),
			Data:      data,
			ExpiresAt: time.Now().Add(-time.Duration(1))}

		if entity.IsValidCache() {
			t.Errorf("Cache is valid with ExpiresAt in the past")
		}
	})

	t.Run("cache_is_invalid_when_missing_ExpiresAt", func(t *testing.T) {
		entity := cache.Cache{
			CacheID: cache.GetCacheId("test a"),
			Data:    data}
		// ExpiresAt: time.Now().Add(-time.Duration(1))}

		if entity.IsValidCache() {
			t.Errorf("Cache is valid with missing ExpiresAt")
		}
	})
}
