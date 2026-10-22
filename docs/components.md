# Components for the pages that ask for input

Styles and markup for the connector, device-code and consent pages,
agreed with Marat (design) in October 2026 and checked against
[`requirements.md`](requirements.md).

They live here rather than in `design/page.css` for two reasons: every
byte of that file is inlined into every page, and these styles are not
used by the result pages; and the examples below contain HTML, which
`page.css` cannot carry — the page tests reject markup a page did not
intend, which is a security check worth keeping.

Move a block into `design/page.css` when you add the page that uses it.

Contrast figures are against `--bg` / `--surface` from `page.css`.
Icons are [Lucide](https://lucide.dev) inline SVG: 24×24 viewBox,
`stroke="currentColor"`, round caps. CSS cannot carry them.

---

## Buttons

```html
<button class="btn btn--primary">Continue</button>
<button class="btn btn--secondary">Cancel</button>
```

```css
.btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  height: 2.9rem;
  padding: 0 1.5rem;
  border: 1.5px solid transparent;
  border-radius: 999px;
  font: inherit;
  font-size: 0.9375rem;
  font-weight: 600;
  line-height: 1;
  cursor: pointer;
  box-sizing: border-box;
}

.btn--primary {
  /* background-origin: without it the gradient repeats under the
     transparent border and a thin line of the opposite end's colour
     shows at the left and right edges. */
  background: var(--gradient-cta);
  background-origin: border-box;
  color: #fff;
  transition: filter 0.15s ease;
}

.btn--primary:hover {
  filter: brightness(1.08);
}

/* A muted colour rather than opacity: opacity on the whole element
   also fades the focus ring, and leaves Cancel with no hover feedback
   while the primary button has some. */
.btn--secondary {
  background: transparent;
  color: var(--text-muted);
  border-color: currentColor;
  transition: color 0.15s ease;
}

.btn--secondary:hover {
  color: var(--text);
}
```

### The primary button does not meet AA, deliberately

White on `--gradient-cta` is **2.66:1** at the orange end and 3.24:1 at
the coral end, against the 4.5:1 AA asks for. This is a documented
exception limited to this one component. It is the only one on these
pages.

The reasoning, so it does not have to be re-argued:

**The gradient is the logomark's gradient.** `#ff7700 → #f65954` is the
same ramp as the mark in the Giant Swarm logo, which sits at the foot of
every one of these pages. It is not decoration applied to a button; it
is the one place the brand appears other than the logo itself, and it is
the single call-to-action treatment across giantswarm.io.

**The label is white everywhere it appears.** A dark label here would
make these pages the only place in the brand where the primary action
looks different.

Four alternatives were considered and rejected:

| Option | Why not |
|---|---|
| Dark label `#1D232C` on the gradient (5.94:1 / 4.87:1) | Passes, and is what the contrast metric prefers — but it would be the only dark-labelled primary in the brand, and to normal vision it reads as *less* legible than white on saturated orange, not more. |
| Solid blue `#005FB8`, white label (6.31:1) | Passes, but blue has no context on this page. In Loom, blue belongs to components built on blue — the calculator, module-level UI. Here the only other blue is a focus ring, leaving the button an orphan. It also fails as a shape in dark mode: 2.44:1 against `#002645`. |
| A darker gradient that carries white | Would have to diverge from the logomark's gradient — either changing the brand everywhere, or keeping two gradients that are nearly the same. |
| Plain dark orange, as the pre-Loom pages used | At that depth it reads as disabled or already hovered. |

**What this costs, plainly:** someone with low vision, or on a dim
screen in bright light, will find this label harder to read than the
rest of the page. The mitigations are real but partial — the form
submits on Enter and the field takes focus on load, so the button is
not the only route forward.

**What it does not cost:** every other element here meets AA, including
the four that were below it in the first round of review — the input
outline, the placeholder, the small error text and the phishing
warning. Those are the ones that decide whether a page can be used at
all, and they are fixed.

If this exception is not acceptable, the decision belongs in
[`requirements.md`](requirements.md) rather than in a stylesheet, and
the honest alternative is the dark label.

---

## Device-code field

```html
<div class="field">
  <label for="code" class="field__label">Device code</label>
  <div class="combo">
    <input id="code" placeholder="XXXX-XXXX" autocomplete="off"
           aria-invalid="true" aria-describedby="code-error">
    <button class="btn btn--primary">
      <span class="sr-only">Continue</span>
      <!-- Lucide arrow-right, 20px -->
    </button>
  </div>
  <p class="inline-error" id="code-error">
    <!-- Lucide circle-alert, 14px -->
    <span>This code is invalid or expired. Check your terminal for a new one.</span>
  </p>
</div>
```

The field and the action are two separate pills. An outlined container
holding a filled pill reads as an iOS segmented control; a circular
action removes that.

The visible label is back. A placeholder disappears the moment someone
types, so it cannot be the only thing naming the field.

The invalid state keys off `aria-invalid="true"` rather than a class, so
the look cannot drift from what a screen reader announces, and it holds
while the field has focus — otherwise the error vanishes exactly as
someone starts correcting it.

```css
.field {
  width: min(28rem, 100%);
  margin: 1.5rem auto;
  text-align: center;
}

.field__label {
  display: block;
  margin-bottom: 0.45rem;
  font-size: 0.8125rem;
  font-weight: 600;
  color: var(--text);
}

.sr-only {
  position: absolute;
  width: 1px;
  height: 1px;
  padding: 0;
  margin: -1px;
  overflow: hidden;
  clip: rect(0 0 0 0);
  white-space: nowrap;
  border: 0;
}

.combo {
  display: flex;
  align-items: center;
  gap: 0.6rem;
}

.combo input {
  flex: 1 1 auto;
  min-width: 0;
  height: 3.1rem;
  padding: 0 1.4rem;
  box-sizing: border-box;
  border: 1.5px solid var(--border-interactive);
  border-radius: 999px;
  background: var(--surface);
  color: var(--text);
  font-family: var(--font-mono);
  font-size: 1.0625rem;
  letter-spacing: 0.12em;
  text-align: center;
  overflow-wrap: anywhere;
  /* A transparent outline rather than none: Windows high-contrast mode
     discards our border and shadow colours and would otherwise show no
     focus at all. */
  outline: 2px solid transparent;
  outline-offset: 2px;
}

/* No opacity. The placeholder carries the format, and at 0.6 it was
   2.4:1. Full strength is 5.34:1. */
.combo input::placeholder {
  color: var(--text-muted);
}

.combo input:focus-visible {
  border-color: var(--accent-strong);
  box-shadow: 0 0 0 3px var(--halo);
}

.combo input[aria-invalid="true"] {
  border-color: var(--error-text);
}

.combo input[aria-invalid="true"]:focus-visible {
  border-color: var(--error-text);
  box-shadow: 0 0 0 3px var(--halo);
}

.combo .btn {
  flex: 0 0 auto;
  width: 3.1rem;
  height: 3.1rem;
  padding: 0;
  border-radius: 999px;
}

/* --error is tuned for the 40px error heading, where 3:1 is the bar.
   This is 13px and needs 4.5:1: --error gives 3.1:1, --error-text
   gives 4.67:1 on --bg and 5.07:1 on --surface. */
.inline-error {
  display: flex;
  gap: 0.5rem;
  align-items: flex-start;
  justify-content: center;
  margin: 0.6rem 0 0;
  font-size: 0.8125rem;
  line-height: 1.5;
  color: var(--error-text);
  text-align: left;
  overflow-wrap: anywhere;
}

.inline-error svg {
  flex: 0 0 auto;
  margin-top: 0.15rem;
}

/* Below 32rem the pair stacks and the action regains its visible label:
   a full-width button showing only an arrow says nothing. 3.1rem is
   49.6px, clearing the 44px touch target either way. */
@media (max-width: 32rem) {
  .combo {
    display: grid;
    gap: 0.6rem;
  }

  .combo input,
  .combo .btn {
    height: 3.25rem;
  }

  .combo .btn {
    width: 100%;
    padding: 0 1.5rem;
    font-size: 1rem;
  }

  .combo .btn svg {
    display: none;
  }

  .combo .btn .sr-only {
    position: static;
    width: auto;
    height: auto;
    margin: 0;
    overflow: visible;
    clip: auto;
  }
}
```

---

## Connector list

```html
<div class="rows">
  <a class="row" href="…">
    <span class="row__icon"><!-- Lucide house / user-round, 18px --></span>
    <span class="row__text">
      <span class="row__name">Customer connector</span>
      <span class="row__id">customer-ad</span>
    </span>
    <span class="row__go"><!-- Lucide chevron-right, 16px --></span>
  </a>
</div>
```

1.5px borders: at this size 1px reads as a hairline and the hover is
easy to miss. The width stays constant and only the colour changes, so
the row never shifts under the cursor. The connector id is mono,
echoing the reference code on the error page — both are machine
identifiers, and both can be long enough to need wrapping.

```css
.rows {
  display: grid;
  gap: 0.6rem;
  width: min(28rem, 100%);
  margin: 1.5rem auto;
  text-align: left;
}

.row {
  display: flex;
  align-items: center;
  gap: 0.8rem;
  padding: 0.9rem 1rem;
  border: 1.5px solid var(--border-interactive);
  border-radius: 8px;
  background: var(--surface);
  color: inherit;
  text-decoration: none;
  transition: border-color 0.15s ease;
}

.row:hover {
  border-color: var(--accent-strong);
}

.row__icon,
.row__go {
  display: flex;
  flex: 0 0 auto;
  color: var(--text-muted);
}

.row__text {
  flex: 1 1 auto;
  min-width: 0;
}

.row__name {
  font-size: 0.9375rem;
  font-weight: 600;
  color: var(--text);
  overflow-wrap: anywhere;
}

.row__id {
  display: block;
  font-family: var(--font-mono);
  font-size: 0.75rem;
  color: var(--text-muted);
  overflow-wrap: anywhere;
}
```

---

## Consent screen

```html
<p class="message"><span class="app">muster</span> would like to</p>
<div class="measure">
  <ul class="scopes">
    <li><!-- Lucide check, 15px --><span>View basic profile information</span></li>
  </ul>
  <p class="warn">Only continue if you started this sign-in yourself.</p>
  <div class="actions">
    <button class="btn btn--secondary">Cancel</button>
    <button class="btn btn--primary">Grant access</button>
  </div>
</div>
```

The interactive block sits in a narrower column than the prose above it.
Cancel and the primary action are a two-column grid, so they stay equal
whatever the labels say. Checkmarks are blue: UI affirmation, not brand
decoration. Scope strings can be URLs, so they wrap.

```css
.measure {
  width: min(24rem, 100%);
  margin-inline: auto;
}

.scopes {
  margin: 1.25rem 0 0;
  padding: 0.25rem 1.1rem;
  list-style: none;
  border: 1.5px solid var(--border);
  border-radius: 8px;
  background: var(--surface);
  text-align: left;
}

.scopes li {
  display: flex;
  align-items: center;
  gap: 0.7rem;
  padding: 0.8rem 0;
  font-size: 0.875rem;
  color: var(--text-muted);
  overflow-wrap: anywhere;
}

.scopes li + li {
  border-top: 1px solid var(--border);
}

.scopes svg {
  flex: 0 0 auto;
  color: var(--accent-strong);
}

/* The phishing warning requirement 9 asks for. No opacity — at 0.9 it
   was 4.0:1, the faintest text on the page, on a safety message. Full
   strength is 4.93:1, and 500 weight gives it some presence. */
.warn {
  margin: 1.1rem 0 0;
  font-size: 0.8125rem;
  font-weight: 500;
  line-height: 1.55;
  color: var(--text-muted);
  overflow-wrap: anywhere;
}

.app {
  font-weight: 600;
  color: var(--text);
}

.actions {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 0.75rem;
  margin: 1.5rem 0;
}
```
