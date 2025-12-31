---
title: "promise/no-promise-in-callback | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/promise/no-promise-in-callback"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/promise/no-promise-in-callback.md for this page in Markdown format

# promise/no-promise-in-callback Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-promise-in-callback.html#promise-no-promise-in-callback)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-promise-in-callback.html#what-it-does)

Disallows the use of Promises within error-first callback functions.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-promise-in-callback.html#why-is-this-bad)

Mixing Promises and callbacks can lead to unclear and inconsistent code. Promises and callbacks are different patterns for handling asynchronous code. Mixing them makes the logic flow harder to follow and complicates error handling, as callbacks rely on an error-first pattern, while Promises use `catch`.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-promise-in-callback.html#examples)

Examples of **incorrect** code for this rule:

js

```
doSomething((err, val) => {
  if (err) console.error(err);
  else doSomethingElse(val).then(console.log);
});
```

Examples of **correct** code for this rule:

js

```
promisify(doSomething)().then(doSomethingElse).then(console.log).catch(console.error);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-promise-in-callback.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["promise"],
  "rules": {
    "promise/no-promise-in-callback": "error"
  }
}
```

bash

```
oxlint --deny promise/no-promise-in-callback --promise-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-promise-in-callback.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/promise/no_promise_in_callback.rs)
