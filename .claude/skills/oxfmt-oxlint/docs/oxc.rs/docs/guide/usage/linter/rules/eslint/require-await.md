---
title: "eslint/require-await | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/require-await"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/require-await.md for this page in Markdown format

# eslint/require-await Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/require-await.html#eslint-require-await)

⚠️🛠️️ A dangerous auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/require-await.html#what-it-does)

Disallow async functions which have no `await` expression.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/require-await.html#why-is-this-bad)

Asynchronous functions in JavaScript behave differently than other functions in two important ways:

1. The return value is always a `Promise`.
2. You can use the `await` operator inside of them.

The primary reason to use asynchronous functions is typically to use the await operator, such as this:

js

```
async function fetchData(processDataItem) {
  const response = await fetch(DATA_URL);
  const data = await response.json();

  return data.map(processDataItem);
}
```

Asynchronous functions that don’t use await might not need to be asynchronous functions and could be the unintentional result of refactoring.

Note: this rule ignores async generator functions. This is because generators yield rather than return a value and async generators might yield all the values of another async generator without ever actually needing to use await.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/require-await.html#examples)

Examples of **incorrect** code for this rule:

js

```
async function foo() {
  doSomething();
}
```

Examples of **correct** code for this rule:

js

```
async function foo() {
  await doSomething();
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/require-await.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "require-await": "error"
  }
}
```

bash

```
oxlint --deny require-await
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/require-await.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/require_await.rs)
