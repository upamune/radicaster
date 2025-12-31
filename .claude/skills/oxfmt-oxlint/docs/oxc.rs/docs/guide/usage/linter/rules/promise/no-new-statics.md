---
title: "promise/no-new-statics | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/promise/no-new-statics"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/promise/no-new-statics.md for this page in Markdown format

# promise/no-new-statics Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-new-statics.html#promise-no-new-statics)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-new-statics.html#what-it-does)

Disallows calling new on static `Promise` methods.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-new-statics.html#why-is-this-bad)

Calling a static `Promise` method with `new` is invalid and will result in a `TypeError` at runtime.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-new-statics.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
const x = new Promise.resolve(value);
```

Examples of **correct** code for this rule:

javascript

```
const x = Promise.resolve(value);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-new-statics.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["promise"],
  "rules": {
    "promise/no-new-statics": "error"
  }
}
```

bash

```
oxlint --deny promise/no-new-statics --promise-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-new-statics.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/promise/no_new_statics.rs)
