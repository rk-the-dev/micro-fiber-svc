package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

// CacheHelper struct wraps the go-cache object
type CacheHelper struct {
	c *cache.Cache
}

// NewCacheHelper initializes a new CacheHelper
func NewCacheHelper(defaultExpiration, cleanupInterval time.Duration) *CacheHelper {
	return &CacheHelper{
		c: cache.New(defaultExpiration, cleanupInterval),
	}
}

// Set adds an item to the cache, replacing any existing item.
func (ch *CacheHelper) Set(key string, value interface{}, duration time.Duration) {
	ch.c.Set(key, value, duration)
}

// Get retrieves an item from the cache. If the item is not found, found is false.
func (ch *CacheHelper) Get(key string) (interface{}, bool) {
	return ch.c.Get(key)
}

// Delete removes an item from the cache. Does nothing if the key is not in the cache.
func (ch *CacheHelper) Delete(key string) {
	ch.c.Delete(key)
}
