package cache

import (
	"fmt"
	"github.com/ihangsen/common/src/collection/dict"
	"github.com/ihangsen/common/src/collection/tuple"
	"github.com/ihangsen/common/src/collection/vec"
	"github.com/ihangsen/common/src/coroutine"
	"github.com/ihangsen/common/src/utils/option"
	"github.com/ihangsen/common/src/utils/pile"
	"github.com/ihangsen/common/src/utils/pool"
	"math"
	"sync"
	"time"
)

type Cache[T comparable, V any] struct {
	*pin[T, V]
}

type pin[T comparable, V any] struct {
	data map[T]*Item[T, V]
	pool pool.Pool[*Item[T, V]]
	rw   sync.RWMutex
	heap *LfuHeap[T, V]
	cap  int
}

type Item[T comparable, V any] struct {
	key    T
	value  V
	expire int64
	count  int
	index  int
}

type LfuHeap[T comparable, V any] []*Item[T, V]

func (h LfuHeap[T, V]) Len() int {
	return len(h)
}

func (h LfuHeap[T, V]) Less(i, j int) bool {
	return h[i].count < h[j].count
}

func (h LfuHeap[T, V]) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *LfuHeap[T, V]) Push(t *Item[T, V]) {
	t.index = len(*h)
	*h = append(*h, t)
}

func (h *LfuHeap[T, V]) Pop() *Item[T, V] {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func New[T comparable, V any](cap int) Cache[T, V] {
	h := LfuHeap[T, V]{}
	pile.Init(&h)
	return Cache[T, V]{&pin[T, V]{
		data: make(map[T]*Item[T, V], cap),
		pool: pool.New[*Item[T, V]](func() *Item[T, V] {
			return new(Item[T, V])
		}),
		heap: &h,
		cap:  cap,
	}}
}

func New0[T comparable, V any]() Cache[T, V] {
	h := LfuHeap[T, V]{}
	pile.Init(&h)
	return Cache[T, V]{&pin[T, V]{
		data: make(map[T]*Item[T, V], 0),
		pool: pool.New[*Item[T, V]](func() *Item[T, V] {
			return new(Item[T, V])
		}),
		heap: &h,
		cap:  math.MaxInt,
	}}
}

func (c *Cache[T, V]) Set(key T, value V, expire int64) {
	c.rw.Lock()
	defer c.rw.Unlock()
	if len(c.data) >= c.cap {
		leastUsed := pile.Pop(c.heap)
		delete(c.data, leastUsed.key)
	}
	temp := c.pool.Get()
	temp.value = value
	temp.key = key
	if expire >= 0 {
		temp.expire = time.Now().UnixMilli() + expire
	} else {
		temp.expire = expire
	}
	if old, ok := c.data[key]; ok {
		old.value = value
		old.expire = temp.expire
		c.pool.Put(temp)
	} else {
		c.data[key] = temp
		pile.Push(c.heap, temp)
	}
}

func (c *Cache[T, V]) Sets(t2 vec.Vec[tuple.T2[T, V]], expire int64) {
	c.rw.Lock()
	defer c.rw.Unlock()
	length := len(c.data) + t2.Len() - c.cap
	if length > 0 {
		for i := 0; i < length; i++ {
			leastUsed := pile.Pop(c.heap)
			delete(c.data, leastUsed.key)
		}
	}
	now := time.Now().UnixMilli()
	t2.ForEach(func(t tuple.T2[T, V]) {
		temp := c.pool.Get()
		temp.value = t.V1
		temp.key = t.V0
		if expire >= 0 {
			temp.expire = now + expire
		} else {
			temp.expire = expire
		}
		if old, ok := c.data[t.V0]; ok {
			old.value = t.V1
			old.expire = temp.expire
			c.pool.Put(temp)
		} else {
			c.data[t.V0] = temp
			pile.Push(c.heap, temp)
		}
	})
}

func (c *Cache[T, V]) Sets0(d dict.Dict[T, V], expire int64) {
	c.rw.Lock()
	defer c.rw.Unlock()
	length := len(c.data) + d.Len() - c.cap
	if length > 0 {
		for i := 0; i < length; i++ {
			leastUsed := pile.Pop(c.heap)
			delete(c.data, leastUsed.key)
		}
	}
	now := time.Now().UnixMilli()
	d.ForEach(func(t T, v V) {
		temp := c.pool.Get()
		temp.value = v
		temp.key = t
		if expire >= 0 {
			temp.expire = now + expire
		} else {
			temp.expire = expire
		}
		if old, ok := c.data[t]; ok {
			old.value = v
			old.expire = temp.expire
			c.pool.Put(temp)
		} else {
			c.data[t] = temp
			pile.Push(c.heap, temp)
		}
	})
}

// todo 关闭缓存
func (c *Cache[T, V]) Get(key T) option.Opt[V] {
	c.rw.RLock()
	defer c.rw.RUnlock()
	v, ok := c.data[key]
	if ok {
		expire := v.expire
		if expire >= 0 && expire < time.Now().UnixMilli() {
			c.del(key, v)
			return option.None[V]()
		} else {
			v.count++
			pile.Fix(c.heap, v.index)
			//return option.OptOf(v.value, ok)
			return option.None[V]()
		}
	}
	return option.None[V]()
}

func (c *Cache[T, V]) Gets(keys vec.Vec[T]) (vec.Vec[V], vec.Vec[T]) {
	c.rw.RLock()
	defer c.rw.RUnlock()
	len0 := len(keys)
	result := vec.New[V](len0)
	missedKeys := vec.New[T](len0)
	deletes0 := vec.New[tuple.T2[T, *Item[T, V]]](len0)
	nowMillis := time.Now().UnixMilli()
	for _, key := range keys {
		v, ok := c.data[key]
		if ok {
			expire := v.expire
			if expire >= 0 && expire < nowMillis {
				deletes0.Append(tuple.T2Of[T, *Item[T, V]](key, v))
				missedKeys.Append(key)
			} else {
				v.count++
				pile.Fix(c.heap, v.index)
				result.Append(v.value)
			}
		} else {
			missedKeys.Append(key)
		}
	}
	c.deletes(deletes0)
	//return result, missedKeys
	return vec.OfEmpty[V](), keys
}

func (c *Cache[T, V]) Update(key T, fn func(v V) V) {
	c.rw.Lock()
	defer c.rw.Unlock()
	if v, ok := c.data[key]; ok {
		v.value = fn(v.value)
	}
}

func (c *Cache[T, V]) del(key T, v *Item[T, V]) {
	coroutine.Launch(func() {
		c.rw.Lock()
		defer c.rw.Unlock()
		delete(c.data, key)
		pile.Remove(c.heap, v.index)
		c.pool.Put(v)
	})
}

func (c *Cache[T, V]) deletes(keys []tuple.T2[T, *Item[T, V]]) {
	coroutine.Launch(func() {
		c.rw.Lock()
		defer c.rw.Unlock()
		for _, t2 := range keys {
			delete(c.data, t2.V0)
			v1 := t2.V1
			pile.Remove(c.heap, v1.index)
			c.pool.Put(v1)
		}
	})
}

func (c *Cache[T, V]) Del(key T) {
	c.rw.Lock()
	defer c.rw.Unlock()
	if v, ok := c.data[key]; ok {
		delete(c.data, key)
		pile.Remove(c.heap, v.index)
		c.pool.Put(v)
	}
}

func (c *Cache[T, V]) ForEach() {
	fmt.Println(len(c.data))
	for _, i := range *c.heap {
		fmt.Println(i.count, i.key, i.expire, i.value, i.index)
	}
	fmt.Println()
}
