# pagekit

Go module and HTML templates for sign-in pages, interstitials and other system pages wearing the Giant Swarm brand.

## What belongs here

Pages that a service or proxy serves on its own, outside a product UI: the page a browser shows after a sign-in, device-code entry, consent screens, error pages, and later perhaps maintenance or "not found" pages. Application UIs are out of scope.

The design is the source of truth. Each kind of consumer gets it in the form it can take in:

- **Go services and CLIs** import `github.com/giantswarm/pagekit`. The design files are embedded in the package, so the pages need no network access at runtime.
- **Dex** and other consumers that are not Go programs take the same sources as files.

Logos come from [giantswarm/brand](https://github.com/giantswarm/brand).

## Status

Early. The repository has the module structure but no pages or API yet. The survey of existing pages and the plan are tracked in giantswarm/giantswarm#38589.

## Layout

The module lives at the repository root, and the root directory is the importable package `pagekit`. Keep it that way:

- Go's `//go:embed` can only reach files inside the module, so the design sources must sit under the root module.
- The generated CI (`architect/go-build`), the Makefile and the release tags (`vX.Y.Z`) all assume a module at the root.
