---
title: "typescript/no-redundant-type-constituents | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-redundant-type-constituents"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-redundant-type-constituents.md for this page in Markdown format

# typescript/no-redundant-type-constituents Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-redundant-type-constituents.html#typescript-no-redundant-type-constituents)

✅ This rule is turned on by default when type-aware linting is enabled.

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-redundant-type-constituents.html#what-it-does)

This rule disallows type constituents of unions and intersections that are redundant.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-redundant-type-constituents.html#why-is-this-bad)

Some constituents of union and intersection types can be redundant due to TypeScript's type system rules. These redundant constituents don't add any value and can make types harder to read and understand.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-redundant-type-constituents.html#examples)

Examples of **incorrect** code for this rule:

ts

```
// unknown is redundant in unions
type T1 = string | unknown;

// any is redundant in unions
type T2 = string | any;

// never is redundant in unions
type T3 = string | never;

// Literal types that are wider than other types
type T4 = string | "hello";

// Object types that are subsets
type T5 = { a: string } | { a: string; b: number };
```

Examples of **correct** code for this rule:

ts

```
type T1 = string | number;

type T2 = "hello" | "world";

type T3 = { a: string } | { b: number };

// unknown in intersections is meaningful
type T4 = string & unknown;

// never in intersections is meaningful
type T5 = string & never;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-redundant-type-constituents.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-redundant-type-constituents": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/no-redundant-type-constituents
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-redundant-type-constituents.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_redundant_type_constituents.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/no_redundant_type_constituents/no_redundant_type_constituents.go)
