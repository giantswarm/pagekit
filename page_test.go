package pagekit_test

import (
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/giantswarm/pagekit"
)

const (
	signedIn     = "Signed in"
	signInFailed = "Sign-in failed"
)

func TestRender_WritesPageContent(t *testing.T) {
	rec := httptest.NewRecorder()

	pagekit.Render(rec, http.StatusOK, pagekit.Page{
		Title:   "Signed in to GitHub",
		Heading: signedIn,
		Summary: "GitHub is now connected",
		Message: "You can close this tab and return to Cursor.",
	})

	require.Equal(t, http.StatusOK, rec.Code)
	require.Equal(t, "text/html; charset=utf-8", rec.Header().Get("Content-Type"))
	body := rec.Body.String()
	require.Contains(t, body, `<html lang="en">`)
	require.Contains(t, body, "<title>Signed in to GitHub</title>")
	require.Contains(t, body, "<h1>Signed in</h1>")
	require.Contains(t, body, "GitHub is now connected")
	require.Contains(t, body, "You can close this tab and return to Cursor.")
}

func TestRender_SetsStrictContentSecurityPolicy(t *testing.T) {
	rec := httptest.NewRecorder()

	pagekit.Render(rec, http.StatusOK, pagekit.Page{Heading: signedIn})

	// A browser hashes the text between <style> and </style> as served.
	m := regexp.MustCompile(`(?s)<style>(.*?)</style>`).FindStringSubmatch(rec.Body.String())
	require.Len(t, m, 2, "page has no inline stylesheet")
	sum := sha256.Sum256([]byte(m[1]))
	styleHash := "'sha256-" + base64.StdEncoding.EncodeToString(sum[:]) + "'"

	require.Equal(t,
		"default-src 'none'; style-src "+styleHash+"; img-src data:; base-uri 'none'; form-action 'none'; frame-ancestors 'none'",
		rec.Header().Get("Content-Security-Policy"))
	require.Equal(t, "nosniff", rec.Header().Get("X-Content-Type-Options"))
	require.Equal(t, "no-referrer", rec.Header().Get("Referrer-Policy"))
	require.Equal(t, "no-store", rec.Header().Get("Cache-Control"))
}

func TestRender_TitleDefaultsToHeading(t *testing.T) {
	rec := httptest.NewRecorder()

	pagekit.Render(rec, http.StatusOK, pagekit.Page{Heading: signedIn})

	require.Contains(t, rec.Body.String(), "<title>Signed in</title>")
}

func TestRender_ShowsDetailAndReference(t *testing.T) {
	testCases := []struct {
		name      string
		page      pagekit.Page
		contains  []string
		notInPage []string
	}{
		{
			name:     "both",
			page:     pagekit.Page{Heading: signInFailed, Detail: "state_expired", Reference: "7f3a9c"},
			contains: []string{"<code>state_expired</code>", "Reference <code>7f3a9c</code>"},
		},
		{
			name:      "detail only",
			page:      pagekit.Page{Heading: signInFailed, Detail: "state_expired"},
			contains:  []string{"<code>state_expired</code>"},
			notInPage: []string{"Reference"},
		},
		{
			name:      "reference only",
			page:      pagekit.Page{Heading: signInFailed, Reference: "7f3a9c"},
			contains:  []string{"Reference <code>7f3a9c</code>"},
			notInPage: []string{" · "},
		},
		{
			name:      "neither",
			page:      pagekit.Page{Heading: signedIn},
			notInPage: []string{`class="detail"`},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			pagekit.Render(rec, http.StatusBadRequest, tc.page)

			require.Equal(t, http.StatusBadRequest, rec.Code)
			for _, s := range tc.contains {
				require.Contains(t, rec.Body.String(), s)
			}
			for _, s := range tc.notInPage {
				require.NotContains(t, rec.Body.String(), s)
			}
		})
	}
}

func TestRender_EscapesText(t *testing.T) {
	rec := httptest.NewRecorder()

	pagekit.Render(rec, http.StatusBadRequest, pagekit.Page{
		Kind:    pagekit.Error,
		Heading: signInFailed,
		Message: `<a href="https://evil.example">Click</a>`,
		Detail:  "<script>alert(1)</script>",
	})

	body := rec.Body.String()
	require.NotContains(t, body, "<script>")
	require.NotContains(t, body, "<a ")
	require.Contains(t, body, "&lt;script&gt;alert(1)&lt;/script&gt;")
}

func TestRender_LoadsNothingExternal(t *testing.T) {
	rec := httptest.NewRecorder()

	pagekit.Render(rec, http.StatusOK, pagekit.Page{Heading: signedIn})

	refs := regexp.MustCompile(`(?:src|srcset|href)="([^"]*)"`).FindAllStringSubmatch(rec.Body.String(), -1)
	require.NotEmpty(t, refs)
	for _, ref := range refs {
		require.True(t, strings.HasPrefix(ref[1], "data:"), "page references %.60q", ref[1])
	}
	require.Contains(t, rec.Body.String(), `alt="Giant Swarm"`)
}
