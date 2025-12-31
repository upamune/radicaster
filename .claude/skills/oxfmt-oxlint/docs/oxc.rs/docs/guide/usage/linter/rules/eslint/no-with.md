---
title: "eslint/no-with | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-with"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-with.md for this page in Markdown format

# eslint/no-with Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-with.html#eslint-no-with)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-with.html#what-it-does)

Disallow [`with`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/with) statements.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-with.html#why-is-this-bad)

The with statement is potentially problematic because it adds members of an object to the current scope, making it impossible to tell what a variable inside the block actually refers to.

It is generally considered a bad practice and is forbidden in strict mode.

This rule is not necessary in TypeScript code if `alwaysStrict` is enabled.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-with.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
with (point) {
  r = Math.sqrt(x * x + y * y); // is r a member of point?
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-with.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-with": "error"
  }
}
```

bash

```
oxlint --deny no-with
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-with.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_with.rs)
