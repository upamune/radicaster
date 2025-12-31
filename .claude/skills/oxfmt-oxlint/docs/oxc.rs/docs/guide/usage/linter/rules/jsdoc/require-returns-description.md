---
title: "jsdoc/require-returns-description | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-returns-description"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsdoc/require-returns-description.md for this page in Markdown format

# jsdoc/require-returns-description Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-returns-description.html#jsdoc-require-returns-description)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-returns-description.html#what-it-does)

Requires that the `@returns` tag has a description value. The error will not be reported if the return value is `void` or `undefined` or if it is `Promise<void>` or `Promise<undefined>`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-returns-description.html#why-is-this-bad)

A `@returns` tag should have a description value.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-returns-description.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
/** @returns */
function quux(foo) {}
```

Examples of **correct** code for this rule:

javascript

```
/** @returns Foo. */
function quux(foo) {}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-returns-description.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsdoc"],
  "rules": {
    "jsdoc/require-returns-description": "error"
  }
}
```

bash

```
oxlint --deny jsdoc/require-returns-description --jsdoc-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-returns-description.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsdoc/require_returns_description.rs)
