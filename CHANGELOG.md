# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).



## [Unreleased]

### Fixed

- The page stays centred when a browser extension adds elements to `<body>`, as Automa, Grammarly and password managers do. `<body>` centred its children as a grid, so an added element got a grid row of its own and pushed the content up; the page wrapper centres the content itself now.

## [0.2.0] - 2026-10-21

### Added

- `OAuthErrorCode`, which keeps an authorization response's error code only when the IANA OAuth Extensions Error Registry lists it for the authorization endpoint (or it is RFC 8707's `invalid_target`), so a page can show it without showing text a third party put into the callback URL.

## [0.1.2] - 2026-10-20

### Changed

- The README shows the success and error pages.

## [0.1.1] - 2026-10-19

### Fixed

- `Render` also sends `X-Frame-Options: DENY`, so browsers that ignore the policy's `frame-ancestors` don't show the pages in a frame either.

## [0.1.0] - 2026-10-19

### Added

- Add `Render`, which writes a success or error result page with the Giant Swarm logo, light and dark mode, and a strict Content Security Policy.

### Changed

- Make the repository root the importable Go package `github.com/giantswarm/pagekit` instead of a `main` package.



[Unreleased]: https://github.com/giantswarm/pagekit/compare/v0.2.0...HEAD
[0.2.0]: https://github.com/giantswarm/pagekit/releases/tag/v0.2.0
[0.1.2]: https://github.com/giantswarm/pagekit/releases/tag/v0.1.2
[0.1.1]: https://github.com/giantswarm/pagekit/releases/tag/v0.1.1
[0.1.0]: https://github.com/giantswarm/pagekit/releases/tag/v0.1.0
