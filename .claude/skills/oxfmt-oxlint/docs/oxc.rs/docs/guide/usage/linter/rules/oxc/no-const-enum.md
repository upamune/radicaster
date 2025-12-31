---
title: "oxc/no-const-enum | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-const-enum"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/oxc/no-const-enum.md for this page in Markdown format

# oxc/no-const-enum Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-const-enum.html#oxc-no-const-enum)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-const-enum.html#what-it-does)

Disallow TypeScript `const enum`

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-const-enum.html#why-is-this-bad)

Const enums are enums that should be inlined at use sites. Const enums are not supported by bundlers and are incompatible with the isolatedModules mode. Their use can lead to import nonexistent values (because const enums are erased).

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-const-enum.html#examples)

Examples of **incorrect** code for this rule:

ts

```
const enum Color {
  Red,
  Green,
  Blue,
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-const-enum.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "oxc/no-const-enum": "error"
  }
}
```

bash

```
oxlint --deny oxc/no-const-enum
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-const-enum.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/oxc/no_const_enum.rs)
