# pagekit

Go module and HTML templates for sign-in pages, interstitials and other system pages wearing the Giant Swarm brand.

| Success, light mode | Error, dark mode |
|---|---|
| ![Success page in light mode: the heading "Signed in", the line "GitHub is now connected", the hint to close the tab and return to Cursor, and the Giant Swarm logo](https://github.com/user-attachments/assets/da281e4b-c578-4601-b4e2-3820cf8174e0) | ![Error page in dark mode: the heading "Sign-in failed", the line "GitHub could not be connected", what to do next, the error code and a reference, and the white Giant Swarm logo](https://github.com/user-attachments/assets/6211d81c-0331-4463-82f5-5e178bdd6719) |

## What belongs here

Pages that a service or proxy serves on its own, outside a product UI: the page a browser shows after a sign-in, device-code entry, consent screens, error pages, and later perhaps maintenance or "not found" pages. Application UIs are out of scope.

The design is the source of truth. Each kind of consumer gets it in the form it can take in:

- **Go services and CLIs** import `github.com/giantswarm/pagekit`. The design files are embedded in the package, so the pages need no network access at runtime.
- **Dex** and other consumers that are not Go programs take the same sources as files.

Logos come from [giantswarm/brand](https://github.com/giantswarm/brand).

What every page has to meet (responsive, light and dark, accessible, no external requests, …) is listed in [docs/requirements.md](docs/requirements.md).

## Status

Early. pagekit renders success and error result pages through `pagekit.Render`. Device-code entry, consent and the files for Dex follow. The survey of existing pages and the plan are tracked in giantswarm/giantswarm#38589.

## Usage

```go
pagekit.Render(w, http.StatusBadRequest, pagekit.Page{
	Title:     "Sign-in to GitHub failed",
	Heading:   "Sign-in failed",
	Summary:   "GitHub could not be connected",
	Message:   "The sign-in link expired before it was used. Start the sign-in again from Cursor.",
	Detail:    "state_expired",
	Reference: requestID,
})
```

A status of 400 or more shows the page as an error. Pass a callback's OAuth `error` parameter through `pagekit.OAuthErrorCode` before showing it: it keeps only the error codes the specifications define, since anyone can put text of their choosing into a callback URL. `Render` sets a Content Security Policy that allows only the page's own inline stylesheet and `data:` images, so the page loads nothing from the network.

## Layout

The module lives at the repository root, and the root directory is the importable package `pagekit`. Keep it that way:

- Go's `//go:embed` can only reach files inside the module, so the design sources must sit under the root module.
- The design sources sit in [`design/`](design/): the page template, the stylesheet and the logos.
- The generated CI (`architect/go-build`), the Makefile and the release tags (`vX.Y.Z`) all assume a module at the root.
