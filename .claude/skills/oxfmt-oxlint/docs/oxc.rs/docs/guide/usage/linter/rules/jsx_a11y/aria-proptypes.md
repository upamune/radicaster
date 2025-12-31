---
title: "jsx_a11y/aria-proptypes | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-proptypes"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsx\_a11y/aria-proptypes.md for this page in Markdown format

# jsx\_a11y/aria-proptypes Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-proptypes.html#jsx-a11y-aria-proptypes)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-proptypes.html#what-it-does)

Enforces that elements do not use invalid ARIA state and property values.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-proptypes.html#why-is-this-bad)

Using invalid ARIA state and property values can mislead screen readers and other assistive technologies. It may cause the accessibility features of the website to fail, making it difficult for users with disabilities to use the site effectively.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-proptypes.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<div aria-level="yes" />
<div aria-relevant="additions removalss" />
```

Examples of **correct** code for this rule:

jsx

```
<div aria-label="foo" />
<div aria-labelledby="foo bar" />
<div aria-checked={false} />
<div aria-invalid="grammar" />
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-proptypes.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsx-a11y"],
  "rules": {
    "jsx-a11y/aria-proptypes": "error"
  }
}
```

bash

```
oxlint --deny jsx-a11y/aria-proptypes --jsx-a11y-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-proptypes.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsx_a11y/aria_proptypes.rs)
