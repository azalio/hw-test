package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(3)

		_ = c.Set("aa", 1) // "aa"
		_ = c.Set("bb", 2) // "bb", "aa"
		_ = c.Set("cc", 3) // "cc", "bb", "aa"
		_ = c.Set("dd", 4) // "dd", "cc", "bb"

		val, ok := c.Get("aa")
		require.False(t, ok)
		require.Nil(t, val)

		val, ok = c.Get("dd") // "dd", "cc", "bb"
		require.True(t, ok)
		require.Equal(t, 4, val)

		val, ok = c.Get("cc") // cc, dd, bb
		require.True(t, ok)
		require.Equal(t, 3, val)

		val, ok = c.Get("bb") // bb, cc, dd
		require.True(t, ok)
		require.Equal(t, 2, val)

		_, _ = c.Get("bb")   // bb, cc, dd
		_, _ = c.Get("dd")   // dd, bb, cc
		_ = c.Set("cc", 300) // cc, dd, bb
		_ = c.Set("aa", 200) // aa, cc, dd

		val, ok = c.Get("bb")
		require.False(t, ok)
		require.Nil(t, val)

		c.Clear()
		val, ok = c.Get("cc")
		require.False(t, ok)
		require.Nil(t, val)
	})
}

func TestCacheMultithreading(t *testing.T) {
	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
