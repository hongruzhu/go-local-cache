package localcache

type Cache interface {
  Get(key string) (value any, ok bool)
  Set(key string, value any)
}
