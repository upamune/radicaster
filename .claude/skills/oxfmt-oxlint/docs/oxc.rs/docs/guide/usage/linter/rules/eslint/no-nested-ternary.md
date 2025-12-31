---
title: "eslint/no-nested-ternary | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-nested-ternary"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-nested-ternary.md for this page in Markdown format

# eslint/no-nested-ternary Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-nested-ternary.html#eslint-no-nested-ternary)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-nested-ternary.html#what-it-does)

Disallows nested ternary expressions to improve code readability and maintainability.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-nested-ternary.html#why-is-this-bad)

Nested ternary expressions make code harder to read and understand. They can lead to complex, difficult-to-debug logic.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-nested-ternary.html#examples)

Examples of **incorrect** code for this rule:

js

```
const result = condition1 ? (condition2 ? "a" : "b") : "c";
```

Examples of **correct** code for this rule:

js

```
let result;
if (condition1) {
  result = condition2 ? "a" : "b";
} else {
  result = "c";
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-nested-ternary.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-nested-ternary": "error"
  }
}
```

bash

```
oxlint --deny no-nested-ternary
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-nested-ternary.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_nested_ternary.rs)
