---
title: "eslint/no-irregular-whitespace | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-irregular-whitespace"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-irregular-whitespace.md for this page in Markdown format

# eslint/no-irregular-whitespace Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-irregular-whitespace.html#eslint-no-irregular-whitespace)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-irregular-whitespace.html#what-it-does)

Disallows the use of irregular whitespace characters in the code.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-irregular-whitespace.html#why-is-this-bad)

Irregular whitespace characters are invisible to most editors and can cause unexpected behavior, making code harder to debug and maintain. They can also cause issues with code formatting and parsing.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-irregular-whitespace.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
// Contains irregular whitespace characters (invisible)
function example() {
  var foo = "bar"; // irregular whitespace before 'bar'
}
```

Examples of **correct** code for this rule:

javascript

```
function example() {
  var foo = "bar"; // regular spaces only
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-irregular-whitespace.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-irregular-whitespace": "error"
  }
}
```

bash

```
oxlint --deny no-irregular-whitespace
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-irregular-whitespace.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_irregular_whitespace.rs)
