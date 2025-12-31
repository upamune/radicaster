---
title: "eslint/no-lone-blocks | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-lone-blocks"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-lone-blocks.md for this page in Markdown format

# eslint/no-lone-blocks Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-lone-blocks.html#eslint-no-lone-blocks)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-lone-blocks.html#what-it-does)

Disallows unnecessary standalone block statements.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-lone-blocks.html#why-is-this-bad)

Standalone blocks can be confusing as they do not provide any meaningful purpose when used unnecessarily. They may introduce extra nesting, reducing code readability, and can mislead readers about scope or intent.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-lone-blocks.html#examples)

Examples of **incorrect** code for this rule:

js

```
{
  var x = 1;
}
```

Examples of **correct** code for this rule:

js

```
if (condition) {
  var x = 1;
}

{
  let x = 1; // Used to create a valid block scope.
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-lone-blocks.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-lone-blocks": "error"
  }
}
```

bash

```
oxlint --deny no-lone-blocks
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-lone-blocks.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_lone_blocks.rs)
