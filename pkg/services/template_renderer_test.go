package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tiny-blob/tinyblob/config"
	"github.com/tiny-blob/tinyblob/templates"
)

func TestTemplateRenderer(t *testing.T) {
	group := "testGroup"
	key := "testKey"

	// Should not exist
	parsed, err := c.TemplateRenderer.Load(group, key)
	assert.Nil(t, parsed)
	assert.Error(t, err)

	// Parse and store template into cache
	tpl, err := c.TemplateRenderer.
		Parse().
		Group(group).
		Key(key).
		Base("htmx").
		Files("htmx", "pages/error").
		Directories("components").
		Store()
	require.NoError(t, err)

	// Parsed template should exist now
	parsed, err = c.TemplateRenderer.Load(group, key)
	require.NoError(t, err)

	// Check that all the expected templates are included.
	expectedTemplate := make(map[string]bool)
	expectedTemplate["htmx"+config.TemplateExt] = true
	expectedTemplate["error"+config.TemplateExt] = true
	components, err := templates.Templates.ReadDir("components")
	require.NoError(t, err)
	for _, f := range components {
		expectedTemplate[f.Name()] = true
	}

	// Remove the expected templates from our parsed templates
	for _, v := range parsed.Template.Templates() {
		delete(expectedTemplate, v.Name())
	}
	assert.Empty(t, expectedTemplate)

	data := struct {
		StatusCode int
	}{
		StatusCode: 500,
	}
	buf, err := tpl.Execute(data)
	require.NoError(t, err)
	require.NotNil(t, buf)
	assert.Contains(t, buf.String(), "Server error. Refresh page and try again.")

	buf, err = c.TemplateRenderer.
		Parse().
		Group(group).
		Key(key).
		Base("htmx").
		Files("htmx", "pages/error").
		Directories("components").
		Execute(data)

	require.NoError(t, err)
	require.NotNil(t, buf)
	assert.Contains(t, buf.String(), "Server error. Refresh page and try again.")
}
