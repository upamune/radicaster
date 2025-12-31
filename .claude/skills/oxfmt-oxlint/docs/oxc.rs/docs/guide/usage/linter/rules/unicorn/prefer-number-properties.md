---
title: "unicorn/prefer-number-properties | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-number-properties"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-number-properties.md for this page in Markdown format

# unicorn/prefer-number-properties Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-number-properties.html#unicorn-prefer-number-properties)

⚠️🛠️️ A dangerous auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-number-properties.html#what-it-does)

Disallows use of `parseInt()`, `parseFloat()`, `isNan()`, `isFinite()`, `Nan`, `Infinity` and `-Infinity` as global variables.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-number-properties.html#why-is-this-bad)

ECMAScript 2015 moved globals onto the `Number` constructor for consistency and to slightly improve them. This rule enforces their usage to limit the usage of globals:

* [`Number.parseInt()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Number/parseInt) over [`parseInt()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/parseInt)
* [`Number.parseFloat()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Number/parseFloat) over [`parseFloat()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/parseFloat)
* [`Number.isNaN()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Number/isNaN) over [`isNaN()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/isNaN) *(they have slightly [different behavior](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Number/isNaN#difference_between_number.isnan_and_global_isnan))*
* [`Number.isFinite()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Number/isFinite) over [`isFinite()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/isFinite) *(they have slightly [different behavior](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Number/isFinite#difference_between_number.isfinite_and_global_isfinite))*
* [`Number.NaN`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Number/NaN) over [`NaN`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/NaN)
* [`Number.POSITIVE_INFINITY`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Number/POSITIVE_INFINITY) over [`Infinity`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Infinity)
* [`Number.NEGATIVE_INFINITY`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Number/NEGATIVE_INFINITY) over [`-Infinity`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Infinity)

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-number-properties.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
const foo = parseInt("10", 2);
const bar = parseFloat("10.5");
```

Examples of **correct** code for this rule:

javascript

```
const foo = Number.parseInt("10", 2);
const bar = Number.parseFloat("10.5");
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-number-properties.html#configuration)

This rule accepts a configuration object with the following properties:

### checkInfinity [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-number-properties.html#checkinfinity)

type: `boolean`

default: `false`

If set to `true`, checks for usage of `Infinity` and `-Infinity` as global variables.

### checkNaN [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-number-properties.html#checknan)

type: `boolean`

default: `true`

If set to `true`, checks for usage of `NaN` as a global variable.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-number-properties.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-number-properties": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-number-properties
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-number-properties.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_number_properties.rs)
