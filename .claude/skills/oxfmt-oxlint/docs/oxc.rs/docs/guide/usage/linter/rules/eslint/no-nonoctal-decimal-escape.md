---
title: "eslint/no-nonoctal-decimal-escape | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-nonoctal-decimal-escape"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-nonoctal-decimal-escape.md for this page in Markdown format

# eslint/no-nonoctal-decimal-escape Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-nonoctal-decimal-escape.html#eslint-no-nonoctal-decimal-escape)

✅ This rule is turned on by default.

💡 A suggestion is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-nonoctal-decimal-escape.html#what-it-does)

This rule disallows \8 and \9 escape sequences in string literals

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-nonoctal-decimal-escape.html#why-is-this-bad)

ECMAScript specification treats \8 and \9 in string literals as a legacy feature

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-nonoctal-decimal-escape.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
let x = "\8";
let y = "\9";
```

Examples of **correct** code for this rule:

javascript

```
let x = "8";
let y = "\\9";
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-nonoctal-decimal-escape.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-nonoctal-decimal-escape": "error"
  }
}
```

bash

```
oxlint --deny no-nonoctal-decimal-escape
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-nonoctal-decimal-escape.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_nonoctal_decimal_escape.rs)
