---
title: "unicorn/prefer-regexp-test | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-regexp-test"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-regexp-test.md for this page in Markdown format

# unicorn/prefer-regexp-test Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-regexp-test.html#unicorn-prefer-regexp-test)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-regexp-test.html#what-it-does)

Prefers `RegExp#test()` over `String#match()` and `String#exec()`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-regexp-test.html#why-is-this-bad)

When you want to know whether a pattern is found in a string, use [`RegExp#test()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/RegExp/test) instead of [`String#match()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/match) or [`RegExp#exec()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/RegExp/exec), as it exclusively returns a boolean and therefore is more efficient.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-regexp-test.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
if (string.match(/unicorn/)) {
}
if (/unicorn/.exec(string)) {
}
```

Examples of **correct** code for this rule:

javascript

```
if (/unicorn/.test(string)) {
}
Boolean(string.match(/unicorn/));
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-regexp-test.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-regexp-test": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-regexp-test
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-regexp-test.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_regexp_test.rs)
