---
title: "promise/spec-only | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/promise/spec-only"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/promise/spec-only.md for this page in Markdown format

# promise/spec-only Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/spec-only.html#promise-spec-only)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/spec-only.html#what-it-does)

Disallow use of non-standard Promise static methods.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/spec-only.html#why-is-this-bad)

Non-standard Promises may cost more maintenance work.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/spec-only.html#examples)

Examples of **incorrect** code for this rule:

js

```
Promise.done();
```

Examples of **correct** code for this rule:

js

```
Promise.resolve();
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/spec-only.html#configuration)

This rule accepts a configuration object with the following properties:

### allowedMethods [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/spec-only.html#allowedmethods)

type: `string[]`

default: `null`

List of Promise static methods that are allowed to be used.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/spec-only.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["promise"],
  "rules": {
    "promise/spec-only": "error"
  }
}
```

bash

```
oxlint --deny promise/spec-only --promise-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/spec-only.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/promise/spec_only.rs)
