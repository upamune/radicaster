---
title: "unicorn/prefer-math-min-max | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-math-min-max"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-math-min-max.md for this page in Markdown format

# unicorn/prefer-math-min-max Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-math-min-max.html#unicorn-prefer-math-min-max)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-math-min-max.html#what-it-does)

Prefers use of `Math.min()` and `Math.max()` instead of ternary expressions when performing simple comparisons.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-math-min-max.html#why-is-this-bad)

Using `Math.min()` and `Math.max()` for simple comparisons is more concise, easier to understand, and less prone to errors than ternary expressions. They clearly express the intent to find the minimum or maximum value.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-math-min-max.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
height > 50 ? 50 : height;
height > 50 ? height : 50;
```

Examples of **correct** code for this rule:

javascript

```
Math.min(height, 50);
Math.max(height, 50);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-math-min-max.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-math-min-max": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-math-min-max
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-math-min-max.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_math_min_max.rs)
