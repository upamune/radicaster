---
title: "eslint/no-useless-concat | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-concat"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-useless-concat.md for this page in Markdown format

# eslint/no-useless-concat Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-concat.html#eslint-no-useless-concat)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-concat.html#what-it-does)

Disallow unnecessary concatenation of literals or template literals

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-concat.html#why-is-this-bad)

It’s unnecessary to concatenate two strings together when they could be combined into a single literal.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-concat.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
var foo = "a" + "b";
```

javascript

```
var foo = "a" + "b" + "c";
```

Examples of **correct** code for this rule:

javascript

```
var foo = "a" + bar;
```

// when the string concatenation is multiline

javascript

```
var foo = "a" + "b" + "c";
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-concat.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-useless-concat": "error"
  }
}
```

bash

```
oxlint --deny no-useless-concat
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-concat.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_useless_concat.rs)
