---
title: "typescript/no-unsafe-function-type | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-function-type"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-unsafe-function-type.md for this page in Markdown format

# typescript/no-unsafe-function-type Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-function-type.html#typescript-no-unsafe-function-type)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-function-type.html#what-it-does)

Disallow using the unsafe built-in Function type.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-function-type.html#why-is-this-bad)

TypeScript's built-in Function type allows being called with any number of arguments and returns type any. Function also allows classes or plain objects that happen to possess all properties of the Function class. It's generally better to specify function parameters and return types with the function type syntax.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-function-type.html#examples)

Examples of **incorrect** code for this rule:

ts

```
let noParametersOrReturn: Function;
noParametersOrReturn = () => {};

let stringToNumber: Function;
stringToNumber = (text: string) => text.length;

let identity: Function;
identity = (value) => value;
```

Examples of **correct** code for this rule:

ts

```
let noParametersOrReturn: () => void;
noParametersOrReturn = () => {};

let stringToNumber: (text: string) => number;
stringToNumber = (text) => text.length;

let identity: <T>(value: T) => T;
identity = (value) => value;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-function-type.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-unsafe-function-type": "error"
  }
}
```

bash

```
oxlint --deny typescript/no-unsafe-function-type
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-function-type.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_unsafe_function_type.rs)
