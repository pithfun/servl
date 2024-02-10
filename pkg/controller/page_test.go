package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/makomarket/mako/pkg/tests"
)

func TestNewPage_IsHome(t *testing.T) {
	ctx, _ := tests.NewContext(c.Web, "/")
	page := NewPage(ctx)

	assert.Equal(t, page.IsHome, true)
}
