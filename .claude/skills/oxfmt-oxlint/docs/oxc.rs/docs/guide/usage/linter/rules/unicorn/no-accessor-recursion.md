---
title: "unicorn/no-accessor-recursion | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-accessor-recursion"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-accessor-recursion.md for this page in Markdown format

# unicorn/no-accessor-recursion Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-accessor-recursion.html#unicorn-no-accessor-recursion)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-accessor-recursion.html#what-it-does)

Disallow recursive access to this within getters and setters

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-accessor-recursion.html#why-is-this-bad)

This rule prevents recursive access to this within getter and setter methods in objects and classes, avoiding infinite recursion and stack overflow errors.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-accessor-recursion.html#examples)

Examples of **incorrect** code for this rule:

js

```
const foo = {
  get bar() {
    return this.bar;
  },
};
```

Examples of **correct** code for this rule:

js

```
const foo = {
  get bar() {
    return this.baz;
  },
};
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-accessor-recursion.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-accessor-recursion": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-accessor-recursion
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-accessor-recursion.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_accessor_recursion.rs)
