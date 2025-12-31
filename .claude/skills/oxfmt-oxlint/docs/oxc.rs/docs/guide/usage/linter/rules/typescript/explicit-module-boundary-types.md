---
title: "typescript/explicit-module-boundary-types | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/explicit-module-boundary-types"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/explicit-module-boundary-types.md for this page in Markdown format

# typescript/explicit-module-boundary-types Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/explicit-module-boundary-types.html#typescript-explicit-module-boundary-types)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/explicit-module-boundary-types.html#what-it-does)

Require explicit return and argument types on exported functions' and classes' public class methods.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/explicit-module-boundary-types.html#why-is-this-bad)

Explicit types for function return values and arguments makes it clear to any calling code what is the module boundary's input and output. Adding explicit type annotations for those types can help improve code readability. It can also improve TypeScript type checking performance on larger codebases.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/explicit-module-boundary-types.html#examples)

Examples of **incorrect** code for this rule:

ts

```
// Should indicate that no value is returned (void)
export function test() {
  return;
}

// Should indicate that a string is returned
export var arrowFn = () => "test";

// All arguments should be typed
export var arrowFn = (arg): string => `test ${arg}`;
export var arrowFn = (arg: any): string => `test ${arg}`;

export class Test {
  // Should indicate that no value is returned (void)
  method() {
    return;
  }
}
```

Examples of **correct** code for this rule:

ts

```
// A function with no return value (void)
export function test(): void {
  return;
}

// A return value of type string
export var arrowFn = (): string => "test";

// All arguments should be typed
export var arrowFn = (arg: string): string => `test ${arg}`;
export var arrowFn = (arg: unknown): string => `test ${arg}`;

export class Test {
  // A class method with no return value (void)
  method(): void {
    return;
  }
}

// The function does not apply because it is not an exported function.
function test() {
  return;
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/explicit-module-boundary-types.html#configuration)

This rule accepts a configuration object with the following properties:

### allowArgumentsExplicitlyTypedAsAny [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/explicit-module-boundary-types.html#allowargumentsexplicitlytypedasany)

type: `boolean`

default: `false`

Whether to ignore arguments that are explicitly typed as `any`.

### allowDirectConstAssertionInArrowFunctions [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/explicit-module-boundary-types.html#allowdirectconstassertioninarrowfunctions)

type: `boolean`

default: `true`

Whether to ignore return type annotations on body-less arrow functions that return an `as const` type assertion. You must still type the parameters of the function.

### allowHigherOrderFunctions [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/explicit-module-boundary-types.html#allowhigherorderfunctions)

type: `boolean`

default: `true`

Whether to ignore return type annotations on functions immediately returning another function expression. You must still type the parameters of the function.

### allowOverloadFunctions [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/explicit-module-boundary-types.html#allowoverloadfunctions)

type: `boolean`

default: `false`

Whether to ignore return type annotations on functions with overload signatures.

### allowTypedFunctionExpressions [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/explicit-module-boundary-types.html#allowtypedfunctionexpressions)

type: `boolean`

default: `true`

Whether to ignore type annotations on the variable of a function expression.

### allowedNames [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/explicit-module-boundary-types.html#allowednames)

type: `string[]`

default: `[]`

An array of function/method names that will not have their arguments or return values checked.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/explicit-module-boundary-types.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/explicit-module-boundary-types": "error"
  }
}
```

bash

```
oxlint --deny typescript/explicit-module-boundary-types
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/explicit-module-boundary-types.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/explicit_module_boundary_types.rs)
