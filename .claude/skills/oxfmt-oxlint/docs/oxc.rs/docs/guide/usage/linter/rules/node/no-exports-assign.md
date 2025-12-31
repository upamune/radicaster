---
title: "node/no-exports-assign | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/node/no-exports-assign"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/node/no-exports-assign.md for this page in Markdown format

# node/no-exports-assign Style [​](https://oxc.rs/docs/guide/usage/linter/rules/node/no-exports-assign.html#node-no-exports-assign)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/node/no-exports-assign.html#what-it-does)

Disallows assignment to `exports`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/node/no-exports-assign.html#why-is-this-bad)

Directly using `exports = {}` can lead to confusion and potential bugs because it reassigns the `exports` object, which may break module exports. It is more predictable and clearer to use `module.exports` directly or in conjunction with `exports`.

This rule is aimed at disallowing `exports = {}`, but allows `module.exports = exports = {}` to avoid conflict with `n/exports-style` rule's `allowBatchAssign` option.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/node/no-exports-assign.html#examples)

Examples of **incorrect** code for this rule:

js

```
exports = {};
```

Examples of **correct** code for this rule:

js

```
module.exports.foo = 1;
exports.bar = 2;
module.exports = {};

// allows `exports = {}` if along with `module.exports =`
module.exports = exports = {};
exports = module.exports = {};
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/node/no-exports-assign.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["node"],
  "rules": {
    "node/no-exports-assign": "error"
  }
}
```

bash

```
oxlint --deny node/no-exports-assign --node-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/node/no-exports-assign.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/node/no_exports_assign.rs)
