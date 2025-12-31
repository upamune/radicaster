---
title: "unicorn/prefer-math-trunc | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-math-trunc"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-math-trunc.md for this page in Markdown format

# unicorn/prefer-math-trunc Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-math-trunc.html#unicorn-prefer-math-trunc)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-math-trunc.html#what-it-does)

Prefers use of [`Math.trunc()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Math/trunc) instead of bitwise operations for clarity and more reliable results.

It prevents the use of the following bitwise operations:

* `x | 0` ([`bitwise OR`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Bitwise_OR) with 0)
* `~~x` (two [`bitwise NOT`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Bitwise_NOT))
* `x >> 0` ([`Signed Right Shift`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Right_shift) with 0)
* `x << 0` ([`Left Shift`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Left_shift) with 0)
* `x ^ 0` ([`bitwise XOR Shift`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Bitwise_XOR) with 0)

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-math-trunc.html#why-is-this-bad)

Using bitwise operations to truncate numbers is not clear and do not work in [some cases](https://stackoverflow.com/a/34706108/11687747).

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-math-trunc.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
const foo = 1.1 | 0;
```

Examples of **correct** code for this rule:

javascript

```
const foo = Math.trunc(1.1);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-math-trunc.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-math-trunc": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-math-trunc
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-math-trunc.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_math_trunc.rs)
