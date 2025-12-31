---
title: "node/no-process-env | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/node/no-process-env"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/node/no-process-env.md for this page in Markdown format

# node/no-process-env Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/node/no-process-env.html#node-no-process-env)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/node/no-process-env.html#what-it-does)

Disallows use of `process.env`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/node/no-process-env.html#why-is-this-bad)

Directly reading `process.env` can lead to implicit runtime configuration, make code harder to test, and bypass configuration validation.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/node/no-process-env.html#examples)

Examples of **incorrect** code for this rule:

js

```
if (process.env.NODE_ENV === "development") {
  // ...
}
```

Examples of **correct** code for this rule:

js

```
import config from "./config";

if (config.env === "development") {
  //...
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/node/no-process-env.html#configuration)

This rule accepts a configuration object with the following properties:

### allowedVariables [​](https://oxc.rs/docs/guide/usage/linter/rules/node/no-process-env.html#allowedvariables)

type: `string[]`

default: `[]`

Variable names which are allowed to be accessed on `process.env`.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/node/no-process-env.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["node"],
  "rules": {
    "node/no-process-env": "error"
  }
}
```

bash

```
oxlint --deny node/no-process-env --node-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/node/no-process-env.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/node/no_process_env.rs)
