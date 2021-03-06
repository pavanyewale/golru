package lru

import (
	"container/list"
	"errors"
)

//LRUInterface to ensure methods of LRU
type lruInterface interface {
	SetItem(key, value interface{})
	GetItem(key interface{}) (interface{}, error)
}

type item struct {
	key   interface{}
	value interface{}
}

//LRU ...
type LRU struct {
	keyIndexMap   map[interface{}]*list.Element
	size          int
	itemList      *list.List
	leastUsedItem *list.Element
	worstUsedItem *list.Element
	lruInterface
}

//NewLRU will return object of LRU
func NewLRU(size int) (*LRU, error) {
	if size < 2 {
		return nil, errors.New("size should be greater than 1")
	}
	return &LRU{
		size:          size,
		keyIndexMap:   make(map[interface{}]*list.Element, size),
		itemList:      list.New(),
		leastUsedItem: nil,
		worstUsedItem: nil,
	}, nil
}

//SetItem will set the item
func (lru *LRU) SetItem(key, value interface{}) {
	if lru.itemList.Len() == lru.size {
		tmp := lru.worstUsedItem
		lru.worstUsedItem = tmp.Prev()
		item := tmp.Value.(item)
		delete(lru.keyIndexMap, item.key)
		lru.itemList.Remove(tmp)
	}
	e := list.Element{Value: item{key: key, value: value}}
	lru.keyIndexMap[key] = &e
	lru.itemList.PushFront(&e)
}

//GetItem will get the item
func (lru *LRU) GetItem(key interface{}) (interface{}, error) {
	k := lru.keyIndexMap[key]
	if k == nil {
		return 0, errors.New("not found")
	}
	e := *k
	lru.itemList.Remove(k)
	lru.itemList.PushFront(&e)
	item := e.Value.(item)
	return item.value, nil
}
