---
title: "typescript/no-unsafe-unary-minus | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-unary-minus"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-unsafe-unary-minus.md for this page in Markdown format

# typescript/no-unsafe-unary-minus Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-unary-minus.html#typescript-no-unsafe-unary-minus)

✅ This rule is turned on by default when type-aware linting is enabled.

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-unary-minus.html#what-it-does)

This rule disallows using the unary minus operator on a value which is not of type 'number' | 'bigint'.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-unary-minus.html#why-is-this-bad)

The unary minus operator should only be used on numeric values. Using it on other types can lead to unexpected behavior due to JavaScript's type coercion rules.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-unary-minus.html#examples)

Examples of **incorrect** code for this rule:

ts

```
declare const value: any;
const result1 = -value; // unsafe on any

declare const str: string;
const result2 = -str; // unsafe on string

declare const bool: boolean;
const result3 = -bool; // unsafe on boolean

declare const obj: object;
const result4 = -obj; // unsafe on object

declare const arr: any[];
const result5 = -arr; // unsafe on array
```

Examples of **correct** code for this rule:

ts

```
declare const num: number;
const result1 = -num; // safe

declare const bigint: bigint;
const result2 = -bigint; // safe

const literal = -42; // safe

const bigintLiteral = -42n; // safe

declare const union: number | bigint;
const result3 = -union; // safe

// Convert to number first if needed
declare const str: string;
const result4 = -Number(str); // safe conversion
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-unary-minus.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-unsafe-unary-minus": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/no-unsafe-unary-minus
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-unary-minus.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_unsafe_unary_minus.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/no_unsafe_unary_minus/no_unsafe_unary_minus.go)
