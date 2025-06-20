package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type cache struct {
	entries map[string]cacheEntry
	mu      sync.Mutex
}

func (c *cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.entries[key] = cacheEntry{time.Now(), val}
	c.mu.Unlock()
}

func (c *cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if entry, ok := c.entries[key]; ok {
		return entry.val, true
	}
	return nil, false
}

func (c *cache) reapLoop(interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	t := <-ticker.C
	for entry := range c.entries {
		if t.After(c.entries[entry].createdAt.Add(interval)) {
			delete(c.entries, entry)
		}
	}
}

func NewCache(interval time.Duration) {

}
