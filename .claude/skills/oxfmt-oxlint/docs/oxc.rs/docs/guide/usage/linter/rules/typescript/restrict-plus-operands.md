---
title: "typescript/restrict-plus-operands | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/restrict-plus-operands"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/restrict-plus-operands.md for this page in Markdown format

# typescript/restrict-plus-operands Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/restrict-plus-operands.html#typescript-restrict-plus-operands)

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/restrict-plus-operands.html#what-it-does)

This rule requires both operands of addition to be the same type and be number, string, or any.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/restrict-plus-operands.html#why-is-this-bad)

JavaScript's + operator can be used for both numeric addition and string concatenation. When the operands are of different types, JavaScript's type coercion rules can lead to unexpected results. This rule helps prevent these issues by requiring both operands to be of compatible types.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/restrict-plus-operands.html#examples)

Examples of **incorrect** code for this rule:

ts

```
declare const num: number;
declare const str: string;
declare const bool: boolean;
declare const obj: object;

// Mixed types
const result1 = num + str; // number + string
const result2 = str + bool; // string + boolean
const result3 = num + bool; // number + boolean
const result4 = obj + str; // object + string

// Literals with different types
const result5 = 42 + "hello"; // number literal + string literal
const result6 = true + 5; // boolean literal + number literal
```

Examples of **correct** code for this rule:

ts

```
declare const num1: number;
declare const num2: number;
declare const str1: string;
declare const str2: string;

// Same types
const sum = num1 + num2; // number + number
const concat = str1 + str2; // string + string

// Explicit conversions
const result1 = num1 + String(num2); // Convert to string first
const result2 = String(num1) + str1; // Convert to string first
const result3 = Number(str1) + num1; // Convert to number first

// Template literals for string concatenation
const result4 = `${num1}${str1}`; // Clear intent to concatenate

// Literals of same type
const numResult = 42 + 58; // number + number
const strResult = "hello" + "world"; // string + string
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/restrict-plus-operands.html#configuration)

This rule accepts a configuration object with the following properties:

### allowAny [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/restrict-plus-operands.html#allowany)

type: `boolean`

default: `true`

Whether to allow `any` type in plus operations.

### allowBoolean [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/restrict-plus-operands.html#allowboolean)

type: `boolean`

default: `true`

Whether to allow `boolean` types in plus operations.

### allowNullish [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/restrict-plus-operands.html#allownullish)

type: `boolean`

default: `true`

Whether to allow nullish types (`null` or `undefined`) in plus operations.

### allowNumberAndString [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/restrict-plus-operands.html#allownumberandstring)

type: `boolean`

default: `true`

Whether to allow mixed number and string operands in plus operations.

### allowRegExp [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/restrict-plus-operands.html#allowregexp)

type: `boolean`

default: `true`

Whether to allow `RegExp` types in plus operations.

### skipCompoundAssignments [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/restrict-plus-operands.html#skipcompoundassignments)

type: `boolean`

default: `false`

Whether to skip compound assignments (e.g., `a += b`).

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/restrict-plus-operands.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/restrict-plus-operands": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/restrict-plus-operands
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/restrict-plus-operands.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/restrict_plus_operands.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/restrict_plus_operands/restrict_plus_operands.go)
