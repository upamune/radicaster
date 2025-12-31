---
title: "import/no-empty-named-blocks | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/import/no-empty-named-blocks"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/import/no-empty-named-blocks.md for this page in Markdown format

# import/no-empty-named-blocks Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-empty-named-blocks.html#import-no-empty-named-blocks)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-empty-named-blocks.html#what-it-does)

Enforces that named import blocks are not empty.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-empty-named-blocks.html#why-is-this-bad)

Empty named imports serve no practical purpose and often result from accidental deletions or tool-generated code.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-empty-named-blocks.html#examples)

Examples of **incorrect** code for this rule:

js

```
import {} from "mod";
import Default from "mod";
```

Examples of **correct** code for this rule:

js

```
import { mod } from "mod";
import Default, { mod } from "mod";
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-empty-named-blocks.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["import"],
  "rules": {
    "import/no-empty-named-blocks": "error"
  }
}
```

bash

```
oxlint --deny import/no-empty-named-blocks --import-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-empty-named-blocks.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/import/no_empty_named_blocks.rs)
