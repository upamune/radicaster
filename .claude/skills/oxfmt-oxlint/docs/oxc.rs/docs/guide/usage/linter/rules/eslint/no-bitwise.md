---
title: "eslint/no-bitwise | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-bitwise"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-bitwise.md for this page in Markdown format

# eslint/no-bitwise Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-bitwise.html#eslint-no-bitwise)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-bitwise.html#what-it-does)

Disallow bitwise operators

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-bitwise.html#why-is-this-bad)

The use of bitwise operators in JavaScript is very rare and often `&` or `|` is simply a mistyped `&&` or `||`, which will lead to unexpected behavior.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-bitwise.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
var x = y | z;
```

javascript

```
var x = y ^ z;
```

javascript

```
var x = y >> z;
```

Examples of **correct** code for this rule:

javascript

```
var x = y || z;
```

javascript

```
var x = y && z;
```

javascript

```
var x = y > z;
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-bitwise.html#configuration)

This rule accepts a configuration object with the following properties:

### allow [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-bitwise.html#allow)

type: `string[]`

default: `[]`

The `allow` option permits the given list of bitwise operators to be used as exceptions to this rule.

For example `{ "allow": ["~"] }` would allow the use of the bitwise operator `~` without restriction. Such as in the following:

javascript

```
~[1, 2, 3].indexOf(1) === -1;
```

### int32Hint [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-bitwise.html#int32hint)

type: `boolean`

default: `false`

When set to `true` the `int32Hint` option allows the use of bitwise OR in |0 pattern for type casting.

For example with `{ "int32Hint": true }` the following is permitted:

javascript

```
const b = a | 0;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-bitwise.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-bitwise": "error"
  }
}
```

bash

```
oxlint --deny no-bitwise
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-bitwise.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_bitwise.rs)
