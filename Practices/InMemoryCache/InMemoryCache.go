package main

import (
	"fmt"
	"sync"
	"time"
)

// InMemoryCache entry
type InMemoryCache struct {
	Key   string
	Field string
	Value string
	TTL   int64 // expiration time in Unix ms
}

// CacheManager manages all cache entries
type CacheManager struct {
	cache map[string]*InMemoryCache
	mu    sync.RWMutex
}

// Create new CacheManager
func NewCacheManager() *CacheManager {
	return &CacheManager{
		cache: make(map[string]*InMemoryCache),
	}
}

// Composite key
func (cm *CacheManager) compositeKey(key, field string) string {
	return key + "::" + field
}

// Check if entry is expired
func (entry *InMemoryCache) isExpired() bool {
	return time.Now().UnixMilli() > entry.TTL
}

// Clear all cache
func (cm *CacheManager) Init() {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.cache = make(map[string]*InMemoryCache)
}

// Add entry
func (cm *CacheManager) Add(key, field, value string, ttlMillis int64) interface{} {
	if key == "" || field == "" || value == "" {
		return "Error: Key, field, and value must not be empty"
	}

	cm.mu.Lock()
	defer cm.mu.Unlock()

	composite := cm.compositeKey(key, field)
	existing, ok := cm.cache[composite]
	if ok && !existing.isExpired() {
		return false // already exists and not expired
	}

	expiration := int64(^uint64(0) >> 1) // max int64 (simulate Long.MAX_VALUE)
	if ttlMillis > 0 {
		expiration = time.Now().UnixMilli() + ttlMillis
	}

	entry := &InMemoryCache{
		Key:   key,
		Field: field,
		Value: value,
		TTL:   expiration,
	}
	cm.cache[composite] = entry
	return entry
}

// Remove entry
func (cm *CacheManager) Remove(key, field string) interface{} {
	if key == "" || field == "" {
		return "Error: Key and field must not be empty"
	}

	cm.mu.Lock()
	defer cm.mu.Unlock()

	composite := cm.compositeKey(key, field)
	existing, ok := cm.cache[composite]
	if !ok || existing.isExpired() {
		delete(cm.cache, composite)
		return false
	}
	delete(cm.cache, composite)
	return existing
}

// Update entry
func (cm *CacheManager) Update(key, field, value string, ttlMillis int64) interface{} {
	if key == "" || field == "" || value == "" {
		return "Error: Key, field, and value must not be empty"
	}

	cm.mu.Lock()
	defer cm.mu.Unlock()

	composite := cm.compositeKey(key, field)
	existing, ok := cm.cache[composite]
	if !ok || existing.isExpired() {
		delete(cm.cache, composite)
		return false
	}

	expiration := int64(^uint64(0) >> 1)
	if ttlMillis > 0 {
		expiration = time.Now().UnixMilli() + ttlMillis
	}
	existing.Value = value
	existing.TTL = expiration
	return existing
}

// Get entry
func (cm *CacheManager) Get(key, field string) interface{} {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	composite := cm.compositeKey(key, field)
	existing, ok := cm.cache[composite]
	if !ok || existing.isExpired() {
		// cleanup expired entry
		cm.mu.RUnlock()
		cm.mu.Lock()
		delete(cm.cache, composite)
		cm.mu.Unlock()
		cm.mu.RLock()
		return false
	}
	return existing
}

func main() {
	cache := NewCacheManager()

	fmt.Println(cache.Add("user1", "name", "Alice", 2000)) // add entry
	time.Sleep(1 * time.Second)

	fmt.Println(cache.Get("user1", "name")) // not expired yet
	time.Sleep(2 * time.Second)

	fmt.Println(cache.Get("user1", "name")) // expired -> false
}
