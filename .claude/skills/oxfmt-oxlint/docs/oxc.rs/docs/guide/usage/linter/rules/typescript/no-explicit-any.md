---
title: "typescript/no-explicit-any | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-explicit-any"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-explicit-any.md for this page in Markdown format

# typescript/no-explicit-any Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-explicit-any.html#typescript-no-explicit-any)

🛠️ An auto-fix is available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-explicit-any.html#what-it-does)

Disallows explicit use of the `any` type.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-explicit-any.html#why-is-this-bad)

The `any` type in TypeScript is a dangerous "escape hatch" from the type system. Using `any` disables many type checking rules and is generally best used only as a last resort or when prototyping code. This rule reports on explicit uses of the `any` keyword as a type annotation.

> TypeScript's `--noImplicitAny` compiler option prevents an implied `any`, but doesn't prevent `any` from being explicitly used the way this rule does.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-explicit-any.html#examples)

Examples of **incorrect** code for this rule:

typescript

```
const age: any = "seventeen";
const ages: any[] = ["seventeen"];
const ages: Array<any> = ["seventeen"];
function greet(): any {}
function greet(): any[] {}
function greet(): Array<any> {}
function greet(): Array<Array<any>> {}
function greet(param: Array<any>): string {}
function greet(param: Array<any>): Array<any> {}
```

Examples of **correct** code for this rule:

typescript

```
const age: number = 17;
const ages: number[] = [17];
const ages: Array<number> = [17];
function greet(): string {}
function greet(): string[] {}
function greet(): Array<string> {}
function greet(): Array<Array<string>> {}
function greet(param: Array<string>): string {}
function greet(param: Array<string>): Array<string> {}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-explicit-any.html#configuration)

This rule accepts a configuration object with the following properties:

### fixToUnknown [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-explicit-any.html#fixtounknown)

type: `boolean`

default: `false`

Whether to enable auto-fixing in which the `any` type is converted to the `unknown` type.

### ignoreRestArgs [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-explicit-any.html#ignorerestargs)

type: `boolean`

default: `false`

Whether to ignore rest parameter arrays.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-explicit-any.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-explicit-any": "error"
  }
}
```

bash

```
oxlint --deny typescript/no-explicit-any
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-explicit-any.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_explicit_any.rs)
