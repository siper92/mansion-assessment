package cache

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"siper92/mansion-assessment/internal/utils"
)

type Cache struct {
	CacheID   string                 `json:"cache_id"`
	Data      map[string]interface{} `json:"data"`
	ExpiresAt time.Time              `json:"expires_at"`
}

func (c Cache) IsValidCache() bool {
	now := time.Now()
	return len(c.CacheID) > 3 && c.ExpiresAt.After(now)
}

func (c Cache) SaveCacheContent() bool {
	if c.CacheID == "" {
		log.Fatal("Cannot save cache with no ID")
	}

	cacheFilePath := utils.GetCacheFilePath(c.CacheID)
	if utils.FileExists(cacheFilePath) {
		err := os.Remove(cacheFilePath)
		if err != nil {
			log.Fatal(err)
		}
	}

	cacheDir := filepath.Dir(cacheFilePath)
	if !utils.DirExists(cacheDir) {
		os.MkdirAll(cacheDir, os.ModePerm)
	}

	file, err := os.Create(cacheFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rawData, _ := json.Marshal(c)
	file.Write(rawData)

	err = file.Sync()
	if err != nil {
		log.Fatal(err)
	}

	return true
}

// get the cache content, return empty cache on missing file
func GetCacheContext(rawId string) Cache {
	id := GetCacheId(rawId)
	cacheFilePath := utils.GetCacheFilePath(id)
	var byteValue []byte
	if utils.FileExists(cacheFilePath) {
		byteValue = utils.GetFileContent(utils.GetCacheFilePath(id))
	} else {
		byteValue = []byte{}
	}

	var res Cache
	json.Unmarshal([]byte(byteValue), &res)
	return res
}

func GetCacheId(id string) string {
	return strings.ReplaceAll(strings.ToLower(id), " ", "_")
}

func GetExpiresAfter(hours int) time.Time {
	return time.Now().Local().Add(time.Hour * time.Duration(hours))
}
