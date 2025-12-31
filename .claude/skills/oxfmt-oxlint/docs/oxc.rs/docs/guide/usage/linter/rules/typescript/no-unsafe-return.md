---
title: "typescript/no-unsafe-return | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-return"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-unsafe-return.md for this page in Markdown format

# typescript/no-unsafe-return Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-return.html#typescript-no-unsafe-return)

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-return.html#what-it-does)

This rule disallows returning a value with type `any` from a function.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-return.html#why-is-this-bad)

The `any` type in TypeScript disables type checking. When you return a value typed as `any` from a function, you're essentially passing the type-safety problem to the caller without providing any guarantees about what the function actually returns.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-return.html#examples)

Examples of **incorrect** code for this rule:

ts

```
declare const anyValue: any;

function getString(): string {
  return anyValue; // unsafe return
}

const getNumber = (): number => anyValue; // unsafe return

function processData(): { name: string; age: number } {
  return anyValue; // unsafe return
}
```

Examples of **correct** code for this rule:

ts

```
declare const stringValue: string;
declare const numberValue: number;
declare const unknownValue: unknown;

function getString(): string {
  return stringValue; // safe
}

const getNumber = (): number => numberValue; // safe

function processUnknown(): unknown {
  return unknownValue; // safe - explicitly returning unknown
}

// Type guard to safely return
function safeGetString(): string | null {
  if (typeof unknownValue === "string") {
    return unknownValue; // safe after type guard
  }
  return null;
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-return.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-unsafe-return": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/no-unsafe-return
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-return.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_unsafe_return.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/no_unsafe_return/no_unsafe_return.go)
