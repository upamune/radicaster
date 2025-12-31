---
title: "jsx_a11y/click-events-have-key-events | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/click-events-have-key-events"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsx\_a11y/click-events-have-key-events.md for this page in Markdown format

# jsx\_a11y/click-events-have-key-events Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/click-events-have-key-events.html#jsx-a11y-click-events-have-key-events)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/click-events-have-key-events.html#what-it-does)

Enforce onClick is accompanied by at least one of the following: onKeyUp, onKeyDown, onKeyPress.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/click-events-have-key-events.html#why-is-this-bad)

Coding for the keyboard is important for users with physical disabilities who cannot use a mouse, AT compatibility, and screenreader users. This does not apply for interactive or hidden elements.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/click-events-have-key-events.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<div onClick={() => void 0} />
```

Examples of **correct** code for this rule:

jsx

```
<div onClick={() => void 0} onKeyDown={() => void 0} />
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/click-events-have-key-events.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsx-a11y"],
  "rules": {
    "jsx-a11y/click-events-have-key-events": "error"
  }
}
```

bash

```
oxlint --deny jsx-a11y/click-events-have-key-events --jsx-a11y-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/click-events-have-key-events.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsx_a11y/click_events_have_key_events.rs)
