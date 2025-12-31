---
title: "eslint/no-setter-return | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-setter-return"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-setter-return.md for this page in Markdown format

# eslint/no-setter-return Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-setter-return.html#eslint-no-setter-return)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-setter-return.html#what-it-does)

Setters cannot return values.

This rule can be disabled for TypeScript code, as the TypeScript compiler enforces this check.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-setter-return.html#why-is-this-bad)

While returning a value from a setter does not produce an error, the returned value is being ignored. Therefore, returning a value from a setter is either unnecessary or a possible error, since the returned value cannot be used.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-setter-return.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
class URL {
  set origin() {
    return true;
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-setter-return.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-setter-return": "error"
  }
}
```

bash

```
oxlint --deny no-setter-return
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-setter-return.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_setter_return.rs)
