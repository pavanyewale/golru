# LRU-in-golang
Implemented a LRU  (least recently used cache) in golang
Used map and list packages to implement this

exposed three methods 
func NewLRU(size int) (*LRU, error) 
func (lru *LRU)	SetItem(key, value interface{})
func (lru *LRU)	GetItem(key interface{}) (interface{}, error)

complexity O(1)
## how to use
    see main.go 
    