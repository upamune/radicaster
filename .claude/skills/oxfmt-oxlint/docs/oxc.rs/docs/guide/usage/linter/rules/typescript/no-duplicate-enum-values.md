---
title: "typescript/no-duplicate-enum-values | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-duplicate-enum-values"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-duplicate-enum-values.md for this page in Markdown format

# typescript/no-duplicate-enum-values Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-duplicate-enum-values.html#typescript-no-duplicate-enum-values)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-duplicate-enum-values.html#what-it-does)

Disallow duplicate enum member values.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-duplicate-enum-values.html#why-is-this-bad)

Although TypeScript supports duplicate enum member values, people usually expect members to have unique values within the same enum. Duplicate values can lead to bugs that are hard to track down.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-duplicate-enum-values.html#examples)

This rule disallows defining an enum with multiple members initialized to the same value. Members without initializers will not be checked.

Example of **incorrect** code:

ts

```
enum E {
  A = 0,
  B = 0,
}
```

ts

```
enum E {
  A = "A",
  B = "A",
}
```

Example of **correct** code:

ts

```
enum E {
  A = 0,
  B = 1,
}
```

ts

```
enum E {
  A = "A",
  B = "B",
}
```

ts

```
enum E {
  A,
  B,
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-duplicate-enum-values.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-duplicate-enum-values": "error"
  }
}
```

bash

```
oxlint --deny typescript/no-duplicate-enum-values
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-duplicate-enum-values.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_duplicate_enum_values.rs)
