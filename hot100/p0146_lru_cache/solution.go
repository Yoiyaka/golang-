package p0146_lru_cache

import "container/list"

type LRUCache struct {
	cap   int
	cache map[int]*list.Element
	list  *list.List
}

type entry struct {
	key, val int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:   capacity,
		cache: make(map[int]*list.Element),
		list:  list.New(),
	}
}

func (c *LRUCache) Get(key int) int {
	if el, ok := c.cache[key]; ok {
		c.list.MoveToFront(el)
		return el.Value.(*entry).val
	}
	return -1
}

func (c *LRUCache) Put(key, value int) {
	if el, ok := c.cache[key]; ok {
		el.Value.(*entry).val = value
		c.list.MoveToFront(el)
		return
	}
	if c.list.Len() == c.cap {
		back := c.list.Back()
		c.list.Remove(back)
		delete(c.cache, back.Value.(*entry).key)
	}
	el := c.list.PushFront(&entry{key, value})
	c.cache[key] = el
}
