package main

import (
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
)

type Cache struct {
	internal *sync.Map
}

func NewCache() *Cache {
	return &Cache{
		internal: &sync.Map{},
	}
}

func (c *Cache) Increment(key string) {
	inc := int32(1)
	count, loaded := c.internal.LoadOrStore(key, inc)
	if loaded {
		atomic.AddInt32(count.(*int32), inc)
	}
}

func (c *Cache) Print() string {
	result := make([]string, 0)
	c.internal.Range(func(key, value any) bool {
		c := fmt.Sprintf("%s:%d", key, value)
		result = append(result, c)
		return true
	})
	return fmt.Sprintf("Key:Value -> [%s]", strings.Join(result, ","))
}
