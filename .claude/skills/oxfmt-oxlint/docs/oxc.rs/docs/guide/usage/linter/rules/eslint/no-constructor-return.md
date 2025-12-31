---
title: "eslint/no-constructor-return | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-constructor-return"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-constructor-return.md for this page in Markdown format

# eslint/no-constructor-return Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-constructor-return.html#eslint-no-constructor-return)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-constructor-return.html#what-it-does)

Disallow returning value from constructor

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-constructor-return.html#why-is-this-bad)

In JavaScript, returning a value in the constructor of a class may be a mistake. Forbidding this pattern prevents mistakes resulting from unfamiliarity with the language or a copy-paste error.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-constructor-return.html#examples)

Examples of **incorrect** code for this rule:

js

```
class C {
  constructor() {
    return 42;
  }
}
```

Examples of **correct** code for this rule:

js

```
class C {
  constructor() {
    this.value = 42;
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-constructor-return.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-constructor-return": "error"
  }
}
```

bash

```
oxlint --deny no-constructor-return
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-constructor-return.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_constructor_return.rs)
