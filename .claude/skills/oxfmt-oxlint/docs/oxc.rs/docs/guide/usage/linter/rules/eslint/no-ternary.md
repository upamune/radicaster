---
title: "eslint/no-ternary | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-ternary"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-ternary.md for this page in Markdown format

# eslint/no-ternary Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-ternary.html#eslint-no-ternary)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-ternary.html#what-it-does)

Disallow ternary operators

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-ternary.html#why-is-this-bad)

The ternary operator is used to conditionally assign a value to a variable. Some believe that the use of ternary operators leads to unclear code.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-ternary.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
var foo = isBar ? baz : qux;
```

javascript

```
function quux() {
  return foo ? bar() : baz();
}
```

Examples of **correct** code for this rule:

javascript

```
let foo;

if (isBar) {
  foo = baz;
} else {
  foo = qux;
}
```

javascript

```
function quux() {
  if (foo) {
    return bar();
  } else {
    return baz();
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-ternary.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-ternary": "error"
  }
}
```

bash

```
oxlint --deny no-ternary
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-ternary.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_ternary.rs)
