package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mu *sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) *Cache {
	output := Cache{
		cache: make(map[string]cacheEntry),
		mu: &sync.Mutex{},
		interval: interval,
	}
	go output.reapLoop()
	return &output
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val: value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, exists := c.cache[key]
	if !exists {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop() {
	for {
		<- time.NewTicker(c.interval).C
		func() {
			c.mu.Lock()
			defer c.mu.Unlock()
			for key, entry := range c.cache {
				if time.Since(entry.createdAt) > c.interval {
					delete(c.cache, key)
				}
			}
		}()
	}
}
