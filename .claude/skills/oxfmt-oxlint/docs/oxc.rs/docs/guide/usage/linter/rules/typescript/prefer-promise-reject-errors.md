---
title: "typescript/prefer-promise-reject-errors | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-promise-reject-errors"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/prefer-promise-reject-errors.md for this page in Markdown format

# typescript/prefer-promise-reject-errors Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-promise-reject-errors.html#typescript-prefer-promise-reject-errors)

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-promise-reject-errors.html#what-it-does)

This rule enforces passing an Error object to Promise.reject().

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-promise-reject-errors.html#why-is-this-bad)

It's considered good practice to only reject promises with Error objects. This is because Error objects automatically capture a stack trace, which is useful for debugging. Additionally, some tools and environments expect rejection reasons to be Error objects.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-promise-reject-errors.html#examples)

Examples of **incorrect** code for this rule:

ts

```
Promise.reject("error"); // rejecting with string

Promise.reject(42); // rejecting with number

Promise.reject(true); // rejecting with boolean

Promise.reject({ message: "error" }); // rejecting with plain object

Promise.reject(null); // rejecting with null

Promise.reject(); // rejecting with undefined

const error = "Something went wrong";
Promise.reject(error); // rejecting with non-Error variable
```

Examples of **correct** code for this rule:

ts

```
Promise.reject(new Error("Something went wrong"));

Promise.reject(new TypeError("Invalid type"));

Promise.reject(new RangeError("Value out of range"));

// Custom Error subclasses
class CustomError extends Error {
  constructor(message: string) {
    super(message);
    this.name = "CustomError";
  }
}
Promise.reject(new CustomError("Custom error occurred"));

// Variables that are Error objects
const error = new Error("Error message");
Promise.reject(error);
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-promise-reject-errors.html#configuration)

This rule accepts a configuration object with the following properties:

### allowEmptyReject [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-promise-reject-errors.html#allowemptyreject)

type: `boolean`

default: `false`

Whether to allow calling `Promise.reject()` with no arguments.

### allowThrowingAny [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-promise-reject-errors.html#allowthrowingany)

type: `boolean`

default: `false`

Whether to allow rejecting Promises with values typed as `any`.

### allowThrowingUnknown [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-promise-reject-errors.html#allowthrowingunknown)

type: `boolean`

default: `false`

Whether to allow rejecting Promises with values typed as `unknown`.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-promise-reject-errors.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/prefer-promise-reject-errors": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/prefer-promise-reject-errors
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-promise-reject-errors.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/prefer_promise_reject_errors.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/prefer_promise_reject_errors/prefer_promise_reject_errors.go)
