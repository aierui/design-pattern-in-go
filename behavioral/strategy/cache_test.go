package strategy

import "testing"

func TestCache(t *testing.T) {
	lfu := &lfu{}
	cache := initCache(lfu)
	cache.add("a", "1")
	cache.add("b", "2")
	cache.add("c", "3")

	lru := &lru{}
	cache.setEvictionAlgo(lru)
	cache.add("d", "4")

	fifo := &fifo{}
	cache.setEvictionAlgo(fifo)
	cache.add("e", "5")

	/*
		Evicting by lfu strtegy
		Evicting by lru strtegy
		Evicting by fifo strtegy
	*/
}
