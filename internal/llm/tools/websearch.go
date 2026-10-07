package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/heliumcode-labs/helium/internal/config"
	"github.com/heliumcode-labs/helium/internal/permission"
)

const (
	WebSearchToolName    = "websearch"
	webSearchDescription = `Searches the web and returns titles, URLs and snippets. No API key is required.

WHEN TO USE THIS TOOL:
- The answer depends on current information: releases, versions, pricing, documentation, incidents
- You know a page exists but not its URL
- You need several candidate pages to compare before reading one with the Fetch tool

HOW TO USE:
- Provide a short, focused query (3-8 keywords work best)
- Optionally set max_results (default 8, max 20)
- Results come back as numbered links with a short snippet; use Fetch on the most promising URL to read it

WORKING BACKENDS:
- DuckDuckGo for general web results
- Wikipedia automatically takes over when the general search returns nothing

LIMITATIONS:
- Results are ranked by the search backend, not by relevance to code
- Snippets are short: always fetch the page when the snippet is not enough
- Sites that block automated requests may fail on the follow-up Fetch call

TIPS:
- Prefer site-scoped queries like "golang slices package docs" over questions
- If the first query returns noise, retry with different keywords instead of fetching bad results`
)

type WebSearchParams struct {
	Query      string `json:"query"`
	MaxResults int    `json:"max_results,omitempty"`
}

type WebSearchPermissionsParams struct {
	Query string `json:"query"`
}

type WebSearchResult struct {
	Title   string `json:"title"`
	URL     string `json:"url"`
	Snippet string `json:"snippet"`
}

type WebSearchResponseMetadata struct {
	Query           string `json:"query"`
	Engine          string `json:"engine"`
	NumberOfResults int    `json:"number_of_results"`
}

type webSearchTool struct {
	client      *http.Client
	permissions permission.Service
}

func NewWebSearchTool(permissions permission.Service) BaseTool {
	return &webSearchTool{
		client: &http.Client{
			Timeout: 20 * time.Second,
		},
		permissions: permissions,
	}
}

func (t *webSearchTool) Info() ToolInfo {
	return ToolInfo{
		Name:        WebSearchToolName,
		Description: webSearchDescription,
		Parameters: map[string]any{
			"query": map[string]any{
				"type":        "string",
				"description": "The search query",
			},
			"max_results": map[string]any{
				"type":        "number",
				"description": "Maximum number of results to return. Defaults to 8, maximum 20.",
			},
		},
		Required: []string{"query"},
	}
}

func (t *webSearchTool) Run(ctx context.Context, call ToolCall) (ToolResponse, error) {
	var params WebSearchParams
	if err := json.Unmarshal([]byte(call.Input), &params); err != nil {
		return NewTextErrorResponse(fmt.Sprintf("error parsing parameters: %s", err)), nil
	}

	query := strings.TrimSpace(params.Query)
	if query == "" {
		return NewTextErrorResponse("query is required"), nil
	}

	limit := params.MaxResults
	if limit <= 0 {
		limit = 8
	}
	if limit > 20 {
		limit = 20
	}

	// Searching only sends the query itself, but it still touches the network
	// so the user gets a chance to approve (and remember) it.
	if sessionID, _ := GetContextValues(ctx); sessionID != "" {
		if allowed := t.permissions.Request(
			permission.CreatePermissionRequest{
				SessionID:   sessionID,
				Path:        config.WorkingDirectory(),
				ToolName:    WebSearchToolName,
				Action:      "search",
				Description: fmt.Sprintf("Search the web for: %s", query),
				Params:      WebSearchPermissionsParams{Query: query},
			},
		); !allowed {
			return ToolResponse{}, permission.ErrorPermissionDenied
		}
	}

	results, engine, err := t.searchDuckDuckGo(ctx, query, limit)
	if err != nil || len(results) == 0 {
		wikiResults, _, wikiErr := t.searchWikipedia(ctx, query, limit)
		if wikiErr == nil && len(wikiResults) > 0 {
			results, engine = wikiResults, "wikipedia"
		} else if err != nil {
			return NewTextErrorResponse(fmt.Sprintf(
				"web search failed: %s (wikipedia fallback: %s)", err, wikiErr)), nil
		} else {
			return NewTextResponse(fmt.Sprintf("No results found for %q", query)), nil
		}
	}

	return WithResponseMetadata(
		NewTextResponse(formatWebSearchResults(results)),
		WebSearchResponseMetadata{
			Query:           query,
			Engine:          engine,
			NumberOfResults: len(results),
		},
	), nil
}

