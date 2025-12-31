---
title: "typescript/no-confusing-non-null-assertion | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-confusing-non-null-assertion"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-confusing-non-null-assertion.md for this page in Markdown format

# typescript/no-confusing-non-null-assertion Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-confusing-non-null-assertion.html#typescript-no-confusing-non-null-assertion)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-confusing-non-null-assertion.html#what-it-does)

Disallow non-null assertion in locations that may be confusing.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-confusing-non-null-assertion.html#why-is-this-bad)

Using a non-null assertion (!) next to an assign or equals check (= or == or ===) creates code that is confusing as it looks similar to a not equals check (!= !==).

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-confusing-non-null-assertion.html#examples)

Examples of **incorrect** code for this rule:

ts

```
a! == b; // a non-null assertions(`!`) and an equals test(`==`)
a !== b; // not equals test(`!==`)
a! === b; // a non-null assertions(`!`) and an triple equals test(`===`)
```

Examples of **correct** code for this rule:

ts

```
a == b;
a !== b;
a === b;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-confusing-non-null-assertion.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-confusing-non-null-assertion": "error"
  }
}
```

bash

```
oxlint --deny typescript/no-confusing-non-null-assertion
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-confusing-non-null-assertion.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_confusing_non_null_assertion.rs)
