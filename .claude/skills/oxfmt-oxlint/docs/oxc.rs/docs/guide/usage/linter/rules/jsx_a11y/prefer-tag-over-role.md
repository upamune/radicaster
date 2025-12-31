---
title: "jsx_a11y/prefer-tag-over-role | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/prefer-tag-over-role"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsx\_a11y/prefer-tag-over-role.md for this page in Markdown format

# jsx\_a11y/prefer-tag-over-role Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/prefer-tag-over-role.html#jsx-a11y-prefer-tag-over-role)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/prefer-tag-over-role.html#what-it-does)

Enforces using semantic HTML tags over `role` attribute.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/prefer-tag-over-role.html#why-is-this-bad)

Using semantic HTML tags can improve accessibility and readability of the code.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/prefer-tag-over-role.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<div role="button" />
```

Examples of **correct** code for this rule:

jsx

```
<button />
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/prefer-tag-over-role.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsx-a11y"],
  "rules": {
    "jsx-a11y/prefer-tag-over-role": "error"
  }
}
```

bash

```
oxlint --deny jsx-a11y/prefer-tag-over-role --jsx-a11y-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/prefer-tag-over-role.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsx_a11y/prefer_tag_over_role.rs)
