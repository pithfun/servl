package funcmap

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tiny-blob/tinyblob/config"
)

func TestHasField(t *testing.T) {
	type example struct {
		name string
	}
	var e example
	e.name = "Sorrento" // Field is not used
	assert.True(t, HasField(e, "name"))
	assert.False(t, HasField(e, "abcd"))
}

func TestLink(t *testing.T) {
	link := string(Link("/abc", "persia", "/abc"))
	expected := `<a class="text-blue-600 hover:underline font-bold" href="/abc">persia</a>`
	assert.Equal(t, expected, link)

	link = string(Link("/abc", "persia", "/abc", "first", "second"))
	expected = `<a class="first second text-blue-600 hover:underline font-bold" href="/abc">persia</a>`
	assert.Equal(t, expected, link)

	link = string(Link("/abc", "persia", "/def"))
	expected = `<a class="text-blue-600 hover:underline" href="/abc">persia</a>`
	assert.Equal(t, expected, link)
}

func TestGetFuncMap(t *testing.T) {
	fileName := "favicon.ico"
	file := File(fileName)
	expected := fmt.Sprintf("/%s/%s?v=%s", config.StaticPrefix, fileName, CacheBuster)
	assert.Equal(t, expected, file)
}
