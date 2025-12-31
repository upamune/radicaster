---
title: "unicorn/no-unreadable-array-destructuring | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unreadable-array-destructuring"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-unreadable-array-destructuring.md for this page in Markdown format

# unicorn/no-unreadable-array-destructuring Style [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unreadable-array-destructuring.html#unicorn-no-unreadable-array-destructuring)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unreadable-array-destructuring.html#what-it-does)

Disallow unreadable array destructuring

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unreadable-array-destructuring.html#why-is-this-bad)

Destructuring is very useful, but it can also make some code harder to read. This rule prevents ignoring consecutive values when destructuring from an array.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unreadable-array-destructuring.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
const [, , foo] = parts;
```

Examples of **correct** code for this rule:

javascript

```
const [foo] = parts;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unreadable-array-destructuring.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-unreadable-array-destructuring": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-unreadable-array-destructuring
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unreadable-array-destructuring.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_unreadable_array_destructuring.rs)
