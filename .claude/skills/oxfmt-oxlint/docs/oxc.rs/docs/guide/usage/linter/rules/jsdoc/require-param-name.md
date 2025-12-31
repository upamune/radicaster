---
title: "jsdoc/require-param-name | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-param-name"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsdoc/require-param-name.md for this page in Markdown format

# jsdoc/require-param-name Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-param-name.html#jsdoc-require-param-name)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-param-name.html#what-it-does)

Requires that all `@param` tags have names.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-param-name.html#why-is-this-bad)

The name of a param should be documented.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-param-name.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
/** @param {SomeType} */
function quux(foo) {}
```

Examples of **correct** code for this rule:

javascript

```
/** @param {SomeType} foo */
function quux(foo) {}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-param-name.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsdoc"],
  "rules": {
    "jsdoc/require-param-name": "error"
  }
}
```

bash

```
oxlint --deny jsdoc/require-param-name --jsdoc-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-param-name.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsdoc/require_param_name.rs)
