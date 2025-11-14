package main

import (
	"fmt"
	"sync"
)

type Cache interface {
    Set(k string, v string)
    Get(k string) (string, bool)
	Size() int
}

type InMemoryCache struct {
    data map[string]string
    mu sync.RWMutex
}

func NewInMemoryCache() *InMemoryCache {
    return &InMemoryCache{
        make(map[string]string),
        sync.RWMutex{},
    }
}

func (c *InMemoryCache) Set(k string, v string) {
    c.mu.Lock()
    defer c.mu.Unlock()

    c.data[k] = v
}

func (c *InMemoryCache) Get(k string) (string, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()

    data, ok := c.data[k]
    if ok {
        fmt.Printf("Cache hit: %s\n", data)
    }
    return data, ok
}

func (c *InMemoryCache) Size() int {
    c.mu.RLock()
    defer c.mu.RUnlock()
    return len(c.data)
}

func main() {
    cache := NewInMemoryCache()
    cache.Set("foo", "bar")
    cache.Set("baz", "qux")

    wg := &sync.WaitGroup{}
    wg.Add(4)

    go func() {
        defer wg.Done()
        cache.Set("foo", "upd_bar")
        fmt.Println("1")
    }()

    go func() {
        defer wg.Done()
        cache.Set("baz", "upd_qux")
        fmt.Println("2")
    }()

    go func() {
        defer wg.Done()
        cache.Get("foo")
        fmt.Println("4")
    }()

    go func() {
        defer wg.Done()
        cache.Get("bar")
        fmt.Println("4")
    }()

    wg.Wait()

    data, ok := cache.Get("foo")
    if !ok {
		fmt.Println("Key: Not found")
        return
    }
	fmt.Println("Key: foo, Value: ", data)
	fmt.Println("Data size is:", cache.Size())
}
