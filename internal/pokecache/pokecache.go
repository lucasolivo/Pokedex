package pokecache

import (
	"time"
	"sync"
)

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	entries map[string]cacheEntry
	mu sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	c := Cache {
		entries: make(map[string]cacheEntry),
	}
	// Start the reapLoop in a separate goroutine
    go c.reapLoop(interval)

    return c

}

func (c *Cache) Add(key string, val []byte) {
    c.mu.Lock()                // Lock the mutex to protect the map
    defer c.mu.Unlock()        // Ensure it always unlocks, even if an error occurs
    c.entries[key] = cacheEntry{
        createdAt: time.Now(), // Record the creation time
        val:       val,        // Store the value
    }
}

func (c *Cache) Get(key string) ([]byte, bool) {
    c.mu.Lock()              // Lock the mutex to ensure safe access
    defer c.mu.Unlock()      // Defer unlocking to safely release the lock after execution
	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval) // Create a ticker that ticks at the given interval
    defer ticker.Stop()                // Ensure the ticker is stopped when the loop ends

    for {
        <-ticker.C // Wait for the next tick

        curTime := time.Now() // Get the current time

        c.mu.Lock()           // Lock the map during modification
        for key, entry := range c.entries {
            if curTime.Sub(entry.createdAt) > interval { // Check entry age
                delete(c.entries, key) // Remove old entries
            }
        }
        c.mu.Unlock()         // Unlock the map after the operation
    }
}