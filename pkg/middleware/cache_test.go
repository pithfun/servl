package middleware

import (
	"goblin/pkg/tests"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCacheControl(t *testing.T) {
	ctx, _ := tests.NewContext(c.Web, "/")

	_ = tests.ExecuteMiddleware(ctx, CacheControl(time.Second*10))
	assert.Equal(t, "public, max-age=10", ctx.Response().Header().Get("Cache-Control"))

	_ = tests.ExecuteMiddleware(ctx, CacheControl(0))
	assert.Equal(t, "no-cache, no-store", ctx.Response().Header().Get("Cache-Control"))
}
