---
title: "eslint/no-return-assign | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-return-assign"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-return-assign.md for this page in Markdown format

# eslint/no-return-assign Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-return-assign.html#eslint-no-return-assign)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-return-assign.html#what-it-does)

Disallows assignment operators in return statements.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-return-assign.html#why-is-this-bad)

Assignment is allowed by js in return expressions, but usually, an expression with only one equal sign is intended to be a comparison. However, because of the missing equal sign, this turns to assignment, which is valid js code Because of this ambiguity, it’s considered a best practice to not use assignment in return statements.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-return-assign.html#examples)

Examples of **incorrect** code for this rule:

js

```
() => (a = b);
function x() {
  return (a = b);
}
```

Examples of **correct** code for this rule:

js

```
() => (a = b);
function x() {
  var result = (a = b);
  return result;
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-return-assign.html#configuration)

This rule accepts one of the following string values:

### `"always"` [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-return-assign.html#always)

Disallow all assignments in return statements.

### `"except-parens"` [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-return-assign.html#except-parens)

Allow assignments in return statements only if they are enclosed in parentheses. This is the default mode.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-return-assign.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-return-assign": "error"
  }
}
```

bash

```
oxlint --deny no-return-assign
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-return-assign.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_return_assign.rs)
