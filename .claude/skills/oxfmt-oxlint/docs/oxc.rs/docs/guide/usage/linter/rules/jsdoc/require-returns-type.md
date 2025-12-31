---
title: "jsdoc/require-returns-type | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-returns-type"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsdoc/require-returns-type.md for this page in Markdown format

# jsdoc/require-returns-type Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-returns-type.html#jsdoc-require-returns-type)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-returns-type.html#what-it-does)

Requires that `@returns` tag has a type value (in curly brackets).

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-returns-type.html#why-is-this-bad)

A `@returns` tag should have a type value.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-returns-type.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
/** @returns */
function quux(foo) {}
```

Examples of **correct** code for this rule:

javascript

```
/** @returns {string} */
function quux(foo) {}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-returns-type.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsdoc"],
  "rules": {
    "jsdoc/require-returns-type": "error"
  }
}
```

bash

```
oxlint --deny jsdoc/require-returns-type --jsdoc-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-returns-type.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsdoc/require_returns_type.rs)
