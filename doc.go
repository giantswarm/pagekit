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
// The package has no API yet; see the README for the plan.
package pagekit
