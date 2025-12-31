---
title: "typescript/return-await | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/return-await"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/return-await.md for this page in Markdown format

# typescript/return-await Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/return-await.html#typescript-return-await)

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/return-await.html#what-it-does)

This rule enforces consistent returning of awaited values from async functions.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/return-await.html#why-is-this-bad)

There are different patterns for returning awaited values from async functions. Sometimes you want to await before returning (to handle errors in the current function), and sometimes you want to return the Promise directly (for better performance). This rule helps enforce consistency.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/return-await.html#examples)

Examples of **incorrect** code for this rule (depending on configuration):

ts

```
// If configured to require await:
async function fetchData() {
  return fetch("/api/data"); // Should be: return await fetch('/api/data');
}

async function processData() {
  return someAsyncOperation(); // Should be: return await someAsyncOperation();
}

// If configured to disallow unnecessary await:
async function fetchData() {
  return await fetch("/api/data"); // Should be: return fetch('/api/data');
}

async function processData() {
  return await someAsyncOperation(); // Should be: return someAsyncOperation();
}
```

Examples of **correct** code for this rule:

ts

```
// When await is required for error handling:
async function fetchData() {
  try {
    return await fetch("/api/data");
  } catch (error) {
    console.error("Fetch failed:", error);
    throw error;
  }
}

// When returning Promise directly for performance:
async function fetchData() {
  return fetch("/api/data");
}

// Processing before return requires await:
async function fetchAndProcess() {
  const response = await fetch("/api/data");
  return response.json();
}

// Multiple async operations:
async function multipleOperations() {
  const data1 = await fetchData1();
  const data2 = await fetchData2();
  return data1 + data2;
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/return-await.html#configuration)

This rule accepts one of the following string values:

### `"in-try-catch"` [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/return-await.html#in-try-catch)

Require `await` when returning Promises inside try/catch/finally blocks. This ensures proper error handling and stack traces.

### `"always"` [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/return-await.html#always)

Require `await` before returning Promises in all cases. Example: `return await Promise.resolve()` is required.

### `"error-handling-correctness-only"` [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/return-await.html#error-handling-correctness-only)

Require `await` only when it affects error handling correctness. Only flags cases where omitting await would change error handling behavior.

### `"never"` [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/return-await.html#never)

Disallow `await` before returning Promises in all cases. Example: `return Promise.resolve()` is required (no await).

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/return-await.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/return-await": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/return-await
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/return-await.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/return_await.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/return_await/return_await.go)
