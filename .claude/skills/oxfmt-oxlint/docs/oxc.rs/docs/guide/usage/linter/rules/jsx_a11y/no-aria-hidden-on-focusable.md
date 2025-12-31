---
title: "jsx_a11y/no-aria-hidden-on-focusable | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-aria-hidden-on-focusable"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsx\_a11y/no-aria-hidden-on-focusable.md for this page in Markdown format

# jsx\_a11y/no-aria-hidden-on-focusable Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-aria-hidden-on-focusable.html#jsx-a11y-no-aria-hidden-on-focusable)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-aria-hidden-on-focusable.html#what-it-does)

Enforces that `aria-hidden="true"` is not set on focusable elements.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-aria-hidden-on-focusable.html#why-is-this-bad)

`aria-hidden="true"` on focusable elements can lead to confusion or unexpected behavior for screen reader users.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-aria-hidden-on-focusable.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<div aria-hidden="true" tabIndex="0" />
```

Examples of **correct** code for this rule:

jsx

```
<div aria-hidden="true" />
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-aria-hidden-on-focusable.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsx-a11y"],
  "rules": {
    "jsx-a11y/no-aria-hidden-on-focusable": "error"
  }
}
```

bash

```
oxlint --deny jsx-a11y/no-aria-hidden-on-focusable --jsx-a11y-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-aria-hidden-on-focusable.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsx_a11y/no_aria_hidden_on_focusable.rs)
