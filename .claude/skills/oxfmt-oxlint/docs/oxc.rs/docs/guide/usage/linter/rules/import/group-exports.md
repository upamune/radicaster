---
title: "import/group-exports | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/import/group-exports"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/import/group-exports.md for this page in Markdown format

# import/group-exports Style [​](https://oxc.rs/docs/guide/usage/linter/rules/import/group-exports.html#import-group-exports)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/import/group-exports.html#what-it-does)

Reports when named exports are not grouped together in a single export declaration or when multiple assignments to CommonJS module.exports or exports object are present in a single file.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/import/group-exports.html#why-is-this-bad)

An export declaration or module.exports assignment can appear anywhere in the code. By requiring a single export declaration all your exports will remain at one place, making it easier to see what exports a module provides.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/import/group-exports.html#examples)

Examples of **incorrect** code for this rule:

js

```
export const first = true;
export const second = true;
```

Examples of **correct** code for this rule:

js

```
const first = true;
const second = true;
export { first, second };
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/import/group-exports.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["import"],
  "rules": {
    "import/group-exports": "error"
  }
}
```

bash

```
oxlint --deny import/group-exports --import-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/import/group-exports.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/import/group_exports.rs)
