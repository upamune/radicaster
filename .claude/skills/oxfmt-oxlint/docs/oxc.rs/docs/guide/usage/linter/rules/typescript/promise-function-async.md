---
title: "typescript/promise-function-async | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/promise-function-async"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/promise-function-async.md for this page in Markdown format

# typescript/promise-function-async Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/promise-function-async.html#typescript-promise-function-async)

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/promise-function-async.html#what-it-does)

This rule requires any function or method that returns a Promise to be marked as async.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/promise-function-async.html#why-is-this-bad)

Functions that return Promises should typically be marked as `async` to make their asynchronous nature clear and to enable the use of `await` within them. This makes the code more readable and helps prevent common mistakes with Promise handling.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/promise-function-async.html#examples)

Examples of **incorrect** code for this rule:

ts

```
// Function returning Promise without async
function fetchData(): Promise<string> {
  return fetch("/api/data").then((res) => res.text());
}

// Method returning Promise without async
class DataService {
  getData(): Promise<any> {
    return fetch("/api/data").then((res) => res.json());
  }
}

// Arrow function returning Promise without async
const processData = (): Promise<void> => {
  return Promise.resolve();
};
```

Examples of **correct** code for this rule:

ts

```
// Async function
async function fetchData(): Promise<string> {
  const response = await fetch("/api/data");
  return response.text();
}

// Async method
class DataService {
  async getData(): Promise<any> {
    const response = await fetch("/api/data");
    return response.json();
  }
}

// Async arrow function
const processData = async (): Promise<void> => {
  await someAsyncOperation();
};

// Functions that don't return Promise are fine
function syncFunction(): string {
  return "hello";
}

// Functions returning Promise-like but not actual Promise
function createThenable(): { then: Function } {
  return { then: () => {} };
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/promise-function-async.html#configuration)

This rule accepts a configuration object with the following properties:

### allowAny [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/promise-function-async.html#allowany)

type: `boolean`

default: `true`

Whether to allow functions returning `any` type without requiring `async`.

### allowedPromiseNames [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/promise-function-async.html#allowedpromisenames)

type: `string[]`

default: `[]`

A list of Promise type names that are allowed without requiring `async`. Example: `["SpecialPromise"]` to allow functions returning `SpecialPromise` without `async`.

### checkArrowFunctions [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/promise-function-async.html#checkarrowfunctions)

type: `boolean`

default: `true`

Whether to check arrow functions for missing `async` keyword.

### checkFunctionDeclarations [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/promise-function-async.html#checkfunctiondeclarations)

type: `boolean`

default: `true`

Whether to check function declarations for missing `async` keyword.

### checkFunctionExpressions [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/promise-function-async.html#checkfunctionexpressions)

type: `boolean`

default: `true`

Whether to check function expressions for missing `async` keyword.

### checkMethodDeclarations [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/promise-function-async.html#checkmethoddeclarations)

type: `boolean`

default: `true`

Whether to check method declarations for missing `async` keyword.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/promise-function-async.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/promise-function-async": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/promise-function-async
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/promise-function-async.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/promise_function_async.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/promise_function_async/promise_function_async.go)
