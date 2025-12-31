---
title: "unicorn/no-array-reduce | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-reduce"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-array-reduce.md for this page in Markdown format

# unicorn/no-array-reduce Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-reduce.html#unicorn-no-array-reduce)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-reduce.html#what-it-does)

Disallow `Array#reduce()` and `Array#reduceRight()`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-reduce.html#why-is-this-bad)

`Array#reduce()` and `Array#reduceRight()` usually result in [hard-to-read](https://twitter.com/jaffathecake/status/1213077702300852224) and [less performant](https://www.richsnapp.com/article/2019/06-09-reduce-spread-anti-pattern) code. In almost every case, it can be replaced by `.map`, `.filter`, or a `for-of` loop.

It's only somewhat useful in the rare case of summing up numbers, which is allowed by default.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-reduce.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
array.reduce(reducer, initialValue);
array.reduceRight(reducer, initialValue);
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-reduce.html#configuration)

This rule accepts a configuration object with the following properties:

### allowSimpleOperations [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-reduce.html#allowsimpleoperations)

type: `boolean`

default: `true`

When set to `true`, allows simple operations (like summing numbers) in `reduce` and `reduceRight` calls.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-reduce.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-array-reduce": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-array-reduce
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-reduce.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_array_reduce.rs)
