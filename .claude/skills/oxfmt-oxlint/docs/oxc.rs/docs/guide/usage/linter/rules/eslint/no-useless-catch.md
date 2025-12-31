---
title: "eslint/no-useless-catch | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-catch"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-useless-catch.md for this page in Markdown format

# eslint/no-useless-catch Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-catch.html#eslint-no-useless-catch)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-catch.html#what-it-does)

Disallow unnecessary catch clauses

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-catch.html#why-is-this-bad)

A catch clause that only rethrows the original error is redundant, and has no effect on the runtime behavior of the program. These redundant clauses can be a source of confusion and code bloat, so it’s better to disallow these unnecessary catch clauses.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-catch.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
try {
  doSomethingThatMightThrow();
} catch (e) {
  throw e;
}
```

Examples of **correct** code for this rule:

javascript

```
doSomethingThatMightThrow();
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-catch.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-useless-catch": "error"
  }
}
```

bash

```
oxlint --deny no-useless-catch
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-catch.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_useless_catch.rs)
