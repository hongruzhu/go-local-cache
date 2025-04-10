/*
Package localcache implements a simple local cache.
*/
package localcache

import "time"

const expireTime = 30 * time.Second

type cache struct{
  items map[string]any
}

// New creates a new cache.
func New() Cache {
  return &cache{
    items: make(map[string]any),
  }
}

// Get gets the value of the key.
func (c *cache) Get(key string) (value any, ok bool) {
  value, ok = c.items[key]
  return
}

// Set sets the value of the key.
func (c *cache) Set(key string, value any) {
  c.items[key] = value
  go func() {
    time.Sleep(expireTime)
    delete(c.items, key)
  }()
}
