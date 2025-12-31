---
title: "eslint/no-redeclare | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-redeclare"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-redeclare.md for this page in Markdown format

# eslint/no-redeclare Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-redeclare.html#eslint-no-redeclare)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-redeclare.html#what-it-does)

This rule disallows redeclaring variables within the same scope, ensuring that each variable is declared only once. It helps avoid confusion and unintended behavior in code.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-redeclare.html#why-is-this-bad)

Redeclaring variables in the same scope can lead to unexpected behavior, overwriting existing values, and making the code harder to understand and maintain.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-redeclare.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
var a = 3;
var a = 10;
```

Examples of **correct** code for this rule:

javascript

```
var a = 3;
a = 10;
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-redeclare.html#configuration)

This rule accepts a configuration object with the following properties:

### builtinGlobals [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-redeclare.html#builtinglobals)

type: `boolean`

default: `true`

When set `true`, it flags redeclaring built-in globals (e.g., `let Object = 1;`).

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-redeclare.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-redeclare": "error"
  }
}
```

bash

```
oxlint --deny no-redeclare
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-redeclare.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_redeclare.rs)
