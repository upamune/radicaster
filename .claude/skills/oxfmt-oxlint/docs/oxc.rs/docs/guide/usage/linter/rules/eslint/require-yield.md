---
title: "eslint/require-yield | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/require-yield"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/require-yield.md for this page in Markdown format

# eslint/require-yield Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/require-yield.html#eslint-require-yield)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/require-yield.html#what-it-does)

This rule generates warnings for generator functions that do not have the yield keyword.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/require-yield.html#why-is-this-bad)

Probably a mistake.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/require-yield.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
function* foo() {
  return 10;
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/require-yield.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "require-yield": "error"
  }
}
```

bash

```
oxlint --deny require-yield
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/require-yield.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/require_yield.rs)
