---
title: "typescript/no-inferrable-types | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-inferrable-types"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-inferrable-types.md for this page in Markdown format

# typescript/no-inferrable-types Style [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-inferrable-types.html#typescript-no-inferrable-types)

💡 A suggestion is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-inferrable-types.html#what-it-does)

Disallow explicit type declarations for variables or parameters initialized to a number, string, or boolean

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-inferrable-types.html#why-is-this-bad)

Explicitly typing variables or parameters that are initialized to a literal value is unnecessary because TypeScript can infer the type from the value.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-inferrable-types.html#examples)

Examples of **incorrect** code for this rule:

ts

```
const a: number = 5;
const b: string = "foo";
const c: boolean = true;
const fn = (a: number = 5, b: boolean = true, c: string = "foo") => {};
```

Examples of **correct** code for this rule:

ts

```
const a = 5;
const b = "foo";
const c = true;
const fn = (a = 5, b = true, c = "foo") => {};
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-inferrable-types.html#configuration)

This rule accepts a configuration object with the following properties:

### ignoreParameters [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-inferrable-types.html#ignoreparameters)

type: `boolean`

default: `false`

When set to `true`, ignores type annotations on function parameters.

### ignoreProperties [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-inferrable-types.html#ignoreproperties)

type: `boolean`

default: `false`

When set to `true`, ignores type annotations on class properties.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-inferrable-types.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-inferrable-types": "error"
  }
}
```

bash

```
oxlint --deny typescript/no-inferrable-types
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-inferrable-types.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_inferrable_types.rs)
