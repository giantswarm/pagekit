package pagekit

import (
	"bytes"
	"crypto/sha256"
	"embed"
	"encoding/base64"
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strings"
)

// Page is the content of a result page. All text is escaped, and every field
// is optional.
type Page struct {
	// Title is the browser tab title, such as "Signed in to GitHub". It
	// defaults to Heading.
	Title string
	// Heading is the outcome in a few words, such as "Signed in". It
	// defaults to "Done", or "Something went wrong" for an error status.
	Heading string
	// Summary says what the outcome applies to, such as "GitHub is now
	// connected".
	Summary string
	// Message says what to do next, such as "You can close this tab and
	// return to Cursor."
	Message string
	// Detail is a technical hint for support, such as an error code. It is
	// shown in monospace.
	Detail string
	// Reference is a code the person can pass on to support, such as a
	// request ID that also appears in the logs.
	Reference string
}

//go:embed design/page.html design/page.css design/logo/*.svg
var design embed.FS

var (
	pageTmpl = template.Must(template.ParseFS(design, "design/page.html"))
	css      = mustRead("design/page.css")
	logo     = svgDataURI(mustRead("design/logo/giant-swarm-logo-standard.svg"))
	logoDark = svgDataURI(mustRead("design/logo/giant-swarm-logo-white.svg"))

	// The page loads nothing: the stylesheet is inline and allowed by its
	// hash, and the logos are data URIs.
	contentSecurityPolicy = "default-src 'none'; style-src " + hashSource(css) +
		"; img-src data:; base-uri 'none'; form-action 'none'; frame-ancestors 'none'"
)

type pageData struct {
	Content  Page
	Error    bool
	CSS      template.CSS
	Logo     template.URL
	LogoDark template.Srcset
}

// Render writes p as an HTML response with the given status code. A status
// of 400 or more shows the page as an error.
func Render(w http.ResponseWriter, status int, p Page) {
	isError := status >= http.StatusBadRequest
	if p.Heading == "" {
		p.Heading = "Done"
		if isError {
			p.Heading = "Something went wrong"
		}
	}
	if p.Title == "" {
		p.Title = p.Heading
	}

	var buf bytes.Buffer
	err := pageTmpl.Execute(&buf, pageData{
		Content:  p,
		Error:    isError,
		CSS:      template.CSS(css),         // #nosec G203 -- embedded at build time.
		Logo:     template.URL(logo),        // #nosec G203 -- embedded at build time.
		LogoDark: template.Srcset(logoDark), // #nosec G203 -- embedded at build time.
	})
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	h := w.Header()
	h.Set("Content-Type", "text/html; charset=utf-8")
	h.Set("Content-Security-Policy", contentSecurityPolicy)
	h.Set("X-Content-Type-Options", "nosniff")
	// frame-ancestors in the policy covers current browsers; this covers
	// older ones.
	h.Set("X-Frame-Options", "DENY")
	h.Set("Referrer-Policy", "no-referrer")
	h.Set("Cache-Control", "no-store")
	w.WriteHeader(status)
	_, _ = w.Write(buf.Bytes())
}

func mustRead(name string) string {
	b, err := design.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func hashSource(s string) string {
	sum := sha256.Sum256([]byte(s))
	return "'sha256-" + base64.StdEncoding.EncodeToString(sum[:]) + "'"
}

var (
	svgPreamble = regexp.MustCompile(`(?s)<\?xml.*?\?>|<!--.*?-->|<!DOCTYPE[^>]*>`)
	whitespace  = regexp.MustCompile(`\s+`)
	betweenTags = regexp.MustCompile(`>\s+<`)
)

// svgDataURI returns svg as a percent-encoded data URI. It drops the XML
// declaration, comments and DOCTYPE and collapses whitespace, which keeps
// the pages small without touching the artwork.
func svgDataURI(svg string) string {
	svg = svgPreamble.ReplaceAllString(svg, "")
	svg = betweenTags.ReplaceAllString(svg, "><")
	svg = strings.TrimSpace(whitespace.ReplaceAllString(svg, " "))

	var b strings.Builder
	b.WriteString("data:image/svg+xml,")
	for i := 0; i < len(svg); i++ {
		c := svg[i]
		if isURISafe(c) {
			b.WriteByte(c)
		} else {
			fmt.Fprintf(&b, "%%%02X", c)
		}
	}
	return b.String()
}

// isURISafe reports whether c can stay unencoded in a data URI inside a
// src or srcset attribute. Whitespace would end a srcset URL early.
func isURISafe(c byte) bool {
	switch {
	case 'a' <= c && c <= 'z', 'A' <= c && c <= 'Z', '0' <= c && c <= '9':
		return true
	}
	return strings.IndexByte("-._~!$()*+,;=:@/?", c) >= 0
}
