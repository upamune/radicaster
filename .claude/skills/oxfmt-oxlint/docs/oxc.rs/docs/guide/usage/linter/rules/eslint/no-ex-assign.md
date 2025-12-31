---
title: "eslint/no-ex-assign | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-ex-assign"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-ex-assign.md for this page in Markdown format

# eslint/no-ex-assign Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-ex-assign.html#eslint-no-ex-assign)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-ex-assign.html#what-it-does)

Disallow reassigning exceptions in catch clauses

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-ex-assign.html#why-is-this-bad)

If a catch clause in a try statement accidentally (or purposely) assigns another value to the exception parameter, it is impossible to refer to the error from that point on. Since there is no arguments object to offer alternative access to this data, assignment of the parameter is absolutely destructive.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-ex-assign.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
try {
  // code
} catch (e) {
  e = 10;
}
```

Examples of **correct** code for this rule:

javascript

```
try {
  // code
} catch (e) {
  let val = 10;
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-ex-assign.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-ex-assign": "error"
  }
}
```

bash

```
oxlint --deny no-ex-assign
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-ex-assign.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_ex_assign.rs)
