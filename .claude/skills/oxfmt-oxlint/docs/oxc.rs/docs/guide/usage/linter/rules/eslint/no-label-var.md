---
title: "eslint/no-label-var | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-label-var"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-label-var.md for this page in Markdown format

# eslint/no-label-var Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-label-var.html#eslint-no-label-var)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-label-var.html#what-it-does)

Disallow labels that share a name with a variable.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-label-var.html#why-is-this-bad)

This rule aims to create clearer code by disallowing the bad practice of creating a label that shares a name with a variable that is in scope.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-label-var.html#examples)

Examples of **incorrect** code for this rule:

js

```
var x = foo;
function bar() {
  x: for (;;) {
    break x;
  }
}
```

Examples of **correct** code for this rule:

js

```
// The variable that has the same name as the label is not in scope.

function foo() {
  var q = t;
}

function bar() {
  q: for (;;) {
    break q;
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-label-var.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-label-var": "error"
  }
}
```

bash

```
oxlint --deny no-label-var
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-label-var.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_label_var.rs)
