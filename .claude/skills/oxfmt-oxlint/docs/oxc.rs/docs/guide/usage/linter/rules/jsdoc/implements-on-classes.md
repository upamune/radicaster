---
title: "jsdoc/implements-on-classes | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/implements-on-classes"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsdoc/implements-on-classes.md for this page in Markdown format

# jsdoc/implements-on-classes Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/implements-on-classes.html#jsdoc-implements-on-classes)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/implements-on-classes.html#what-it-does)

Reports an issue with any non-constructor function using `@implements`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/implements-on-classes.html#why-is-this-bad)

Constructor functions should be whether marked with `@class`, `@constructs`, or being a class constructor.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/implements-on-classes.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
/**
 * @implements {SomeClass}
 */
function quux() {}
```

Examples of **correct** code for this rule:

javascript

```
class Foo {
  /**
   * @implements {SomeClass}
   */
  constructor() {}
}
/**
 * @implements {SomeClass}
 * @class
 */
function quux() {}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/implements-on-classes.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsdoc"],
  "rules": {
    "jsdoc/implements-on-classes": "error"
  }
}
```

bash

```
oxlint --deny jsdoc/implements-on-classes --jsdoc-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/implements-on-classes.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsdoc/implements_on_classes.rs)
