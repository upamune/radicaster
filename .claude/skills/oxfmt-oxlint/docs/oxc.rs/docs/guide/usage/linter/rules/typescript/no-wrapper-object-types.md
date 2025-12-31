---
title: "typescript/no-wrapper-object-types | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-wrapper-object-types"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-wrapper-object-types.md for this page in Markdown format

# typescript/no-wrapper-object-types Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-wrapper-object-types.html#typescript-no-wrapper-object-types)

✅ This rule is turned on by default.

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-wrapper-object-types.html#what-it-does)

Disallow the use of wrapper object types.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-wrapper-object-types.html#why-is-this-bad)

Wrapper object types are types that are defined in the global scope and are not primitive types. These types are not recommended to be used in TypeScript code.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-wrapper-object-types.html#examples)

Examples of **incorrect** code for this rule:

ts

```
let myBigInt: BigInt;
let myBoolean: Boolean;
let myNumber: Number;
let myString: String;
let mySymbol: Symbol;

let myObject: Object = "allowed by TypeScript";
```

Examples of **correct** code for this rule:

ts

```
let myBigint: bigint;
let myBoolean: boolean;
let myNumber: number;
let myString: string;
let mySymbol: symbol;

let myObject: object = "Type 'string' is not assignable to type 'object'.";
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-wrapper-object-types.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-wrapper-object-types": "error"
  }
}
```

bash

```
oxlint --deny typescript/no-wrapper-object-types
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-wrapper-object-types.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_wrapper_object_types.rs)
