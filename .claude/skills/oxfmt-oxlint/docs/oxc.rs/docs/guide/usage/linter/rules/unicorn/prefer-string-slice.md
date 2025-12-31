---
title: "unicorn/prefer-string-slice | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-string-slice"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-string-slice.md for this page in Markdown format

# unicorn/prefer-string-slice Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-string-slice.html#unicorn-prefer-string-slice)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-string-slice.html#what-it-does)

Prefer [`String#slice()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/slice) over [`String#substr()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/substr) and [`String#substring()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/substring).

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-string-slice.html#why-is-this-bad)

[`String#substr()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/substr) and [`String#substring()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/substring) are the two lesser known legacy ways to slice a string. It's better to use [`String#slice()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/slice) as it's a more popular option with clearer behavior that has a consistent [`Array` counterpart](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/slice).

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-string-slice.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
"foo".substr(1, 2);
```

Examples of **correct** code for this rule:

javascript

```
"foo".slice(1, 2);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-string-slice.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-string-slice": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-string-slice
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-string-slice.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_string_slice.rs)
