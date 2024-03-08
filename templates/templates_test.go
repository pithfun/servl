package templates

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGo(t *testing.T) {
	_, err := Get().Open("pages/home.tmpl")
	require.NoError(t, err)
}

func TestGetOS(t *testing.T) {
	_, err := GetOS().Open("pages/home.tmpl")
	require.NoError(t, err)
}
