package gocache

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Test_get(t *testing.T) {
	NewCache(5)
	Set("key1", "value1")
	fmt.Printf("cache: %#v\n\n\n", _cache)
	cache := Cache()
	fmt.Printf("compare: %#v\n", _cache == cache)
	time.Sleep(6 * time.Second)
	fmt.Printf("cache: %#v\n\n\n", cache)
}

func Test_set(t *testing.T) {
	NewCache(10)
	Set("key1", "value1")
	Set("key2", "value2")
	for i := 0; i < 10; i++ {
		go Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}
	for i := 0; i < 10; i++ {
		go Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}
	item := Cache()
	fmt.Printf("cache: %#v\n", item)
	time.Sleep(time.Second)
}

func Test_clear(t *testing.T) {
	NewCache(10)
	for i := 0; i < 10; i++ {
		Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}
	time.Sleep(time.Second)
	Clear()
	item := Cache()
	fmt.Printf("c item: %#v\n", item)
}

func Benchmark_set(b *testing.B) {
	NewCache(10)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m := rand.Intn(100) + 1
			n := rand.Intn(m)
			Set(fmt.Sprintf("%d", m), fmt.Sprintf("%d", n))
			Get(fmt.Sprintf("%d", m))
		}
	})
}

func Benchmark_clear(b *testing.B) {
	NewCache(10)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m := rand.Intn(100) + 1
			n := rand.Intn(m)
			Set(fmt.Sprintf("%d", m), fmt.Sprintf("%d", n))
			Clear()
		}
	})
	item := Cache()
	fmt.Printf("c item: %#v\n", item)
}

func Benchmark_delete(b *testing.B) {
	NewCache(10)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m := rand.Intn(100) + 1
			n := rand.Intn(m)
			Set(fmt.Sprintf("%d", m), fmt.Sprintf("%d", n))
			Delete(fmt.Sprintf("%d", m))
		}
	})
	item := Cache()
	fmt.Printf("c item: %#v\n", item)
}

func Benchmark_len(b *testing.B) {
	NewCache(10)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m := rand.Intn(100) + 1
			n := rand.Intn(m)
			Set(fmt.Sprintf("%d", m), fmt.Sprintf("%d", n))
			l := Len()
			fmt.Printf("l: %d\n", l)
		}
	})
	item := Cache()
	fmt.Printf("item: %#v\n", item)
}
