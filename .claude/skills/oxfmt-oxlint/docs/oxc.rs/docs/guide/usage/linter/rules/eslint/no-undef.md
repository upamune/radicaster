---
title: "eslint/no-undef | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-undef"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-undef.md for this page in Markdown format

# eslint/no-undef Nursery [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-undef.html#eslint-no-undef)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-undef.html#what-it-does)

Disallow the use of undeclared variables.

This rule can be disabled for TypeScript code, as the TypeScript compiler enforces this check.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-undef.html#why-is-this-bad)

It is most likely a potential ReferenceError caused by a misspelling of a variable or parameter name.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-undef.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
var foo = someFunction();
var bar = a + 1;
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-undef.html#configuration)

This rule accepts a configuration object with the following properties:

### typeof [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-undef.html#typeof)

type: `boolean`

default: `false`

When set to `true`, warns on undefined variables used in a `typeof` expression.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-undef.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-undef": "error"
  }
}
```

bash

```
oxlint --deny no-undef
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-undef.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_undef.rs)
