---
title: "typescript/no-duplicate-type-constituents | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-duplicate-type-constituents"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-duplicate-type-constituents.md for this page in Markdown format

# typescript/no-duplicate-type-constituents Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-duplicate-type-constituents.html#typescript-no-duplicate-type-constituents)

✅ This rule is turned on by default when type-aware linting is enabled.

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-duplicate-type-constituents.html#what-it-does)

This rule disallows duplicate constituents of union or intersection types.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-duplicate-type-constituents.html#why-is-this-bad)

Duplicate constituents in union and intersection types serve no purpose and can make code harder to read. They are likely a mistake.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-duplicate-type-constituents.html#examples)

Examples of **incorrect** code for this rule:

ts

```
type T1 = "A" | "A";

type T2 = A | A | B;

type T3 = { a: string } & { a: string };

type T4 = [A, A];

type T5 = "foo" | "bar" | "foo";
```

Examples of **correct** code for this rule:

ts

```
type T1 = "A" | "B";

type T2 = A | B | C;

type T3 = { a: string } & { b: string };

type T4 = [A, B];

type T5 = "foo" | "bar" | "baz";
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-duplicate-type-constituents.html#configuration)

This rule accepts a configuration object with the following properties:

### ignoreIntersections [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-duplicate-type-constituents.html#ignoreintersections)

type: `boolean`

default: `false`

Whether to ignore duplicate types in intersection types. When true, allows `type T = A & A`.

### ignoreUnions [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-duplicate-type-constituents.html#ignoreunions)

type: `boolean`

default: `false`

Whether to ignore duplicate types in union types. When true, allows `type T = A | A`.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-duplicate-type-constituents.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-duplicate-type-constituents": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/no-duplicate-type-constituents
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-duplicate-type-constituents.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_duplicate_type_constituents.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/no_duplicate_type_constituents/no_duplicate_type_constituents.go)
