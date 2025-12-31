---
title: "typescript/no-unsafe-enum-comparison | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-enum-comparison"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-unsafe-enum-comparison.md for this page in Markdown format

# typescript/no-unsafe-enum-comparison Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-enum-comparison.html#typescript-no-unsafe-enum-comparison)

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-enum-comparison.html#what-it-does)

This rule disallows comparing an enum value with a non-enum value.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-enum-comparison.html#why-is-this-bad)

Enum values should only be compared with other values of the same enum type or their underlying literal values in a type-safe manner. Comparing enums with unrelated values can lead to unexpected behavior and defeats the purpose of using enums for type safety.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-enum-comparison.html#examples)

Examples of **incorrect** code for this rule:

ts

```
enum Status {
  Open = "open",
  Closed = "closed",
}

enum Color {
  Red = "red",
  Blue = "blue",
}

declare const status: Status;
declare const color: Color;
declare const str: string;

// Comparing enum with different enum
if (status === color) {
} // unsafe

// Comparing enum with string (unless it's a literal that matches)
if (status === str) {
} // unsafe

// Comparing with arbitrary value
if (status === "unknown") {
} // unsafe
```

Examples of **correct** code for this rule:

ts

```
enum Status {
  Open = "open",
  Closed = "closed",
}

declare const status: Status;

// Comparing with same enum values
if (status === Status.Open) {
} // safe

// Comparing with the correct literal type
if (status === "open") {
} // safe

// Using enum methods
if (Object.values(Status).includes(someValue)) {
} // safe way to check
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-enum-comparison.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-unsafe-enum-comparison": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/no-unsafe-enum-comparison
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-enum-comparison.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_unsafe_enum_comparison.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/no_unsafe_enum_comparison/no_unsafe_enum_comparison.go)
