---
title: "eslint/preserve-caught-error | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/preserve-caught-error"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/preserve-caught-error.md for this page in Markdown format

# eslint/preserve-caught-error Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/preserve-caught-error.html#eslint-preserve-caught-error)

🛠️ An auto-fix is available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/preserve-caught-error.html#what-it-does)

Enforces that when re-throwing an error in a catch block, the original error is preserved using the 'cause' property.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/preserve-caught-error.html#why-is-this-bad)

Re-throwing an error without preserving the original error loses important debugging information and makes it harder to trace the root cause of issues.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/preserve-caught-error.html#examples)

Examples of **incorrect** code for this rule:

js

```
try {
  doSomething();
} catch (err) {
  throw new Error("Something failed");
}
```

Examples of **correct** code for this rule:

js

```
try {
  doSomething();
} catch (err) {
  throw new Error("Something failed", { cause: err });
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/preserve-caught-error.html#configuration)

This rule accepts a configuration object with the following properties:

### requireCatchParameter [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/preserve-caught-error.html#requirecatchparameter)

type: `boolean`

default: `false`

When set to `true`, requires that catch clauses always have a parameter.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/preserve-caught-error.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "preserve-caught-error": "error"
  }
}
```

bash

```
oxlint --deny preserve-caught-error
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/preserve-caught-error.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/preserve_caught_error.rs)
