---
title: "typescript/no-unnecessary-type-arguments | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-type-arguments"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-unnecessary-type-arguments.md for this page in Markdown format

# typescript/no-unnecessary-type-arguments Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-type-arguments.html#typescript-no-unnecessary-type-arguments)

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-type-arguments.html#what-it-does)

This rule disallows type arguments that are identical to the default type parameter.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-type-arguments.html#why-is-this-bad)

Explicit type arguments that are the same as their default values are unnecessary and add visual noise to the code. TypeScript will infer these types automatically.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-type-arguments.html#examples)

Examples of **incorrect** code for this rule:

ts

```
function identity<T = string>(arg: T): T {
  return arg;
}

// Unnecessary type argument - string is the default
const result = identity<string>("hello");

interface Container<T = number> {
  value: T;
}

// Unnecessary type argument - number is the default
const container: Container<number> = { value: 42 };

class MyClass<T = boolean> {
  constructor(public value: T) {}
}

// Unnecessary type argument - boolean is the default
const instance = new MyClass<boolean>(true);
```

Examples of **correct** code for this rule:

ts

```
function identity<T = string>(arg: T): T {
  return arg;
}

// Using default type
const result1 = identity("hello");

// Using different type
const result2 = identity<number>(42);

interface Container<T = number> {
  value: T;
}

// Using default type
const container1: Container = { value: 42 };

// Using different type
const container2: Container<string> = { value: "hello" };
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-type-arguments.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-unnecessary-type-arguments": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/no-unnecessary-type-arguments
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-type-arguments.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_unnecessary_type_arguments.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/no_unnecessary_type_arguments/no_unnecessary_type_arguments.go)
