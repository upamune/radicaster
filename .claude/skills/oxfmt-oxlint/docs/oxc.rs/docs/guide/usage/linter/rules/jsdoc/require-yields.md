---
title: "jsdoc/require-yields | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-yields"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsdoc/require-yields.md for this page in Markdown format

# jsdoc/require-yields Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-yields.html#jsdoc-require-yields)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-yields.html#what-it-does)

Requires that yields are documented. Will also report if multiple `@yields` tags are present.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-yields.html#why-is-this-bad)

The rule is intended to prevent the omission of `@yields` tags when they are necessary.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-yields.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
function* quux(foo) {
  yield foo;
}

/**
 * @yields {undefined}
 * @yields {void}
 */
function* quux(foo) {}
```

Examples of **correct** code for this rule:

javascript

```
/** * @yields Foo */
function* quux(foo) {
  yield foo;
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-yields.html#configuration)

This rule accepts a configuration object with the following properties:

### exemptedBy [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-yields.html#exemptedby)

type: `string[]`

default: `["inheritdoc"]`

Functions with these tags will be exempted from the lint rule.

### forceRequireYields [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-yields.html#forcerequireyields)

type: `boolean`

default: `false`

When `true`, all generator functions must have a `@yields` tag, even if they don't yield a value or have an empty body.

### withGeneratorTag [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-yields.html#withgeneratortag)

type: `boolean`

default: `false`

When `true`, require `@yields` when a `@generator` tag is present.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-yields.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsdoc"],
  "rules": {
    "jsdoc/require-yields": "error"
  }
}
```

bash

```
oxlint --deny jsdoc/require-yields --jsdoc-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-yields.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsdoc/require_yields.rs)
