---
title: "eslint/no-unreachable | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unreachable"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-unreachable.md for this page in Markdown format

# eslint/no-unreachable Nursery [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unreachable.html#eslint-no-unreachable)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unreachable.html#what-it-does)

Disallow unreachable code after `return`, `throw`, `continue`, and `break` statements.

This rule can be disabled for TypeScript code if `allowUnreachableCode: false` is configured in the `tsconfig.json`, as the TypeScript compiler enforces this check.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unreachable.html#why-is-this-bad)

Unreachable code after a `return`, `throw`, `continue`, or `break` statement can never be run.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unreachable.html#examples)

Examples of **incorrect** code for this rule:

ts

```
function foo() {
  return 2;
  console.log("this will never be executed");
}
```

Examples of **correct** code for this rule:

ts

```
function foo() {
  console.log("this will be executed");
  return 2;
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unreachable.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-unreachable": "error"
  }
}
```

bash

```
oxlint --deny no-unreachable
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unreachable.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_unreachable.rs)
