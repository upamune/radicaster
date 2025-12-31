---
title: "eslint/no-array-constructor | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-array-constructor"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-array-constructor.md for this page in Markdown format

# eslint/no-array-constructor Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-array-constructor.html#eslint-no-array-constructor)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-array-constructor.html#what-it-does)

Disallows creating arrays with the `Array` constructor.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-array-constructor.html#why-is-this-bad)

Use of the `Array` constructor to construct a new array is generally discouraged in favor of array literal notation because of the single-argument pitfall and because the `Array` global may be redefined. The exception is when the `Array` constructor is used to intentionally create sparse arrays of a specified size by giving the constructor a single numeric argument.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-array-constructor.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
let arr = new Array();
```

Examples of **correct** code for this rule:

javascript

```
let arr = [];
let arr2 = Array.from(iterable);
let arr3 = new Array(9);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-array-constructor.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-array-constructor": "error"
  }
}
```

bash

```
oxlint --deny no-array-constructor
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-array-constructor.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_array_constructor.rs)
