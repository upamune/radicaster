---
title: "eslint/no-delete-var | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-delete-var"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-delete-var.md for this page in Markdown format

# eslint/no-delete-var Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-delete-var.html#eslint-no-delete-var)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-delete-var.html#what-it-does)

The purpose of the `delete` operator is to remove a property from an object.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-delete-var.html#why-is-this-bad)

Using the `delete` operator on a variable might lead to unexpected behavior.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-delete-var.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
var x;
delete x;
```

Examples of **correct** code for this rule:

javascript

```
var x;

var y;
delete y.prop;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-delete-var.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-delete-var": "error"
  }
}
```

bash

```
oxlint --deny no-delete-var
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-delete-var.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_delete_var.rs)
