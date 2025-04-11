/*
Package localcache implements a simple local cache.
*/
package localcache

import (
	"sync"
	"time"
)

const defaultDuration = 30 * time.Second

type localCache struct {
	mu    sync.RWMutex
	items map[string]*cacheItem
}

type cacheItem struct {
	val   any
	timer *time.Timer
}

// New creates a new cache.
func New() Cache {
	return &localCache{
		items: make(map[string]*cacheItem),
	}
}

// Get gets the value of the key.
func (c *localCache) Get(key string) (value any, ok bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, ok := c.items[key]
	if !ok {
		return nil, false
	}
	return item.val, true
}

// Set sets the value of the key.
func (c *localCache) Set(key string, value any, duration *time.Duration) {
	expiry := defaultDuration
	if duration != nil {
		expiry = *duration
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if oldItem, exists := c.items[key]; exists {
		oldItem.timer.Stop()
	}

	c.items[key] = &cacheItem{
		val: value,
		timer: time.AfterFunc(expiry, func() {
			c.mu.Lock()
			defer c.mu.Unlock()

			delete(c.items, key)
		}),
	}
}
