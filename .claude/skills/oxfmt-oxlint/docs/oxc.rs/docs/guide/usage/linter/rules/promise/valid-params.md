---
title: "promise/valid-params | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/promise/valid-params"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/promise/valid-params.md for this page in Markdown format

# promise/valid-params Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/valid-params.html#promise-valid-params)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/valid-params.html#what-it-does)

Enforces the proper number of arguments are passed to Promise functions.

This rule is generally unnecessary if using TypeScript.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/valid-params.html#why-is-this-bad)

Calling a Promise function with the incorrect number of arguments can lead to unexpected behavior or hard to spot bugs.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/valid-params.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
Promise.resolve(1, 2);
```

Examples of **correct** code for this rule:

javascript

```
Promise.resolve(1);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/valid-params.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["promise"],
  "rules": {
    "promise/valid-params": "error"
  }
}
```

bash

```
oxlint --deny promise/valid-params --promise-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/valid-params.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/promise/valid_params.rs)
