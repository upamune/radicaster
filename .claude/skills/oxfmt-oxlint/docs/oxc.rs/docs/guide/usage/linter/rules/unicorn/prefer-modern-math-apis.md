---
title: "unicorn/prefer-modern-math-apis | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-modern-math-apis"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-modern-math-apis.md for this page in Markdown format

# unicorn/prefer-modern-math-apis Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-modern-math-apis.html#unicorn-prefer-modern-math-apis)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-modern-math-apis.html#what-it-does)

Checks for usage of legacy patterns for mathematical operations.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-modern-math-apis.html#why-is-this-bad)

Modern JavaScript provides more concise and readable alternatives to legacy patterns.

Currently, the following cases are checked:

* Prefer `Math.log10(x)` over alternatives
* Prefer `Math.hypot(…)` over alternatives

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-modern-math-apis.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
Math.log(x) * Math.LOG10E;
Math.sqrt(a * a + b * b);
```

Examples of **correct** code for this rule:

javascript

```
Math.log10(x);
Math.hypot(a, b);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-modern-math-apis.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-modern-math-apis": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-modern-math-apis
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-modern-math-apis.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_modern_math_apis.rs)
