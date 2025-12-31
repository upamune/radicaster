---
title: "typescript/explicit-function-return-type | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/explicit-function-return-type"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/explicit-function-return-type.md for this page in Markdown format

# typescript/explicit-function-return-type Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/explicit-function-return-type.html#typescript-explicit-function-return-type)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/explicit-function-return-type.html#what-it-does)

This rule enforces that functions have an explicit return type annotation.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/explicit-function-return-type.html#why-is-this-bad)

Explicit return types make it clearer what type is returned by a function. Making the type returned by a function obvious allows the reader to infer what the function does and how it can be used from a quick glance.

Another benefit of explicit return types is the potential for a speed up of type checking in large codebases with many large functions.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/explicit-function-return-type.html#examples)

Examples of **incorrect** code for this rule:

ts

```
// Should indicate that no value is returned (void)
function test() {
  return;
}

// Should indicate that a number is returned
var fn = function () {
  return 1;
};

// Should indicate that a string is returned
var arrowFn = () => "test";

class Test {
  // Should indicate that no value is returned (void)
  method() {
    return;
  }
}
```

Examples of **correct** code for this rule:

ts

```
// No return value should be expected (void)
function test(): void {
  return;
}

// A return value of type number
var fn = function (): number {
  return 1;
};

// A return value of type string
var arrowFn = (): string => "test";

class Test {
  // No return value should be expected (void)
  method(): void {
    return;
  }
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/explicit-function-return-type.html#configuration)

This rule accepts a configuration object with the following properties:

### allowConciseArrowFunctionExpressionsStartingWithVoid [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/explicit-function-return-type.html#allowconcisearrowfunctionexpressionsstartingwithvoid)

type: `boolean`

default: `false`

Whether to allow concise arrow functions that start with the `void` keyword.

### allowDirectConstAssertionInArrowFunctions [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/explicit-function-return-type.html#allowdirectconstassertioninarrowfunctions)

type: `boolean`

default: `true`

Whether to allow arrow functions that use `as const` assertion on their return value.

### allowExpressions [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/explicit-function-return-type.html#allowexpressions)

type: `boolean`

default: `false`

Whether to allow expressions as function return types. When `true`, allows functions that immediately return an expression without a return type annotation.

### allowFunctionsWithoutTypeParameters [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/explicit-function-return-type.html#allowfunctionswithouttypeparameters)

type: `boolean`

default: `false`

Whether to allow functions that do not have generic type parameters.

### allowHigherOrderFunctions [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/explicit-function-return-type.html#allowhigherorderfunctions)

type: `boolean`

default: `true`

Whether to allow higher-order functions (functions that return another function) without return type annotations.

### allowIIFEs [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/explicit-function-return-type.html#allowiifes)

type: `boolean`

default: `false`

Whether to allow immediately invoked function expressions (IIFEs) without return type annotations.

### allowTypedFunctionExpressions [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/explicit-function-return-type.html#allowtypedfunctionexpressions)

type: `boolean`

default: `true`

Whether to allow typed function expressions. When `true`, allows function expressions that are assigned to a typed variable or parameter.

### allowedNames [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/explicit-function-return-type.html#allowednames)

type: `string[]`

default: `[]`

Array of function names that are exempt from requiring return type annotations.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/explicit-function-return-type.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/explicit-function-return-type": "error"
  }
}
```

bash

```
oxlint --deny typescript/explicit-function-return-type
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/explicit-function-return-type.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/explicit_function_return_type.rs)
