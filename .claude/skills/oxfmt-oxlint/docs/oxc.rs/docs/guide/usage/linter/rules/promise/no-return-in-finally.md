---
title: "promise/no-return-in-finally | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/promise/no-return-in-finally"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/promise/no-return-in-finally.md for this page in Markdown format

# promise/no-return-in-finally Nursery [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-return-in-finally.html#promise-no-return-in-finally)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-return-in-finally.html#what-it-does)

Disallow return statements in a finally() callback of a promise.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-return-in-finally.html#why-is-this-bad)

Disallow return statements inside a callback passed to finally(), since nothing would consume what's returned.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-return-in-finally.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
myPromise.finally(function (val) {
  return val;
});
```

Examples of **correct** code for this rule:

javascript

```
Promise.resolve(1).finally(() => {
  console.log(2);
});
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-return-in-finally.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["promise"],
  "rules": {
    "promise/no-return-in-finally": "error"
  }
}
```

bash

```
oxlint --deny promise/no-return-in-finally --promise-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-return-in-finally.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/promise/no_return_in_finally.rs)
