package services

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCacheClient(t *testing.T) {
	type cacheTest struct {
		Value string
	}

	// Cache some test data
	data := cacheTest{Value: "test_data"}
	group := "test_group"
	key := "test_key"
	err := c.Cache.
		Set().
		Group(group).
		Key(key).
		Data(data).
		Save(context.Background())
	require.NoError(t, err)

	// Retrieve the cached data
	cachedData, err := c.Cache.
		Get().
		Group(group).
		Key(key).
		Type(new(cacheTest)).
		Fetch(context.Background())
	require.NoError(t, err)
	cast, ok := cachedData.(*cacheTest)
	require.True(t, ok)
	require.Equal(t, data, *cast)

	// The same key without a group should return an error
	_, err = c.Cache.
		Get().
		Key(key).
		Type(new(cacheTest)).
		Fetch(context.Background())
	assert.Error(t, err)

	// Flush the cache
	err = c.Cache.
		Flush().
		Group(group).
		Key(key).
		Execute(context.Background())
	require.NoError(t, err)

	// The data should no longer be cached
	assertFlushed := func() {
		_, err = c.Cache.
			Get().
			Group(group).
			Key(key).
			Type(new(cacheTest)).
			Fetch(context.Background())
		assert.Equal(t, redis.Nil, err)
	}
	assertFlushed()

	// Cache some test data with tags
	err = c.Cache.
		Set().
		Group(group).
		Key(key).
		Data(data).
		Tags("tag1", "tag2").
		Save(context.Background())
	require.NoError(t, err)

	// Flush the cache by tag
	err = c.Cache.
		Flush().
		Tags("tag1").
		Execute(context.Background())
	require.NoError(t, err)

	// The data should no longer be cached
	assertFlushed()

	// Cache some data with an expiration
	err = c.Cache.
		Set().
		Group(group).
		Key(key).
		Data(data).
		Expiration(time.Millisecond).
		Save(context.Background())
	require.NoError(t, err)

	time.Sleep(time.Millisecond * 2)

	// The data should no longer be cached
	assertFlushed()
}
