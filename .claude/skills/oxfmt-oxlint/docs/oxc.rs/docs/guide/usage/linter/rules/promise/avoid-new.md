---
title: "promise/avoid-new | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/promise/avoid-new"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/promise/avoid-new.md for this page in Markdown format

# promise/avoid-new Style [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/avoid-new.html#promise-avoid-new)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/avoid-new.html#what-it-does)

Disallow creating promises with `new Promise()`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/avoid-new.html#why-is-this-bad)

Many cases that use `new Promise()` could be refactored to use an `async` function. `async` is considered more idiomatic in modern JavaScript.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/avoid-new.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
function foo() {
  return new Promise((resolve, reject) => {
    /* ... */
  });
}
```

Examples of **correct** code for this rule:

javascript

```
async function foo() {
  // ...
}
const bar = await Promise.all([baz(), bang()]);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/avoid-new.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["promise"],
  "rules": {
    "promise/avoid-new": "error"
  }
}
```

bash

```
oxlint --deny promise/avoid-new --promise-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/avoid-new.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/promise/avoid_new.rs)
