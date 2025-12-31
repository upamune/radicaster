---
title: "typescript/prefer-reduce-type-parameter | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-reduce-type-parameter"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/prefer-reduce-type-parameter.md for this page in Markdown format

# typescript/prefer-reduce-type-parameter Style [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-reduce-type-parameter.html#typescript-prefer-reduce-type-parameter)

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-reduce-type-parameter.html#what-it-does)

This rule prefers using a type parameter for the accumulator in Array.reduce instead of casting.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-reduce-type-parameter.html#why-is-this-bad)

Array.reduce can be called with a generic type parameter to specify the type of the accumulator. This is preferred over casting the result because it provides better type safety and is more explicit about the intended type.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-reduce-type-parameter.html#examples)

Examples of **incorrect** code for this rule:

ts

```
const numbers = [1, 2, 3];

// Casting the result
const sum = numbers.reduce((acc, val) => acc + val, 0) as number;

// Using type assertion on accumulator
const result = [1, 2, 3].reduce((acc: string[], curr) => {
  acc.push(curr.toString());
  return acc;
}, [] as string[]);
```

Examples of **correct** code for this rule:

ts

```
const numbers = [1, 2, 3];

// Using type parameter
const sum = numbers.reduce<number>((acc, val) => acc + val, 0);

// Type parameter for complex types
const result = [1, 2, 3].reduce<string[]>((acc, curr) => {
  acc.push(curr.toString());
  return acc;
}, []);

// When TypeScript can infer the type, no parameter needed
const simpleSum = numbers.reduce((acc, val) => acc + val, 0);

// Object accumulator with type parameter
interface Count {
  [key: string]: number;
}

const counts = ["a", "b", "a"].reduce<Count>((acc, item) => {
  acc[item] = (acc[item] || 0) + 1;
  return acc;
}, {});
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-reduce-type-parameter.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/prefer-reduce-type-parameter": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/prefer-reduce-type-parameter
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-reduce-type-parameter.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/prefer_reduce_type_parameter.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/prefer_reduce_type_parameter/prefer_reduce_type_parameter.go)
