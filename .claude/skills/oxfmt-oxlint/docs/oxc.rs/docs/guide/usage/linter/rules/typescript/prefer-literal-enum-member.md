---
title: "typescript/prefer-literal-enum-member | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-literal-enum-member"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/prefer-literal-enum-member.md for this page in Markdown format

# typescript/prefer-literal-enum-member Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-literal-enum-member.html#typescript-prefer-literal-enum-member)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-literal-enum-member.html#what-it-does)

Explicit enum value must only be a literal value (string, number, boolean, etc).

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-literal-enum-member.html#why-is-this-bad)

TypeScript allows the value of an enum member to be many different kinds of valid JavaScript expressions. However, because enums create their own scope whereby each enum member becomes a variable in that scope, developers are often surprised at the resultant values.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-literal-enum-member.html#examples)

Examples of **incorrect** code for this rule:

ts

```
const imOutside = 2;
const b = 2;
enum Foo {
  outer = imOutside,
  a = 1,
  b = a,
  c = b,
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-literal-enum-member.html#configuration)

This rule accepts a configuration object with the following properties:

### allowBitwiseExpressions [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-literal-enum-member.html#allowbitwiseexpressions)

type: `boolean`

default: `false`

When set to `true`, allows bitwise expressions in enum member initializers. This includes bitwise NOT (`~`), AND (`&`), OR (`|`), XOR (`^`), and shift operators (`<<`, `>>`, `>>>`).

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-literal-enum-member.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/prefer-literal-enum-member": "error"
  }
}
```

bash

```
oxlint --deny typescript/prefer-literal-enum-member
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-literal-enum-member.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/prefer_literal_enum_member.rs)
