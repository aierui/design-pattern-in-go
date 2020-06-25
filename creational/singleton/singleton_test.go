package singleton

import (
	"sync"
	"testing"
)

func TestSingleton(t *testing.T) {
	bar := getInstance()
	baz := getInstance()

	if bar != baz {
		t.Fatal("getInstance is not equal")
	}
}

func TestParalleSingleton(t *testing.T) {
	wg := sync.WaitGroup{}
	num := 100
	instances := make([]*single, num)
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func(idx int) {
			instances[idx] = getInstance()
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 1; i < num; i++ {
		if instances[i-1] != instances[i] {
			t.Fatal("instance is not equal")
		}
	}
}
