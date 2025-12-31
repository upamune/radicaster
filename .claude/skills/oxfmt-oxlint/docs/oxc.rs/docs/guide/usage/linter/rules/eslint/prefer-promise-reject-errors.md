---
title: "eslint/prefer-promise-reject-errors | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-promise-reject-errors"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/prefer-promise-reject-errors.md for this page in Markdown format

# eslint/prefer-promise-reject-errors Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-promise-reject-errors.html#eslint-prefer-promise-reject-errors)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-promise-reject-errors.html#what-it-does)

Require using Error objects as Promise rejection reasons.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-promise-reject-errors.html#why-is-this-bad)

It is considered good practice to only pass instances of the built-in `Error` object to the `reject()` function for user-defined errors in Promises. `Error` objects automatically store a stack trace, which can be used to debug an error by determining where it came from. If a Promise is rejected with a non-`Error` value, it can be difficult to determine where the rejection occurred.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-promise-reject-errors.html#examples)

Examples of **incorrect** code for this rule:

js

```
Promise.reject("something bad happened");

Promise.reject(5);

Promise.reject();

new Promise(function (resolve, reject) {
  reject("something bad happened");
});

new Promise(function (resolve, reject) {
  reject();
});
```

Examples of **correct** code for this rule:

js

```
Promise.reject(new Error("something bad happened"));

Promise.reject(new TypeError("something bad happened"));

new Promise(function (resolve, reject) {
  reject(new Error("something bad happened"));
});

var foo = getUnknownValue();
Promise.reject(foo);
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-promise-reject-errors.html#configuration)

This rule accepts a configuration object with the following properties:

### allowEmptyReject [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-promise-reject-errors.html#allowemptyreject)

type: `boolean`

default: `false`

Whether to allow calls to `Promise.reject()` with no arguments.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-promise-reject-errors.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "prefer-promise-reject-errors": "error"
  }
}
```

bash

```
oxlint --deny prefer-promise-reject-errors
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-promise-reject-errors.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/prefer_promise_reject_errors.rs)
