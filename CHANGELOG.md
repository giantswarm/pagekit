# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).



## [Unreleased]

### Added

- Add `Render`, which writes a success or error result page with the Giant Swarm logo, light and dark mode, and a strict Content Security Policy.

### Changed

- Make the repository root the importable Go package `github.com/giantswarm/pagekit` instead of a `main` package.

### Fixed

- `Render` also sends `X-Frame-Options: DENY`, so browsers that ignore the policy's `frame-ancestors` don't show the pages in a frame either.



[Unreleased]: https://github.com/giantswarm/pagekit/tree/main
