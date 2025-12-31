---
title: "import/no-named-default | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/import/no-named-default"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/import/no-named-default.md for this page in Markdown format

# import/no-named-default Style [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-named-default.html#import-no-named-default)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-named-default.html#what-it-does)

Reports use of a default export as a locally named import.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-named-default.html#why-is-this-bad)

Rationale: the syntax exists to import default exports expressively, let's use it.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-named-default.html#examples)

Examples of **incorrect** code for this rule:

js

```
// message: Using exported name 'bar' as identifier for default export.
import { default as foo } from "./foo.js";
import { default as foo, bar } from "./foo.js";
```

Examples of **correct** code for this rule:

js

```
import foo from "./foo.js";
import foo, { bar } from "./foo.js";
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-named-default.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["import"],
  "rules": {
    "import/no-named-default": "error"
  }
}
```

bash

```
oxlint --deny import/no-named-default --import-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-named-default.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/import/no_named_default.rs)
