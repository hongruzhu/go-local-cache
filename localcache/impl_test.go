package localcache

import (
  "testing"
  "time"

  "github.com/stretchr/testify/assert"
)

func TestNewCache(t *testing.T) {
  cache := New()
  assert.NotNil(t, cache)
}

func TestGetCache(t *testing.T) {
  cache := New()
  cache.Set("key1", "value1")
  value, ok := cache.Get("key1")
  assert.Equal(t, "value1", value)
  assert.True(t, ok)
}

func TestSetCache(t *testing.T) {
  cache := New()
  cache.Set("key1", "value1")
  value, ok := cache.Get("key1")
  assert.Equal(t, "value1", value)
  assert.True(t, ok)
}

func TestOverwriteCache(t *testing.T) {
  cache := New()
  cache.Set("key1", "value1")
  cache.Set("key1", "value2")
  value, ok := cache.Get("key1")
  assert.Equal(t, "value2", value)
  assert.True(t, ok)
}

func TestCacheExpire(t *testing.T) {
  cache := New()
  cache.Set("key1", "value1")
  value, ok := cache.Get("key1")
  assert.Equal(t, "value1", value)
  assert.True(t, ok)

  time.Sleep(30 * time.Second)
  value, ok = cache.Get("key1")
  assert.Nil(t, value)
  assert.False(t, ok)
}