---
title: "eslint/no-async-promise-executor | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-async-promise-executor"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-async-promise-executor.md for this page in Markdown format

# eslint/no-async-promise-executor Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-async-promise-executor.html#eslint-no-async-promise-executor)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-async-promise-executor.html#what-it-does)

Disallow using an async function as a Promise executor.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-async-promise-executor.html#why-is-this-bad)

The `new Promise` constructor accepts an executor function as an argument, which has `resolve` and `reject` parameters that can be used to control the state of the created Promise. For example:

javascript

```
const result = new Promise(function executor(resolve, reject) {
  readFile("foo.txt", function (err, result) {
    if (err) {
      reject(err);
    } else {
      resolve(result);
    }
  });
});
```

The executor function can also be an `async function`. However, this is usually a mistake, for a few reasons:

* If an async executor function throws an error, the error will be lost and won’t cause the newly-constructed `Promise` to reject.This could make it difficult to debug and handle some errors.
* If a `Promise` executor function is using `await`, this is usually a sign that it is not actually necessary to use the new `Promise` constructor, or the scope of the new `Promise` constructor can be reduced.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-async-promise-executor.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
const foo = new Promise(async (resolve, reject) => {
  readFile("foo.txt", function (err, result) {
    if (err) {
      reject(err);
    } else {
      resolve(result);
    }
  });
});

const result = new Promise(async (resolve, reject) => {
  resolve(await foo);
});
```

Examples of **correct** code for this rule:

javascript

```
const foo = new Promise((resolve, reject) => {
  readFile("foo.txt", function (err, result) {
    if (err) {
      reject(err);
    } else {
      resolve(result);
    }
  });
});

const result = Promise.resolve(foo);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-async-promise-executor.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-async-promise-executor": "error"
  }
}
```

bash

```
oxlint --deny no-async-promise-executor
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-async-promise-executor.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_async_promise_executor.rs)
