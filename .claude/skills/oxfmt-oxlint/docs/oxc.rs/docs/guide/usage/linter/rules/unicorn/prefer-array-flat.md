---
title: "unicorn/prefer-array-flat | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-array-flat"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-array-flat.md for this page in Markdown format

# unicorn/prefer-array-flat Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-array-flat.html#unicorn-prefer-array-flat)

⚠️🛠️️ A dangerous auto-fix is available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-array-flat.html#what-it-does)

Prefers `Array#flat()` over legacy techniques to flatten arrays.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-array-flat.html#why-is-this-bad)

ES2019 introduced a new method [`Array#flat()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/flat) that flatten arrays.

This rule aims to standardize the use of `Array#flat()` over legacy techniques to flatten arrays.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-array-flat.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
const foo = array.flatMap((x) => x);
const foo = array.reduce((a, b) => a.concat(b), []);
const foo = array.reduce((a, b) => [...a, ...b], []);
const foo = [].concat(maybeArray);
const foo = [].concat(...array);
const foo = [].concat.apply([], array);
const foo = Array.prototype.concat.apply([], array);
const foo = Array.prototype.concat.call([], maybeArray);
const foo = Array.prototype.concat.call([], ...array);
```

Examples of **correct** code for this rule:

javascript

```
const foo = array.flat();
const foo = [maybeArray].flat();
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-array-flat.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-array-flat": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-array-flat
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-array-flat.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_array_flat.rs)
