---
title: "typescript/prefer-function-type | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-function-type"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/prefer-function-type.md for this page in Markdown format

# typescript/prefer-function-type Style [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-function-type.html#typescript-prefer-function-type)

🛠️ An auto-fix is available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-function-type.html#what-it-does)

Enforce using function types instead of interfaces with call signatures.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-function-type.html#why-is-this-bad)

TypeScript allows for two common ways to declare a type for a function:

* Function type: `() => string`
* Object type with a signature: `{ (): string }`

The function type form is generally preferred when possible for being more succinct and readable. Interfaces with only call signatures add unnecessary verbosity without providing additional functionality.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-function-type.html#examples)

Examples of **incorrect** code for this rule:

typescript

```
interface Example {
  (): string;
}

function foo(example: { (): number }): number {
  return example();
}

interface ReturnsSelf {
  (arg: string): this;
}
```

Examples of **correct** code for this rule:

typescript

```
type Example = () => string;

function foo(example: () => number): number {
  return example();
}

// Returns the function itself, not the `this` argument
type ReturnsSelf = (arg: string) => ReturnsSelf;

// Multiple properties are allowed
function foo(bar: { (): string; baz: number }): string {
  return bar();
}

// Multiple call signatures (overloads) are allowed
interface Overloaded {
  (data: string): number;
  (id: number): string;
}

// this is equivalent to Overloaded interface.
type Intersection = ((data: string) => number) & ((id: number) => string);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-function-type.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/prefer-function-type": "error"
  }
}
```

bash

```
oxlint --deny typescript/prefer-function-type
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-function-type.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/prefer_function_type.rs)
