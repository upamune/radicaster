---
title: "typescript/no-unsafe-argument | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-argument"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-unsafe-argument.md for this page in Markdown format

# typescript/no-unsafe-argument Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-argument.html#typescript-no-unsafe-argument)

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-argument.html#what-it-does)

This rule disallows calling a function with an argument which is typed as `any`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-argument.html#why-is-this-bad)

The `any` type in TypeScript is a dangerous "escape hatch" from the type system. Using `any` disables most type checking rules and is generally unsafe. When you pass a value typed as `any` to a function, you lose type safety for that function call.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-argument.html#examples)

Examples of **incorrect** code for this rule:

ts

```
declare const anyValue: any;

function takesString(str: string): void {
  console.log(str.length);
}

takesString(anyValue); // unsafe

declare function takesNumber(num: number): number;
const result = takesNumber(anyValue); // unsafe
```

Examples of **correct** code for this rule:

ts

```
declare const stringValue: string;
declare const numberValue: number;
declare const unknownValue: unknown;

function takesString(str: string): void {
  console.log(str.length);
}

takesString(stringValue); // safe

// Type guard to safely use unknown
if (typeof unknownValue === "string") {
  takesString(unknownValue); // safe after type guard
}

// Type assertion if you're sure about the type
takesString(unknownValue as string); // explicitly unsafe, but intentional
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-argument.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-unsafe-argument": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/no-unsafe-argument
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-argument.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_unsafe_argument.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/no_unsafe_argument/no_unsafe_argument.go)
