package cache

import "time"

type Cache struct {
	store map[string]string
}

func NewCache() Cache {
	return Cache{store: make(map[string]string)}
}

func (c *Cache) Get(key string) (string, bool) {
	value, found := c.store[key]
	return value, found

}

func (c *Cache) Put(key, value string) {
	c.store[key] = value
}

func (c Cache) Keys() []string {
	var keys []string
	for key := range c.store {
		keys = append(keys, key)
	}
	return keys
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
}
