---
title: "eslint/no-self-compare | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-self-compare"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-self-compare.md for this page in Markdown format

# eslint/no-self-compare Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-self-compare.html#eslint-no-self-compare)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-self-compare.html#what-it-does)

Disallow comparisons where both sides are exactly the same

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-self-compare.html#why-is-this-bad)

Comparing a variable against itself is usually an error, either a typo or refactoring error. It is confusing to the reader and may potentially introduce a runtime error.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-self-compare.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
var x = 10;
if (x === x) {
  x = 20;
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-self-compare.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-self-compare": "error"
  }
}
```

bash

```
oxlint --deny no-self-compare
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-self-compare.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_self_compare.rs)
