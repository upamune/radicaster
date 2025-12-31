---
title: "jsdoc/require-property | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-property"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsdoc/require-property.md for this page in Markdown format

# jsdoc/require-property Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-property.html#jsdoc-require-property)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-property.html#what-it-does)

Requires that all `@typedef` and `@namespace` tags have `@property` tags when their type is a plain `object`, `Object`, or `PlainObject`.

Note: this rule can be configured via [jsdoc settings](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#settings) option.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-property.html#why-is-this-bad)

Object type should have properties defined.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-property.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
/**
 * @typedef {Object} SomeTypedef
 */

/**
 * @namespace {Object} SomeNamesoace
 */
```

Examples of **correct** code for this rule:

javascript

```
/**
 * @typedef {Object} SomeTypedef
 * @property {SomeType} propName Prop description
 */

/**
 * @typedef {object} Foo
 * @property someProp
 */
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-property.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsdoc"],
  "rules": {
    "jsdoc/require-property": "error"
  }
}
```

bash

```
oxlint --deny jsdoc/require-property --jsdoc-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-property.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsdoc/require_property.rs)
