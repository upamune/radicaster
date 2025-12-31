---
title: "typescript/no-unsafe-assignment | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-assignment"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-unsafe-assignment.md for this page in Markdown format

# typescript/no-unsafe-assignment Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-assignment.html#typescript-no-unsafe-assignment)

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-assignment.html#what-it-does)

This rule disallows assigning a value with type `any` to variables and properties.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-assignment.html#why-is-this-bad)

The `any` type in TypeScript disables type checking and can lead to runtime errors. When you assign an `any` value to a typed variable, you're essentially bypassing TypeScript's type safety without any guarantees about the actual value.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-assignment.html#examples)

Examples of **incorrect** code for this rule:

ts

```
declare const anyValue: any;

const str: string = anyValue; // unsafe assignment

let num: number;
num = anyValue; // unsafe assignment

const obj = {
  prop: anyValue as any, // unsafe assignment
};

interface User {
  name: string;
  age: number;
}

const user: User = anyValue; // unsafe assignment
```

Examples of **correct** code for this rule:

ts

```
declare const stringValue: string;
declare const numberValue: number;
declare const unknownValue: unknown;

const str: string = stringValue; // safe

let num: number;
num = numberValue; // safe

// Use type guards with unknown
if (typeof unknownValue === "string") {
  const str2: string = unknownValue; // safe after type guard
}

// Explicit any assignment (still not recommended, but intentional)
const anything: any = unknownValue;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-assignment.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-unsafe-assignment": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/no-unsafe-assignment
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-assignment.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_unsafe_assignment.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/no_unsafe_assignment/no_unsafe_assignment.go)
