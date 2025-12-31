---
title: "unicorn/prefer-structured-clone | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-structured-clone"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-structured-clone.md for this page in Markdown format

# unicorn/prefer-structured-clone Style [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-structured-clone.html#unicorn-prefer-structured-clone)

💡 A suggestion is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-structured-clone.html#what-it-does)

Prefer using structuredClone to create a deep clone.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-structured-clone.html#why-is-this-bad)

structuredClone is the modern way to create a deep clone of a value.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-structured-clone.html#examples)

Examples of **incorrect** code for this rule:

js

```
const clone = JSON.parse(JSON.stringify(foo));

const clone = _.cloneDeep(foo);
```

Examples of **correct** code for this rule:

js

```
const clone = structuredClone(foo);
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-structured-clone.html#configuration)

This rule accepts a configuration object with the following properties:

### functions [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-structured-clone.html#functions)

type: `string[]`

default: `["cloneDeep", "utils.clone"]`

List of functions that are allowed to be used for deep cloning instead of structuredClone.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-structured-clone.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-structured-clone": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-structured-clone
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-structured-clone.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_structured_clone.rs)
