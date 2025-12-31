---
title: "eslint/sort-vars | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/sort-vars"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/sort-vars.md for this page in Markdown format

# eslint/sort-vars Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/sort-vars.html#eslint-sort-vars)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/sort-vars.html#what-it-does)

When declaring multiple variables within the same block, sorting variable names make it easier to find necessary variable easier at a later time.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/sort-vars.html#why-is-this-bad)

Unsorted variable declarations can make the code harder to read and maintain.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/sort-vars.html#examples)

Examples of **incorrect** code for this rule:

js

```
var b, a;
var a, B, c;
```

Examples of **correct** code for this rule:

js

```
var a, b, c, d;
var B, a, c;
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/sort-vars.html#configuration)

This rule accepts a configuration object with the following properties:

### ignoreCase [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/sort-vars.html#ignorecase)

type: `boolean`

default: `false`

When `true`, the rule ignores case-sensitivity when sorting variables.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/sort-vars.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "sort-vars": "error"
  }
}
```

bash

```
oxlint --deny sort-vars
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/sort-vars.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/sort_vars.rs)
