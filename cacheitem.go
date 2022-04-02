package cache_db

import (
	"sync"
)

type CacheItem struct {
	sync.RWMutex
	key  interface{}
	data interface{}
}

func NewCacheItem(key interface{}, data interface{}) *CacheItem {
	return &CacheItem{
		key:  key,
		data: data,
	}
}

func (item *CacheItem) KeepAlive() {
	item.Lock()
	defer item.Unlock()
}

func (item *CacheItem) Key() interface{} {
	return item.key
}

func (item *CacheItem) Data() interface{} {
	return item.data
}
