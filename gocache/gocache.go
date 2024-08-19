package gocache

import (
	"sync"
	"time"
)

var (
	// default cache type
	_cache *cache
	// default cache expired time
	// _duration = 5 * time.Second
)

type cache struct {
	item                 map[string]item
	mutex                sync.RWMutex
	cleanupExpiredPeriod time.Duration
	// add file or sql
}

type item struct {
	value   interface{}
	expired time.Time
}

// newcache single mod
func NewCache(t time.Duration) {
	_cache = &cache{
		item:                 make(map[string]item),
		cleanupExpiredPeriod: t,
	}
	go cacheCleanUpExpired(_cache)
}

func Get(key string) interface{} {
	return _cache.get(key)
}

func Set(key string, value interface{}) {
	_cache.set(key, value)
}

func Clear() {
	_cache.clear()
}

func Delete(key string) {
	_cache.delete(key)
}

func Len() int {
	return _cache.len()
}

func AddExpired(t time.Duration) {
	_cache.addExpired(t)
}

func Cache() *cache {
	return _cache
}

// get cache value
//
//	c.get(key)
func (c *cache) get(key string) interface{} {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	item, ok := c.item[key]
	if !ok {
		return nil
	}
	return item
}

// set cache key and value
//
//	c.set(key, value)
func (c *cache) set(key string, value interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.item[key] = item{
		value:   value,
		expired: time.Now().Add(c.cleanupExpiredPeriod),
	}
}

// clear cache
func (c *cache) clear() {
	c.item = make(map[string]item)
}

// delete key
//
//	c.delete(key)
func (c *cache) delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	_, ok := c.item[key]
	if ok {
		delete(c.item, key)
	}
}

// modify expired period
func (c *cache) addExpired(t time.Duration) {
	c.cleanupExpiredPeriod = t
	for _, v := range c.item {
		v.expired = time.Now().Add(c.cleanupExpiredPeriod)
	}
}

// get cache length
func (c *cache) len() int {
	return len(c.item)
}

func cacheCleanUpExpired(cache *cache) {
	for {
		<-time.After(cache.cleanupExpiredPeriod)
		cache.mutex.Lock()
		now := time.Now()
		for k, v := range cache.item {
			if now.After(v.expired) {
				delete(cache.item, k)
			}
		}
		cache.mutex.Unlock()
	}
}
