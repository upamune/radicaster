---
title: "typescript/non-nullable-type-assertion-style | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/non-nullable-type-assertion-style"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/non-nullable-type-assertion-style.md for this page in Markdown format

# typescript/non-nullable-type-assertion-style Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/non-nullable-type-assertion-style.html#typescript-non-nullable-type-assertion-style)

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/non-nullable-type-assertion-style.html#what-it-does)

This rule prefers a non-null assertion over an explicit type cast for non-nullable types.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/non-nullable-type-assertion-style.html#why-is-this-bad)

When you know that a value cannot be null or undefined, you can use either a non-null assertion (`!`) or a type assertion (`as Type`). The non-null assertion is more concise and clearly communicates the intent that you're asserting the value is not null/undefined.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/non-nullable-type-assertion-style.html#examples)

Examples of **incorrect** code for this rule:

ts

```
declare const value: string | null;

// Type assertion when non-null assertion would be clearer
const result1 = value as string;

declare const maybe: number | undefined;
const result2 = maybe as number;

// In function calls
function takesString(s: string) {
  console.log(s);
}

takesString(value as string);
```

Examples of **correct** code for this rule:

ts

```
declare const value: string | null;

// Non-null assertion for non-nullable types
const result1 = value!;

declare const maybe: number | undefined;
const result2 = maybe!;

// In function calls
function takesString(s: string) {
  console.log(s);
}

takesString(value!);

// Type assertion for actual type changes is still fine
declare const unknown: unknown;
const str = unknown as string; // This is a different type, not just removing null
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/non-nullable-type-assertion-style.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/non-nullable-type-assertion-style": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/non-nullable-type-assertion-style
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/non-nullable-type-assertion-style.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/non_nullable_type_assertion_style.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/non_nullable_type_assertion_style/non_nullable_type_assertion_style.go)
