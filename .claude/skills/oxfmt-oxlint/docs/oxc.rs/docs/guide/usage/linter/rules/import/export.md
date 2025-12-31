---
title: "import/export | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/import/export"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/import/export.md for this page in Markdown format

# import/export Nursery [​](https://oxc.rs/docs/guide/usage/linter/rules/import/export.html#import-export)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/import/export.html#what-it-does)

Reports funny business with exports, like repeated exports of names or defaults.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/import/export.html#why-is-this-bad)

Having multiple exports of the same name can lead to ambiguity and confusion in the codebase. It makes it difficult to track which export is being used and can result in runtime errors if the wrong export is referenced.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/import/export.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
let foo;
export { foo }; // Multiple exports of name 'foo'.
export * from "./export-all"; // Conflicts if export-all.js also exports foo
```

Examples of **correct** code for this rule:

javascript

```
let foo;
export { foo as foo1 }; // Renamed export to avoid conflict
export * from "./export-all"; // No conflict if export-all.js also exports foo
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/import/export.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["import"],
  "rules": {
    "import/export": "error"
  }
}
```

bash

```
oxlint --deny import/export --import-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/import/export.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/import/export.rs)
