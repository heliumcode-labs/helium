package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnwrapDDGLink(t *testing.T) {
	t.Run("resolves the redirect wrapper", func(t *testing.T) {
		got := unwrapDDGLink("//duckduckgo.com/l/?uddg=https%3A%2F%2Fgo.dev%2Fdoc%2F&rut=abc")
		assert.Equal(t, "https://go.dev/doc/", got)
	})

	t.Run("keeps direct links", func(t *testing.T) {
		got := unwrapDDGLink("https://example.com/a?b=1")
		assert.Equal(t, "https://example.com/a?b=1", got)
	})

	t.Run("rejects non http targets", func(t *testing.T) {
		assert.Empty(t, unwrapDDGLink("javascript:alert(1)"))
		assert.Empty(t, unwrapDDGLink("//duckduckgo.com/l/?uddg=ftp%3A%2F%2Fexample.com"))
	})

	t.Run("ignores garbage", func(t *testing.T) {
		assert.Empty(t, unwrapDDGLink("::::"))
	})
}

func TestHTMLText(t *testing.T) {
	assert.Equal(t, "a b", htmlText("<span class=\"searchmatch\">a</span> <span>b</span>"))
	assert.Equal(t, "", htmlText(""))
}

func TestFormatWebSearchResults(t *testing.T) {
	got := formatWebSearchResults([]WebSearchResult{
		{Title: "Go 1.30 released", URL: "https://go.dev/blog/go1.30", Snippet: "The Go team ships the release."},
		{Title: "no snippet", URL: "https://example.com"},
	})

	assert.Contains(t, got, "1. [Go 1.30 released](https://go.dev/blog/go1.30)")
	assert.Contains(t, got, "   The Go team ships the release.")
	assert.Contains(t, got, "2. [no snippet](https://example.com)")
}

func TestSanitizeLinkText(t *testing.T) {
	assert.Equal(t, "title (v2) now", sanitizeLinkText("title [v2] now"))
}
