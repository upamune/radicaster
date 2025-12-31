---
title: "eslint/prefer-numeric-literals | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-numeric-literals"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/prefer-numeric-literals.md for this page in Markdown format

# eslint/prefer-numeric-literals Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-numeric-literals.html#eslint-prefer-numeric-literals)

🛠️ An auto-fix is available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-numeric-literals.html#what-it-does)

Disallow parseInt() and Number.parseInt() in favor of binary, octal, and hexadecimal literals.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-numeric-literals.html#why-is-this-bad)

The parseInt() and Number.parseInt() functions can be used to turn binary, octal, and hexadecimal strings into integers. As binary, octal, and hexadecimal literals are supported in ES2015, this rule encourages use of those numeric literals instead of parseInt() or Number.parseInt().

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-numeric-literals.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
parseInt("111110111", 2) === 503;
parseInt(`111110111`, 2) === 503;
parseInt("767", 8) === 503;
parseInt("1F7", 16) === 503;
Number.parseInt("111110111", 2) === 503;
Number.parseInt("767", 8) === 503;
Number.parseInt("1F7", 16) === 503;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-numeric-literals.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "prefer-numeric-literals": "error"
  }
}
```

bash

```
oxlint --deny prefer-numeric-literals
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-numeric-literals.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/prefer_numeric_literals.rs)
