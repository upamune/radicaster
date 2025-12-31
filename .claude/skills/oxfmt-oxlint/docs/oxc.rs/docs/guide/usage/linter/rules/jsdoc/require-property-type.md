---
title: "jsdoc/require-property-type | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-property-type"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsdoc/require-property-type.md for this page in Markdown format

# jsdoc/require-property-type Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-property-type.html#jsdoc-require-property-type)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-property-type.html#what-it-does)

Requires that each `@property` tag has a type value (within curly brackets).

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-property-type.html#why-is-this-bad)

The type of a property should be documented.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-property-type.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
/**
 * @typedef {SomeType} SomeTypedef
 * @property foo
 */
```

Examples of **correct** code for this rule:

javascript

```
/**
 * @typedef {SomeType} SomeTypedef
 * @property {number} foo
 */
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-property-type.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsdoc"],
  "rules": {
    "jsdoc/require-property-type": "error"
  }
}
```

bash

```
oxlint --deny jsdoc/require-property-type --jsdoc-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-property-type.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsdoc/require_property_type.rs)
