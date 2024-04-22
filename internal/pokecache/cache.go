package pokecache

import (
	"fmt"
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
	go cache.reapLoop()
	return cache, nil
}

func (c Cache) Add(key string, val []byte) error {
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

func (c Cache) reapLoop() error {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for {
		select {
		case t := <-ticker.C:
			for key, value := range c.Data {
				if (t.Sub(value.CreatedAt)) >= c.interval {
					c.mu.Lock()
					delete(c.Data, key)
					c.mu.Unlock()
				}

			}
		default:
			fmt.Println("Checking cache")
		}
	}
}
