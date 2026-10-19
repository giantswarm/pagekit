# Requirements

Every page pagekit provides meets these requirements. The survey of the existing pages and the reasoning behind the list are in giantswarm/giantswarm#38589.

1. **Responsive.** Pages work from phone width (320 CSS px) to desktop without horizontal scrolling.
2. **Follows the light/dark preference.** Colours follow `prefers-color-scheme`. There is no toggle.
3. **Carries the Giant Swarm logo.** The artwork comes from [giantswarm/brand](https://github.com/giantswarm/brand).
4. **Loads and renders fast.** A page makes no external requests: CSS is inline, images are inline SVG or `data:` URIs, and text uses the system font stack. This also allows a strict Content Security Policy and lets pages work on loopback servers and air-gapped installations.
5. **WCAG 2.2 AA.** In particular:
   - text and controls meet the AA contrast ratios in both light and dark mode
   - focus is always visible
   - every input has a label
   - touch targets are at least 44 × 44 px
   - animation respects `prefers-reduced-motion`
6. **Works without JavaScript.** Scripts may only add convenience, such as closing the tab or redirecting to an app, and the page must make sense without them.
7. *(best effort)* **Every error is an HTML page.** It says what happened, what to do next, and gives a reference code a person can pass on to support. A browser never lands on raw JSON or plain text.
8. *(best effort)* **Consistent wording.** Pages use the same phrasing for the same situation, show display names (the `ui.giantswarm.io/display-name` annotation) instead of resource names, and name the client the person should return to.
9. **Input pages warn about phishing.** Any page that asks for input or confirmation, such as device-code entry or consent, tells the person to continue only if they started the sign-in themselves.
10. **English only.** Pages declare `lang="en"`. There is no localisation.

Requirements 7 and 8 depend on each consumer adopting them, for example by passing a display name or a reference code. pagekit makes them easy to take up one consumer at a time and does not block a consumer that hasn't yet.
