---
title: "jsdoc/check-property-names | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/check-property-names"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsdoc/check-property-names.md for this page in Markdown format

# jsdoc/check-property-names Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/check-property-names.html#jsdoc-check-property-names)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/check-property-names.html#what-it-does)

Ensures that property names in JSDoc are not duplicated on the same block and that nested properties have defined roots.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/check-property-names.html#why-is-this-bad)

`@property` tags with the same name can be confusing and may indicate a mistake.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/check-property-names.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
/**
 * @typedef {object} state
 * @property {number} foo
 * @property {string} foo
 */

/**
 * @typedef {object} state
 * @property {number} foo.bar
 */
```

Examples of **correct** code for this rule:

javascript

```
/**
 * @typedef {object} state
 * @property {number} foo
 */

/**
 * @typedef {object} state
 * @property {object} foo
 * @property {number} foo.bar
 */
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/check-property-names.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsdoc"],
  "rules": {
    "jsdoc/check-property-names": "error"
  }
}
```

bash

```
oxlint --deny jsdoc/check-property-names --jsdoc-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/check-property-names.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsdoc/check_property_names.rs)
