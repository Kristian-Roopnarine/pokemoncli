package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

type Cache struct {
	Data     map[string]cacheEntry
	interval time.Duration
	mu       *sync.RWMutex
}

// should this return the pointer to the new item?
func NewCache(interval time.Duration) (Cache, error) {
	cache := Cache{Data: make(map[string]cacheEntry), interval: interval, mu: &sync.RWMutex{}}
	go cache.reapLoop(interval)
	return cache, nil
}

func (c Cache) Add(key string, val []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Data[key] = cacheEntry{
		CreatedAt: time.Now(),
		Val:       val,
	}
	return nil
}

func (c Cache) Get(key string) ([]byte, bool, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, ok := c.Data[key]
	if ok {
		return entry.Val, true, nil
	}
	// should this return error or nil?
	return nil, false, nil
}

func (c Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.Data {
		if v.CreatedAt.Before(now.Add(-last)) {
			delete(c.Data, k)
		}
	}
}
