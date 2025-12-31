---
title: "eslint/no-new | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-new"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-new.md for this page in Markdown format

# eslint/no-new Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-new.html#eslint-no-new)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-new.html#what-it-does)

Disallow new operators outside of assignments or comparisons.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-new.html#why-is-this-bad)

Calling new without assigning or comparing it the reference is thrown away and in many cases the constructor can be replaced with a function.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-new.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
new Person();

() => {
  new Date();
};
```

Examples of **correct** code for this rule:

javascript

```
var a = new Date()(() => new Date());
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-new.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-new": "error"
  }
}
```

bash

```
oxlint --deny no-new
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-new.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_new.rs)
