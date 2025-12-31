---
title: "typescript/prefer-enum-initializers | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-enum-initializers"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/prefer-enum-initializers.md for this page in Markdown format

# typescript/prefer-enum-initializers Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-enum-initializers.html#typescript-prefer-enum-initializers)

💡 A suggestion is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-enum-initializers.html#what-it-does)

Require each enum member value to be explicitly initialized.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-enum-initializers.html#why-is-this-bad)

In projects where the value of `enum` members are important, allowing implicit values for enums can cause bugs if enums are modified over time.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-enum-initializers.html#examples)

Examples of **incorrect** code for this rule:

typescript

```
// wrong, the value of `Close` is not constant
enum Status {
  Open = 1,
  Close,
}
```

Examples of **correct** code for this rule:

typescript

```
enum Status {
  Open = 1,
  Close = 2,
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-enum-initializers.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/prefer-enum-initializers": "error"
  }
}
```

bash

```
oxlint --deny typescript/prefer-enum-initializers
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-enum-initializers.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/prefer_enum_initializers.rs)
