---
title: "unicorn/prefer-type-error | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-type-error"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-type-error.md for this page in Markdown format

# unicorn/prefer-type-error Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-type-error.html#unicorn-prefer-type-error)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-type-error.html#what-it-does)

Enforce throwing a `TypeError` instead of a generic `Error` after a type checking if-statement.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-type-error.html#why-is-this-bad)

Throwing a `TypeError` instead of a generic `Error` after a type checking if-statement is more specific and helps to catch bugs.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-type-error.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
if (Array.isArray(foo)) {
  throw new Error("Expected foo to be an array");
}
```

Examples of **correct** code for this rule:

javascript

```
if (Array.isArray(foo)) {
  throw new TypeError("Expected foo to be an array");
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-type-error.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-type-error": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-type-error
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-type-error.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_type_error.rs)
