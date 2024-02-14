package middleware

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ghostyinc/ghosty/pkg/tests"
)

func TestServeCachedPage(t *testing.T) {
	// Cache a page
	cp := CachedPage{
		URL:        "/cache",
		HTML:       []byte("html"),
		Headers:    make(map[string]string),
		StatusCode: http.StatusCreated,
	}
	cp.Headers["a"] = "b"
	cp.Headers["c"] = "d"

	err := c.Cache.
		Set().
		Group(CachedPageGroup).
		Key(cp.URL).
		Data(cp).
		Save(context.Background())
	require.NoError(t, err)

	// Request the URL of the cached page
	ctx, rec := tests.NewContext(c.Web, cp.URL)
	err = tests.ExecuteMiddleware(ctx, ServeCachedPage(c.Cache))
	assert.NoError(t, err)
	assert.Equal(t, cp.StatusCode, ctx.Response().Status)
	assert.Equal(t, cp.Headers["a"], ctx.Response().Header().Get("a"))
	assert.Equal(t, cp.Headers["c"], ctx.Response().Header().Get("c"))
	assert.Equal(t, cp.HTML, rec.Body.Bytes())

	// TODO Authentication tests
}

func TestCacheControl(t *testing.T) {
	ctx, _ := tests.NewContext(c.Web, "/")
	_ = tests.ExecuteMiddleware(ctx, CacheControl(time.Second*10))
	assert.Equal(t, "public, max-age=10", ctx.Response().Header().Get("Cache-Control"))
	_ = tests.ExecuteMiddleware(ctx, CacheControl(0))
	assert.Equal(t, "no-cache, no-store", ctx.Response().Header().Get("Cache-Control"))
}
