// Package pagekit holds the shared design of the standalone pages that
// Giant Swarm services serve outside a product UI: sign-in results,
// device-code entry, consent and error pages.
//
// The design sources (page layout, stylesheet, logos and wording) live in
// this repository next to this package, so the package can embed them and
// consumers ship the pages inside their own binaries, with nothing loaded
// at runtime. Consumers that are not Go programs, such as Dex, are meant to
// take the same sources as files.
//
// Render writes a result page, such as the page a browser shows after a
// sign-in:
//
//	pagekit.Render(w, http.StatusOK, pagekit.Page{
//		Title:   "Signed in to GitHub",
//		Heading: "Signed in",
//		Summary: "GitHub is now connected",
//		Message: "You can close this tab and return to Cursor.",
//	})
package pagekit
