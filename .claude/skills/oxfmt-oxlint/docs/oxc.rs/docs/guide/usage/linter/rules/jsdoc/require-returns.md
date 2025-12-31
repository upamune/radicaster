---
title: "jsdoc/require-returns | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-returns"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsdoc/require-returns.md for this page in Markdown format

# jsdoc/require-returns Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-returns.html#jsdoc-require-returns)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-returns.html#what-it-does)

Requires that return statements are documented. Will also report if multiple `@returns` tags are present.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-returns.html#why-is-this-bad)

The rule is intended to prevent the omission of `@returns` tag when necessary.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-returns.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
/** Foo. */
function quux() {
  return foo;
}

/**
 * @returns Foo!
 * @returns Foo?
 */
function quux() {
  return foo;
}
```

Examples of **correct** code for this rule:

javascript

```
/** @returns Foo. */
function quux() {
  return foo;
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-returns.html#configuration)

This rule accepts a configuration object with the following properties:

### checkConstructors [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-returns.html#checkconstructors)

type: `boolean`

default: `false`

Whether to check constructor methods.

### checkGetters [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-returns.html#checkgetters)

type: `boolean`

default: `true`

Whether to check getter methods.

### exemptedBy [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-returns.html#exemptedby)

type: `string[]`

default: `["inheritdoc"]`

Tags that exempt functions from requiring `@returns`.

### forceRequireReturn [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-returns.html#forcerequirereturn)

type: `boolean`

default: `false`

Whether to require a `@returns` tag even if the function doesn't return a value.

### forceReturnsWithAsync [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-returns.html#forcereturnswithasync)

type: `boolean`

default: `false`

Whether to require a `@returns` tag for async functions.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-returns.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsdoc"],
  "rules": {
    "jsdoc/require-returns": "error"
  }
}
```

bash

```
oxlint --deny jsdoc/require-returns --jsdoc-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-returns.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsdoc/require_returns.rs)
