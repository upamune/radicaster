---
title: "unicorn/prefer-array-index-of | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-array-index-of"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-array-index-of.md for this page in Markdown format

# unicorn/prefer-array-index-of Style [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-array-index-of.html#unicorn-prefer-array-index-of)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-array-index-of.html#what-it-does)

Enforces using `indexOf` or `lastIndexOf` instead of `findIndex` or `findLastIndex` when the callback is a simple strict equality comparison.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-array-index-of.html#why-is-this-bad)

Using `findIndex(x => x === value)` is unnecessarily verbose when `indexOf(value)` accomplishes the same thing more concisely and clearly. It also avoids the overhead of creating a callback function.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-array-index-of.html#examples)

Examples of **incorrect** code for this rule:

js

```
values.findIndex((x) => x === "foo");
values.findLastIndex((x) => x === "bar");
```

Examples of **correct** code for this rule:

js

```
values.indexOf("foo");
values.lastIndexOf("bar");
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-array-index-of.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-array-index-of": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-array-index-of
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-array-index-of.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_array_index_of.rs)
