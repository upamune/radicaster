---
title: "promise/prefer-await-to-callbacks | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/promise/prefer-await-to-callbacks"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/promise/prefer-await-to-callbacks.md for this page in Markdown format

# promise/prefer-await-to-callbacks Style [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/prefer-await-to-callbacks.html#promise-prefer-await-to-callbacks)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/prefer-await-to-callbacks.html#what-it-does)

The rule encourages the use of `async/await` for handling asynchronous code instead of traditional callback functions. `async/await`, introduced in ES2017, provides a clearer and more concise syntax for writing asynchronous code, making it easier to read and maintain.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/prefer-await-to-callbacks.html#why-is-this-bad)

Using callbacks can lead to complex, nested structures known as "callback hell," which make code difficult to read and maintain. Additionally, error handling can become cumbersome with callbacks, whereas `async/await` allows for more straightforward try/catch blocks for managing errors.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/prefer-await-to-callbacks.html#examples)

Examples of **incorrect** code for this rule:

js

```
cb();
callback();
doSomething(arg, (err) => {});
function doSomethingElse(cb) {}
```

Examples of **correct** code for this rule:

js

```
await doSomething(arg);
async function doSomethingElse() {}
function* generator() {
  yield yieldValue((err) => {});
}
eventEmitter.on("error", (err) => {});
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/prefer-await-to-callbacks.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["promise"],
  "rules": {
    "promise/prefer-await-to-callbacks": "error"
  }
}
```

bash

```
oxlint --deny promise/prefer-await-to-callbacks --promise-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/prefer-await-to-callbacks.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/promise/prefer_await_to_callbacks.rs)
