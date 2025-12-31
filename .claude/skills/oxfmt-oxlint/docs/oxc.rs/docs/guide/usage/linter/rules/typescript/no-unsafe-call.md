---
title: "typescript/no-unsafe-call | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-call"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-unsafe-call.md for this page in Markdown format

# typescript/no-unsafe-call Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-call.html#typescript-no-unsafe-call)

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-call.html#what-it-does)

This rule disallows calling a value with type `any`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-call.html#why-is-this-bad)

The `any` type in TypeScript disables type checking. When you call a value typed as `any`, TypeScript cannot verify that it's actually a function, what parameters it expects, or what it returns. This can lead to runtime errors.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-call.html#examples)

Examples of **incorrect** code for this rule:

ts

```
declare const anyValue: any;

anyValue(); // unsafe call

anyValue(1, 2, 3); // unsafe call

const result = anyValue("hello"); // unsafe call

// Chained unsafe calls
anyValue().then().catch(); // unsafe
```

Examples of **correct** code for this rule:

ts

```
declare const fn: () => void;
declare const fnWithParams: (a: number, b: string) => boolean;
declare const unknownValue: unknown;

fn(); // safe

const result = fnWithParams(1, "hello"); // safe

// Type guard for unknown
if (typeof unknownValue === "function") {
  unknownValue(); // safe after type guard
}

// Explicit type assertion if you're certain
(anyValue as () => void)(); // explicitly unsafe but intentional
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-call.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-unsafe-call": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/no-unsafe-call
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-call.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_unsafe_call.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/no_unsafe_call/no_unsafe_call.go)
