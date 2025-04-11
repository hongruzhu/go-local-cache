package localcache

import "time"

type Cache interface {
	Get(key string) (value any, ok bool)
	Set(key string, value any, duration *time.Duration)
}
