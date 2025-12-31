---
title: "unicorn/no-useless-length-check | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-length-check"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-useless-length-check.md for this page in Markdown format

# unicorn/no-useless-length-check Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-length-check.html#unicorn-no-useless-length-check)

✅ This rule is turned on by default.

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-length-check.html#what-it-does)

It checks for an unnecessary array length check in a logical expression.

The cases are:

* `array.length === 0 || array.every(Boolean)` (`array.every` returns `true` if array is has elements)
* `array.length > 0 && array.some(Boolean)` (`array.some` returns `false` if array is empty)

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-length-check.html#why-is-this-bad)

An extra unnecessary length check is done.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-length-check.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
if (array.length === 0 || array.every(Boolean)) {
  // do something!
}
```

Examples of **correct** code for this rule:

javascript

```
if (array.every(Boolean)) {
  // do something!
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-length-check.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-useless-length-check": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-useless-length-check
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-length-check.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_useless_length_check.rs)
