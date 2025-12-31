---
title: "unicorn/no-nested-ternary | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-nested-ternary"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-nested-ternary.md for this page in Markdown format

# unicorn/no-nested-ternary Style [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-nested-ternary.html#unicorn-no-nested-ternary)

🛠️ An auto-fix is available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-nested-ternary.html#what-it-does)

This rule disallows deeply nested ternary expressions. Nested ternary expressions that are only one level deep and wrapped in parentheses are allowed.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-nested-ternary.html#why-is-this-bad)

Nesting ternary expressions can make code more difficult to understand.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-nested-ternary.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
const foo = i > 5 ? (i < 100 ? true : false) : true;
const foo = i > 5 ? true : i < 100 ? true : i < 1000 ? true : false;
```

Examples of **correct** code for this rule:

javascript

```
const foo = i > 5 ? (i < 100 ? true : false) : true;
const foo = i > 5 ? (i < 100 ? true : false) : i < 100 ? true : false;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-nested-ternary.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-nested-ternary": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-nested-ternary
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-nested-ternary.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_nested_ternary.rs)
