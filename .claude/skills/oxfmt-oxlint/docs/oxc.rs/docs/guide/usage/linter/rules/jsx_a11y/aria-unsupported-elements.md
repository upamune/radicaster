---
title: "jsx_a11y/aria-unsupported-elements | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-unsupported-elements"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsx\_a11y/aria-unsupported-elements.md for this page in Markdown format

# jsx\_a11y/aria-unsupported-elements Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-unsupported-elements.html#jsx-a11y-aria-unsupported-elements)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-unsupported-elements.html#what-it-does)

Enforces that reserved DOM elements do not contain ARIA roles, states, or properties.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-unsupported-elements.html#why-is-this-bad)

Certain reserved DOM elements do not support ARIA roles, states and properties. This is often because they are not visible, for example `meta`, `html`, `script`, `style`. Adding ARIA attributes to these elements is meaningless and can create confusion for screen readers.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-unsupported-elements.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<meta charset="UTF-8" aria-hidden="false" />
```

Examples of **correct** code for this rule:

jsx

```
<meta charset="UTF-8" />
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-unsupported-elements.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsx-a11y"],
  "rules": {
    "jsx-a11y/aria-unsupported-elements": "error"
  }
}
```

bash

```
oxlint --deny jsx-a11y/aria-unsupported-elements --jsx-a11y-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-unsupported-elements.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsx_a11y/aria_unsupported_elements.rs)
