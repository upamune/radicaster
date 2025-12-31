---
title: "typescript/no-non-null-assertion | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-non-null-assertion"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-non-null-assertion.md for this page in Markdown format

# typescript/no-non-null-assertion Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-non-null-assertion.html#typescript-no-non-null-assertion)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-non-null-assertion.html#what-it-does)

Disallow non-null assertions using the ! postfix operator.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-non-null-assertion.html#why-is-this-bad)

TypeScript's ! non-null assertion operator asserts to the type system that an expression is non-nullable, as in not null or undefined. Using assertions to tell the type system new information is often a sign that code is not fully type-safe. It's generally better to structure program logic so that TypeScript understands when values may be nullable.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-non-null-assertion.html#examples)

Examples of **incorrect** code for this rule:

ts

```
x!;
x!.y;
x.y!;
```

Examples of **correct** code for this rule:

ts

```
x;
x?.y;
x.y;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-non-null-assertion.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-non-null-assertion": "error"
  }
}
```

bash

```
oxlint --deny typescript/no-non-null-assertion
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-non-null-assertion.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_non_null_assertion.rs)
