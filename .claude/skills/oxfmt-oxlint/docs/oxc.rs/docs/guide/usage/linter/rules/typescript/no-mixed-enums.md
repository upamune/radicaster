---
title: "typescript/no-mixed-enums | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-mixed-enums"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-mixed-enums.md for this page in Markdown format

# typescript/no-mixed-enums Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-mixed-enums.html#typescript-no-mixed-enums)

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-mixed-enums.html#what-it-does)

This rule disallows enums from having both string and numeric members.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-mixed-enums.html#why-is-this-bad)

TypeScript enums can have string, numeric, or computed members. Having mixed string and numeric members in the same enum can lead to confusion and unexpected runtime behavior due to how TypeScript compiles enums.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-mixed-enums.html#examples)

Examples of **incorrect** code for this rule:

ts

```
enum Status {
  Open = 1,
  Closed = "closed",
}

enum Direction {
  Up = "up",
  Down = 2,
  Left = "left",
  Right = 4,
}
```

Examples of **correct** code for this rule:

ts

```
// All numeric
enum Status {
  Open = 1,
  Closed = 2,
}

// All string
enum Direction {
  Up = "up",
  Down = "down",
  Left = "left",
  Right = "right",
}

// Auto-incremented numeric
enum Color {
  Red,
  Green,
  Blue,
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-mixed-enums.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-mixed-enums": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/no-mixed-enums
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-mixed-enums.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_mixed_enums.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/no_mixed_enums/no_mixed_enums.go)
