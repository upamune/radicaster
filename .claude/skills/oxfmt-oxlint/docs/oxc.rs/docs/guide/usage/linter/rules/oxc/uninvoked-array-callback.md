---
title: "oxc/uninvoked-array-callback | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/oxc/uninvoked-array-callback"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/oxc/uninvoked-array-callback.md for this page in Markdown format

# oxc/uninvoked-array-callback Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/uninvoked-array-callback.html#oxc-uninvoked-array-callback)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/uninvoked-array-callback.html#what-it-does)

This rule applies when an Array function has a callback argument used for an array with empty slots.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/uninvoked-array-callback.html#why-is-this-bad)

When the Array constructor is called with a single number argument, an array with the specified number of empty slots (not actual undefined values) is constructed. If a callback function is passed to the function of this array, the callback function is never invoked because the array has no actual elements.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/uninvoked-array-callback.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
const list = new Array(5).map((_) => createElement());
```

Examples of **correct** code for this rule:

javascript

```
const list = new Array(5).fill().map((_) => createElement());
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/uninvoked-array-callback.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "oxc/uninvoked-array-callback": "error"
  }
}
```

bash

```
oxlint --deny oxc/uninvoked-array-callback
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/uninvoked-array-callback.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/oxc/uninvoked_array_callback.rs)
