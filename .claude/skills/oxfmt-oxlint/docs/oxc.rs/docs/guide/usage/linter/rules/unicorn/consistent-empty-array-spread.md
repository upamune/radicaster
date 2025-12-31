---
title: "unicorn/consistent-empty-array-spread | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-empty-array-spread"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/consistent-empty-array-spread.md for this page in Markdown format

# unicorn/consistent-empty-array-spread Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-empty-array-spread.html#unicorn-consistent-empty-array-spread)

💡 A suggestion is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-empty-array-spread.html#what-it-does)

When spreading a ternary in an array, we can use both [] and '' as fallbacks, but it's better to have consistent types in both branches.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-empty-array-spread.html#why-is-this-bad)

Having consistent types in both branches makes the code easier to read and understand.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-empty-array-spread.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
const array = [a, ...(foo ? [b, c] : "")];

const array = [a, ...(foo ? "bc" : [])];
```

Examples of **correct** code for this rule:

javascript

```
const array = [a, ...(foo ? [b, c] : [])];

const array = [a, ...(foo ? "bc" : "")];
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-empty-array-spread.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/consistent-empty-array-spread": "error"
  }
}
```

bash

```
oxlint --deny unicorn/consistent-empty-array-spread
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-empty-array-spread.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/consistent_empty_array_spread.rs)
