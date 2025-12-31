---
title: "eslint/no-var | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-var"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-var.md for this page in Markdown format

# eslint/no-var Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-var.html#eslint-no-var)

🛠️ An auto-fix is available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-var.html#what-it-does)

ECMAScript 2015 allows programmers to create variables with block scope instead of function scope using the `let` and `const` keywords. Block scope is common in many other programming languages and helps programmers avoid mistakes.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-var.html#why-is-this-bad)

Using `var` in an ES2015 environment triggers this error

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-var.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
var x = "y";
var CONFIG = {};
```

Examples of **correct** code for this rule:

javascript

```
let x = "y";
const CONFIG = {};
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-var.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-var": "error"
  }
}
```

bash

```
oxlint --deny no-var
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-var.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_var.rs)
