---
title: "eslint/use-isnan | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/use-isnan"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/use-isnan.md for this page in Markdown format

# eslint/use-isnan Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/use-isnan.html#eslint-use-isnan)

✅ This rule is turned on by default.

🛠️ An auto-fix is available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/use-isnan.html#what-it-does)

Disallows checking against NaN without using `isNaN()` call.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/use-isnan.html#why-is-this-bad)

In JavaScript, NaN is a special value of the Number type. It’s used to represent any of the “not-a-number” values represented by the double-precision 64-bit format as specified by the IEEE Standard for Binary Floating-Point Arithmetic.

Because NaN is unique in JavaScript by not being equal to anything, including itself, the results of comparisons to NaN are confusing:

* `NaN === NaN` or `NaN == NaN` evaluate to false
* `NaN !== NaN` or `NaN != NaN` evaluate to true

Therefore, use `Number.isNaN()` or global `isNaN()` functions to test whether a value is NaN.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/use-isnan.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
foo == NaN;
foo === NaN;
foo <= NaN;
foo > NaN;
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/use-isnan.html#configuration)

This rule accepts a configuration object with the following properties:

### enforceForIndexOf [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/use-isnan.html#enforceforindexof)

type: `boolean`

default: `false`

Whether to disallow NaN as arguments of `indexOf` and `lastIndexOf`

### enforceForSwitchCase [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/use-isnan.html#enforceforswitchcase)

type: `boolean`

default: `true`

Whether to disallow NaN in switch cases and discriminants

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/use-isnan.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "use-isnan": "error"
  }
}
```

bash

```
oxlint --deny use-isnan
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/use-isnan.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/use_isnan.rs)
