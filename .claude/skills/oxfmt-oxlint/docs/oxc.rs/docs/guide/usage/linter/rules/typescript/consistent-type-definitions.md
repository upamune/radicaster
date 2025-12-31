---
title: "typescript/consistent-type-definitions | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-type-definitions"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/consistent-type-definitions.md for this page in Markdown format

# typescript/consistent-type-definitions Style [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-type-definitions.html#typescript-consistent-type-definitions)

⚠️🛠️️ A dangerous auto-fix is available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-type-definitions.html#what-it-does)

Enforce type definitions to consistently use either `interface` or `type`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-type-definitions.html#why-is-this-bad)

TypeScript provides two common ways to define an object type: `interface` and `type`. The two are generally very similar, and can often be used interchangeably. Using the same type declaration style consistently helps with code readability.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-type-definitions.html#examples)

By default this rule enforces the use of `interface` for defining object types.

Examples of **incorrect** code for this rule:

typescript

```
type T = { x: number };
```

Examples of **correct** code for this rule:

typescript

```
type T = string;
type Foo = string | {};

interface T {
  x: number;
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-type-definitions.html#configuration)

This rule accepts one of the following string values:

### `"interface"` [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-type-definitions.html#interface)

Prefer `interface` over `type` for object type definitions:

typescript

```
interface T {
  x: number;
}
```

### `"type"` [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-type-definitions.html#type)

Prefer `type` over `interface` for object type definitions:

typescript

```
type T = { x: number };
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-type-definitions.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/consistent-type-definitions": "error"
  }
}
```

bash

```
oxlint --deny typescript/consistent-type-definitions
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-type-definitions.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/consistent_type_definitions.rs)
