---
title: "unicorn/prefer-code-point | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-code-point"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-code-point.md for this page in Markdown format

# unicorn/prefer-code-point Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-code-point.html#unicorn-prefer-code-point)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-code-point.html#what-it-does)

Prefers usage of `String.prototype.codePointAt` over `String.prototype.charCodeAt`. Prefers usage of `String.fromCodePoint` over `String.fromCharCode`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-code-point.html#why-is-this-bad)

Unicode is better supported in [`String#codePointAt()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/codePointAt) and [`String.fromCodePoint()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/fromCodePoint).

[Difference between `String.fromCodePoint()` and `String.fromCharCode()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/fromCodePoint#compared_to_fromcharcode)

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-code-point.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
"🦄".charCodeAt(0);
String.fromCharCode(0x1f984);
```

Examples of **correct** code for this rule:

javascript

```
"🦄".codePointAt(0);
String.fromCodePoint(0x1f984);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-code-point.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-code-point": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-code-point
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-code-point.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_code_point.rs)
