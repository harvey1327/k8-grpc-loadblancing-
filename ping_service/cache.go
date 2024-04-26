package main

import (
	"fmt"
	"strings"
	"sync"
)

type Cache struct {
	internal map[string]int
	mutex    sync.Mutex
}

func NewCache() *Cache {
	return &Cache{
		internal: make(map[string]int),
	}
}

func (c *Cache) Increment(key string) {
	c.mutex.Lock()
	if v, ok := c.internal[key]; ok {
		c.internal[key] = v + 1
	} else {
		c.internal[key] = 1
	}
	c.mutex.Unlock()
}

func (c *Cache) Print() string {
	c.mutex.Lock()

	result := make([]string, 0)
	for k, v := range c.internal {
		c := fmt.Sprintf("%s:%d", k, v)
		result = append(result, c)
	}

	c.mutex.Unlock()
	return fmt.Sprintf("Key:Value -> [%s]", strings.Join(result, ","))
}
