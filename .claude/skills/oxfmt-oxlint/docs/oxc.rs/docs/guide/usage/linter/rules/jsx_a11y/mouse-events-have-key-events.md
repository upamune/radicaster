---
title: "jsx_a11y/mouse-events-have-key-events | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/mouse-events-have-key-events"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsx\_a11y/mouse-events-have-key-events.md for this page in Markdown format

# jsx\_a11y/mouse-events-have-key-events Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/mouse-events-have-key-events.html#jsx-a11y-mouse-events-have-key-events)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/mouse-events-have-key-events.html#what-it-does)

Enforce onMouseOver/onMouseOut are accompanied by onFocus/onBlur.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/mouse-events-have-key-events.html#why-is-this-bad)

Coding for the keyboard is important for users with physical disabilities who cannot use a mouse, AT compatibility, and screen reader users.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/mouse-events-have-key-events.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<div onMouseOver={() => void 0} />
```

Examples of **correct** code for this rule:

jsx

```
<div onMouseOver={() => void 0} onFocus={() => void 0} />
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/mouse-events-have-key-events.html#configuration)

This rule accepts a configuration object with the following properties:

### hoverInHandlers [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/mouse-events-have-key-events.html#hoverinhandlers)

type: `string[]`

default: `["onMouseOver"]`

List of hover-in mouse event handlers that require corresponding keyboard event handlers.

### hoverOutHandlers [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/mouse-events-have-key-events.html#hoverouthandlers)

type: `string[]`

default: `["onMouseOut"]`

List of hover-out mouse event handlers that require corresponding keyboard event handlers.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/mouse-events-have-key-events.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsx-a11y"],
  "rules": {
    "jsx-a11y/mouse-events-have-key-events": "error"
  }
}
```

bash

```
oxlint --deny jsx-a11y/mouse-events-have-key-events --jsx-a11y-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/mouse-events-have-key-events.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsx_a11y/mouse_events_have_key_events.rs)
