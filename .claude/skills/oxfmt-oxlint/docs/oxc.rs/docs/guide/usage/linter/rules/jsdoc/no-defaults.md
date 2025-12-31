---
title: "jsdoc/no-defaults | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/no-defaults"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsdoc/no-defaults.md for this page in Markdown format

# jsdoc/no-defaults Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/no-defaults.html#jsdoc-no-defaults)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/no-defaults.html#what-it-does)

This rule reports defaults being used on the relevant portion of `@param` or `@default`. It also optionally reports the presence of the square-bracketed optional arguments at all.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/no-defaults.html#why-is-this-bad)

The rule is intended to prevent the indication of defaults on tags where this would be redundant with ES2015 default parameters.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/no-defaults.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
/** @param {number} [foo="7"] */
function quux(foo) {}
```

Examples of **correct** code for this rule:

javascript

```
/** @param {number} foo */
function quux(foo) {}

/** @param foo */
function quux(foo) {}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/no-defaults.html#configuration)

This rule accepts a configuration object with the following properties:

### noOptionalParamNames [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/no-defaults.html#nooptionalparamnames)

type: `boolean`

default: `false`

If true, report the presence of optional param names (square brackets) on `@param` tags.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/no-defaults.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsdoc"],
  "rules": {
    "jsdoc/no-defaults": "error"
  }
}
```

bash

```
oxlint --deny jsdoc/no-defaults --jsdoc-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/no-defaults.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsdoc/no_defaults.rs)
