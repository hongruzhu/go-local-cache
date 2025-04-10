package localcache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

func TestLocalCacheSuite(t *testing.T) {
	suite.Run(t, new(LocalCacheSuite))
}

type LocalCacheSuite struct {
	suite.Suite
	cache Cache
}

func (lcs *LocalCacheSuite) SetupSuite() {
	lcs.cache = New()
}

func (lcs *LocalCacheSuite) TestGet() {
	testCase := []struct {
		Desc      string
		key       string
		value     any
		ExpResult any
	}{
		{
			Desc:      "test get cache",
			key:       "key1",
			value:     "value1",
			ExpResult: "value1",
		},
	}

	for _, tc := range testCase {
		lcs.cache.Set(tc.key, tc.value)

		value, ok := lcs.cache.Get(tc.key)
		lcs.Equal(tc.ExpResult, value, tc.Desc)
		lcs.True(ok, tc.Desc)
	}
}

func (lcs *LocalCacheSuite) TestSet() {
	testCase := []struct {
		Desc      string
		key       string
		value     any
		ExpResult any
	}{
		{
			Desc:      "test set cache",
			key:       "key2",
			value:     "value2",
			ExpResult: "value2",
		},
		{
			Desc:      "test overwrite cache",
			key:       "key3",
			value:     "value3",
			ExpResult: "value3",
		},
	}

	for _, tc := range testCase {
		lcs.cache.Set(tc.key, tc.value)
		value, ok := lcs.cache.Get(tc.key)
		lcs.Equal(tc.ExpResult, value, tc.Desc)
		lcs.True(ok, tc.Desc)
	}
}

func (lcs *LocalCacheSuite) TestCacheExpire() {
	testCase := []struct {
		Desc     string
		key      string
		value    any
		duration time.Duration
	}{
		{
			Desc:     "test cache expire",
			key:      "key4",
			value:    "value4",
			duration: 3 * time.Second,
		},
	}

	for _, tc := range testCase {
		lcs.cache.Set(tc.key, tc.value, tc.duration)
		value, ok := lcs.cache.Get(tc.key)
		lcs.Equal(tc.value, value, tc.Desc)
		lcs.True(ok, tc.Desc)

		time.Sleep(tc.duration + 1*time.Millisecond)
		value, ok = lcs.cache.Get(tc.key)
		lcs.Equal(nil, value, tc.Desc)
		lcs.False(ok, tc.Desc)
	}
}

func (lcs *LocalCacheSuite) TestCacheTimerReset() {
	testCase := []struct {
		Desc     string
		key      string
		value    any
	}{
		{
			Desc:     "test cache timer reset",
			key:      "key5",
			value:    "value5",
		},
	}

	for _, tc := range testCase {
		lcs.cache.Set(tc.key, tc.value, 3*time.Second)
		lcs.cache.Set(tc.key, tc.value, 1*time.Second)

		time.Sleep(2 * time.Second)
		value, ok := lcs.cache.Get(tc.key)
		lcs.Equal(nil, value, tc.Desc)
		lcs.False(ok, tc.Desc)
	}
}