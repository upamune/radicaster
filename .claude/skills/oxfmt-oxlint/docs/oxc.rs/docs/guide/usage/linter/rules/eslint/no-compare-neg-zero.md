---
title: "eslint/no-compare-neg-zero | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-compare-neg-zero"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-compare-neg-zero.md for this page in Markdown format

# eslint/no-compare-neg-zero Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-compare-neg-zero.html#eslint-no-compare-neg-zero)

✅ This rule is turned on by default.

🛠️💡 An auto-fix and a suggestion are available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-compare-neg-zero.html#what-it-does)

Disallow comparing against `-0`

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-compare-neg-zero.html#why-is-this-bad)

The rule should warn against code that tries to compare against `-0`, since that will not work as intended. That is, code like `x === -0` will pass for both `+0` and `-0`. The author probably intended `Object.is(x, -0)`.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-compare-neg-zero.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
if (x === -0) {
  // doSomething()...
}
```

javascript

```
if (-0 > x) {
  // doSomething()...
}
```

Examples of **correct** code for this rule:

javascript

```
if (x === 0) {
  // doSomething()...
}
```

javascript

```
if (Object.is(x, -0)) {
  // doSomething()...
}
```

javascript

```
if (0 > x) {
  // doSomething()...
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-compare-neg-zero.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-compare-neg-zero": "error"
  }
}
```

bash

```
oxlint --deny no-compare-neg-zero
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-compare-neg-zero.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_compare_neg_zero.rs)
