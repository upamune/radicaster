---
title: "jsdoc/require-property-description | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-property-description"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsdoc/require-property-description.md for this page in Markdown format

# jsdoc/require-property-description Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-property-description.html#jsdoc-require-property-description)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-property-description.html#what-it-does)

Requires that all `@property` tags have descriptions.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-property-description.html#why-is-this-bad)

The description of a property should be documented.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-property-description.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
/**
 * @typedef {SomeType} SomeTypedef
 * @property {number} foo
 */
```

Examples of **correct** code for this rule:

javascript

```
/**
 * @typedef {SomeType} SomeTypedef
 * @property {number} foo Foo.
 */
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-property-description.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsdoc"],
  "rules": {
    "jsdoc/require-property-description": "error"
  }
}
```

bash

```
oxlint --deny jsdoc/require-property-description --jsdoc-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-property-description.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsdoc/require_property_description.rs)
