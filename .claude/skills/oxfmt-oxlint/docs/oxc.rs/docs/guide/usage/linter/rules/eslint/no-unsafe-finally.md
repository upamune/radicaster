---
title: "eslint/no-unsafe-finally | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unsafe-finally"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-unsafe-finally.md for this page in Markdown format

# eslint/no-unsafe-finally Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unsafe-finally.html#eslint-no-unsafe-finally)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unsafe-finally.html#what-it-does)

Disallow control flow statements in `finally` blocks.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unsafe-finally.html#why-is-this-bad)

JavaScript suspends the control flow statements of try and catch blocks until the execution of finally block finishes. So, when return, throw, break, or continue is used in finally, control flow statements inside try and catch are overwritten, which is considered as unexpected behavior.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unsafe-finally.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
// We expect this function to return 1;
(() => {
  try {
    return 1; // 1 is returned but suspended until finally block ends
  } catch (err) {
    return 2;
  } finally {
    return 3; // 3 is returned before 1, which we did not expect
  }
})();

// > 3
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unsafe-finally.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-unsafe-finally": "error"
  }
}
```

bash

```
oxlint --deny no-unsafe-finally
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unsafe-finally.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_unsafe_finally.rs)
