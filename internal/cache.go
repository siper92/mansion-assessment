package cache

import (
	"fmt"
	"log"
	"os"
)

type Cache struct {
	cacheID string
}

func GetCacheContext(id string) Cache {

	return Cache{cacheID: id}
}

func getCacheFilePath(id string) string {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
		path = "."
	}

	return fmt.Sprintf("%s/var/cache/%s", path, id)
}
