---
title: "eslint/no-dupe-keys | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-dupe-keys"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-dupe-keys.md for this page in Markdown format

# eslint/no-dupe-keys Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-dupe-keys.html#eslint-no-dupe-keys)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-dupe-keys.html#what-it-does)

Disallow duplicate keys in object literals.

This rule can be disabled for TypeScript code, as the TypeScript compiler enforces this check.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-dupe-keys.html#why-is-this-bad)

Multiple properties with the same key in object literals can cause unexpected behavior in your application.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-dupe-keys.html#examples)

Examples of **incorrect** code for this rule:

js

```
var foo = {
  bar: "baz",
  bar: "qux",
};

var foo = {
  bar: "baz",
  bar: "qux",
};

var foo = {
  0x1: "baz",
  1: "qux",
};
```

Examples of **correct** code for this rule:

js

```
var foo = {
  bar: "baz",
  qux: "qux",
};
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-dupe-keys.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-dupe-keys": "error"
  }
}
```

bash

```
oxlint --deny no-dupe-keys
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-dupe-keys.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_dupe_keys.rs)
