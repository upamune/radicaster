---
title: "typescript/no-var-requires | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-var-requires"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-var-requires.md for this page in Markdown format

# typescript/no-var-requires Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-var-requires.html#typescript-no-var-requires)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-var-requires.html#what-it-does)

Disallow `require` statements except in import statements

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-var-requires.html#why-is-this-bad)

In other words, the use of forms such as var foo = require("foo") are banned. Instead use ES module imports or import foo = require("foo") imports.

typescript

```
var foo = require("foo");
const foo = require("foo");
let foo = require("foo");
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-var-requires.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-var-requires": "error"
  }
}
```

bash

```
oxlint --deny typescript/no-var-requires
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-var-requires.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_var_requires.rs)
