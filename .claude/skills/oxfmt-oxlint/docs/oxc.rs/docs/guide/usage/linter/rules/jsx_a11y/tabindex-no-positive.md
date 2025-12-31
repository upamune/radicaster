---
title: "jsx_a11y/tabindex-no-positive | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/tabindex-no-positive"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsx\_a11y/tabindex-no-positive.md for this page in Markdown format

# jsx\_a11y/tabindex-no-positive Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/tabindex-no-positive.html#jsx-a11y-tabindex-no-positive)

⚠️💡 A dangerous suggestion is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/tabindex-no-positive.html#what-it-does)

Enforces that positive values for the `tabIndex` attribute are not used in JSX.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/tabindex-no-positive.html#why-is-this-bad)

Using `tabIndex` values greater than `0` can make navigation and interaction difficult for keyboard and assistive technology users, disrupting the logical order of content.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/tabindex-no-positive.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<span tabIndex="1">foo</span>
```

Examples of **correct** code for this rule:

jsx

```
<span tabIndex="0">foo</span>
<span tabIndex="-1">bar</span>
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/tabindex-no-positive.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsx-a11y"],
  "rules": {
    "jsx-a11y/tabindex-no-positive": "error"
  }
}
```

bash

```
oxlint --deny jsx-a11y/tabindex-no-positive --jsx-a11y-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/tabindex-no-positive.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsx_a11y/tabindex_no_positive.rs)
