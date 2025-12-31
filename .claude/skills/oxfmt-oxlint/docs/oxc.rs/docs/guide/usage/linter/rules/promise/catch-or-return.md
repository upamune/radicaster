---
title: "promise/catch-or-return | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/promise/catch-or-return"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/promise/catch-or-return.md for this page in Markdown format

# promise/catch-or-return Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/catch-or-return.html#promise-catch-or-return)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/catch-or-return.html#what-it-does)

Ensure that each time a `then()` is applied to a promise, a `catch()` must be applied as well. Exceptions are made for promises returned from a function.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/catch-or-return.html#why-is-this-bad)

Not catching errors in a promise can cause hard to debug problems or missing handling of error conditions. In the worst case, unhandled promise rejections can cause your application to crash.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/catch-or-return.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
myPromise.then(doSomething);
myPromise.then(doSomething, catchErrors); // catch() may be a little better
```

Examples of **correct** code for this rule:

javascript

```
myPromise.then(doSomething).catch(errors);
function doSomethingElse() {
  return myPromise.then(doSomething);
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/catch-or-return.html#configuration)

This rule accepts a configuration object with the following properties:

### allowFinally [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/catch-or-return.html#allowfinally)

type: `boolean`

default: `false`

Whether to allow `finally()` as a termination method.

### allowThen [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/catch-or-return.html#allowthen)

type: `boolean`

default: `false`

Whether to allow `then()` with two arguments as a termination method.

### terminationMethod [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/catch-or-return.html#terminationmethod)

type: `string[]`

default: `["catch"]`

List of allowed termination methods (e.g., `catch`, `done`).

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/catch-or-return.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["promise"],
  "rules": {
    "promise/catch-or-return": "error"
  }
}
```

bash

```
oxlint --deny promise/catch-or-return --promise-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/catch-or-return.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/promise/catch_or_return.rs)
