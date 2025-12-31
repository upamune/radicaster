---
title: "unicorn/no-magic-array-flat-depth | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-magic-array-flat-depth"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-magic-array-flat-depth.md for this page in Markdown format

# unicorn/no-magic-array-flat-depth Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-magic-array-flat-depth.html#unicorn-no-magic-array-flat-depth)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-magic-array-flat-depth.html#what-it-does)

Disallow magic numbers for `Array.prototype.flat` depth.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-magic-array-flat-depth.html#why-is-this-bad)

Magic numbers are hard to understand and maintain. When calling `Array.prototype.flat`, it is usually called with `1` or infinity. If you are using a different number, it is better to add a comment explaining the depth.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-magic-array-flat-depth.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
array.flat(2);
array.flat(20);
```

Examples of **correct** code for this rule:

javascript

```
array.flat(2 /* explanation */);
array.flat(1);
array.flat();
array.flat(Infinity);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-magic-array-flat-depth.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-magic-array-flat-depth": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-magic-array-flat-depth
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-magic-array-flat-depth.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_magic_array_flat_depth.rs)
