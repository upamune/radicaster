---
title: "unicorn/no-unnecessary-array-splice-count | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unnecessary-array-splice-count"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-unnecessary-array-splice-count.md for this page in Markdown format

# unicorn/no-unnecessary-array-splice-count Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unnecessary-array-splice-count.html#unicorn-no-unnecessary-array-splice-count)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unnecessary-array-splice-count.html#what-it-does)

Disallows passing `.length` or `Infinity` as the `deleteCount` or `skipCount` argument of `Array#splice()` or `Array#toSpliced()`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unnecessary-array-splice-count.html#why-is-this-bad)

When calling `Array#splice(start, deleteCount)` or `Array#toSpliced(start, skipCount)`, omitting the `deleteCount` or `skipCount` argument will delete or skip all elements after `start`. Using `.length` or `Infinity` is unnecessary and makes the code more verbose.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unnecessary-array-splice-count.html#examples)

Examples of **incorrect** code for this rule:

js

```
array.splice(1, array.length);
array.splice(1, Infinity);
array.splice(1, Number.POSITIVE_INFINITY);
array.toSpliced(1, array.length);
```

Examples of **correct** code for this rule:

js

```
array.splice(1);
array.toSpliced(1);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unnecessary-array-splice-count.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-unnecessary-array-splice-count": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-unnecessary-array-splice-count
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unnecessary-array-splice-count.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_unnecessary_array_splice_count.rs)
