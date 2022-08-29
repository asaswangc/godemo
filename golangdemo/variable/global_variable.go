package variable

import (
	"sync"
)

// Global SyncMap 全局容器[这个主要存全局通用的变量]
var Global = syncMap{}

// SyncMap 全局容器
type syncMap struct {
	syncMap sync.Map
}

// Get 从全局容器中获取值
func (self *syncMap) Get(key string) interface{} {
	value, ok := self.syncMap.Load(key)
	if ok {
		return value
	}
	return nil
}

// Set 向全局容器中添加
func (self *syncMap) Set(key string, value interface{}) {
	self.syncMap.Store(key, value)
}
