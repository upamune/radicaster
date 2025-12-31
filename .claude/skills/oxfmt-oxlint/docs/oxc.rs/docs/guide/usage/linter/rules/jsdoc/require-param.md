---
title: "jsdoc/require-param | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-param"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsdoc/require-param.md for this page in Markdown format

# jsdoc/require-param Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-param.html#jsdoc-require-param)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-param.html#what-it-does)

Requires that all function parameters are documented with JSDoc `@param` tags.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-param.html#why-is-this-bad)

The rule is aimed at enforcing code quality and maintainability by requiring that all function parameters are documented.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-param.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
/** @param foo */
function quux(foo, bar) {}
```

Examples of **correct** code for this rule:

javascript

```
/** @param foo */
function quux(foo) {}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-param.html#configuration)

This rule accepts a configuration object with the following properties:

### checkConstructors [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-param.html#checkconstructors)

type: `boolean`

default: `false`

Whether to check constructor methods.

### checkDestructured [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-param.html#checkdestructured)

type: `boolean`

default: `true`

Whether to check destructured parameters.

### checkDestructuredRoots [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-param.html#checkdestructuredroots)

type: `boolean`

default: `true`

Whether to check destructured parameters when you have code like `function doSomething({ a, b }) { ... }`. Because there is no named parameter in this example, when this option is `true` you must have a `@param` tag that corresponds to `{a, b}`.

### checkGetters [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-param.html#checkgetters)

type: `boolean`

default: `true`

Whether to check getter methods.

### checkRestProperty [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-param.html#checkrestproperty)

type: `boolean`

default: `false`

Whether to check rest properties.

### checkSetters [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-param.html#checksetters)

type: `boolean`

default: `true`

Whether to check setter methods.

### checkTypesPattern [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-param.html#checktypespattern)

type: `string`

default: `"^(?:[oO]bject|[aA]rray|PlainObject|Generic(?:Object|Array))$"`

Regex pattern to match types that exempt parameters from checking.

### exemptedBy [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-param.html#exemptedby)

type: `string[]`

default: `["inheritdoc"]`

List of JSDoc tags that exempt functions from `@param` checking.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-param.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsdoc"],
  "rules": {
    "jsdoc/require-param": "error"
  }
}
```

bash

```
oxlint --deny jsdoc/require-param --jsdoc-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-param.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsdoc/require_param.rs)
