---
title: "import/exports-last | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/import/exports-last"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/import/exports-last.md for this page in Markdown format

# import/exports-last Style [​](https://oxc.rs/docs/guide/usage/linter/rules/import/exports-last.html#import-exports-last)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/import/exports-last.html#what-it-does)

This rule enforces that all exports are declared at the bottom of the file. This rule will report any export declarations that comes before any non-export statements.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/import/exports-last.html#why-is-this-bad)

Exports scattered throughout the file can lead to poor code readability and increase the cost of locating the export quickly

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/import/exports-last.html#examples)

Examples of **incorrect** code for this rule:

js

```
const bool = true;
export const foo = "bar";
const str = "foo";
```

Examples of **correct** code for this rule:

js

```
const arr = ["bar"];
export const bool = true;
export const str = "foo";
export function func() {
  console.log("Hello World");
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/import/exports-last.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["import"],
  "rules": {
    "import/exports-last": "error"
  }
}
```

bash

```
oxlint --deny import/exports-last --import-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/import/exports-last.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/import/exports_last.rs)
