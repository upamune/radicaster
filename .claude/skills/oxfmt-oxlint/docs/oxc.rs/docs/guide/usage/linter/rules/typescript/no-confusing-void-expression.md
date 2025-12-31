---
title: "typescript/no-confusing-void-expression | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-confusing-void-expression"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-confusing-void-expression.md for this page in Markdown format

# typescript/no-confusing-void-expression Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-confusing-void-expression.html#typescript-no-confusing-void-expression)

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-confusing-void-expression.html#what-it-does)

This rule forbids using void expressions in confusing locations such as arrow function returns.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-confusing-void-expression.html#why-is-this-bad)

The void operator is useful when you want to execute an expression while evaluating to undefined. However, it can be confusing when used in places where the return value is meaningful, particularly in arrow functions and conditional expressions.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-confusing-void-expression.html#examples)

Examples of **incorrect** code for this rule:

ts

```
// arrow function returning void expression
const foo = () => void bar();

// conditional expression
const result = condition ? void foo() : bar();

// void in conditional
if (void foo()) {
  // ...
}
```

Examples of **correct** code for this rule:

ts

```
// proper use of void
void foo();

// explicit return statement
const foo = () => {
  bar();
  return;
};

// statement expression
foo();

// IIFE with void
void (function () {
  console.log("immediately invoked");
})();
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-confusing-void-expression.html#configuration)

This rule accepts a configuration object with the following properties:

### ignoreArrowShorthand [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-confusing-void-expression.html#ignorearrowshorthand)

type: `boolean`

default: `false`

Whether to ignore arrow function shorthand that returns void. When true, allows expressions like `() => someVoidFunction()`.

### ignoreVoidOperator [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-confusing-void-expression.html#ignorevoidoperator)

type: `boolean`

default: `false`

Whether to ignore expressions using the void operator. When true, allows `void someExpression`.

### ignoreVoidReturningFunctions [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-confusing-void-expression.html#ignorevoidreturningfunctions)

type: `boolean`

default: `false`

Whether to ignore calling functions that are declared to return void. When true, allows expressions like `x = voidReturningFunction()`.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-confusing-void-expression.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-confusing-void-expression": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/no-confusing-void-expression
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-confusing-void-expression.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_confusing_void_expression.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/no_confusing_void_expression/no_confusing_void_expression.go)
