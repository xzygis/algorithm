package _146_lru_cache

import (
	"testing"
)

func Test_LRUCache(t *testing.T) {
	lRUCache := Constructor(2)
	lRUCache.Put(1, 1) // 缓存是 {1=1}
	lRUCache.Info()
	lRUCache.Put(2, 2) // 缓存是 {1=1, 2=2}
	lRUCache.Info()

	// 返回 1
	if lRUCache.Get(1) != 1 {
		t.Fail()
	}
	lRUCache.Info()

	lRUCache.Put(3, 3) // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	lRUCache.Info()

	// 返回 -1 (未找到)
	if lRUCache.Get(2) != -1 {
		t.Fail()
	}
	lRUCache.Info()

	lRUCache.Put(4, 4) // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	lRUCache.Info()

	// 返回 -1 (未找到)
	if lRUCache.Get(1) != -1 {
		t.Fail()
	}
	lRUCache.Info()

	// 返回 3
	if lRUCache.Get(3) != 3 {
		t.Fail()
	}
	lRUCache.Info()
	if lRUCache.Get(4) != 4 {
		t.Fail()
	}
	lRUCache.Info()

}
