---
title: "jsx_a11y/aria-props | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-props"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsx\_a11y/aria-props.md for this page in Markdown format

# jsx\_a11y/aria-props Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-props.html#jsx-a11y-aria-props)

🛠️ An auto-fix is available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-props.html#what-it-does)

Enforces that elements do not use invalid ARIA attributes.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-props.html#why-is-this-bad)

Using invalid ARIA attributes can mislead screen readers and other assistive technologies. It may cause the accessibility features of the website to fail, making it difficult for users with disabilities to use the site effectively.

This rule includes fixes for some common typos.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-props.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<input aria-labeledby="address_label" />
```

Examples of **correct** code for this rule:

jsx

```
<input aria-labelledby="address_label" />
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-props.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsx-a11y"],
  "rules": {
    "jsx-a11y/aria-props": "error"
  }
}
```

bash

```
oxlint --deny jsx-a11y/aria-props --jsx-a11y-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-props.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsx_a11y/aria_props.rs)