func (t *webSearchTool) searchDuckDuckGo(ctx context.Context, query string, limit int) ([]WebSearchResult, string, error) {
	endpoint := "https://html.duckduckgo.com/html/?" + url.Values{"q": {query}}.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml")

	resp, err := t.client.Do(req)
	if err != nil {
		return nil, "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("search backend returned status %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, "", fmt.Errorf("failed to parse response: %w", err)
	}

	results := []WebSearchResult{}
	doc.Find(".result").EachWithBreak(func(_ int, s *goquery.Selection) bool {
		link := s.Find("a.result__a")
		if link.Length() == 0 {
			return true
		}

		target := unwrapDDGLink(link.AttrOr("href", ""))
		if target == "" {
			return true
		}

		results = append(results, WebSearchResult{
			Title:   strings.TrimSpace(link.Text()),
			URL:     target,
			Snippet: strings.TrimSpace(s.Find(".result__snippet").Text()),
		})
		return len(results) < limit
	})

	return results, "duckduckgo", nil
}

func (t *webSearchTool) searchWikipedia(ctx context.Context, query string, limit int) ([]WebSearchResult, string, error) {
	endpoint := "https://en.wikipedia.org/w/api.php?" + url.Values{
		"action":   {"query"},
		"list":     {"search"},
		"srsearch": {query},
		"srlimit":  {fmt.Sprint(limit)},
		"format":   {"json"},
		"utf8":     {"1"},
	}.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("User-Agent", "heliumcode/1.0 (https://github.com/heliumcode-labs/helium)")

	resp, err := t.client.Do(req)
	if err != nil {
		return nil, "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("wikipedia returned status %d", resp.StatusCode)
	}

	var payload struct {
		Query struct {
			Search []struct {
				Title   string `json:"title"`
				Snippet string `json:"snippet"`
			} `json:"search"`
		} `json:"query"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return nil, "", fmt.Errorf("failed to parse response: %w", err)
	}

	results := make([]WebSearchResult, 0, len(payload.Query.Search))
	for _, hit := range payload.Query.Search {
		results = append(results, WebSearchResult{
			Title:   hit.Title,
			URL:     "https://en.wikipedia.org/wiki/" + strings.ReplaceAll(hit.Title, " ", "_"),
			Snippet: htmlText(hit.Snippet),
		})
	}

	return results, "wikipedia", nil
}

// unwrapDDGLink resolves DuckDuckGo's redirect links (/l/?uddg=…) to the
// destination URL and rejects anything that is not plain http(s).
func unwrapDDGLink(href string) string {
	if strings.HasPrefix(href, "//") {
		href = "https:" + href
	}

	u, err := url.Parse(strings.TrimSpace(href))
	if err != nil {
		return ""
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return ""
	}

	if strings.Contains(u.Host, "duckduckgo.com") && strings.HasSuffix(u.Path, "/l/") {
		target := u.Query().Get("uddg")
		parsed, err := url.Parse(target)
		if err != nil || (parsed.Scheme != "http" && parsed.Scheme != "https") {
			return ""
		}
		return parsed.String()
	}

	return u.String()
}

var htmlTagPattern = regexp.MustCompile(`<[^>]*>`)

// htmlText strips markup (Wikipedia highlights snippets with <span>) and
// decodes the entities that are left over.
func htmlText(s string) string {
	if s == "" {
		return ""
	}

	if doc, err := goquery.NewDocumentFromReader(strings.NewReader("<div>" + s + "</div>")); err == nil {
		return strings.Join(strings.Fields(doc.Text()), " ")
	}

	return strings.Join(strings.Fields(htmlTagPattern.ReplaceAllString(s, " ")), " ")
}

func formatWebSearchResults(results []WebSearchResult) string {
	var b strings.Builder
	for i, r := range results {
		title := sanitizeLinkText(r.Title)
		if title == "" {
			title = r.URL
		}
		fmt.Fprintf(&b, "%d. [%s](%s)\n", i+1, title, r.URL)
		if snippet := strings.TrimSpace(r.Snippet); snippet != "" {
			fmt.Fprintf(&b, "   %s\n", snippet)
		}
		b.WriteString("\n")
	}

	return strings.TrimRight(b.String(), "\n")
}

// sanitizeLinkText keeps markdown links parseable when titles contain
// brackets or other markup.
func sanitizeLinkText(s string) string {
	s = strings.ReplaceAll(s, "[", "(")
	s = strings.ReplaceAll(s, "]", ")")
	return strings.Join(strings.Fields(s), " ")
}
