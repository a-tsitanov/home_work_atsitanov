package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (c *lruCache) Set(key Key, value interface{}) bool {

	if el, ok := c.items[key]; ok {
		el.Value = &cacheItem{key, value}
		c.queue.MoveToFront(el)
		return true
	}

	if c.queue.Len() == c.capacity {
		var backItem = c.queue.Back()
		c.queue.Remove(backItem)
		delete(c.items, backItem.Value.(*cacheItem).key)
	}

	var cacheItem = &cacheItem{key, value}
	c.items[key] = c.queue.PushFront(cacheItem)
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if el, ok := c.items[key]; ok {
		c.queue.MoveToFront(el)
		return el.Value.(*cacheItem).value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem)
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
