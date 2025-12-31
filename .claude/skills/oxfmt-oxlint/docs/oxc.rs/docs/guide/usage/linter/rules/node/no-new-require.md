---
title: "node/no-new-require | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/node/no-new-require"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/node/no-new-require.md for this page in Markdown format

# node/no-new-require Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/node/no-new-require.html#node-no-new-require)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/node/no-new-require.html#what-it-does)

Warn about calling `new` on `require`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/node/no-new-require.html#why-is-this-bad)

The `require` function is used to include modules and might return a constructor. As this is not always the case this can be confusing.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/node/no-new-require.html#examples)

Examples of **incorrect** code for this rule:

js

```
var appHeader = new require("app-header");
```

Examples of **correct** code for this rule:

js

```
var AppHeader = require("app-header");
var appHeader = new AppHeader();
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/node/no-new-require.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["node"],
  "rules": {
    "node/no-new-require": "error"
  }
}
```

bash

```
oxlint --deny node/no-new-require --node-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/node/no-new-require.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/node/no_new_require.rs)
