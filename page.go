package pagekit

import (
	"bytes"
	"crypto/sha256"
	"embed"
	"encoding/base64"
	"html/template"
	"net/http"
	"regexp"
)

// Kind is the outcome a page reports.
type Kind int

const (
	// Success reports that the person's action worked.
	Success Kind = iota
	// Error reports that it did not. The heading takes the error colour.
	Error
)

// Page is the content of a result page. All text is escaped; only Heading
// is required.
type Page struct {
	Kind Kind
	// Title is the browser tab title, such as "Signed in to GitHub". It
	// defaults to Heading.
	Title string
	// Heading is the outcome in a few words, such as "Signed in".
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
	Page
	Error    bool
	CSS      template.CSS
	Logo     template.URL
	LogoDark template.Srcset
}

// Render writes p as an HTML response with the given status code.
func Render(w http.ResponseWriter, status int, p Page) {
	if p.Title == "" {
		p.Title = p.Heading
	}

	var buf bytes.Buffer
	err := pageTmpl.Execute(&buf, pageData{
		Page:     p,
		Error:    p.Kind == Error,
		CSS:      template.CSS(css),         // #nosec G203 -- embedded at build time.
		Logo:     template.URL(logo),        // #nosec G203 -- embedded at build time.
		LogoDark: template.Srcset(logoDark), // #nosec G203 -- embedded at build time.
	})
	if err != nil {
		http.Error(w, http.StatusText(status), status)
		return
	}

	h := w.Header()
	h.Set("Content-Type", "text/html; charset=utf-8")
	h.Set("Content-Security-Policy", contentSecurityPolicy)
	h.Set("X-Content-Type-Options", "nosniff")
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

var whitespace = regexp.MustCompile(`\s+`)

// svgDataURI returns svg as a data URI. Collapsing whitespace keeps the
// pages small without touching the artwork.
func svgDataURI(svg string) string {
	compact := whitespace.ReplaceAllString(svg, " ")
	return "data:image/svg+xml;base64," + base64.StdEncoding.EncodeToString([]byte(compact))
}
