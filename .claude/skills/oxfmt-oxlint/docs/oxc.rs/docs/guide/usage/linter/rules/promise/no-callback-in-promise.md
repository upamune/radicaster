---
title: "promise/no-callback-in-promise | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/promise/no-callback-in-promise"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/promise/no-callback-in-promise.md for this page in Markdown format

# promise/no-callback-in-promise Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-callback-in-promise.html#promise-no-callback-in-promise)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-callback-in-promise.html#what-it-does)

Disallows calling a callback function (`cb()`) inside a `Promise.prototype.then()` or `Promise.prototype.catch()`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-callback-in-promise.html#why-is-this-bad)

Directly invoking a callback inside a `then()` or `catch()` method can lead to unexpected behavior, such as the callback being called multiple times. Additionally, mixing the callback and promise paradigms in this way can make the code confusing and harder to maintain.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-callback-in-promise.html#examples)

Examples of **incorrect** code for this rule:

js

```
function callback(err, data) {
  console.log("Callback got called with:", err, data);
  throw new Error("My error");
}

Promise.resolve()
  .then(() => callback(null, "data"))
  .catch((err) => callback(err.message, null));
```

Examples of **correct** code for this rule:

js

```
Promise.resolve()
  .then((data) => {
    console.log(data);
  })
  .catch((err) => {
    console.error(err);
  });
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-callback-in-promise.html#configuration)

This rule accepts a configuration object with the following properties:

### callbacks [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-callback-in-promise.html#callbacks)

type: `string[]`

default: `["callback", "cb", "done", "next"]`

List of callback function names to check for within Promise `then` and `catch` methods.

### exceptions [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-callback-in-promise.html#exceptions)

type: `string[]`

default: `[]`

List of callback function names to allow within Promise `then` and `catch` methods.

### timeoutsErr [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-callback-in-promise.html#timeoutserr)

type: `boolean`

default: `false`

Boolean as to whether callbacks in timeout functions like `setTimeout` will err.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-callback-in-promise.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["promise"],
  "rules": {
    "promise/no-callback-in-promise": "error"
  }
}
```

bash

```
oxlint --deny promise/no-callback-in-promise --promise-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-callback-in-promise.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/promise/no_callback_in_promise.rs)
