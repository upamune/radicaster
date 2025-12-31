---
title: "eslint/array-callback-return | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/array-callback-return"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/array-callback-return.md for this page in Markdown format

# eslint/array-callback-return Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/array-callback-return.html#eslint-array-callback-return)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/array-callback-return.html#what-it-does)

Enforce return statements in callbacks of array methods

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/array-callback-return.html#why-is-this-bad)

Array has several methods for filtering, mapping, and folding. If we forget to write return statement in a callback of those, it’s probably a mistake. If you don’t want to use a return or don’t need the returned results, consider using .forEach instead.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/array-callback-return.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
let foo = [1, 2, 3, 4];
foo.map((a) => {
  console.log(a);
});
```

Examples of **correct** code for this rule:

javascript

```
let foo = [1, 2, 3, 4];
foo.map((a) => {
  console.log(a);
  return a;
});
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/array-callback-return.html#configuration)

This rule accepts a configuration object with the following properties:

### allowImplicit [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/array-callback-return.html#allowimplicit)

type: `boolean`

default: `false`

When set to true, allows callbacks of methods that require a return value to implicitly return undefined with a return statement containing no expression.

### checkForEach [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/array-callback-return.html#checkforeach)

type: `boolean`

default: `false`

When set to true, rule will also report forEach callbacks that return a value.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/array-callback-return.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "array-callback-return": "error"
  }
}
```

bash

```
oxlint --deny array-callback-return
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/array-callback-return.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/array_callback_return/mod.rs)
