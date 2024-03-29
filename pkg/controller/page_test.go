package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tiny-blob/tinyblob/pkg/tests"
)

func TestNewPage_IsHome(t *testing.T) {
	// TODO: Implement proper tests
	ctx, _ := tests.NewContext(c.Web, "/")
	page := NewPage(ctx)

	assert.Equal(t, page.IsHome, true)
}
