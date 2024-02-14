package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/eko/gocache/v2/cache"
	"github.com/eko/gocache/v2/marshaler"
	"github.com/eko/gocache/v2/store"
	"github.com/go-redis/redis/v8"

	"github.com/ghostyinc/ghosty/config"
)

type (
	// CacheClient is the interface that allows us to interact with the cache client
	CacheClient struct {
		// Client stores the client to the underlying cache service
		Client *redis.Client

		// Cache stores the cache interface
		cache *cache.Cache
	}

	// cacheFlush handles chainable cache flush operations
	cacheFlush struct {
		client *CacheClient
		group  string
		key    string
		tags   []string
	}

	// cacheGet handles chainable cache get operations
	cacheGet struct {
		client   *CacheClient
		group    string
		key      string
		dataType interface{}
	}

	// cacheSet handles chainable cache set operations
	cacheSet struct {
		client     *CacheClient
		data       interface{}
		expiration time.Duration
		group      string
		key        string
		tags       []string
	}
)

// NewCacheClient creates a new cache client
func NewCacheClient(cfg *config.Config) (*CacheClient, error) {
	db := cfg.Cache.Database
	// Check if we are in a test environment
	if cfg.App.Environment == config.EnvTest {
		db = cfg.Cache.TestDatabase
	}

	// Connect to the client
	c := &CacheClient{}
	c.Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Cache.Hostname, cfg.Cache.Port),
		DB:       db,
		Password: cfg.Cache.Password,
	})
	if _, err := c.Client.Ping(context.Background()).Result(); err != nil {
		return c, err
	}

	// If we are in a test environment, flush the database
	if cfg.App.Environment == config.EnvTest {
		if err := c.Client.FlushDB(context.Background()).Err(); err != nil {
			return c, err
		}
	}

	cacheStore := store.NewRedis(c.Client, nil)
	c.cache = cache.New(cacheStore)
	return c, nil
}

// cacheKey formats the cache key with an optional group
func (c *CacheClient) cacheKey(group, key string) string {
	if group != "" {
		return fmt.Sprintf("%s::%s", group, key)
	}
	return key
}

// Close closes the connection to the client
func (c *CacheClient) Close() error {
	return c.Client.Close()
}

// Data sets the data to the cache
func (c *cacheSet) Data(data interface{}) *cacheSet {
	c.data = data
	return c
}

// Fetch retrieves the data from the cache
func (c *cacheGet) Fetch(ctx context.Context) (interface{}, error) {
	if c.key == "" {
		return nil, errors.New("no cache key provided")
	}

	if c.dataType == nil {
		return nil, errors.New("no data type provided")
	}

	return marshaler.
		New(c.client.cache).
		Get(ctx, c.client.cacheKey(c.group, c.key), c.dataType)
}

// Execute executes the cache flush operation
func (c *cacheFlush) Execute(ctx context.Context) error {
	if len(c.tags) > 0 {
		if err := c.client.cache.Invalidate(ctx, store.InvalidateOptions{
			Tags: c.tags,
		}); err != nil {
			return err
		}
	}

	if c.key != "" {
		return c.client.cache.Delete(ctx, c.client.cacheKey(c.group, c.key))
	}

	return nil
}

// Expiration sets the cache expiration on the set operation
func (c *cacheSet) Expiration(expiration time.Duration) *cacheSet {
	c.expiration = expiration
	return c
}

// Flush creates a cache flush operation
func (c *CacheClient) Flush() *cacheFlush {
	return &cacheFlush{
		client: c,
	}
}

// Get creates a cache get operation
func (c *CacheClient) Get() *cacheGet {
	return &cacheGet{
		client: c,
	}
}

// Group sets the cache group on the flush operation
func (c *cacheFlush) Group(group string) *cacheFlush {
	c.group = group
	return c
}

// Group sets the cache group on the get operation
func (c *cacheGet) Group(group string) *cacheGet {
	c.group = group
	return c
}

// Group sets the cache group on the set operation
func (c *cacheSet) Group(group string) *cacheSet {
	c.group = group
	return c
}

// Key sets the cache key on the flush operation
func (c *cacheFlush) Key(key string) *cacheFlush {
	c.key = key
	return c
}

// Key sets the cache key on the set operation
func (c *cacheSet) Key(key string) *cacheSet {
	c.key = key
	return c
}

// Key sets the cache key on the get operation
func (c *cacheGet) Key(key string) *cacheGet {
	c.key = key
	return c
}

// Save saves the data to the cache
func (c *cacheSet) Save(ctx context.Context) error {
	if c.key == "" {
		return errors.New("no cache key provided")
	}

	opts := &store.Options{
		Expiration: c.expiration,
		Tags:       c.tags,
	}

	return marshaler.
		New(c.client.cache).
		Set(ctx, c.client.cacheKey(c.group, c.key), c.data, opts)
}

// Set creates a cache set operation
func (c *CacheClient) Set() *cacheSet {
	return &cacheSet{
		client: c,
	}
}

// Tags sets the cache tags on the flush operation
func (c *cacheFlush) Tags(tags ...string) *cacheFlush {
	c.tags = tags
	return c
}

// Tags sets the cache tags on the set operation
func (c *cacheSet) Tags(tags ...string) *cacheSet {
	c.tags = tags
	return c
}

// Type set the cache type for the expected data
func (c *cacheGet) Type(expectedType interface{}) *cacheGet {
	c.dataType = expectedType
	return c
}
