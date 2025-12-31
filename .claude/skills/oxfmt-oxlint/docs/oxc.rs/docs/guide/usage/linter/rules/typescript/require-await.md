---
title: "typescript/require-await | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/require-await"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/require-await.md for this page in Markdown format

# typescript/require-await Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/require-await.html#typescript-require-await)

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/require-await.html#what-it-does)

This rule disallows async functions which do not have an await expression.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/require-await.html#why-is-this-bad)

Async functions that don't use await are usually a mistake. They return a Promise unnecessarily and can often be converted to regular functions. This can improve performance and make the code clearer.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/require-await.html#examples)

Examples of **incorrect** code for this rule:

ts

```
// Async function without await
async function fetchData() {
  return fetch("/api/data");
}

// Async arrow function without await
const processData = async () => {
  return someData.map((x) => x * 2);
};

// Async method without await
class DataService {
  async getData() {
    return this.data;
  }
}

// Async function that returns Promise but doesn't await
async function getPromise() {
  return Promise.resolve("value");
}
```

Examples of **correct** code for this rule:

ts

```
// Async function with await
async function fetchData() {
  const response = await fetch("/api/data");
  return response.json();
}

// Regular function returning Promise
function fetchDataSync() {
  return fetch("/api/data");
}

// Async function with await in conditional
async function conditionalAwait(condition: boolean) {
  if (condition) {
    return await someAsyncOperation();
  }
  return "default";
}

// Async function with await in loop
async function processItems(items: string[]) {
  const results = [];
  for (const item of items) {
    results.push(await processItem(item));
  }
  return results;
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/require-await.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/require-await": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/require-await
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/require-await.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/require_await.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/require_await/require_await.go)
