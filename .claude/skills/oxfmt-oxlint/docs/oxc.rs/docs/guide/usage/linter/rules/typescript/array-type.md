---
title: "typescript/array-type | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/array-type"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/array-type.md for this page in Markdown format

# typescript/array-type Style [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/array-type.html#typescript-array-type)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/array-type.html#what-it-does)

Require consistently using either `T[]` or `Array<T>` for arrays.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/array-type.html#why-is-this-bad)

Using the `Array` type directly is not idiomatic. Instead, use the array type `T[]` or `Array<T>`.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/array-type.html#examples)

Examples of **incorrect** code for this rule:

typescript

```
/*oxlint array-type: ["error", { "default": "array" }] */
const arr: Array<number> = new Array<number>();
```

typescript

```
/*oxlint array-type: ["error", { "default": "generic" }] */
const arr: number[] = new Array<number>();
```

typescript

```
/*oxlint array-type: ["error", { "default": "array-simple" }] */
const a: (string | number)[] = ["a", "b"];
const b: { prop: string }[] = [{ prop: "a" }];
const c: Array<MyType> = ["a", "b"];
const d: Array<string> = ["a", "b"];
```

Examples of **correct** code for this rule:

typescript

```
/*oxlint array-type: ["error", { "default": "array" }] */
const arr: number[] = new Array<number>();
```

typescript

```
/*oxlint array-type: ["error", { "default": "generic" }] */
const arr: Array<number> = new Array<number>();
```

typescript

```
/*oxlint array-type: ["error", { "default": "array-simple" }] */
const a: Array<string | number> = ["a", "b"];
const b: Array<{ prop: string }> = [{ prop: "a" }];
const c: string[] = ["a", "b"];
const d: MyType[] = ["a", "b"];
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/array-type.html#configuration)

This rule accepts a configuration object with the following properties:

### default [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/array-type.html#default)

type: `"array" | "array-simple" | "generic"`

default: `"array"`

The array type expected for mutable cases.

### readonly [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/array-type.html#readonly)

type: `"array" | "array-simple" | "generic"`

default: `null`

The array type expected for readonly cases. If omitted, the value for `default` will be used.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/array-type.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/array-type": "error"
  }
}
```

bash

```
oxlint --deny typescript/array-type
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/array-type.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/array_type.rs)
