---
title: "eslint/no-this-before-super | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-this-before-super"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-this-before-super.md for this page in Markdown format

# eslint/no-this-before-super Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-this-before-super.html#eslint-no-this-before-super)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-this-before-super.html#what-it-does)

Requires calling `super()` before using `this` or `super`.

This rule can be disabled for TypeScript code, as the TypeScript compiler enforces this check.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-this-before-super.html#why-is-this-bad)

In the constructor of derived classes, if `this`/`super` are used before `super()` calls, it raises a `ReferenceError`.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-this-before-super.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
class A1 extends B {
  constructor() {
    // super() needs to be called first
    this.a = 0;
    super();
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-this-before-super.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-this-before-super": "error"
  }
}
```

bash

```
oxlint --deny no-this-before-super
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-this-before-super.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_this_before_super.rs)
